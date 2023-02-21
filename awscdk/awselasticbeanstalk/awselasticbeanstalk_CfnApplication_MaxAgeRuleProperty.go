package awselasticbeanstalk


// Use the `MaxAgeRule` property type to specify a max age rule to restrict the length of time that application versions are retained for an AWS Elastic Beanstalk application when defining an AWS::ElasticBeanstalk::Application resource in an AWS CloudFormation template.
//
// A lifecycle rule that deletes application versions after the specified number of days.
//
// `MaxAgeRule` is a property of the [ApplicationVersionLifecycleConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maxAgeRuleProperty := &maxAgeRuleProperty{
//   	deleteSourceFromS3: jsii.Boolean(false),
//   	enabled: jsii.Boolean(false),
//   	maxAgeInDays: jsii.Number(123),
//   }
//
type CfnApplication_MaxAgeRuleProperty struct {
	// Set to `true` to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
	DeleteSourceFromS3 interface{} `field:"optional" json:"deleteSourceFromS3" yaml:"deleteSourceFromS3"`
	// Specify `true` to apply the rule, or `false` to disable it.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specify the number of days to retain an application versions.
	MaxAgeInDays *float64 `field:"optional" json:"maxAgeInDays" yaml:"maxAgeInDays"`
}

