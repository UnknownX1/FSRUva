package api

import (
	"github.com/gin-gonic/gin"
	"main/Model"
)

// GetNotUsedDrones 获取空闲的无人机设备
func GetNotUsedDrones(c *gin.Context) {
	//获取设备信息
	uav := Model.GetUavByStates("free", "drone")

	//JSON格式返回
	c.JSON(200, &uav)
}

// GetNotUsedBattery 获取空闲的电池设备
func GetNotUsedBattery(c *gin.Context) {
	//获取设备信息
	Battery := Model.GetUavByStates("free", "Battery")

	//JSON格式返回
	c.JSON(200, &Battery)
}

// GetNotUsedControl 获取空闲的遥控设备
func GetNotUsedControl(c *gin.Context) {
	//获取设备信息
	Control := Model.GetUavByStates("free", "Control")

	//JSON格式返回
	c.JSON(200, &Control)
}

// GetDrones 获取所有无人机设备
func GetDrones(c *gin.Context) {
	uav := Model.GetUavByStates("", "drone")

	c.JSON(200, &uav)
}

// GetBattery 获取所有电池设备
func GetBattery(c *gin.Context) {
	Battery := Model.GetUavByStates("", "battery")

	c.JSON(200, &Battery)
}

// GetControl 获取所有遥控设备
func GetControl(c *gin.Context) {
	Control := Model.GetUavByStates("", "control")

	c.JSON(200, &Control)
}

// GetUnderReview 获取借用审核的设备
func GetUnderReview(c *gin.Context) {
	uav := Model.GetUavByStates("Get under review", "")

	c.JSON(200, &uav)
}

// BackUnderReview 获取归还审核设备
func BackUnderReview(c *gin.Context) {
	uav := Model.GetUavByStates("Back under review", "")

	c.JSON(200, &uav)
}
