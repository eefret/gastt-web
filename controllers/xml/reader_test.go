package xml

import "testing"

var testData = []byte(`
<?xml version="1.0" encoding="utf-8"?>
<resources>
    <string name="app_name">GASTT</string>
    <string name="title_section1">Section 1</string>
    <string name="title_section2">Section 2</string>
    <string name="title_section3">Section 3</string>
    <string name="hello_world">Hello world!</string>
    <string name="action_settings">Settings</string>
</resources>`)

func TestReader(t *testing.T) {
	r := Read(testData)
	if r["app_name"] != "GASTT" {
		t.Error("Unmarshall is not ok!")
	}
}
