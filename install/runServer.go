package install

import "github.com/onism68/helmOnap/boot"

func RunServer() {
	go func() {
		boot.Boot()
	}()
}