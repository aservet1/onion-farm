// package requestWebRapl
package main

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"time"
	//"os"
	// "strconv"
)

func EnergyStatCheck() string {

	Url, _ := url.Parse("http://localhost:8080")
	Url.Path += "/energy/stats"

	parameters := url.Values{}
	// parameters.Add("duration", strconv.Itoa(duration))
	Url.RawQuery = parameters.Encode()

	urL := Url.String()

	request, _ := http.NewRequest("GET", urL, nil)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return "error happened"
	}

	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)

	return string(bodyBytes[:])
	// umsh := make(map[string]interface{})
	// json.Unmarshal([]byte(body), &umsh)
	// list := umsh["list"].([]interface{})

}

func main() {

	// duration, _ := strconv.Atoi(os.Args[1])
	fmt.Println( "<< start :: " + time.Now().String() );
	fmt.Printf( "%s\n", EnergyStatCheck() );
	fmt.Println( "<< stop  :: " + time.Now().String() );
}
