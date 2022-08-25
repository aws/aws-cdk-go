// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Options for the 'reverse()' operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   reverseOptions := &reverseOptions{
//   	failConcat: jsii.Boolean(false),
//   }
//
type ReverseOptions struct {
	// Fail if the given string is a concatenation.
	//
	// If `false`, just return `undefined`.
	FailConcat *bool `field:"optional" json:"failConcat" yaml:"failConcat"`
}

