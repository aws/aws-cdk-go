package awscdkamplifyalpha


// Custom response header of an Amplify App.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import amplify_alpha "github.com/aws/aws-cdk-go/awscdkamplifyalpha"
//
//   customResponseHeader := &CustomResponseHeader{
//   	Headers: map[string]*string{
//   		"headersKey": jsii.String("headers"),
//   	},
//   	Pattern: jsii.String("pattern"),
//
//   	// the properties below are optional
//   	AppRoot: jsii.String("appRoot"),
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
	// If the app uses a monorepo structure, the appRoot from the build spec to apply the custom headers to.
	// Default: - The appRoot is omitted in the custom headers output.
	//
	// Experimental.
	AppRoot *string `field:"optional" json:"appRoot" yaml:"appRoot"`
}

