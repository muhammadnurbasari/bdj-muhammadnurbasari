package registry

import "bdj-muhammadnurbasari/routes/bdjRoutes"

//initializeDomainModules calls the domain module routes
//in folder routes/*
func (reg *AppRegistry) initializeDomainModules() {
	bdjRoutes.BdjRoutes(reg.serverHttp.GetRouteEngine(), reg.ConnPublic)
}
