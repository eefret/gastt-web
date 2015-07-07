package xml

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"log"

	"github.com/eefret/gastt-web/controllers/lang"
)

//Builder struct will handle the xml creation
type Builder struct {
	language lang.Lang
	strings  map[string]string
}

type resources struct {
	XMLName xml.Name `xml:resources`
	Strings []tag    `xml:"string"`
}

type tag struct {
	Key string `xml:"name,attr"`
	Val string `xml:",chardata"`
}

//NewBuilder creates a new Builder
func NewBuilder(strings map[string]string) (*Builder, error) {
	return &Builder{strings: strings}, nil
}

//SetLang sets a language for the XMl
func (b *Builder) SetLang(language lang.Lang) {
	b.language = language
}

func (b *Builder) build() ([]byte, error) {
	r := &resources{Strings: make([]tag, 0)}
	for k, v := range b.strings {
		r.Strings = append(r.Strings, tag{Key: k, Val: v})
	}
	return xml.MarshalIndent(r, "  ", "    ")
}

//BuildStr will create the xml format and return the contents of the file in string
func (b *Builder) BuildStr() string {
	byteArr, err := b.build()
	check(err)
	return string(byteArr[:])
}

//BuildFile will create the xml File and return it named strings.xml
func (b *Builder) BuildFile() []byte {
	buf := &bytes.Buffer{}
	buf.WriteString(b.BuildStr())
	return buf.Bytes()
}

//BuildZip will create the file and the default path inside a zip file
//default path is res/values-{lang}/strings.xml
func (b *Builder) BuildZip() []byte {
	//Creating writer
	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	//Creating directories
	wr, err := w.Create("res/values-" + string(b.language) + "/strings.xml")
	check(err)
	//Building xml and writing to file
	s, err := b.build()
	check(err)
	wr.Write(s)
	//Closing writers
	w.Close()
	return buf.Bytes()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
