package awsappsync


// the base properties for a channel namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var appSyncBackedDataSource AppSyncBackedDataSource
//   var code Code
//
//   baseChannelNamespaceProps := &BaseChannelNamespaceProps{
//   	AuthorizationConfig: &NamespaceAuthConfig{
//   		PublishAuthModeTypes: []AppSyncAuthorizationType{
//   			awscdk.Aws_appsync.AppSyncAuthorizationType_API_KEY,
//   		},
//   		SubscribeAuthModeTypes: []AppSyncAuthorizationType{
//   			awscdk.*Aws_appsync.AppSyncAuthorizationType_API_KEY,
//   		},
//   	},
//   	ChannelNamespaceName: jsii.String("channelNamespaceName"),
//   	Code: code,
//   	PublishHandlerConfig: &HandlerConfig{
//   		DataSource: appSyncBackedDataSource,
//   		Direct: jsii.Boolean(false),
//   		LambdaInvokeType: awscdk.*Aws_appsync.LambdaInvokeType_EVENT,
//   	},
//   	SubscribeHandlerConfig: &HandlerConfig{
//   		DataSource: appSyncBackedDataSource,
//   		Direct: jsii.Boolean(false),
//   		LambdaInvokeType: awscdk.*Aws_appsync.LambdaInvokeType_EVENT,
//   	},
//   }
//
type BaseChannelNamespaceProps struct {
	// Authorization config for channel namespace.
	// Default: - defaults to Event API default auth config.
	//
	AuthorizationConfig *NamespaceAuthConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// the name of the channel namespace.
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

