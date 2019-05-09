package aggFuncs

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"os"

	structs "pkg/structPrototypes"
)

func GetConfigJSON() structs.JsonConfigStruct {

	var config structs.JsonConfigStruct

	raw, err := ioutil.ReadFile("Config/config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &config)

	fmt.Println(config)
	return config
}

/*/ Test Function
func main() {
    config := getConfigJSON()
    fmt.Println(config)
    for _, test := range config {
        fmt.Println(test.toString())
    }
//*/
