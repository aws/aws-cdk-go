// An experiment to bundle the entire CDK into a single module
package awscdk


// Options for creating a lazy list token.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   lazyListValueOptions := &LazyListValueOptions{
//   	DisplayHint: jsii.String("displayHint"),
//   	OmitEmpty: jsii.Boolean(false),
//   }
//
// Experimental.
type LazyListValueOptions struct {
	// Use the given name as a display hint.
	// Experimental.
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
	// If the produced list is empty, return 'undefined' instead.
	// Experimental.
	OmitEmpty *bool `field:"optional" json:"omitEmpty" yaml:"omitEmpty"`
}

