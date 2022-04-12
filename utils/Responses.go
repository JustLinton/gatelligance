package utils

import "gatelligance/entity"

type RegisterResponse struct {
	IsSuccess bool
	ErrorMsg  string
}

type LoginResponse struct {
	Token     string
	IsSuccess bool
	ErrorMsg  string
}

type WorkSubmitResponse struct {
	// TaskList  string
	TransactionID string
	IsSuccess     bool
	ErrorMsg      string
}

type FetchUserInfoResponse struct {
	IsSuccess bool
	ErrorMsg  string
	Email     string
	NickName  string
	Avatar    string
	Gender    string
	Activated bool
}

type SetUserInfoResponse struct {
	IsSuccess bool
	ErrorMsg  string
}

type RefreshTokenResponse struct {
	IsSuccess bool
	ErrorMsg  string
	Token     string
}

type StandardResponse struct {
	IsSuccess bool
	ErrorMsg  string
}

type TransactionListResponse struct {
	IsSuccess bool
	ErrorMsg  string
	TaskList  []TaskListRow
}

type AvatarListResponse struct {
	IsSuccess bool
	ErrorMsg  string
	Data      []entity.AvatarResourceTable
}

type CheckLinkTransactionResponse_OutputStruct struct {
	OriginalText string
	SummaryText  string
}

type CheckLinkTransactionResponse struct {
	IsSuccess bool
	ErrorMsg  string

	Progress string
	Status   string
	Output   CheckLinkTransactionResponse_OutputStruct

	Avatar string
	Title  string
	Type   string
}

type TaskListRow struct {
	Progress      string
	Status        string
	Type          string
	TransactionID string

	Avatar string
	Title  string
}

type TaskCheckReturn struct {
	Progress string
	Status   string
	Type     string

	Avatar string
	Title  string

	Output string
}
