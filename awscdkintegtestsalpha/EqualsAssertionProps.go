// CDK Integration Testing Constructs
package awscdkintegtestsalpha


// Options for an EqualsAssertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   var actualResult actualResult
//   var expectedResult expectedResult
//
//   equalsAssertionProps := &EqualsAssertionProps{
//   	Actual: actualResult,
//   	Expected: expectedResult,
//
//   	// the properties below are optional
//   	FailDeployment: jsii.Boolean(false),
//   }
//
// Experimental.
type EqualsAssertionProps struct {
	// The actual results to compare.
	// Experimental.
	Actual ActualResult `field:"required" json:"actual" yaml:"actual"`
	// The expected result to assert.
	// Experimental.
	Expected ExpectedResult `field:"required" json:"expected" yaml:"expected"`
	// Set this to true if a failed assertion should result in a CloudFormation deployment failure.
	//
	// This is only necessary if assertions are being
	// executed outside of `integ-runner`.
	// Experimental.
	FailDeployment *bool `field:"optional" json:"failDeployment" yaml:"failDeployment"`
}

