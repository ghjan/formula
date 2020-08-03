package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Id       int
	Name     string
	Address  string
	Email    string
	School   string
	City     string
	Company  string
	Age      int
	Sex      string
	Proviece string
	Com      string
	PostTo   string
	Buys     string
	Hos      string
}

var person = Person{
	Id:       98439,
	Name:     "zhaondifnei",
	Address:  "大沙地",
	Email:    "dashdisnin@126.com",
	School:   "广州第十五中学",
	City:     "zhongguoguanzhou",
	Company:  "sndifneinsifnienisn",
	Age:      23,
	Sex:      "F",
	Proviece: "jianxi",
	Com:      "广州兰博基尼",
	PostTo:   "蓝鲸XXXXXXXX",
	Buys:     "shensinfienisnfieni",
	Hos:      "zhonsndifneisnidnfie",
}

func BenchmarkStructToMapViaJson(b *testing.B) {
	a := assert.New(b)
	m := StructToMapViaJson(&person)
	a.Equal(person.Id, int(m["Id"].(float64)))
	a.Equal(person.Name, m["Name"])
	a.Equal(person.Address, m["Address"])
	a.Equal(person.Email, m["Email"])
}

func BenchmarkStructToMapViaReflect(b *testing.B) {
	a := assert.New(b)
	p := StructToMapViaReflect(&person)
	a.Equal(person.Id, p["Id"])
	a.Equal(person.Name, p["Name"])
	a.Equal(person.Address, p["Address"])
	a.Equal(person.Email, p["Email"])
}

func TestStructToMapViaJson(t *testing.T) {
	a := assert.New(t)
	p := StructToMapViaJson(&person)
	a.Equal(person.Id, int(p["Id"].(float64)))
	a.Equal(person.Name, p["Name"])
	a.Equal(person.Address, p["Address"])
	a.Equal(person.Email, p["Email"])
}

func TestStructToMapViaReflect(t *testing.T) {
	a := assert.New(t)
	p := StructToMapViaReflect(&person)
	a.Equal(person.Id, p["Id"])
	a.Equal(person.Name, p["Name"])
	a.Equal(person.Address, p["Address"])
	a.Equal(person.Email, p["Email"])
}
