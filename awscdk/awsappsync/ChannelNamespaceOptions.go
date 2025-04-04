package awsappsync


// Option configuration for channel namespace.
//
// Example:
//   var api eventApi
//
//
//   // create a channel namespace
//   // create a channel namespace
//   appsync.NewChannelNamespace(this, jsii.String("Namespace"), &ChannelNamespaceProps{
//   	Api: Api,
//   })
//
//   // You can also create a namespace through the addChannelNamespace method
//   api.AddChannelNamespace(jsii.String("AnotherNameSpace"), &ChannelNamespaceOptions{
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
}

