package service

import (
	"encoding/json"
	"fmt"
	"gatelligance/entity"
	"net/url"
	"time"

	"github.com/jinzhu/gorm"
)

type checkLinkTransactionStruct struct {
	Progress string `json:"Progress"`
	Status   string `json:"Status"`
	Output   string `json:"Output"`
}

func GetTransactionServerNumber(db *gorm.DB, tuid string) int {
	var transactions []entity.Transaction

	db.Find(&transactions, "id=?", tuid)

	if len(transactions) == 0 {
		fmt.Printf("transcation: not found\n")
		return -1
	}

	return transactions[0].Server
}

func CheckLinkTransactionService(tuid string, db *gorm.DB) (string, string, string) {
	sid := GetTransactionServerNumber(db, tuid)
	saddr := GetSlaveServerAddress(db, sid)
	resBody := SendPostRequest(saddr+"/checkLinkWork",
		url.Values{"uuid": {tuid}, "test": {"false"}})

	var retStruct checkLinkTransactionStruct

	err := json.Unmarshal([]byte(string(resBody)), &retStruct)

	if err != nil {
		return "-1", "-1", "nil"
	}

	return retStruct.Progress, retStruct.Status, retStruct.Output
}

// v2.0
func CreateLinkTransaction(videoLink string, db *gorm.DB, creator string) string {
	sid, saddr := GetNextUseableSlaveServer(db)
	resBody := SendPostRequest(saddr+"/addLinkWork",
		url.Values{"addr": {videoLink}, "id": {"123"}})

	// resBody := SendPostRequest(saddr, "/addLinkWork", videoLink)

	var nt = new(entity.Transaction)
	nt.CreatedAt = time.Now()
	nt.Owner = creator
	nt.ID = resBody
	nt.Server = sid
	db.Create(nt)

	return resBody
}

type getUsersTransactionListDBResult struct {
	Uuid string `json:"uuid"`
	Tuid string `json:"tuid"`
}

func GetUsersTransactionList(db *gorm.DB, userID string) {
	var results []getUsersTransactionListDBResult
	db.Raw("SELECT users.id as uuid,transactions.id as tuid FROM users, transactions WHERE  users.id= ? AND users.id=transactions.owner", userID).Scan(&results)
	for _, value := range results {
		println(value.Uuid + "_" + value.Tuid)
	}
	// println(MapToJson(results))
}

//v1.0
// func CreateLinkTransaction(addr string) string {
// 	resp, err := http.PostForm("http://114.215.149.199:8091/addLinkWork",
// 		url.Values{"addr": {addr}, "id": {"123"}})

// 	// print(addr)

// 	if err != nil {
// 		return "nil"
// 	}

// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "nil"
// 	}

// 	// fmt.Println(string(body))
// 	return string(body)
// }
