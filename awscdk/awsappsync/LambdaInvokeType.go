package awsappsync


// Invoke types for direct Lambda data sources.
//
// Example:
//   var api eventApi
//   var lambdaDataSource appSyncLambdaDataSource
//
//
//   // Lambda data source for publish handler
//   api.AddChannelNamespace(jsii.String("lambda-ns"), &ChannelNamespaceOptions{
//   	Code: appsync.Code_FromInline(jsii.String("/* event handler code here.*/")),
//   	PublishHandlerConfig: &HandlerConfig{
//   		DataSource: lambdaDataSource,
//   	},
//   })
//
//   // Direct Lambda data source for publish handler
//   api.AddChannelNamespace(jsii.String("lambda-direct-ns"), &ChannelNamespaceOptions{
//   	PublishHandlerConfig: &HandlerConfig{
//   		DataSource: lambdaDataSource,
//   		Direct: jsii.Boolean(true),
//   	},
//   })
//
//   api.AddChannelNamespace(jsii.String("lambda-direct-async-ns"), &ChannelNamespaceOptions{
//   	PublishHandlerConfig: &HandlerConfig{
//   		DataSource: lambdaDataSource,
//   		Direct: jsii.Boolean(true),
//   		LambdaInvokeType: appsync.LambdaInvokeType_EVENT,
//   	},
//   })
//
type LambdaInvokeType string

const (
	// Invoke function asynchronously.
	LambdaInvokeType_EVENT LambdaInvokeType = "EVENT"
	// Invoke function synchronously.
	LambdaInvokeType_REQUEST_RESPONSE LambdaInvokeType = "REQUEST_RESPONSE"
)

