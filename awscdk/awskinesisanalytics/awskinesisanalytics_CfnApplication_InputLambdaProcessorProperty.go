package awskinesisanalytics


// An object that contains the Amazon Resource Name (ARN) of the [AWS Lambda](https://docs.aws.amazon.com/lambda/) function that is used to preprocess records in the stream, and the ARN of the IAM role that is used to access the AWS Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputLambdaProcessorProperty := &inputLambdaProcessorProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnApplication_InputLambdaProcessorProperty struct {
	// The ARN of the [AWS Lambda](https://docs.aws.amazon.com/lambda/) function that operates on records in the stream.
	//
	// > To specify an earlier version of the Lambda function than the latest, include the Lambda function version in the Lambda function ARN. For more information about Lambda ARNs, see [Example ARNs: AWS Lambda](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-lambda)
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The ARN of the IAM role that is used to access the AWS Lambda function.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

