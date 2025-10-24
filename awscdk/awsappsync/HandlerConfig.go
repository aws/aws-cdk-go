package awsappsync


// Handler configuration construct for onPublish and onSubscribe.
//
// Example:
//   var api EventApi
//   var ddbDataSource AppSyncDynamoDbDataSource
//   var ebDataSource AppSyncEventBridgeDataSource
//
//
//   // DynamoDB data source for publish handler
//   api.AddChannelNamespace(jsii.String("ddb-eb-ns"), &ChannelNamespaceOptions{
//   	Code: appsync.Code_FromInline(jsii.String("/* event handler code here.*/")),
//   	PublishHandlerConfig: &HandlerConfig{
//   		DataSource: ddbDataSource,
//   	},
//   	SubscribeHandlerConfig: &HandlerConfig{
//   		DataSource: ebDataSource,
//   	},
//   })
//
type HandlerConfig struct {
	// The Event Handler data source.
	// Default: - no data source is used.
	//
	DataSource AppSyncBackedDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	// If the Event Handler should invoke the data source directly.
	// Default: - false.
	//
	Direct *bool `field:"optional" json:"direct" yaml:"direct"`
	// The Lambda invocation type for direct integrations.
	// Default: - LambdaInvokeType.REQUEST_RESPONSE
	//
	LambdaInvokeType LambdaInvokeType `field:"optional" json:"lambdaInvokeType" yaml:"lambdaInvokeType"`
}

