package awskinesisanalytics


// When you configure a SQL-based Kinesis Data Analytics application's output, identifies an Amazon Lambda function as the destination.
//
// You provide the function Amazon Resource Name (ARN) of the Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaOutputProperty := &lambdaOutputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnApplicationOutputV2_LambdaOutputProperty struct {
	// The Amazon Resource Name (ARN) of the destination Lambda function to write to.
	//
	// > To specify an earlier version of the Lambda function than the latest, include the Lambda function version in the Lambda function ARN. For more information about Lambda ARNs, see [Example ARNs: Amazon Lambda](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-lambda)
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

