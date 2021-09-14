package main

import(
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func main(){

	githubusername := "enesbuyuk"

	type GitApi []struct {
		Ad   	string `json:"name"`
		Aciklama 	string `json:"description"`
		Url     	string `json:"html_url"`
	}

	gitrepurl, err := http.Get("https://api.github.com/users/"+githubusername+"/repos")
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(gitrepurl.Body)

	var f GitApi

	err = json.Unmarshal(body, &f)

	fmt.Printf("%+v", f[0].Ad)
	for i := 0; i < len(f); i++ {
		fmt.Println(f[i].Ad)
	}
}
