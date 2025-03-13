package awsec2


// Specifies an elastic inference accelerator.
//
// `LaunchTemplateElasticInferenceAccelerator` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateElasticInferenceAcceleratorProperty := &LaunchTemplateElasticInferenceAcceleratorProperty{
//   	Count: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplateelasticinferenceaccelerator.html
//
type CfnLaunchTemplate_LaunchTemplateElasticInferenceAcceleratorProperty struct {
	// The number of elastic inference accelerators to attach to the instance.
	//
	// Default: 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplateelasticinferenceaccelerator.html#cfn-ec2-launchtemplate-launchtemplateelasticinferenceaccelerator-count
	//
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The type of elastic inference accelerator.
	//
	// The possible values are eia1.medium, eia1.large, and eia1.xlarge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplateelasticinferenceaccelerator.html#cfn-ec2-launchtemplate-launchtemplateelasticinferenceaccelerator-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

