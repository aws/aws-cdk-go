package awswafv2


// Inspect the body of the web request as JSON. The body immediately follows the request headers.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// Use the specifications in this object to indicate which parts of the JSON body to inspect using the rule's inspection criteria. AWS WAF inspects only the parts of the JSON that result from the matches that you indicate.
//
// Example JSON: `"JsonBody": { "MatchPattern": { "All": {} }, "MatchScope": "ALL" }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   jsonBodyProperty := &jsonBodyProperty{
//   	matchPattern: &jsonMatchPatternProperty{
//   		all: all,
//   		includedPaths: []*string{
//   			jsii.String("includedPaths"),
//   		},
//   	},
//   	matchScope: jsii.String("matchScope"),
//
//   	// the properties below are optional
//   	invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   	oversizeHandling: jsii.String("oversizeHandling"),
//   }
//
type CfnRuleGroup_JsonBodyProperty struct {
	// The patterns to look for in the JSON body.
	//
	// AWS WAF inspects the results of these pattern matches against the rule inspection criteria.
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// The parts of the JSON to match against using the `MatchPattern` .
	//
	// If you specify `All` , AWS WAF matches against keys and values.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// What AWS WAF should do if it fails to completely parse the JSON body. The options are the following:.
	//
	// - `EVALUATE_AS_STRING` - Inspect the body as plain text. AWS WAF applies the text transformations and inspection criteria that you defined for the JSON inspection to the body text string.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	//
	// If you don't provide this setting, AWS WAF parses and evaluates the content only up to the first parsing failure that it encounters.
	//
	// AWS WAF does its best to parse the entire JSON body, but might be forced to stop for reasons such as invalid characters, duplicate keys, truncation, and any content whose root node isn't an object or an array.
	//
	// AWS WAF parses the JSON in the following examples as two valid key, value pairs:
	//
	// - Missing comma: `{"key1":"value1""key2":"value2"}`
	// - Missing colon: `{"key1":"value1","key2""value2"}`
	// - Extra colons: `{"key1"::"value1","key2""value2"}`.
	InvalidFallbackBehavior *string `field:"optional" json:"invalidFallbackBehavior" yaml:"invalidFallbackBehavior"`
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

