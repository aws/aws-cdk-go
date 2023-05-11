package awsautoscaling


// Specifies a launch template to use when provisioning EC2 instances for an Auto Scaling group.
//
// You must specify the following:
//
// - The ID or the name of the launch template, but not both.
// - The version of the launch template.
//
// `LaunchTemplateSpecification` is property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html) resource. It is also a property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html) and [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html) property types.
//
// For information about creating a launch template, see [AWS::EC2::LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) and [Create a launch template for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// For examples of launch templates, see [Auto scaling template snippets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-autoscaling.html) and the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#aws-resource-ec2-launchtemplate--examples) section in the `AWS::EC2::LaunchTemplate` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateSpecificationProperty := &LaunchTemplateSpecificationProperty{
//   	Version: jsii.String("version"),
//
//   	// the properties below are optional
//   	LaunchTemplateId: jsii.String("launchTemplateId"),
//   	LaunchTemplateName: jsii.String("launchTemplateName"),
//   }
//
type CfnAutoScalingGroup_LaunchTemplateSpecificationProperty struct {
	// The version number of the launch template.
	//
	// Specifying `$Latest` or `$Default` for the template version number is not supported. However, you can specify `LatestVersionNumber` or `DefaultVersionNumber` using the `Fn::GetAtt` intrinsic function. For more information, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html) .
	//
	// > For an example of using the `Fn::GetAtt` function, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#aws-properties-as-group--examples) section of the `AWS::AutoScaling::AutoScalingGroup` resource.
	Version *string `field:"required" json:"version" yaml:"version"`
	// The ID of the launch template.
	//
	// You must specify the `LaunchTemplateID` or the `LaunchTemplateName` , but not both.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The name of the launch template.
	//
	// You must specify the `LaunchTemplateName` or the `LaunchTemplateID` , but not both.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
}

