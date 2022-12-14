package awsec2


// Specifies the launch template to use for an EC2 Fleet.
//
// You must specify either the launch template ID or launch template name in the request.
//
// `FleetLaunchTemplateSpecificationRequest` is a property of the [FleetLaunchTemplateConfigRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-fleetlaunchtemplateconfigrequest.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fleetLaunchTemplateSpecificationRequestProperty := &fleetLaunchTemplateSpecificationRequestProperty{
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	launchTemplateId: jsii.String("launchTemplateId"),
//   	launchTemplateName: jsii.String("launchTemplateName"),
//   }
//
type CfnEC2Fleet_FleetLaunchTemplateSpecificationRequestProperty struct {
	// The launch template version number, `$Latest` , or `$Default` . You must specify a value, otherwise the request fails.
	//
	// If the value is `$Latest` , Amazon EC2 uses the latest version of the launch template.
	//
	// If the value is `$Default` , Amazon EC2 uses the default version of the launch template.
	Version *string `field:"required" json:"version" yaml:"version"`
	// The ID of the launch template.
	//
	// If you specify the template ID, you can't specify the template name.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The name of the launch template.
	//
	// If you specify the template name, you can't specify the template ID.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
}

