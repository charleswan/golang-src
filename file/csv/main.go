package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Payment ...
type Payment struct {
	TopNumber      string
	PaymentMethond string
	CustomerID     string
	CountryCode    string
	Currency       string
	Amount         int64
	Fee            int64
	CreateTime     int64
	OCreateTime    int64
	CreateBy       string
	Status         string
	Flag           bool
}

func doFile() (payments []*Payment) {
	filePath := "./test.csv"
	fileObj, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}
	csvReader := csv.NewReader(fileObj)
	var headers []string
	m := make(map[string]string)

	for {
		ctx, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			break
		} else {
			if headers == nil {
				headers = ctx
			} else {
				for i, v := range ctx {
					m[headers[i]] = v
				}
				if m["Status"] != "paid" {
					continue
				}

				amount, err := strconv.ParseFloat(m["Amount"], 64)
				if err != nil {
					panic(err)
				}
				tm, err := time.ParseInLocation("2006-01-02 15:04:05", m["Trx Date Time"], time.Local)
				if err != nil {
					panic(err)
				}
				country := "SG"
				if m["Currency"] == "MYR" {
					country = "MY"
				}

				orderS := strings.Split(m["Order No"], "_")

				p := &Payment{
					TopNumber:      orderS[0],
					PaymentMethond: "Molpay",
					CountryCode:    country,
					Amount:         int64((amount + 0.005) * 100),
					Currency:       m["Currency"],
					Status:         "success",
					CreateTime:     tm.Unix(),
					CreateBy:       m["Customer Email"],
				}
				payments = append(payments, p)

				_ = amount
				_ = tm
				_ = country
				_ = orderS
			}
		}
	}
	return
}

func main() {
	a := doFile()
	length := len(a)
	log.Println(length)
	for i := 0; i < length; i++ {
		log.Printf("%v", a[i])
	}
}
