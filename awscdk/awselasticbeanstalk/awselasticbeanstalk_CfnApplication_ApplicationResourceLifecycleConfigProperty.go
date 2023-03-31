package awselasticbeanstalk


// Use the `ApplicationResourceLifecycleConfig` property type to specify lifecycle settings for resources that belong to an AWS Elastic Beanstalk application when defining an AWS::ElasticBeanstalk::Application resource in an AWS CloudFormation template.
//
// The resource lifecycle configuration for an application. Defines lifecycle settings for resources that belong to the application, and the service role that Elastic Beanstalk assumes in order to apply lifecycle settings. The version lifecycle configuration defines lifecycle settings for application versions.
//
// `ApplicationResourceLifecycleConfig` is a property of the [AWS::ElasticBeanstalk::Application](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationResourceLifecycleConfigProperty := &applicationResourceLifecycleConfigProperty{
//   	serviceRole: jsii.String("serviceRole"),
//   	versionLifecycleConfig: &applicationVersionLifecycleConfigProperty{
//   		maxAgeRule: &maxAgeRuleProperty{
//   			deleteSourceFromS3: jsii.Boolean(false),
//   			enabled: jsii.Boolean(false),
//   			maxAgeInDays: jsii.Number(123),
//   		},
//   		maxCountRule: &maxCountRuleProperty{
//   			deleteSourceFromS3: jsii.Boolean(false),
//   			enabled: jsii.Boolean(false),
//   			maxCount: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnApplication_ApplicationResourceLifecycleConfigProperty struct {
	// The ARN of an IAM service role that Elastic Beanstalk has permission to assume.
	//
	// The `ServiceRole` property is required the first time that you provide a `ResourceLifecycleConfig` for the application. After you provide it once, Elastic Beanstalk persists the Service Role with the application, and you don't need to specify it again. You can, however, specify it in subsequent updates to change the Service Role to another value.
	ServiceRole *string `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// Defines lifecycle settings for application versions.
	VersionLifecycleConfig interface{} `field:"optional" json:"versionLifecycleConfig" yaml:"versionLifecycleConfig"`
}

