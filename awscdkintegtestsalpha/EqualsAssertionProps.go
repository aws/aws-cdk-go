package awscdkintegtestsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Options for an EqualsAssertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var actualResult ActualResult
//   var expectedResult ExpectedResult
//
//   equalsAssertionProps := &EqualsAssertionProps{
//   	Actual: actualResult,
//   	Expected: expectedResult,
//
//   	// the properties below are optional
//   	FailDeployment: jsii.Boolean(false),
//   	ProviderLogLevel: awscdk.Aws_lambda.ApplicationLogLevel_INFO,
//   }
//
// Experimental.
type EqualsAssertionProps struct {
	// The log level of the provider lambda function.
	// Default: ApplicationLogLevel.FATAL
	//
	// Experimental.
	ProviderLogLevel awslambda.ApplicationLogLevel `field:"optional" json:"providerLogLevel" yaml:"providerLogLevel"`
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
	// Default: false.
	//
	// Experimental.
	FailDeployment *bool `field:"optional" json:"failDeployment" yaml:"failDeployment"`
}

