package awsappsync


// Option configuration for channel namespace.
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
type ChannelNamespaceOptions struct {
	// Authorization config for channel namespace.
	// Default: - defaults to Event API default auth config.
	//
	AuthorizationConfig *NamespaceAuthConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The Channel Namespace name.
	// Default: - the construct's id will be used.
	//
	ChannelNamespaceName *string `field:"optional" json:"channelNamespaceName" yaml:"channelNamespaceName"`
	// The Event Handler code.
	// Default: - no code is used.
	//
	Code Code `field:"optional" json:"code" yaml:"code"`
	// onPublish handler config.
	// Default: - no handler config.
	//
	PublishHandlerConfig *HandlerConfig `field:"optional" json:"publishHandlerConfig" yaml:"publishHandlerConfig"`
	// onSubscribe handler config.
	// Default: - no handler config.
	//
	SubscribeHandlerConfig *HandlerConfig `field:"optional" json:"subscribeHandlerConfig" yaml:"subscribeHandlerConfig"`
}

