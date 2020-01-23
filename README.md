# metals rates
**build**
go mod tindy
go mod vender
go build
**run**
./merals-rates

**API**

    GET: /latest
    GET: /latest/:date
    PUT: /latest/:date
    DELETE: /delete/:date

example.

    http://localhost:1323/delete/20200124
    http://localhost:1323/latest/20200124
    http://localhost:1323/latest
