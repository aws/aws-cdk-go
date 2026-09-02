package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFpgaImage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFpgaImageProps := &CfnFpgaImageProps{
//   	Description: jsii.String("description"),
//   	InputStorageLocation: &StorageLocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//   	LogsStorageLocation: &StorageLocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fpgaimage.html
//
type CfnFpgaImageProps struct {
	// A description for the AFI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fpgaimage.html#cfn-ec2-fpgaimage-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Describes a storage location in Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fpgaimage.html#cfn-ec2-fpgaimage-inputstoragelocation
	//
	InputStorageLocation interface{} `field:"optional" json:"inputStorageLocation" yaml:"inputStorageLocation"`
	// Describes a storage location in Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fpgaimage.html#cfn-ec2-fpgaimage-logsstoragelocation
	//
	LogsStorageLocation interface{} `field:"optional" json:"logsStorageLocation" yaml:"logsStorageLocation"`
	// A name for the AFI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fpgaimage.html#cfn-ec2-fpgaimage-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags assigned to the FPGA image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fpgaimage.html#cfn-ec2-fpgaimage-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

