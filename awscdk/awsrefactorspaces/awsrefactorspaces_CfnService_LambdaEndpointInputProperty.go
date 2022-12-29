package awsrefactorspaces


// The input for the AWS Lambda endpoint type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaEndpointInputProperty := &lambdaEndpointInputProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnService_LambdaEndpointInputProperty struct {
	// The Amazon Resource Name (ARN) of the Lambda endpoint.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

