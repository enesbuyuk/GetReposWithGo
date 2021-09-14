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

	for i := 0; i < len(f); i++ {
		fmt.Printf(" \n------------------\nRepo Name:"+f[i].Ad +"\nRepo Desc: "+ f[i].Aciklama +"\nRepo Full Url:"+ f[i].Url)
	}
}
