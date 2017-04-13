package main

//import "fmt"
import "encoding/json"
import "os"

type AjaxJson struct {
	Status string
	Msg    string
}

func main() {

	aj := AjaxJson{Status: "true", Msg: "my god"}
	b, err := json.Marshal(aj)
	if err != nil {
	}
	os.Stdout.Write(b)
}
