package model

func School() {
	type School struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
		Desc    string `json:"desc"`
	}
}
