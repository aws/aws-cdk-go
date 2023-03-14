package awswafv2


// Inspect all headers in the web request.
//
// You can specify the parts of the headers to inspect and you can narrow the set of headers to inspect by including or excluding specific keys.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// If you want to inspect just the value of a single header, use the `SingleHeader` `FieldToMatch` setting instead.
//
// Example JSON: `"Headers": { "MatchPattern": { "All": {} }, "MatchScope": "KEY", "OversizeHandling": "MATCH" }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   headersProperty := &headersProperty{
//   	matchPattern: &headerMatchPatternProperty{
//   		all: all,
//   		excludedHeaders: []*string{
//   			jsii.String("excludedHeaders"),
//   		},
//   		includedHeaders: []*string{
//   			jsii.String("includedHeaders"),
//   		},
//   	},
//   	matchScope: jsii.String("matchScope"),
//   	oversizeHandling: jsii.String("oversizeHandling"),
//   }
//
type CfnWebACL_HeadersProperty struct {
	// The filter to use to identify the subset of headers to inspect in a web request.
	//
	// You must specify exactly one setting: either `All` , `IncludedHeaders` , or `ExcludedHeaders` .
	//
	// Example JSON: `"MatchPattern": { "ExcludedHeaders": {"KeyToExclude1", "KeyToExclude2"} }`.
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// The parts of the headers to match with the rule inspection criteria.
	//
	// If you specify `All` , AWS WAF inspects both keys and values.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// What AWS WAF should do if the headers of the request are larger than AWS WAF can inspect.
	//
	// AWS WAF does not support inspecting the entire contents of request headers when they exceed 8 KB (8192 bytes) or 200 total headers. The underlying host service forwards a maximum of 200 headers and at most 8 KB of header contents to AWS WAF .
	//
	// The options for oversize handling are the following:
	//
	// - `CONTINUE` - Inspect the headers normally, according to the rule inspection criteria.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	OversizeHandling *string `field:"required" json:"oversizeHandling" yaml:"oversizeHandling"`
}

