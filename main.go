package main

import (
	"net/http"
	"log"
)



func main() {

	http.HandleFunc("/get-devices" , HandleGetDevices)

	http.HandleFunc("/run-command" , HandleRunCommand)

	http.HandleFunc("/test-shell", func(w http.ResponseWriter, r *http.Request){

		
		// fmt.Println(jsonParsed.Path("Samsung.package").Data())
		// var command5 = TVCommand {Command: "sdb", Args:[]string{"-s","192.168.50.211:26101","shell","0","was_execute","JU3eZLUoMR.datazoom"},Next: nil}
		// var command4 = TVCommand {Command: "sdb", Args:[]string{"-s","192.168.50.211:26101","shell","0","vd_appinstall","JU3eZLUoMR.datazoom","/home/owner/share/tmp/sdk_tools/tmp/datazoom.wgt"},Next: &command5}
		// var command3 = TVCommand {Command: "sdb", Args: []string{"-s","192.168.50.211:26101","push","C:/Users/Popa/Desktop/dz/datazoom/tizen-datazoom/wgt/datazoom.wgt","/home/owner/share/tmp/sdk_tools/tmp/"},Next: &command4}
		// var command2 = TVCommand {Command: "tizen", Args: []string{"package","--type","wgt","--sign","samsungtv","--output","C:/Users/Popa/Desktop/dz/datazoom/tizen-datazoom/wgt/datazoom.wgt","--","C:/Users/Popa/Desktop/dz/datazoom/tizen-datazoom"}, Next: &command3}
		// var command1 = TVCommand {Command: "sdb" , Args : []string {"connect","192.168.50.211"}, Next: &command2}

		// command1.exec()
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
