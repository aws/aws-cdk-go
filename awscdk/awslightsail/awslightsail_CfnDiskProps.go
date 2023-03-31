package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDisk`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDiskProps := &cfnDiskProps{
//   	diskName: jsii.String("diskName"),
//   	sizeInGb: jsii.Number(123),
//
//   	// the properties below are optional
//   	addOns: []interface{}{
//   		&addOnProperty{
//   			addOnType: jsii.String("addOnType"),
//
//   			// the properties below are optional
//   			autoSnapshotAddOnRequest: &autoSnapshotAddOnProperty{
//   				snapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   			},
//   			status: jsii.String("status"),
//   		},
//   	},
//   	availabilityZone: jsii.String("availabilityZone"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDiskProps struct {
	// The name of the disk.
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	// The size of the disk in GB.
	SizeInGb *float64 `field:"required" json:"sizeInGb" yaml:"sizeInGb"`
	// An array of add-ons for the disk.
	//
	// > If the disk has an add-on enabled when performing a delete disk request, the add-on is automatically disabled before the disk is deleted.
	AddOns interface{} `field:"optional" json:"addOns" yaml:"addOns"`
	// The AWS Region and Availability Zone location for the disk (for example, `us-east-1a` ).
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

