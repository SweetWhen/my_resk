package my_resk

import (
	"myresk/infra"
	"myresk/infra/base"
)

func init()  {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
}