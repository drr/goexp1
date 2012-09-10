/* A little fiddling around with Go
 *
 * Install a couple different http handlers
 * Try some reflection
 * Try some json encoding
 */

package main

import (
	"fmt"
	"net/http"
	"reflect"
        "encoding/json"
        "os"
        "log"
)

func rootdoc(w http.ResponseWriter, req *http.Request) {
        dump(req)
	fmt.Fprintln(w, "Hello World!")
}

func jsondoc(w http.ResponseWriter, req *http.Request) {
        enc := json.NewEncoder(w)
        if err := enc.Encode(req); err != nil {
            log.Println(err)
        }
        b, err := json.MarshalIndent(req, "", "  ")
        if err != nil {
            log.Println(err)
        }
        os.Stdout.Write(b)
}

func dump(obj interface{}) {
        v := reflect.ValueOf(obj)

	fmt.Println("type: ", reflect.TypeOf(obj))
	fmt.Println("value: ", v)
	fmt.Println("kind: ", v.Kind())
        fmt.Println("is struct? ", v.Kind() == reflect.Struct)
}

func main() {
	http.HandleFunc("/", rootdoc)
	http.HandleFunc("/json", jsondoc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
