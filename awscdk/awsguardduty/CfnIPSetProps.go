package awsguardduty

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIPSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPSetProps := &CfnIPSetProps{
//   	Activate: jsii.Boolean(false),
//   	DetectorId: jsii.String("detectorId"),
//   	Format: jsii.String("format"),
//   	Location: jsii.String("location"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnIPSetProps struct {
	// Indicates whether or not  uses the `IPSet` .
	Activate interface{} `field:"required" json:"activate" yaml:"activate"`
	// The unique ID of the detector of the GuardDuty account that you want to create an IPSet for.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The format of the file that contains the IPSet.
	Format *string `field:"required" json:"format" yaml:"format"`
	// The URI of the file that contains the IPSet.
	Location *string `field:"required" json:"location" yaml:"location"`
	// The user-friendly name to identify the IPSet.
	//
	// Allowed characters are alphanumeric, whitespace, dash (-), and underscores (_).
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags to be added to a new IP set resource.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

