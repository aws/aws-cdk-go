package awskinesisvideo

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnSignalingChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSignalingChannelProps := &cfnSignalingChannelProps{
//   	messageTtlSeconds: jsii.Number(123),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnSignalingChannelProps struct {
	// The period of time a signaling channel retains undelivered messages before they are discarded.
	MessageTtlSeconds *float64 `field:"optional" json:"messageTtlSeconds" yaml:"messageTtlSeconds"`
	// A name for the signaling channel that you are creating.
	//
	// It must be unique for each AWS account and AWS Region .
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A type of the signaling channel that you are creating.
	//
	// Currently, `SINGLE_MASTER` is the only supported channel type.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

