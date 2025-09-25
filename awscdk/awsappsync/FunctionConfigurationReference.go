package awsappsync


// A reference to a FunctionConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionConfigurationReference := &FunctionConfigurationReference{
//   	FunctionArn: jsii.String("functionArn"),
//   }
//
type FunctionConfigurationReference struct {
	// The FunctionArn of the FunctionConfiguration resource.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
}

