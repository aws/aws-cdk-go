package previewawss3mixins


// Describes the AWS Lambda functions to invoke and the events for which to invoke them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lambdaConfigurationProperty := &LambdaConfigurationProperty{
//   	Event: jsii.String("event"),
//   	Filter: &NotificationFilterProperty{
//   		S3Key: &S3KeyFilterProperty{
//   			Rules: []interface{}{
//   				&FilterRuleProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	Function: jsii.String("function"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lambdaconfiguration.html
//
type CfnBucketPropsMixin_LambdaConfigurationProperty struct {
	// The Amazon S3 bucket event for which to invoke the AWS Lambda function.
	//
	// For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lambdaconfiguration.html#cfn-s3-bucket-lambdaconfiguration-event
	//
	Event *string `field:"optional" json:"event" yaml:"event"`
	// The filtering rules that determine which objects invoke the AWS Lambda function.
	//
	// For example, you can create a filter so that only image files with a `.jpg` extension invoke the function when they are added to the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lambdaconfiguration.html#cfn-s3-bucket-lambdaconfiguration-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// The Amazon Resource Name (ARN) of the AWS Lambda function that Amazon S3 invokes when the specified event type occurs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lambdaconfiguration.html#cfn-s3-bucket-lambdaconfiguration-function
	//
	Function *string `field:"optional" json:"function" yaml:"function"`
}

