package service

import (
	"fmt"
	"gatelligance/entity"
	"gatelligance/utils"
	"net/url"
	"strconv"

	"github.com/jinzhu/gorm"
)

// type checkLinkTransactionStruct struct {
// 	Progress string `json:"Progress"`
// 	Status   string `json:"Status"`
// 	Output   string `json:"Output"`
// }

func GetTransactionServerNumber(db *gorm.DB, tuid string) int {
	var transactions []entity.Transaction

	db.Find(&transactions, "id=?", tuid)

	if len(transactions) == 0 {
		fmt.Printf("transcation: not found\n")
		return -1
	}

	return transactions[0].Server
}

//v3.0
func CheckLinkTransactionService(tuid string, db *gorm.DB) utils.TaskCheckReturn {
	var results []getUsersTransactionListDBResult
	db.Raw("SELECT transactions.title as title,transactions.avatar as avatar,link_transactions.progress as progress,link_transactions.output as output,link_transactions.status as status, transactions.type as type, users.id as uuid,transactions.id as tuid FROM users, transactions,link_transactions WHERE  transactions.id= ? AND users.id=transactions.owner AND link_transactions.id=transactions.id", tuid).Scan(&results)

	if len(results) == 0 {
		fmt.Printf("transcation: not found\n")
		return utils.TaskCheckReturn{}
	}

	return utils.TaskCheckReturn{
		Output: results[0].Output,
		Avatar: results[0].Avatar,
		Title:  results[0].Title,

		Progress: results[0].Progress,
		Status:   results[0].Status,
		Type:     results[0].Type,
	}
}

// v3.0
func CreateLinkTransaction(videoLink string, db *gorm.DB, creator string) string {
	sid, saddr := GetNextUseableSlaveServer(db)
	resBody := SendPostRequest(saddr+"/addLinkWork",
		url.Values{"addr": {videoLink}, "owner": {creator}, "sid": {strconv.Itoa(sid)}})

	return resBody
}

type getUsersTransactionListDBResult struct {
	Uuid     string `json:"uuid"`
	Tuid     string `json:"tuid"`
	Progress string
	Status   string
	Type     string
	Output   string
	Avatar   string
	Title    string
}

func GetUsersTransactionList(db *gorm.DB, userID string, page int) []utils.TaskListRow {
	var results []getUsersTransactionListDBResult
	var ret []utils.TaskListRow
	db.Raw("SELECT transactions.title as title,transactions.avatar as avatar, link_transactions.output as output, link_transactions.progress as progress,link_transactions.status as status, transactions.type as type, users.id as uuid,transactions.id as tuid FROM users, transactions,link_transactions WHERE  users.id= ? AND users.id=transactions.owner AND link_transactions.id=transactions.id", userID).Scan(&results)
	// db_algo.Raw("SELECT link_transactions.progress as progress,link_transactions.status as status, transactions.type as type FROM transactions, link_transactions WHERE  link_transactions.id= ? AND link_transactions.id=transactions.id", value.Tuid).Scan(&algoResults)Z
	var si = (page - 1) * 10
	var i = 0
	for _, value := range results {

		//分页.页大小是10.
		i++
		// println(i)
		if i < si {
			continue
		}
		if i > si+10 {
			break
		}

		ret = append(ret, utils.TaskListRow{
			Progress: value.Progress,
			Status:   value.Status,
			Avatar:   value.Avatar,
			Title:    value.Title,
			Type:     value.Type,
			// TaskList: Service.CreateLinkTransaction(link),
		})

	}
	// println(MapToJson(results))
	return ret
}

//v2.0
// func CheckLinkTransactionService(tuid string, db *gorm.DB) (string, string, string) {
// 	sid := GetTransactionServerNumber(db, tuid)
// 	saddr := GetSlaveServerAddress(db, sid)
// 	resBody := SendPostRequest(saddr+"/checkLinkWork",
// 		url.Values{"uuid": {tuid}, "test": {"false"}})

// 	var retStruct checkLinkTransactionStruct

// 	err := json.Unmarshal([]byte(string(resBody)), &retStruct)

// 	if err != nil {
// 		return "-1", "-1", "nil"
// 	}

// 	return retStruct.Progress, retStruct.Status, retStruct.Output
// }

// v2.0
// func CreateLinkTransaction(videoLink string, db *gorm.DB, creator string) string {
// 	sid, saddr := GetNextUseableSlaveServer(db)
// 	resBody := SendPostRequest(saddr+"/addLinkWork",
// 		url.Values{"addr": {videoLink}, "id": {"123"}})

// 	// resBody := SendPostRequest(saddr, "/addLinkWork", videoLink)

// 	var nt = new(entity.Transaction)
// 	nt.CreatedAt = time.Now()
// 	nt.Owner = creator
// 	nt.ID = resBody
// 	nt.Server = sid
// 	nt.Type = "1"
// 	db.Create(nt)

// 	return resBody
// }

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
