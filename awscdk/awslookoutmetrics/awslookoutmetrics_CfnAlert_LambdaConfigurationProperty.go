package awslookoutmetrics


// Contains information about a Lambda configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaConfigurationProperty := &lambdaConfigurationProperty{
//   	lambdaArn: jsii.String("lambdaArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnAlert_LambdaConfigurationProperty struct {
	// The ARN of the Lambda function.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// The ARN of an IAM role that has permission to invoke the Lambda function.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

