package cmd

import (
	"context"
	"fmt"
	"kc-configurator/internal"
	"log"
	"os"

	"github.com/Nerzal/gocloak/v11"
)

type KcConfigurator struct {
	info   *internal.Info
	client gocloak.GoCloak
	token  string
}

func NewKcConfigurator() *KcConfigurator {
	user := parseEnv("USER")
	password := parseEnv("PASSWORD")
	KcURL := parseEnv("KC_URL")
	RpURL := parseEnv("RP_URL")

	client := gocloak.NewClient(KcURL, gocloak.SetAuthAdminRealms("admin/realms"), gocloak.SetAuthRealms("realms"))

	jwt, err := client.LoginAdmin(context.Background(), user, password, "master")
	if err != nil {
		panic(err)
	}

	info := internal.GetInfo(RpURL)

	return &KcConfigurator{
		info:   info,
		client: client,
		token:  jwt.AccessToken,
	}
}

func (c *KcConfigurator) Run() error {
	if err := c.createRealm(); err != nil {
		return err
	}

	if err := c.createClient(); err != nil {
		return err
	}

	if err := c.createUser(); err != nil {
		return err
	}

	return nil
}

func (c *KcConfigurator) createRealm() error {
	realmRepr := gocloak.RealmRepresentation{
		Realm:   &c.info.Realm,
		Enabled: gocloak.BoolP(true),
	}

	_, err := c.client.CreateRealm(context.Background(), c.token, realmRepr)
	if err != nil {
		return err
	}

	log.Printf("Created realm %s\n", c.info.Realm)
	return nil
}

func (c *KcConfigurator) createClient() error {
	clientRepr := gocloak.Client{
		ClientID:     gocloak.StringP(c.info.Client),
		Name:         &c.info.Client,
		Secret:       &c.info.ClientSecret,
		RedirectURIs: &[]string{"*"},
	}

	_, err := c.client.CreateClient(context.Background(), c.token, c.info.Realm, clientRepr)
	if err != nil {
		return err
	}

	log.Printf("Created client %s\n", c.info.Client)
	return nil
}

func (c *KcConfigurator) createUser() error {
	userRepr := gocloak.User{
		Username: &c.info.Username,
		Credentials: &[]gocloak.CredentialRepresentation{
			{
				Temporary: gocloak.BoolP(false),
				Type:      gocloak.StringP("password"),
				Value:     gocloak.StringP(c.info.Password),
			},
		},
		Enabled: gocloak.BoolP(true),
	}

	_, err := c.client.CreateUser(context.Background(), c.token, c.info.Realm, userRepr)
	if err != nil {
		return err
	}

	log.Printf("Created user %s\n", c.info.Username)
	return nil
}

func parseEnv(name string) string {
	env := os.Getenv(name)

	if env == "" {
		panic(fmt.Sprintf("Missing ENV: %s", name))
	}

	return env
}
