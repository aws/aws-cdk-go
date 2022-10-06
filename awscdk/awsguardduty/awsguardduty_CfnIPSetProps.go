package awsguardduty

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnIPSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPSetProps := &cfnIPSetProps{
//   	activate: jsii.Boolean(false),
//   	detectorId: jsii.String("detectorId"),
//   	format: jsii.String("format"),
//   	location: jsii.String("location"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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
	// Allowed characters are alphanumerics, spaces, hyphens (-), and underscores (_).
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::GuardDuty::IPSet.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

