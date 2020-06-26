package api

/*
{"title":"TestFooBar","description":"hello world","index_prefix":"testingtest","writable":true,"shards":"1","replicas":0,"retention_strategy_class":"org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy","retention_strategy":{"max_number_of_indices":"1","type":"org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"},"rotation_strategy_class":"org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy","rotation_strategy":{"max_docs_per_index":20000000,"type":"org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"},"index_analyzer":"standard","index_optimization_max_num_segments":1,"index_optimization_disabled":false,"field_type_refresh_interval":1000000,"creation_date":"2020-06-26T22:47:32.039Z"}
*/

type GraylogIndexSetRetentionStrategy struct {
	MaxNumberIndices string `json:"max_number_of_indices"`
	Type             string `json:"type"`
}

type GraylogNewIndexSetInput struct {
	Title                     string                           `json:"title"`
	Description               string                           `json:"description"`
	IndexPrefix               string                           `json:"index_prefix"`
	Writable                  bool                             `json:"writable"`
	Shards                    string                           `json:"shards"`
	Replicas                  int                              `json:"replicas"`
	RetentionStrategyClass    string                           `json:"retention_strategy_class"`
	RetentionStrategy         GraylogIndexSetRetentionStrategy `json:"retention_strategy"`
	IndexAnalyzer             string                           `json:"index_analyzer"`
	IndexOptimizationDisabled bool                             `json:"index_optimization_disabled"`
	FieldTypeRefreshInterval  int                              `json:"field_type_refresh_interval"`
}
