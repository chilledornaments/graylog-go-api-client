package api

type GraylogGELFUDPGlobalInputConfiguration struct {
	BindAddress         string `json:"bind_address"`
	DecompressSizeLimit int    `json:"decompress_size_limit"`
	NumberWorkerThreads int    `json:"number_worker_threads"`
	OverrideSource      string `json:"override_source,omitempty"`
	Port                int    `json:"port"`
	RecvBufferSize      int    `json:"recv_buffer_size"`
}
type GraylogGELFUDPGlobalInput struct {
	Title         string                                 `json:"title"`
	Type          string                                 `json:"type"`
	Global        bool                                   `json:"global"`
	Configuration GraylogGELFUDPGlobalInputConfiguration `json:"configuration"`
}

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
