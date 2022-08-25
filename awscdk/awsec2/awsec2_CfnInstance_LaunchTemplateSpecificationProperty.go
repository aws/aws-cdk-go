package awsec2


// Specifies a launch template. You must specify either the launch template ID or launch template name.
//
// `LaunchTemplateSpecification` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateSpecificationProperty := &launchTemplateSpecificationProperty{
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	launchTemplateId: jsii.String("launchTemplateId"),
//   	launchTemplateName: jsii.String("launchTemplateName"),
//   }
//
type CfnInstance_LaunchTemplateSpecificationProperty struct {
	// The version number of the launch template.
	//
	// AWS CloudFormation does not support specifying `$Latest` , or `$Default` for the template version number.
	Version *string `field:"required" json:"version" yaml:"version"`
	// The ID of the launch template.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The name of the launch template.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
}

