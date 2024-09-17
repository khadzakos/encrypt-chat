package models

type Messange struct {
	ID        int    `json:"id"`
	Chatroom  string `json:"chatroom"`
	Username  string `json:"username"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type Chatroom struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
