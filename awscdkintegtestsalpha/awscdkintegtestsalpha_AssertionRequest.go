// CDK Integration Testing Constructs
package awscdkintegtestsalpha


// A request to make an assertion that the actual value matches the expected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   var actual interface{}
//   var expected interface{}
//
//   assertionRequest := &assertionRequest{
//   	actual: actual,
//   	expected: expected,
//
//   	// the properties below are optional
//   	failDeployment: jsii.Boolean(false),
//   }
//
// Experimental.
type AssertionRequest struct {
	// The actual value received.
	// Experimental.
	Actual interface{} `field:"required" json:"actual" yaml:"actual"`
	// The expected value to assert.
	// Experimental.
	Expected interface{} `field:"required" json:"expected" yaml:"expected"`
	// Set this to true if a failed assertion should result in a CloudFormation deployment failure.
	//
	// This is only necessary if assertions are being
	// executed outside of `integ-runner`.
	// Experimental.
	FailDeployment *bool `field:"optional" json:"failDeployment" yaml:"failDeployment"`
}

