package main

import (
	"github.com/DnsUnlock/Dpanel/backend/app"
	"github.com/DnsUnlock/Dpanel/backend/db/sql"
	"github.com/DnsUnlock/Dpanel/backend/model/user"
)

func main() {
	sql.Get().AutoMigrate(user.User{})
	sql.Get().AutoMigrate(user.Email{})
	app.Run()
}
