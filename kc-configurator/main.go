package main

import "kc-configurator/cmd"

func main() {
	if err := cmd.NewKcConfigurator().Run(); err != nil {
		panic(err)
	}
}
