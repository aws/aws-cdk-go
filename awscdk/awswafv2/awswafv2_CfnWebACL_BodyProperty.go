package awswafv2


// Inspect the body of the web request. The body immediately follows the request headers.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bodyProperty := &bodyProperty{
//   	oversizeHandling: jsii.String("oversizeHandling"),
//   }
//
type CfnWebACL_BodyProperty struct {
	// What AWS WAF should do if the body is larger than AWS WAF can inspect.
	//
	// AWS WAF does not support inspecting the entire contents of the body of a web request when the body exceeds 8 KB (8192 bytes). Only the first 8 KB of the request body are forwarded to AWS WAF by the underlying host service.
	//
	// The options for oversize handling are the following:
	//
	// - `CONTINUE` - Inspect the body normally, according to the rule inspection criteria.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	//
	// You can combine the `MATCH` or `NO_MATCH` settings for oversize handling with your rule and web ACL action settings, so that you block any request whose body is over 8 KB.
	//
	// Default: `CONTINUE`.
	OversizeHandling *string `field:"optional" json:"oversizeHandling" yaml:"oversizeHandling"`
}

