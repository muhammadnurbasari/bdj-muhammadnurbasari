package registry

import (
	"bdj-muhammadnurbasari/routes/bdjRoutes"
	"bdj-muhammadnurbasari/routes/usersRoutes"
)

//initializeDomainModules calls the domain module routes
//in folder routes/*
func (reg *AppRegistry) initializeDomainModules() {
	bdjRoutes.BdjRoutes(reg.serverHttp.GetRouteEngine(), reg.ConnPublic)
	usersRoutes.UsersRoutes(reg.serverHttp.GetRouteEngine(), reg.ConnPublic, reg.ConnSql)
}
