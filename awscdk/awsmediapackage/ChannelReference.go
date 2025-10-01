package awsmediapackage


// A reference to a Channel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelReference := &ChannelReference{
//   	ChannelArn: jsii.String("channelArn"),
//   	ChannelId: jsii.String("channelId"),
//   }
//
type ChannelReference struct {
	// The ARN of the Channel resource.
	ChannelArn *string `field:"required" json:"channelArn" yaml:"channelArn"`
	// The Id of the Channel resource.
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
}

