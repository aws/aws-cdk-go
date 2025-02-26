package awscdk


// Options for creating a lazy string token.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   lazyStringValueOptions := &LazyStringValueOptions{
//   	DisplayHint: jsii.String("displayHint"),
//   }
//
type LazyStringValueOptions struct {
	// Use the given name as a display hint.
	// Default: - No hint.
	//
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
}

