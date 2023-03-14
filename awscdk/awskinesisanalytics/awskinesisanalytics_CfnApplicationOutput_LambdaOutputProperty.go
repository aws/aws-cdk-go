package awskinesisanalytics


// When configuring application output, identifies an AWS Lambda function as the destination.
//
// You provide the function Amazon Resource Name (ARN) and also an IAM role ARN that Amazon Kinesis Analytics can use to write to the function on your behalf.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaOutputProperty := &lambdaOutputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnApplicationOutput_LambdaOutputProperty struct {
	// Amazon Resource Name (ARN) of the destination Lambda function to write to.
	//
	// > To specify an earlier version of the Lambda function than the latest, include the Lambda function version in the Lambda function ARN. For more information about Lambda ARNs, see [Example ARNs: AWS Lambda](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-lambda)
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to write to the destination function on your behalf.
	//
	// You need to grant the necessary permissions to this role.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

