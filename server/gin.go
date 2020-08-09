package server

import (
	"github.com/gin-gonic/gin"
)

var(
	router = gin.Default()
)
func StartGinApplication() {
	mapUrls()
	_ = router.Run(":9090")
}
