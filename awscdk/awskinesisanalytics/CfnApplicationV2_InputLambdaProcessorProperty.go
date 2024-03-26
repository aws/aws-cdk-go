package awskinesisanalytics


// An object that contains the Amazon Resource Name (ARN) of the Amazon Lambda function that is used to preprocess records in the stream in a SQL-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputLambdaProcessorProperty := &InputLambdaProcessorProperty{
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputlambdaprocessor.html
//
type CfnApplicationV2_InputLambdaProcessorProperty struct {
	// The ARN of the Amazon Lambda function that operates on records in the stream.
	//
	// > To specify an earlier version of the Lambda function than the latest, include the Lambda function version in the Lambda function ARN. For more information about Lambda ARNs, see [Example ARNs: Amazon Lambda](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-lambda)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-inputlambdaprocessor.html#cfn-kinesisanalyticsv2-application-inputlambdaprocessor-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

