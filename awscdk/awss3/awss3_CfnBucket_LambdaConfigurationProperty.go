package awss3


// Describes the AWS Lambda functions to invoke and the events for which to invoke them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaConfigurationProperty := &lambdaConfigurationProperty{
//   	event: jsii.String("event"),
//   	function: jsii.String("function"),
//
//   	// the properties below are optional
//   	filter: &notificationFilterProperty{
//   		s3Key: &s3KeyFilterProperty{
//   			rules: []interface{}{
//   				&filterRuleProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_LambdaConfigurationProperty struct {
	// The Amazon S3 bucket event for which to invoke the AWS Lambda function.
	//
	// For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
	Event *string `field:"required" json:"event" yaml:"event"`
	// The Amazon Resource Name (ARN) of the AWS Lambda function that Amazon S3 invokes when the specified event type occurs.
	Function *string `field:"required" json:"function" yaml:"function"`
	// The filtering rules that determine which objects invoke the AWS Lambda function.
	//
	// For example, you can create a filter so that only image files with a `.jpg` extension invoke the function when they are added to the Amazon S3 bucket.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

