package main 

import (
    "io/ioutil"
	"github.com/Jeffail/gabs"
	"os"
	"reflect"
	"fmt"
	"path/filepath"
	"archive/zip"
	"io"
    "net/http"
	"strings"

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

func getEvnField(field string) string {

    var key_words = strings.Split(field,"_")[1:];


    
    for i,key_word := range key_words {
        
        key_words[i] = strings.ToUpper(key_word)
        
    }

    return os.Getenv(strings.Join(key_words,"_"))

}

func Unzip(src string, dest string) ([]string, error) {

    var filenames []string

    r, err := zip.OpenReader(src)
    if err != nil {
        return filenames, err
    }
    defer r.Close()

    for _, f := range r.File {

        // Store filename/path for returning and using later on
        fpath := filepath.Join(dest, f.Name)

        // Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
        if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
            return filenames, fmt.Errorf("%s: illegal file path", fpath)
        }

        filenames = append(filenames, fpath)

        if f.FileInfo().IsDir() {
            os.MkdirAll(fpath, os.ModePerm)
            continue
        }

        // Make File
        if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
            return filenames, err
        }

        outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
        if err != nil {
            return filenames, err
        }

        rc, err := f.Open()
        if err != nil {
            return filenames, err
        }

        _, err = io.Copy(outFile, rc)

        outFile.Close()
        rc.Close()

        if err != nil {
            return filenames, err
        }
    }
    return filenames, nil
}

func ResolveDeviceTypeUploadDist(DeviceType int) []string {

	switch DeviceType {
		case 0:
			return []string {"Samsung"}
		case 1:
			return []string {"LG"}
		case 2:
			return []string {"Samsung","LG"}
		default:
			return []string {"Samsung","LG"}
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}