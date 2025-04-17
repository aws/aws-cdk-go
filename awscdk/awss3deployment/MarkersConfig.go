package awss3deployment


// A configuration for markers substitution strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   markersConfig := &MarkersConfig{
//   	JsonEscape: jsii.Boolean(false),
//   }
//
type MarkersConfig struct {
	// If set to `true`, the marker substitution will make ure the value inserted in the file will be a valid JSON string.
	// Default: - false.
	//
	JsonEscape *bool `field:"optional" json:"jsonEscape" yaml:"jsonEscape"`
}

