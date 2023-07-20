package awselasticbeanstalk


// Use the `ApplicationVersionLifecycleConfig` property type to specify application version lifecycle settings for an AWS Elastic Beanstalk application when defining an AWS::ElasticBeanstalk::Application resource in an AWS CloudFormation template.
//
// The application version lifecycle settings for an application. Defines the rules that Elastic Beanstalk applies to an application's versions in order to avoid hitting the per-region limit for application versions.
//
// When Elastic Beanstalk deletes an application version from its database, you can no longer deploy that version to an environment. The source bundle remains in S3 unless you configure the rule to delete it.
//
// `ApplicationVersionLifecycleConfig` is a property of the [ApplicationResourceLifecycleConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationresourcelifecycleconfig.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationVersionLifecycleConfigProperty := &ApplicationVersionLifecycleConfigProperty{
//   	MaxAgeRule: &MaxAgeRuleProperty{
//   		DeleteSourceFromS3: jsii.Boolean(false),
//   		Enabled: jsii.Boolean(false),
//   		MaxAgeInDays: jsii.Number(123),
//   	},
//   	MaxCountRule: &MaxCountRuleProperty{
//   		DeleteSourceFromS3: jsii.Boolean(false),
//   		Enabled: jsii.Boolean(false),
//   		MaxCount: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.html
//
type CfnApplication_ApplicationVersionLifecycleConfigProperty struct {
	// Specify a max age rule to restrict the length of time that application versions are retained for an application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.html#cfn-elasticbeanstalk-application-applicationversionlifecycleconfig-maxagerule
	//
	MaxAgeRule interface{} `field:"optional" json:"maxAgeRule" yaml:"maxAgeRule"`
	// Specify a max count rule to restrict the number of application versions that are retained for an application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.html#cfn-elasticbeanstalk-application-applicationversionlifecycleconfig-maxcountrule
	//
	MaxCountRule interface{} `field:"optional" json:"maxCountRule" yaml:"maxCountRule"`
}

