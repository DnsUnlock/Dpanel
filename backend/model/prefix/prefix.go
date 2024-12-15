package prefix

import "github.com/DnsUnlock/Dpanel/backend/config"

func Prefix() string {
	return config.Config.Sql.Prefix + "_"
}
