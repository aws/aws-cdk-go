package awsappsync


// Valid data source types for AppSync.
type AppSyncDataSourceType string

const (
	// Lambda data source type.
	AppSyncDataSourceType_LAMBDA AppSyncDataSourceType = "LAMBDA"
	// DynamoDB data source type.
	AppSyncDataSourceType_DYNAMODB AppSyncDataSourceType = "DYNAMODB"
	// EventBridge data source type.
	AppSyncDataSourceType_EVENTBRIDGE AppSyncDataSourceType = "EVENTBRIDGE"
	// OpenSearch service data source type.
	AppSyncDataSourceType_OPENSEARCH_SERVICE AppSyncDataSourceType = "OPENSEARCH_SERVICE"
	// HTTP data source type.
	AppSyncDataSourceType_HTTP AppSyncDataSourceType = "HTTP"
	// Relational DB data source type.
	AppSyncDataSourceType_RELATIONAL_DATABASE AppSyncDataSourceType = "RELATIONAL_DATABASE"
	// Bedrock runtime data source type.
	AppSyncDataSourceType_BEDROCK AppSyncDataSourceType = "BEDROCK"
)

