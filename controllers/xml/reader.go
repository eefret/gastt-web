package xml

import (
	"encoding/xml"
	"log"
)

//Read will parse a strings XML file to a map[string]string
func Read(file []byte) map[string]string {
	var res resources
	var ret = make(map[string]string)
	if err := xml.Unmarshal(file, &res); err != nil {
		log.Fatalln(err)
	}

	for _, v := range res.Strings {
		ret[v.Key] = v.Val
	}

	return ret
}
