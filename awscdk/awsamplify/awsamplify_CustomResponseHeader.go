package awsamplify


// Custom response header of an Amplify App.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customResponseHeader := &CustomResponseHeader{
//   	Headers: map[string]*string{
//   		"headersKey": jsii.String("headers"),
//   	},
//   	Pattern: jsii.String("pattern"),
//   }
//
// Experimental.
type CustomResponseHeader struct {
	// The map of custom headers to be applied.
	// Experimental.
	Headers *map[string]*string `field:"required" json:"headers" yaml:"headers"`
	// These custom headers will be applied to all URL file paths that match this pattern.
	// Experimental.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
}

