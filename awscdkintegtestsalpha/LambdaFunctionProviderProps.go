package awscdkintegtestsalpha


// Properties for a lambda function provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   lambdaFunctionProviderProps := &LambdaFunctionProviderProps{
//   	Handler: jsii.String("handler"),
//   }
//
// Experimental.
type LambdaFunctionProviderProps struct {
	// The handler to use for the lambda function.
	// Experimental.
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
}

