package Routers

import (
	"main/api"
)

func UavRoute() {
	g := r.Group("/Uav")

	//可借用无人机展示
	g.GET("/GetUav", api.GetNotUsedDrones)

	//可借用电池展示
	g.GET("/GetBattery", api.GetNotUsedBattery)

	//可借用遥控展示
	g.GET("/GetControl", api.GetNotUsedControl)

}
