package internal

import (
	"encoding/json"
	"net/http"
)

type Info struct {
	Realm, Client, ClientSecret, Username, Password string
}

func GetInfo(url string) *Info {
	info := &Info{}

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(info)
	if err != nil {
		panic(err)
	}

	return info
}
