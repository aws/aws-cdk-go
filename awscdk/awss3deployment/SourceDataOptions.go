package awss3deployment


// Options for `Source.data()` sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceDataOptions := &SourceDataOptions{
//   	JsonEscape: jsii.Boolean(false),
//   }
//
type SourceDataOptions struct {
	// Default marker encoding for markers that don't have a `markerEncodings` entry.
	//
	// - If `false` (default), the default marker encoding is `MarkerEncoding.LITERAL`.
	// - If `true`, the default marker encoding is `MarkerEncoding.UNQUOTED_JSON_VALUE`.
	//
	// See `Source.data()` for a full description of the effect of this setting.
	//
	// Unless you need fine-grained control over marker encoding, you should
	// use `Source.jsonData()` or `Source.yamlData()`.
	// Default: - false.
	//
	JsonEscape *bool `field:"optional" json:"jsonEscape" yaml:"jsonEscape"`
}

