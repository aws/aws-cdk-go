package awselasticbeanstalk


// Use the `MaxAgeRule` property type to specify a max count rule to restrict the number of application versions that are retained for an AWS Elastic Beanstalk application when defining an AWS::ElasticBeanstalk::Application resource in an AWS CloudFormation template.
//
// A lifecycle rule that deletes the oldest application version when the maximum count is exceeded.
//
// `MaxCountRule` is a property of the [ApplicationVersionLifecycleConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maxCountRuleProperty := &MaxCountRuleProperty{
//   	DeleteSourceFromS3: jsii.Boolean(false),
//   	Enabled: jsii.Boolean(false),
//   	MaxCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html
//
type CfnApplication_MaxCountRuleProperty struct {
	// Set to `true` to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-deletesourcefroms3
	//
	DeleteSourceFromS3 interface{} `field:"optional" json:"deleteSourceFromS3" yaml:"deleteSourceFromS3"`
	// Specify `true` to apply the rule, or `false` to disable it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specify the maximum number of application versions to retain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-maxcount
	//
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
}

