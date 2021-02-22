package main

import (
	"github.com/jinzhu/gorm"
	"github.com/vitor9/codepix/tree/master/codepix/application/grpc"
	"github.com/vitor9/codepix/tree/master/codepix/infrastructure/db"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051) // 50051 eh a porta padrao do gRPC
}
