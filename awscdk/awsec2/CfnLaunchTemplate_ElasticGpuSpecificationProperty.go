package awsec2


// > Amazon Elastic Graphics reached end of life on January 8, 2024.
//
// Specifies a specification for an Elastic GPU for an Amazon EC2 launch template.
//
// `ElasticGpuSpecification` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticGpuSpecificationProperty := &ElasticGpuSpecificationProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-elasticgpuspecification.html
//
type CfnLaunchTemplate_ElasticGpuSpecificationProperty struct {
	// The type of Elastic Graphics accelerator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-elasticgpuspecification.html#cfn-ec2-launchtemplate-elasticgpuspecification-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

