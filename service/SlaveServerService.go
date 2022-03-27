package service

import (
	"fmt"
	"gatelligance/entity"

	"github.com/jinzhu/gorm"
)

var serverInd = 0

func GetSlaveServerAddress(db *gorm.DB, serverID int) string {
	var server []entity.SlaveServer

	db.Find(&server, "server_id=?", serverID)

	if len(server) == 0 {
		fmt.Printf("transcation: not found\n")
		return "nil"
	}

	return server[0].Address
}

func CheckIfSlaveServerUsable(db *gorm.DB, serverID int) bool {
	var server []entity.SlaveServer

	db.Find(&server, "server_id=?", serverID)

	if len(server) == 0 {
		fmt.Printf("transcation: not found\n")
		return false
	}

	if server[0].Usable == 1 {
		return true
	}

	return false
}

func GetNextUseableSlaveServer(db *gorm.DB) (int, string) {
	var server []entity.SlaveServer
	db.Find(&server)

	serverCnt := len(server)

	if serverCnt == 0 {
		fmt.Printf("transcation: not found\n")
		return -1, "nil"
	}

	if serverCnt == 1 {
		if server[0].Usable != 1 {
			return -2, "nil"
		}
	}

	for {
		serverInd = (serverInd + 1) % serverCnt
		if server[serverInd].Usable == 1 {
			break
		}
	}

	return server[serverInd].ServerID, server[serverInd].Address
}
