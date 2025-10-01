package awsmediapackagev2


// A reference to a ChannelGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelGroupReference := &ChannelGroupReference{
//   	ChannelGroupArn: jsii.String("channelGroupArn"),
//   }
//
type ChannelGroupReference struct {
	// The Arn of the ChannelGroup resource.
	ChannelGroupArn *string `field:"required" json:"channelGroupArn" yaml:"channelGroupArn"`
}

