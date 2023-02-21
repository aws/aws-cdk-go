// An experiment to bundle the entire CDK into a single module
package awscdk


// Options for creating lazy untyped tokens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   lazyAnyValueOptions := &lazyAnyValueOptions{
//   	displayHint: jsii.String("displayHint"),
//   	omitEmptyArray: jsii.Boolean(false),
//   }
//
// Experimental.
type LazyAnyValueOptions struct {
	// Use the given name as a display hint.
	// Experimental.
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
	// If the produced value is an array and it is empty, return 'undefined' instead.
	// Experimental.
	OmitEmptyArray *bool `field:"optional" json:"omitEmptyArray" yaml:"omitEmptyArray"`
}

