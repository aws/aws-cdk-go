package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnStreamKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStreamKeyProps := &cfnStreamKeyProps{
//   	channelArn: jsii.String("channelArn"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStreamKeyProps struct {
	// Channel ARN for the stream.
	ChannelArn *string `field:"required" json:"channelArn" yaml:"channelArn"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

