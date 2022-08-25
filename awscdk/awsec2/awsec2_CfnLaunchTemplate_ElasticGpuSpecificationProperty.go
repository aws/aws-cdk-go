package awsec2


// Specifies a specification for an Elastic GPU for an Amazon EC2 launch template.
//
// `ElasticGpuSpecification` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticGpuSpecificationProperty := &elasticGpuSpecificationProperty{
//   	type: jsii.String("type"),
//   }
//
type CfnLaunchTemplate_ElasticGpuSpecificationProperty struct {
	// The type of Elastic Graphics accelerator.
	//
	// For more information about the values to specify for `Type` , see [Elastic Graphics Basics](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html#elastic-graphics-basics) , specifically the Elastic Graphics accelerator column, in the *Amazon Elastic Compute Cloud User Guide for Windows Instances* .
	Type *string `field:"optional" json:"type" yaml:"type"`
}

