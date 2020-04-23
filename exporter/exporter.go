package exporter

import(
	"os"
	"fmt"
	"Go/social"
	"encoding/json"
	"encoding/xml"
	"io"
)

//JSON file of social media
func JSON (data social.SocialMedia, filename string) error{
	// var i int
	// // i = 1
	// // func(I int)int {i++; return i}
	
	// }
		f, err := json.MarshalIndent(data.Feed(), "", "" ) //(data.Feed())
		// _ = ioutil.WriteFile(filename, f, 0644)
		if err != nil {
			panic(err)
		}
		// enc.Encode(f)
		fmt.Println(string(f))
		jsonFile, err := os.Create(filename)

         if err != nil {
                 panic(err)
         }
         defer jsonFile.Close()

         jsonFile.Write(f)
         jsonFile.Close()
		 fmt.Println("JSON data written to ", jsonFile.Name())
		 
		 return nil
	}
	// return nil
// }

// XML file of social media
func XML(data social.SocialMedia, filename string) error {
	xmlString, err := xml.MarshalIndent(data.Feed(), "", " ")

	if err != nil {
	fmt.Println(err)
	}

	fmt.Printf("%s \n", string(xmlString))

	// everything ok now, write to file.

	file, _ := os.Create(filename)

	xmlWriter := io.Writer(file)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", " ")
	if err := enc.Encode(data.Feed()); err != nil {
	fmt.Printf("error: %v\n", err)
 }
 return nil
}
