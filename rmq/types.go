package rmq

type OnlineUser struct {
	UserId   int64
	DataTime int64
}

type OnlineWeb struct {
	AccessToken string
	DataTime    int64
}
