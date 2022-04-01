package utils

type FetchListPostForm struct {
	Token string `form:"token"`
	Page  int    `form:"page"`
}
