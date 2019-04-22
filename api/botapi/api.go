package botapi

type VKMsg struct {
	Type    string      `json:"type"`
	GroupId int         `json:"group_id"`
	MSG     interface{} `json:"object"`
}

type MessageNew struct {
	Id        int    `json:"id"`
	Date      int    `json:"date"`
	PeerId    int    `json:"peer_id"`
	FromId    int    `json:"from_id"`
	Text      string `json:"text"`
	RandomId  int    `json:"random_id"`
	Ref       string `json:"ref"`
	RefSource string `json:"ref_source"`
	Important bool   `json:"important"`
	Payload   string `json:"payload"`
}
