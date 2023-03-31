package awskinesisvideo

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStreamProps := &cfnStreamProps{
//   	dataRetentionInHours: jsii.Number(123),
//   	deviceName: jsii.String("deviceName"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	mediaType: jsii.String("mediaType"),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStreamProps struct {
	// How long the stream retains data, in hours.
	DataRetentionInHours *float64 `field:"optional" json:"dataRetentionInHours" yaml:"dataRetentionInHours"`
	// The name of the device that is associated with the stream.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// The ID of the AWS Key Management Service ( AWS KMS ) key that Kinesis Video Streams uses to encrypt data on the stream.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The `MediaType` of the stream.
	MediaType *string `field:"optional" json:"mediaType" yaml:"mediaType"`
	// The name of the stream.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

