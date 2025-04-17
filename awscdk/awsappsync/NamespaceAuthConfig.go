package awsappsync


// Authorization configuration for the Channel Namespace.
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
type NamespaceAuthConfig struct {
	// The publish auth modes for this Event Api.
	// Default: - API Key authorization.
	//
	PublishAuthModeTypes *[]AppSyncAuthorizationType `field:"optional" json:"publishAuthModeTypes" yaml:"publishAuthModeTypes"`
	// The subscribe auth modes for this Event Api.
	// Default: - API Key authorization.
	//
	SubscribeAuthModeTypes *[]AppSyncAuthorizationType `field:"optional" json:"subscribeAuthModeTypes" yaml:"subscribeAuthModeTypes"`
}

