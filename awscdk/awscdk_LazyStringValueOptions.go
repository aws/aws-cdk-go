// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Options for creating a lazy string token.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   lazyStringValueOptions := &lazyStringValueOptions{
//   	displayHint: jsii.String("displayHint"),
//   }
//
type LazyStringValueOptions struct {
	// Use the given name as a display hint.
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
}

