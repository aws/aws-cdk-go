package awsec2


// Describes the Amazon EC2 launch template and the launch template version that can be used by a Spot Fleet request to configure Amazon EC2 instances.
//
// For information about launch templates, see [Launching an instance from a launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html) in the *Amazon EC2 User Guide for Linux Instances* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fleetLaunchTemplateSpecificationProperty := &fleetLaunchTemplateSpecificationProperty{
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	launchTemplateId: jsii.String("launchTemplateId"),
//   	launchTemplateName: jsii.String("launchTemplateName"),
//   }
//
type CfnSpotFleet_FleetLaunchTemplateSpecificationProperty struct {
	// The version number of the launch template. You must specify a version number.
	//
	// Minimum length of 1. Maximum length of 255. Versions must fit the following pattern: `[\ u0020-\ uD7FF\ uE000-\ uFFFD\ uD800\ uDC00-\ uDBFF\ uDFFF\r\n\t]*`
	Version *string `field:"required" json:"version" yaml:"version"`
	// The ID of the launch template.
	//
	// If you specify the template ID, you can't specify the template name.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The name of the launch template. You must specify either a template name or a template ID.
	//
	// Minimum length of 3. Maximum length of 128. Names must match the following pattern: `[a-zA-Z0-9\(\)\.-/_]+`
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
}

