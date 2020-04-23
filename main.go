package main

import (
	// "encoding/json"
	// "fmt"
	// "os"
	// "Go/social"
	"Go/Task3/fb"
	"Go/Task3/twtr"
	"Go/Task3/lnkdn"
	"Go/Task3/exporter"
	// "io/ioutil"
)

func main() {
	fb := new (facebook.Facebook)
	twtr := new (twitter.Twitter)
	lnkdn := new (linkedin.Linkedin)
	// err := exporter.Text(fb, "fbData.txt")
	// err = exporter.Text(twtr, "twtrData.txt")
	// err = exporter.Text(lnkdn, "lnkdnData.txt")

	// var jsonData []byte
	// jsonData, err = json.Marshal(fb.Feed())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(jsonData))
	err := exporter.JSON(fb, "fbdata.json")
	err = exporter.JSON(twtr, "twtrdata.json")
	err = exporter.JSON(lnkdn, "lnkdndata.json")
	
	err = exporter.XML(fb, "fbdata.xml")
	err = exporter.XML(twtr, "twtrdata.xml")
	err = exporter.XML(lnkdn, "lnkdndata.xml")

	if err != nil {
		panic(err)
	}
	
	// ScrollFeeds(fb, twtr, lnkdin)
}
