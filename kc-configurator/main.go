package main

import "kc-configurator/cmd"

func main() {
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
