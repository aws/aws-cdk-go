package awsappsync


// The attributes for imported AppSync Functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appsyncFunctionAttributes := &AppsyncFunctionAttributes{
//   	FunctionArn: jsii.String("functionArn"),
//   }
//
// Experimental.
type AppsyncFunctionAttributes struct {
	// the ARN of the AppSync function.
	// Experimental.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
}

