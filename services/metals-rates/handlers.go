package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

var url = "https://data-asg.goldprice.org/dbXRates/USD"

// get all
func getAllHandler(c echo.Context) error {
	root := "./tmp/data/"
	files, err := ioutil.ReadDir(root)
	filesname := []string{}
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
		filesname = append(filesname, f.Name())
	}
	urlsJson, _ := json.Marshal(filesname)
	return c.JSON(http.StatusOK, string(urlsJson))
}

// handler get date
func getHandler(c echo.Context) error {
	var byteValue []byte
	var golds Golds
	// genarate file name
	date := c.Param("date")
	filename := "./tmp/data/" + date + ".json"
	// check file data
	if _, err := os.Stat(filename); err == nil {
		// exists
		byteValue = getLocalFile(filename)
		json.Unmarshal(byteValue, &golds)
		print(golds)
	} else if os.IsNotExist(err) {
		// dose not exits
		byteValue = getGoldPrice()
		json.Unmarshal(byteValue, &golds)
		print(golds)
		err := ioutil.WriteFile(filename, byteValue, 0644)
		if checkErr(err) {
			return c.JSON(http.StatusFound, "")
		}
	}

	return c.JSON(http.StatusOK, golds)
}

func updateHandler(c echo.Context) error {
	var byteValue []byte
	var golds Golds

	date := c.Param("date")
	filename := "./tmp/data/" + date + ".json"
	byteValue = getGoldPrice()
	json.Unmarshal(byteValue, &golds)
	print(golds)
	err := ioutil.WriteFile(filename, byteValue, 0644)
	if checkErr(err) {
		return c.JSON(http.StatusFound, "")
	}

	msg := Message{
		Message: "success",
	}
	return c.JSON(http.StatusAccepted, msg)
}

func deteteHandle(c echo.Context) error {
	date := c.Param("date")
	root := "./tmp/data/"
	filename := root + date + ".json"
	fmt.Println(filename)

	if _, err := os.Stat(filename); err == nil {
		// exists
		var err = os.Remove(filename)
		if checkErr(err) {
			return c.JSON(http.StatusFound, "")
		}
	} else if os.IsNotExist(err) {
		// dose not exits
		msg := Message{
			Message: "file not found",
		}
		return c.JSON(http.StatusFound, msg)
	}

	fmt.Println("File Deleted")
	msg := Message{
		Message: "success",
	}
	return c.JSON(http.StatusAccepted, msg)
}

func print(golds Golds) {
	fmt.Println("Date:", golds.Date)
	for i := 0; i < len(golds.Items); i++ {
		fmt.Println("Currency:", golds.Items[i].Curr)
		fmt.Println("Xau price:", golds.Items[i].XauPrice)
		fmt.Println("xag Price:", golds.Items[i].XagPrice)
	}
	fmt.Println("-----------")
}

func getLocalFile(filename string) []byte {
	jsonFile, err := os.Open(filename)
	checkErr(err)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func getGoldPrice() []byte {
	url := "https://data-asg.goldprice.org/dbXRates/USD"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	byteValue, err := ioutil.ReadAll(res.Body)
	return byteValue
}

func checkErr(e error) bool {
	if e != nil {
		log.Fatal(e)
		return true
	}
	return false
}

type Golds struct {
	Ts    int     `json:"ts"`
	Tsj   int     `json:"tsj"`
	Date  string  `json:"date"`
	Items []Items `json:"items"`
}

type Items struct {
	Curr     string  `json:"curr"`
	XauPrice float32 `json:"xauPrice"`
	XagPrice float32 `json:"xagPrice"`
	ChgXau   float32 `json:"chgXau"`
	ChgXag   float32 `json:"chgXag"`
	PcXau    float32 `json:"pcXau"`
	PcXag    float32 `json:"pcXag"`
	XauClose float32 `json:"xauClose"`
	XagClose float32 `json:"xagClose"`
}

type Message struct {
	Message string
}
