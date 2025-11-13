package interfacesawsappsync


// A reference to a ChannelNamespace resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelNamespaceReference := &ChannelNamespaceReference{
//   	ChannelNamespaceArn: jsii.String("channelNamespaceArn"),
//   }
//
type ChannelNamespaceReference struct {
	// The ChannelNamespaceArn of the ChannelNamespace resource.
	ChannelNamespaceArn *string `field:"required" json:"channelNamespaceArn" yaml:"channelNamespaceArn"`
}

