package api

type GraylogSyslogUDPGlobalInputConfiguration struct {
	BindAddress          string `json:"bind_address"`
	NumberWorkerThreads  int    `json:"number_worker_threads"`
	OverrideSource       string `json:"override_source,omitempty"`
	Port                 int    `json:"port"`
	RecvBufferSize       int    `json:"recv_buffer_size"`
	ExpandStructuredData bool   `json:"expand_structured_data"`
	ForceRDNS            bool   `json:"force_rdns"`
	AllowDateOverride    bool   `json:"allow_override_date"`
	StoreFullMessage     bool   `json:"store_full_message"`
}

type GraylogSyslogUDPGlobalInput struct {
	Title         string                                   `json:"title"`
	Type          string                                   `json:"type"`
	Global        bool                                     `json:"global"`
	Configuration GraylogSyslogUDPGlobalInputConfiguration `json:"configuration"`
}

type GraylogSyslogUDPGlobalInputGetResponse struct {
	ID          string                                   `json:"id"`
	Node        string                                   `json:"node"`
	Title       string                                   `json:"title"`
	Type        string                                   `json:"type"`
	Global      bool                                     `json:"global"`
	Attributes  GraylogSyslogUDPGlobalInputConfiguration `json:"attributes"`
	ContentPack string                                   `json:"content_pack"`
	CreatedBy   string                                   `json:"creator_user_id"`
}
