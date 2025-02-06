package awsappsync


// Additional property for an AppSync channel namespace for an Event API reference.
//
// Example:
//   var api eventApi
//
//
//   appsync.NewChannelNamespace(this, jsii.String("Namespace"), &ChannelNamespaceProps{
//   	Api: Api,
//   	AuthorizationConfig: &NamespaceAuthConfig{
//   		// Override publishing authorization to API Key
//   		PublishAuthModeTypes: []appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_API_KEY,
//   		},
//   		// Override subscribing authorization to Lambda
//   		SubscribeAuthModeTypes: []*appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_LAMBDA,
//   		},
//   	},
//   })
//
type ChannelNamespaceProps struct {
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
	// The API this channel namespace is associated with.
	Api IEventApi `field:"required" json:"api" yaml:"api"`
}

