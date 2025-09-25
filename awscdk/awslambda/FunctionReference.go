package awslambda


// A reference to a Function resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionReference := &FunctionReference{
//   	FunctionArn: jsii.String("functionArn"),
//   	FunctionName: jsii.String("functionName"),
//   }
//
type FunctionReference struct {
	// The ARN of the Function resource.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The FunctionName of the Function resource.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
}

