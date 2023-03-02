// An experiment to bundle the entire CDK into a single module
package awscdk


// Properties to string encodings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   encodingOptions := &encodingOptions{
//   	displayHint: jsii.String("displayHint"),
//   }
//
// Experimental.
type EncodingOptions struct {
	// A hint for the Token's purpose when stringifying it.
	// Experimental.
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
}

