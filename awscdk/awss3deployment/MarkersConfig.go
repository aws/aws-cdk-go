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
//   	MarkerEncodings: map[string]MarkerEncoding{
//   		"markerEncodingsKey": awscdk.aws_s3_deployment.MarkerEncoding_LITERAL,
//   	},
//   }
//
type MarkersConfig struct {
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
	// For every marker specified in the `markers` property, this defines how the marker value should be encoded.
	//
	// The key is the marker name and the value is the encoding to use for that marker.
	// Default: - markers with no entry use the default encoding derived from jsonEscape.
	//
	MarkerEncodings *map[string]MarkerEncoding `field:"optional" json:"markerEncodings" yaml:"markerEncodings"`
}

