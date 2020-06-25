package internal

type GraylogInputCreateResponse struct {
	ID string `json:"id"`
}

type GraylogGELFUDPGlobalInputGetResponse struct {
	ID          string                                 `json:"id"`
	Node        string                                 `json:"node"`
	Title       string                                 `json:"title"`
	Type        string                                 `json:"type"`
	Global      bool                                   `json:"global"`
	Attributes  GraylogGELFUDPGlobalInputConfiguration `json:"attributes"`
	ContentPack string                                 `json:"content_pack"`
	CreatedBy   string                                 `json:"creator_user_id"`
}
