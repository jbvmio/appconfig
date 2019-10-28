package appconfig

import (
	"crypto/sha1"
	"fmt"
)

// DataType defines the type of data.
type DataType int

func (d DataType) String() string {
	return DataTypeString[d]
}

// Datatypes Defined:
const (
	TypeInvalid DataType = iota // 0
	TypeSimple
	TypeParameter
	TypeEndpoint
)

// DataTypeString enables a way to identify a Datatype with a string.
var DataTypeString = [...]string{
	TypeInvalid:   "invalid",
	TypeSimple:    "simple",
	TypeParameter: "parameter",
	TypeEndpoint:  "endpoint",
}

// DataTypeMap maps the types given by string to a DataType.
var DataTypeMap = map[string]DataType{
	"invalid":   TypeInvalid,
	"simple":    TypeSimple,
	"parameter": TypeParameter,
	"endpoint":  TypeEndpoint,
}

// Data contains details using consistent fields.
type Data struct {
	T         string   `json:"type"`
	Pkg       string   `json:"pkg"`
	Tpls      []string `json:"tpls"`
	Src       string   `json:"src"`
	Key       string   `json:"k"`
	Value     string   `json:"v"`
	AppDomain string   `json:"appdomain"`
}

// SHA returns the Sha1 string of the data.
func (d Data) SHA() string {
	b := []byte(fmt.Sprintf("%+v", d))
	return fmt.Sprintf("%x", sha1.Sum(b))
}

// DataType returns the DataType.
func (d *Data) DataType() DataType {
	val, ok := DataTypeMap[d.T]
	if ok {
		return val
	}
	return TypeInvalid
}

// AssignAD extracts and stores the AppDomain value if applicable, otherwise will assign the default value given.
// If an array is given, then the first element is used.
func (d *Data) AssignAD(defaultValue ...string) {
	var dv string
	if len(defaultValue) > 0 {
		dv = defaultValue[0]
	}
	switch d.DataType() {
	case TypeEndpoint:
		d.AppDomain = epAD(d.Value)
	default:
		d.AppDomain = dv
	}
}

// HasKey returns true if the entered string matches the data key, false otherwise.
func (d *Data) HasKey(key string) bool {
	return d.Key == key
}

// HasValue returns true if the entered string matches the data value, false otherwise.
func (d *Data) HasValue(value string) bool {
	return d.Value == value
}

// HasPkg returns true if the entered string matches the data pkg, false otherwise.
func (d *Data) HasPkg(pkg string) bool {
	return d.Pkg == pkg
}

// HasAD returns true if the entered string matches the data AppDomain, false otherwise.
func (d *Data) HasAD(ad string) bool {
	return d.AppDomain == ad
}
