package awskinesisvideo


// A reference to a SignalingChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signalingChannelReference := &SignalingChannelReference{
//   	SignalingChannelArn: jsii.String("signalingChannelArn"),
//   	SignalingChannelName: jsii.String("signalingChannelName"),
//   }
//
type SignalingChannelReference struct {
	// The ARN of the SignalingChannel resource.
	SignalingChannelArn *string `field:"required" json:"signalingChannelArn" yaml:"signalingChannelArn"`
	// The Name of the SignalingChannel resource.
	SignalingChannelName *string `field:"required" json:"signalingChannelName" yaml:"signalingChannelName"`
}

