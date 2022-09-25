package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Info struct {
	Realm, Client, ClientSecret, Username, Password string
}

func GetInfo(URL string) *Info {
	info := &Info{}

	res, err := http.Get(fmt.Sprintf("%s/public/authInfo.json", URL))
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(info); err != nil {
		panic(err)
	}

	return info
}
