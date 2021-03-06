package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Contact struct {
	Email  string
	Github string
}

type Info struct {
	Version     string
	Description string
	Contact     Contact
}

func main() {
	programInfo := Info{
		"1.0.0",
		"This script checks for new version",
		Contact{
			"example@gmail.com",
			"http://github.com/example.com",
		},
	}

	needUpdate, err := needUpdate(programInfo)
	if err != nil {
		fmt.Println(err)
	}
	if needUpdate {
		fmt.Println("Start update")
		// call update function
	} else {
		fmt.Println("Update not need")
	}
}

func needUpdate(cerVer Info) (bool, error) {
	resp, err := http.Get("http://meromen.github.io/go-tasks/updater/getInfo.json")
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var getInfo Info

	err = json.Unmarshal(body, &getInfo)
	if err != nil {
		return false, err
	}

	intCurVer, _ := strconv.ParseInt(strings.Replace(cerVer.Version, ".", "", -1), 10, 64)
	intGetVer, _ := strconv.ParseInt(strings.Replace(getInfo.Version, ".", "", -1), 10, 64)

	anwser := intGetVer > intCurVer

	return anwser, nil
}
