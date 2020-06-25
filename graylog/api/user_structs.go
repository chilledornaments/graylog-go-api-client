package api

type UserItem struct {
	Username    string
	Roles       []interface{}
	Permissions []interface{}
	Password    string
	Email       string
	FullName    string
}

type graylogUserItem struct {
	Username         string   `json:"username"`
	Password         string   `json:"password,omitempty"`
	Email            string   `json:"email,omitempty"`
	FullName         string   `json:"full_name,omitempty"`
	Permissions      []string `json:"permissions,omitempty"`
	Timezone         string   `json:"timezone,omitempty"`
	SessionTimeoutMS int      `json:"session_timeout_ms,omitempty"`
	StartPage        string   `json:"start_page,omitempty"`
	Roles            []string `json:"roles,omitempty"`
	ClientAddress    string   `json:"client_address,omitempty"`
	External         bool     `json:"external,omitempty"`
	/*
		Preferences      struct {
			UpdateUnfocussed  bool `json:"updateUnfocussed,omitempty"`
			EnableSmartSearch bool `json:"enableSmartSearch,omitempty"`
		} `json:"preferences,omitempty"`
	*/
}
