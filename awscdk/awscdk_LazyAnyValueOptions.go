// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Options for creating lazy untyped tokens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   lazyAnyValueOptions := &lazyAnyValueOptions{
//   	displayHint: jsii.String("displayHint"),
//   	omitEmptyArray: jsii.Boolean(false),
//   }
//
type LazyAnyValueOptions struct {
	// Use the given name as a display hint.
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
	// If the produced value is an array and it is empty, return 'undefined' instead.
	OmitEmptyArray *bool `field:"optional" json:"omitEmptyArray" yaml:"omitEmptyArray"`
}

