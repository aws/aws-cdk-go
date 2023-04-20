// An experiment to bundle the entire CDK into a single module
package awscdk


// An error returned during the validation phase.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var construct construct
//
//   validationError := &ValidationError{
//   	Message: jsii.String("message"),
//   	Source: construct,
//   }
//
// Experimental.
type ValidationError struct {
	// The error message.
	// Experimental.
	Message *string `field:"required" json:"message" yaml:"message"`
	// The construct which emitted the error.
	// Experimental.
	Source Construct `field:"required" json:"source" yaml:"source"`
}

