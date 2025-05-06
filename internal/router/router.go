package router

import (
	"github.com/huynhthanhthao/hrm_hr_service/ent"

	"github.com/gin-gonic/gin"
)

func SetupRouter(client *ent.Client) *gin.Engine {
	r := gin.Default()

	return r
}
