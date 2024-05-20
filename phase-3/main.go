package main

import (
	"github.com/gnolang/gno/gno.me/apps"
	"github.com/gnolang/gno/gno.me/gno"
	"github.com/gnolang/gno/gno.me/http"
)

func main() {
	vm, _ := gno.NewVM()
	apps.CreateInstaller(vm)
	server := http.NewServer(vm)
	if err := server.ListenAndServe(); err != nil {
		panic("could not start server: " + err.Error())
	}
}
