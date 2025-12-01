package previewawskinesisvideomixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnStreamPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStreamMixinProps := &CfnStreamMixinProps{
//   	DataRetentionInHours: jsii.Number(123),
//   	DeviceName: jsii.String("deviceName"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MediaType: jsii.String("mediaType"),
//   	Name: jsii.String("name"),
//   	StreamStorageConfiguration: &StreamStorageConfigurationProperty{
//   		DefaultStorageTier: jsii.String("defaultStorageTier"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisvideo-stream.html
//
type CfnStreamMixinProps struct {
	// How long the stream retains data, in hours.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisvideo-stream.html#cfn-kinesisvideo-stream-dataretentioninhours
	//
	DataRetentionInHours *float64 `field:"optional" json:"dataRetentionInHours" yaml:"dataRetentionInHours"`
	// The name of the device that is associated with the stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisvideo-stream.html#cfn-kinesisvideo-stream-devicename
	//
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// The ID of the AWS Key Management Service ( AWS  ) key that Kinesis Video Streams uses to encrypt data on the stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisvideo-stream.html#cfn-kinesisvideo-stream-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The `MediaType` of the stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisvideo-stream.html#cfn-kinesisvideo-stream-mediatype
	//
	MediaType *string `field:"optional" json:"mediaType" yaml:"mediaType"`
	// The name of the stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisvideo-stream.html#cfn-kinesisvideo-stream-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The configuration for stream storage, including the default storage tier for stream data.
	//
	// This configuration determines how stream data is stored and accessed, with different tiers offering varying levels of performance and cost optimization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisvideo-stream.html#cfn-kinesisvideo-stream-streamstorageconfiguration
	//
	StreamStorageConfiguration interface{} `field:"optional" json:"streamStorageConfiguration" yaml:"streamStorageConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisvideo-stream.html#cfn-kinesisvideo-stream-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

