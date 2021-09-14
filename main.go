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

	var x GitApi

	err = json.Unmarshal(body, &x)

	for i := 0; i < len(x); i++ {
		fmt.Printf(" \n------------------\nRepo Name: "+x[i].Ad +"\nRepo Desc: "+ x[i].Aciklama +"\nRepo Full Url: "+ x[i].Url)
	}
}
