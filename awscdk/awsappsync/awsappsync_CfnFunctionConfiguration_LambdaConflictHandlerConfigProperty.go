package awsappsync


// The `LambdaConflictHandlerConfig` object when configuring `LAMBDA` as the Conflict Handler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaConflictHandlerConfigProperty := &lambdaConflictHandlerConfigProperty{
//   	lambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   }
//
type CfnFunctionConfiguration_LambdaConflictHandlerConfigProperty struct {
	// The Amazon Resource Name (ARN) for the Lambda function to use as the Conflict Handler.
	LambdaConflictHandlerArn *string `field:"optional" json:"lambdaConflictHandlerArn" yaml:"lambdaConflictHandlerArn"`
}

