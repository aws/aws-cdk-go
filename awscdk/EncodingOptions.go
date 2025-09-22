package awscdk


// Properties to string encodings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   encodingOptions := &EncodingOptions{
//   	DisplayHint: jsii.String("displayHint"),
//   }
//
type EncodingOptions struct {
	// A hint for the Token's purpose when stringifying it.
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
}

