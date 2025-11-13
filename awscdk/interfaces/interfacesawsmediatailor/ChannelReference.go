package interfacesawsmediatailor


// A reference to a Channel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelReference := &ChannelReference{
//   	ChannelArn: jsii.String("channelArn"),
//   	ChannelName: jsii.String("channelName"),
//   }
//
type ChannelReference struct {
	// The ARN of the Channel resource.
	ChannelArn *string `field:"required" json:"channelArn" yaml:"channelArn"`
	// The ChannelName of the Channel resource.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
}

