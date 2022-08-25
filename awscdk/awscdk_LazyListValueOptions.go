// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Options for creating a lazy list token.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   lazyListValueOptions := &lazyListValueOptions{
//   	displayHint: jsii.String("displayHint"),
//   	omitEmpty: jsii.Boolean(false),
//   }
//
type LazyListValueOptions struct {
	// Use the given name as a display hint.
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
	// If the produced list is empty, return 'undefined' instead.
	OmitEmpty *bool `field:"optional" json:"omitEmpty" yaml:"omitEmpty"`
}

