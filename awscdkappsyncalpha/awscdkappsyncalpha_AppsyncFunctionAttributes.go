// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// The attributes for imported AppSync Functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appsync_alpha "github.com/aws/aws-cdk-go/awscdkappsyncalpha"
//
//   appsyncFunctionAttributes := &appsyncFunctionAttributes{
//   	functionArn: jsii.String("functionArn"),
//   }
//
// Experimental.
type AppsyncFunctionAttributes struct {
	// the ARN of the AppSync function.
	// Experimental.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
}

