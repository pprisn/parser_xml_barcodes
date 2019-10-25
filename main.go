// Парсер файла barcodes.xml
// Цель получить значения полей
// ID;NAME;REGEX;JSVALIDEXPR;JSPREEXPR;DATE_CREAT;LAST_UPDATE
//
//Струкура xml получена с применением https://www.onlinetool.io/xmltogo/

package main

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Datapacket struct {
	XMLName     xml.Name `xml:"datapacket"`
	Text        string   `xml:",chardata"`
	Version     string   `xml:"version,attr"`
	Destination string   `xml:"destination,attr"`
	Day         string   `xml:"day,attr"`
	Month       string   `xml:"month,attr"`
	Year        string   `xml:"year,attr"`
	Hour        string   `xml:"hour,attr"`
	Min         string   `xml:"min,attr"`
	Sec         string   `xml:"sec,attr"`
	Msec        string   `xml:"msec,attr"`
	Metadata    struct {
		Text          string `xml:",chardata"`
		AttrFielddefs string `xml:"fielddefs,attr"`
		AttrIndexdefs string `xml:"indexdefs,attr"`
		Recorddata    string `xml:"recorddata,attr"`
		Changes       string `xml:"changes,attr"`
		Fielddefs     struct {
			Text     string `xml:",chardata"`
			Count    string `xml:"count,attr"`
			Fielddef []struct {
				Text         string `xml:",chardata"`
				Name         string `xml:"name,attr"`
				Fieldkind    string `xml:"fieldkind,attr"`
				Datatype     string `xml:"datatype,attr"`
				Fieldsize    string `xml:"fieldsize,attr"`
				Displaylabel string `xml:"displaylabel,attr"`
				Editmask     string `xml:"editmask,attr"`
				Displaywidth string `xml:"displaywidth,attr"`
				Fieldindex   string `xml:"fieldindex,attr"`
				Required     string `xml:"required,attr"`
				Readonly     string `xml:"readonly,attr"`
			} `xml:"fielddef"`
		} `xml:"fielddefs"`
		Indexdefs struct {
			Text  string `xml:",chardata"`
			Count string `xml:"count,attr"`
		} `xml:"indexdefs"`
	} `xml:"metadata"`
	Recorddata struct {
		Text  string `xml:",chardata"`
		Count string `xml:"count,attr"`
		Row   []struct {
			Text  string `xml:",chardata"`
			Field []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value,attr"`
				Size  string `xml:"size,attr"`
			} `xml:"field"`
		} `xml:"row"`
	} `xml:"recorddata"`
}

func main() {
	var c Datapacket
	r, err := os.Open("barcodes.xml")
	if err != nil {
		panic(err)
	}
	defer r.Close()
	stat, err := r.Stat()
	if err != nil {
		panic(err)
	}

	//Вычислим размер файла
	sizeFile := stat.Size()
	fmt.Printf("File size is %d \n", sizeFile)

	//Инициализация объема памяти []byte для считывания данных из файла
	h := make([]byte, sizeFile)

	//Чтение из потока r в массив h[:] размером sizeFile
	_, err = io.ReadFull(r, h[:])

	err = xml.Unmarshal(h, &c)
	//	fmt.Printf("%+v %v", c, err)

	for inx, _ := range c.Metadata.Fielddefs.Fielddef {
		fmt.Printf(" Name[%s]\n", c.Metadata.Fielddefs.Fielddef[inx].Name)
	}
	var data []byte
	for idx, _ := range c.Recorddata.Row {

		for idx2, _ := range c.Recorddata.Row[idx].Field {
			fmt.Printf(" %s\n", c.Recorddata.Row[idx].Field[idx2].Name)
			fmt.Printf(" %s\n", c.Recorddata.Row[idx].Field[idx2].Value)

			if c.Recorddata.Row[idx].Field[idx2].Name == "NAME" {
                 	       data, _ = base64.StdEncoding.DecodeString(c.Recorddata.Row[idx].Field[idx2].Text)
                               fmt.Printf("%v\n", string(data))

			} else {
				data, err = base64.StdEncoding.DecodeString(c.Recorddata.Row[idx].Field[idx2].Text)
				if err != nil {
					fmt.Println("error:", err)
					return
				}
                                fmt.Printf("%q\n", data)
			}
			

			
		}
	}

}
