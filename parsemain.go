package main

import (
	"fmt"
	parser "github.com/oucema001/OutlookMessageParser-go"
	"github.com/oucema001/OutlookMessageParser-go/models"
	"github.com/p2c2e/eml"
	"io/ioutil"
	"os"
)

func AnalyzeEmlFile(file string) (res *models.Message, err error) {
	res = &models.Message{}

	ba, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	raw, err := eml.ParseRaw(ba)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	m, err := eml.Process(raw)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		fmt.Println(m.ContentType)
	}

	res.BodyText = string(m.Body)
	res.BodyHTML = res.BodyText
	for _, rh := range raw.RawHeaders {
		//fmt.Println(string(rh.Key)+" = "+string(rh.Value))
		key := string(rh.Key)
		if key == "From" {
			res.FromName = string(rh.Value)
		} else if key == "To" {
			res.ToName = string(rh.Value)
		} else if key == "Subject" {
			res.Subject = string(rh.Value)
		}
		res.BodyHTML = m.Html
		res.BodyText = m.Text
	}
	return res, nil
}

func main(){
	filename := os.Args[1]
	res, _ := parser.AnalyzeMsgFile(filename)
	if res == nil {
		res, _ := AnalyzeEmlFile(filename)
		if res == nil {
			fmt.Println("Sorry! Unable to parse the file")
		} else {
			fmt.Println("EML CONTENT START ######################")
			fmt.Println(
				"FROM: "+res.FromName+" <"+res.FromEmail+ ">\n"+
					"TO: "+res.ToName+" <"+res.ToEmail+ ">\n"+
					"TO2: "+res.DisplayTo+"\n"+
					"SUBJECT: "+res.Subject+"\nBODY: "+res.BodyText+"\nDone")
			fmt.Println("EML CONTENT END ######################")
		}
	} else {
		fmt.Println("MSG CONTENT START ######################")
		fmt.Println(
			"FROM: "+res.FromName+" <"+res.FromEmail+ ">\n"+
				"TO: "+res.ToName+" <"+res.ToEmail+ ">\n"+
				"TO2: "+res.DisplayTo+"\n"+
				"SUBJECT: "+res.Subject+"\nBODY: "+res.BodyText+"\n"+
				res.BodyHTML+"\nMSG CONTENT END ###################")
	}
}
