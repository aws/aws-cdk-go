package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The launch template to apply the distributed AMI to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var launchTemplate LaunchTemplate
//
//   launchTemplateConfiguration := &LaunchTemplateConfiguration{
//   	LaunchTemplate: launchTemplate,
//
//   	// the properties below are optional
//   	AccountId: jsii.String("accountId"),
//   	SetDefaultVersion: jsii.Boolean(false),
//   }
//
// Experimental.
type LaunchTemplateConfiguration struct {
	// The launch template to apply the distributed AMI to.
	//
	// A new launch template version will be created for the
	// provided launch template with the distributed AMI applied.
	//
	// *Note:* The launch template should expose a `launchTemplateId`. Templates
	// imported by name only are not supported.
	// Experimental.
	LaunchTemplate awsec2.ILaunchTemplate `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
	// The AWS account ID that owns the launch template.
	// Default: The current account is used.
	//
	// Experimental.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Whether to set the new launch template version that is created as the default launch template version.
	//
	// After
	// creation of the launch template version containing the distributed AMI, it will be automatically set as the
	// default version for the launch template.
	// Default: false.
	//
	// Experimental.
	SetDefaultVersion *bool `field:"optional" json:"setDefaultVersion" yaml:"setDefaultVersion"`
}

