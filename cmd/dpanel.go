package main

import (
	"github.com/DnsUnlock/Dpanel/backend/app"
	"gitoo.icu/Nexus/Nexus"
)

func main() {
	Nexus.New()

	app.Run()
}
