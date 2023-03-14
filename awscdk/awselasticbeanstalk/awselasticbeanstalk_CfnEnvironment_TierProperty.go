package awselasticbeanstalk


// Use the `Tier` property type to specify the environment tier for an AWS Elastic Beanstalk environment when defining an AWS::ElasticBeanstalk::Environment resource in an AWS CloudFormation template.
//
// Describes the environment tier for an [AWS::ElasticBeanstalk::Environment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html) resource. For more information, see [Environment Tiers](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/using-features-managing-env-tiers.html) in the *AWS Elastic Beanstalk Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tierProperty := &tierProperty{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//   	version: jsii.String("version"),
//   }
//
type CfnEnvironment_TierProperty struct {
	// The name of this environment tier.
	//
	// Valid values:
	//
	// - For *Web server tier* – `WebServer`
	// - For *Worker tier* – `Worker`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of this environment tier.
	//
	// Valid values:
	//
	// - For *Web server tier* – `Standard`
	// - For *Worker tier* – `SQS/HTTP`.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The version of this environment tier.
	//
	// When you don't set a value to it, Elastic Beanstalk uses the latest compatible worker tier version.
	//
	// > This member is deprecated. Any specific version that you set may become out of date. We recommend leaving it unspecified.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

