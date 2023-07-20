package awsec2


// Specifies the launch template to be used by the EC2 Fleet for configuring Amazon EC2 instances.
//
// You must specify the following:
//
// - The ID or the name of the launch template, but not both.
// - The version of the launch template.
//
// `FleetLaunchTemplateSpecificationRequest` is a property of the [FleetLaunchTemplateConfigRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-fleetlaunchtemplateconfigrequest.html) property type.
//
// For information about creating a launch template, see [AWS::EC2::LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) and [Create a launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#create-launch-template) in the *Amazon EC2 User Guide* .
//
// For examples of launch templates, see [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#aws-resource-ec2-launchtemplate--examples) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fleetLaunchTemplateSpecificationRequestProperty := &FleetLaunchTemplateSpecificationRequestProperty{
//   	Version: jsii.String("version"),
//
//   	// the properties below are optional
//   	LaunchTemplateId: jsii.String("launchTemplateId"),
//   	LaunchTemplateName: jsii.String("launchTemplateName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-fleetlaunchtemplatespecificationrequest.html
//
type CfnEC2Fleet_FleetLaunchTemplateSpecificationRequestProperty struct {
	// The launch template version number, `$Latest` , or `$Default` . You must specify a value, otherwise the request fails.
	//
	// If the value is `$Latest` , Amazon EC2 uses the latest version of the launch template.
	//
	// If the value is `$Default` , Amazon EC2 uses the default version of the launch template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-fleetlaunchtemplatespecificationrequest.html#cfn-ec2-ec2fleet-fleetlaunchtemplatespecificationrequest-version
	//
	Version *string `field:"required" json:"version" yaml:"version"`
	// The ID of the launch template.
	//
	// You must specify the `LaunchTemplateId` or the `LaunchTemplateName` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-fleetlaunchtemplatespecificationrequest.html#cfn-ec2-ec2fleet-fleetlaunchtemplatespecificationrequest-launchtemplateid
	//
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The name of the launch template.
	//
	// You must specify the `LaunchTemplateName` or the `LaunchTemplateId` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-fleetlaunchtemplatespecificationrequest.html#cfn-ec2-ec2fleet-fleetlaunchtemplatespecificationrequest-launchtemplatename
	//
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
}

