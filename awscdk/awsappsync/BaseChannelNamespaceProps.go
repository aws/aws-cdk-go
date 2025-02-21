package awsappsync


// the base properties for a channel namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var code code
//
//   baseChannelNamespaceProps := &BaseChannelNamespaceProps{
//   	AuthorizationConfig: &NamespaceAuthConfig{
//   		PublishAuthModeTypes: []appSyncAuthorizationType{
//   			awscdk.Aws_appsync.*appSyncAuthorizationType_API_KEY,
//   		},
//   		SubscribeAuthModeTypes: []*appSyncAuthorizationType{
//   			awscdk.*Aws_appsync.*appSyncAuthorizationType_API_KEY,
//   		},
//   	},
//   	ChannelNamespaceName: jsii.String("channelNamespaceName"),
//   	Code: code,
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
}

