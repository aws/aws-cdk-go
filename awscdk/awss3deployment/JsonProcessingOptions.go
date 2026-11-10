package awss3deployment


// Define options which can be passed using the method `Source.jsonData()`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonProcessingOptions := &JsonProcessingOptions{
//   	Escape: jsii.Boolean(false),
//   }
//
type JsonProcessingOptions struct {
	// Whether deploy-time values should be treated as values and escaped.
	//
	// The default behavior (`true`) treats deploy-time values as values and
	// serializes them accordingly.
	//
	// The only reason to set this to `false` would be if all Token values are
	// full JSON strings that you want to include into your JSON content as-is. You
	// must use `Token.asNumber()` on those to serialize them into an unquoted
	// context. This should be an exceedingly rare use case and hence it
	// is not the default.
	// Default: - true.
	//
	Escape *bool `field:"optional" json:"escape" yaml:"escape"`
}

