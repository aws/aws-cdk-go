// CDK Integration Testing Constructs
package awscdkintegtestsalpha


// Properties for defining an AssertionsProvider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   assertionsProviderProps := &assertionsProviderProps{
//   	handler: jsii.String("handler"),
//   	uuid: jsii.String("uuid"),
//   }
//
// Experimental.
type AssertionsProviderProps struct {
	// The handler to use for the lambda function.
	// Experimental.
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// This determines the uniqueness of each AssertionsProvider.
	//
	// You should only need to provide something different here if you
	// _know_ that you need a separate provider.
	// Experimental.
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

