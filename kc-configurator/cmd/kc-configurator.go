package cmd

import (
	"context"
	"fmt"
	"kc-configurator/internal"
	"os"

	"github.com/Nerzal/gocloak/v11"
)

func Run() error {
	user := parseEnv("USER")
	password := parseEnv("PASSWORD")
	URL := parseEnv("URL")

	client := gocloak.NewClient(URL)

	jwt, err := client.LoginAdmin(context.Background(), user, password, "master")
	if err != nil {
		return err
	}

	info := internal.GetInfo(URL)

	_, err = client.CreateRealm(context.Background(), jwt.AccessToken, gocloak.RealmRepresentation{Realm: &info.Realm})
	if err != nil {
		return err
	}

	clientRepr := gocloak.Client{Name: &info.Client, Secret: &info.ClientSecret, RedirectURIs: &[]string{URL}}

	_, err = client.CreateClient(context.Background(), jwt.AccessToken, info.Realm, clientRepr)
	if err != nil {
		return err
	}

	userRepr := gocloak.User{
		Username: &info.Username,
		Credentials: &[]gocloak.CredentialRepresentation{
			{
				Temporary: gocloak.BoolP(false),
				Type:      gocloak.StringP("password"),
				Value:     gocloak.StringP(info.Password),
			},
		},
	}

	_, err = client.CreateUser(context.Background(), jwt.AccessToken, info.Realm, userRepr)

	return nil
}

func parseEnv(name string) string {
	env := os.Getenv(name)

	if env == "" {
		panic(fmt.Sprintf("Missing ENV: %s", name))
	}

	return env
}
