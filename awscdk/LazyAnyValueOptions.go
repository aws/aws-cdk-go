package awscdk


// Options for creating lazy untyped tokens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   lazyAnyValueOptions := &LazyAnyValueOptions{
//   	DisplayHint: jsii.String("displayHint"),
//   	OmitEmptyArray: jsii.Boolean(false),
//   }
//
type LazyAnyValueOptions struct {
	// Use the given name as a display hint.
	// Default: - No hint.
	//
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
	// If the produced value is an array and it is empty, return 'undefined' instead.
	// Default: false.
	//
	OmitEmptyArray *bool `field:"optional" json:"omitEmptyArray" yaml:"omitEmptyArray"`
}

