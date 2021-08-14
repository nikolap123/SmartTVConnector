package smarttv

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkError(e error) (empty string, err error) {
	if e != nil {
		return
	}

	return
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

func getField(c *ConnectorDTO, fields []string) reflect.Value {

	r := reflect.ValueOf(c)
	var value = reflect.Indirect(r)

	for _,p_key := range fields {
		value = value.FieldByName(p_key)
	}

	return value
}

func getEvnField(field string) string {

	var key_words = strings.Split(field,"_")[1:];



	for i,key_word := range key_words {

		key_words[i] = strings.ToUpper(key_word)

	}

	return os.Getenv(strings.Join(key_words,"_"))

}