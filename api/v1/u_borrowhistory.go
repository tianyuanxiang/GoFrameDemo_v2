package v1

import "github.com/gogf/gf/v2/frame/g"

type UHistoryBorrowReq struct {
	g.Meta   `path:"/UserHistoryBorrow/Query" tags:"HistoryBorrowInfromation"  method:"get" summary:"查询图书历史借阅信息"`
	UserIP   string `json:"UserIP" v:"required"`
	UserName string `json:"UserName"`
}

type UHistoryBorrowRes struct {
	Message       string                 `json:"Message"`
	HistoryBorrow *[]UHistoryinformation `json:"HistoryBorrow"`
}

type UHistoryinformation struct {
	ID             int    `json:"ID"`
	BookName       string `json:"BookName" `
	ISBN           string `json:"BookISBN"`
	UserIP         string `json:"UserIP"`
	UserName       string `json:"UserName"`
	BorrowingOrder int    `json:"BorrowingOrder" v:"required"`
	ReturnedOrNot  string `json:"Status"`
}
