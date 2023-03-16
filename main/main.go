package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"mongo_data/core/data_operator"
	"mongo_data/servers"
	"strconv"
)

func main() {
	port := flag.Int("port", 8081, "端口")
	data_operator.MongoUri = *flag.String("mongo_uri", "mongodb://1.13.197.114:27017", "数据库连接")
	flag.Parse()
	portStr := strconv.Itoa(*port)
	r := gin.Default()
	r.POST("/data/create", servers.DataCreateHandle)
	r.Run(":" + portStr)
}
