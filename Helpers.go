package main 

import (
    "io/ioutil"
	"github.com/Jeffail/gabs"
	"os"
	"reflect"
	"fmt"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


func parseJson(fileName string) *gabs.Container {

	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonData, _ := ioutil.ReadAll(jsonFile)

	jsonParsed, err := gabs.ParseJSON(jsonData)

	if err != nil {
		panic(err)
	}

	return jsonParsed

}

func getField(c *Connector, fields []string) reflect.Value {

    r := reflect.ValueOf(c)
	var value = reflect.Indirect(r)

	for _,p_key := range fields {
		value = value.FieldByName(p_key)
	}

	return value
}
