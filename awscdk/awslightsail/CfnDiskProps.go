package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDisk`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDiskProps := &CfnDiskProps{
//   	DiskName: jsii.String("diskName"),
//   	SizeInGb: jsii.Number(123),
//
//   	// the properties below are optional
//   	AddOns: []interface{}{
//   		&AddOnProperty{
//   			AddOnType: jsii.String("addOnType"),
//
//   			// the properties below are optional
//   			AutoSnapshotAddOnRequest: &AutoSnapshotAddOnProperty{
//   				SnapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   			},
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	Location: &LocationProperty{
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		RegionName: jsii.String("regionName"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disk.html
//
type CfnDiskProps struct {
	// The name of the disk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disk.html#cfn-lightsail-disk-diskname
	//
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	// The size of the disk in GB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disk.html#cfn-lightsail-disk-sizeingb
	//
	SizeInGb *float64 `field:"required" json:"sizeInGb" yaml:"sizeInGb"`
	// An array of add-ons for the disk.
	//
	// > If the disk has an add-on enabled when performing a delete disk request, the add-on is automatically disabled before the disk is deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disk.html#cfn-lightsail-disk-addons
	//
	AddOns interface{} `field:"optional" json:"addOns" yaml:"addOns"`
	// The AWS Region and Availability Zone location for the disk (for example, `us-east-1a` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disk.html#cfn-lightsail-disk-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AWS Region and Availability Zone where the disk is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disk.html#cfn-lightsail-disk-location
	//
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disk.html#cfn-lightsail-disk-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

