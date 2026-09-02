package interfacesawschime


// A reference to a ChannelFlow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelFlowReference := &ChannelFlowReference{
//   	ChannelFlowArn: jsii.String("channelFlowArn"),
//   }
//
type ChannelFlowReference struct {
	// The Arn of the ChannelFlow resource.
	ChannelFlowArn *string `field:"required" json:"channelFlowArn" yaml:"channelFlowArn"`
}

