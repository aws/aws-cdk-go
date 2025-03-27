package awswafv2


// Inspect the cookies in the web request.
//
// You can specify the parts of the cookies to inspect and you can narrow the set of cookies to inspect by including or excluding specific keys.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// Example JSON: `"Cookies": { "MatchPattern": { "All": {} }, "MatchScope": "KEY", "OversizeHandling": "MATCH" }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   cookiesProperty := &CookiesProperty{
//   	MatchPattern: &CookieMatchPatternProperty{
//   		All: all,
//   		ExcludedCookies: []*string{
//   			jsii.String("excludedCookies"),
//   		},
//   		IncludedCookies: []*string{
//   			jsii.String("includedCookies"),
//   		},
//   	},
//   	MatchScope: jsii.String("matchScope"),
//   	OversizeHandling: jsii.String("oversizeHandling"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-cookies.html
//
type CfnWebACL_CookiesProperty struct {
	// The filter to use to identify the subset of cookies to inspect in a web request.
	//
	// You must specify exactly one setting: either `All` , `IncludedCookies` , or `ExcludedCookies` .
	//
	// Example JSON: `"MatchPattern": { "IncludedCookies": [ "session-id-time", "session-id" ] }`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-cookies.html#cfn-wafv2-webacl-cookies-matchpattern
	//
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// The parts of the cookies to inspect with the rule inspection criteria.
	//
	// If you specify `ALL` , AWS WAF inspects both keys and values.
	//
	// `All` does not require a match to be found in the keys and a match to be found in the values. It requires a match to be found in the keys or the values or both. To require a match in the keys and in the values, use a logical `AND` statement to combine two match rules, one that inspects the keys and another that inspects the values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-cookies.html#cfn-wafv2-webacl-cookies-matchscope
	//
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// What AWS WAF should do if the cookies of the request are more numerous or larger than AWS WAF can inspect.
	//
	// AWS WAF does not support inspecting the entire contents of request cookies when they exceed 8 KB (8192 bytes) or 200 total cookies. The underlying host service forwards a maximum of 200 cookies and at most 8 KB of cookie contents to AWS WAF .
	//
	// The options for oversize handling are the following:
	//
	// - `CONTINUE` - Inspect the available cookies normally, according to the rule inspection criteria.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-cookies.html#cfn-wafv2-webacl-cookies-oversizehandling
	//
	OversizeHandling *string `field:"required" json:"oversizeHandling" yaml:"oversizeHandling"`
}

