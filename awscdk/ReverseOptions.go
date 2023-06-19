package awscdk


// Options for the 'reverse()' operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   reverseOptions := &ReverseOptions{
//   	FailConcat: jsii.Boolean(false),
//   }
//
// Experimental.
type ReverseOptions struct {
	// Fail if the given string is a concatenation.
	//
	// If `false`, just return `undefined`.
	// Experimental.
	FailConcat *bool `field:"optional" json:"failConcat" yaml:"failConcat"`
}

