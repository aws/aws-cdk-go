package awslex


// Specifies the text input specifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textInputSpecificationProperty := &textInputSpecificationProperty{
//   	startTimeoutMs: jsii.Number(123),
//   }
//
type CfnBot_TextInputSpecificationProperty struct {
	// Time for which a bot waits before re-prompting a customer for text input.
	StartTimeoutMs *float64 `field:"required" json:"startTimeoutMs" yaml:"startTimeoutMs"`
}

