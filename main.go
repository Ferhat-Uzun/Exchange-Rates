package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.tcmb.gov.tr/kurlar/today.xml")

	if err != nil {
		log.Fatal("Hata $s", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	var s date
	xml.Unmarshal(data, &s)
	fmt.Println(s.Currencies)
}

type currency struct {
	XmlName      xml.Name `xml:"Currency"`
	CurrencyName string   `xml:"CurrencyName"`
	ForexBuying  string   `xml:"ForexBuying"`
	ForexSelling string   `xml:"ForexSelling"`
}

type date struct {
	XmlName    xml.Name   `xml:"Tarih_Date"`
	Currencies []currency `xml:"Currency"`
}

func (c currency) String() string {
	return fmt.Sprintf("\t Name: %s , Forex Buying: %s , Forex Selling: %s \n", c.CurrencyName, c.ForexBuying, c.ForexSelling)
}
