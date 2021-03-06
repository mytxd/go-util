package jsonutil

import (
	"testing"
	"time"
	"encoding/json"
	"bytes"
)

type TestStruct struct {
	String    string  `json:"string"`
	Float     float64 `json:"float"`
	Int       int     `json:"int"`
	BoolTrue  bool    `json:"bool_true"`
	BoolFalse bool    `json:"bool_false"`
	Nested struct {
		Number float64 `json:"number"`
		String string  `json:"string"`
	} `json:"nested"`
	Time time.Time `json:"time"`
}

func TestConvert(t *testing.T) {
	jsonStr := `{
        "string": "abc",
        "float": 2.721,
        "int": 215,
        "bool_true": true,
        "bool_false": false,
        "nested": {
          "number": 53.47273,
          "string": "abc"
        },
        "time": "2016-09-19T18:32:19Z"
}`

	var jsonStruct interface{}

	err := json.Unmarshal([]byte(jsonStr), &jsonStruct)
	if err != nil {
		t.Error(err)
	}

	asMap := jsonStruct.(map[string]interface{})
	if asMap["float"] != 2.721 {
		t.Error("Wrong data asMap")
	}

	var parsed TestStruct
	bytes, err := json.Marshal(jsonStruct)
	err = json.Unmarshal(bytes, &parsed)

	if err != nil {
		t.Error("Can not convert jsonStruct", err)
	}
	if parsed.Float != 2.721 {
		t.Error("Wrong Float data", parsed.Float)
	}
	if parsed.Nested.Number != 53.47273 {
		t.Error("Wrong Nested.Number data", parsed.Nested.Number)
	}
	if parsed.BoolTrue != true {
		t.Error("Wrong BoolTrue data", parsed.BoolTrue)
	}
}

type ByteString struct {
	Str ByteJsonString
}

func TestMarshalByteString(t *testing.T) {
	o := ByteString{
		Str: []byte("lobaro"),
	}

	data, err := json.Marshal(o)

	if err != nil {
		t.Error(err)
	}

	expected := `{"Str":"lobaro"}`
	if string(data) != expected {
		t.Error("Unexpected json result: " + string(data))
	}
}

func TestUnMarshalByteString(t *testing.T) {
	j := `{"Str":"lobaro"}`
	o := &ByteString{}

	err := json.Unmarshal([]byte(j), o)

	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(o.Str, []byte("lobaro")) {
		t.Error("Expected o.Str to be 'lobaro' but is: '" + string(o.Str)+"'")
	}
}

type ByteArray struct {
	Bytes ByteJsonArray
}

func TestMarshalByteArray(t *testing.T) {
	o := ByteArray{
		Bytes: []byte{1, 2, 3},
	}

	data, err := json.Marshal(o)

	if err != nil {
		t.Error(err)
	}

	expected := `{"Bytes":[1,2,3]}`
	if string(data) != expected {
		t.Error("Unexpected json result: " + string(data))
	}
}

func TestUnMarshalByteArray(t *testing.T) {
	j := `{"Bytes":[1, 2,3]}`
	o := &ByteArray{}

	err := json.Unmarshal([]byte(j), o)

	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(o.Bytes, []byte{1,2,3}) {
		t.Errorf("Expected o.Bytes to be '[1,2,3]' but is: '%v' (%s)", o.Bytes, o.Bytes)
	}
}