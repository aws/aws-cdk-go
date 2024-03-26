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
//   jsonBodyProperty := &JsonBodyProperty{
//   	MatchPattern: &JsonMatchPatternProperty{
//   		All: all,
//   		IncludedPaths: []*string{
//   			jsii.String("includedPaths"),
//   		},
//   	},
//   	MatchScope: jsii.String("matchScope"),
//
//   	// the properties below are optional
//   	InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   	OversizeHandling: jsii.String("oversizeHandling"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-jsonbody.html
//
type CfnRuleGroup_JsonBodyProperty struct {
	// The patterns to look for in the JSON body.
	//
	// AWS WAF inspects the results of these pattern matches against the rule inspection criteria.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-jsonbody.html#cfn-wafv2-rulegroup-jsonbody-matchpattern
	//
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// The parts of the JSON to match against using the `MatchPattern` .
	//
	// If you specify `ALL` , AWS WAF matches against keys and values.
	//
	// `All` does not require a match to be found in the keys and a match to be found in the values. It requires a match to be found in the keys or the values or both. To require a match in the keys and in the values, use a logical `AND` statement to combine two match rules, one that inspects the keys and another that inspects the values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-jsonbody.html#cfn-wafv2-rulegroup-jsonbody-matchscope
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-jsonbody.html#cfn-wafv2-rulegroup-jsonbody-invalidfallbackbehavior
	//
	InvalidFallbackBehavior *string `field:"optional" json:"invalidFallbackBehavior" yaml:"invalidFallbackBehavior"`
	// What AWS WAF should do if the body is larger than AWS WAF can inspect.
	//
	// AWS WAF does not support inspecting the entire contents of the web request body if the body exceeds the limit for the resource type. When a web request body is larger than the limit, the underlying host service only forwards the contents that are within the limit to AWS WAF for inspection.
	//
	// - For Application Load Balancer and AWS AppSync , the limit is fixed at 8 KB (8,192 bytes).
	// - For CloudFront, API Gateway, Amazon Cognito, App Runner, and Verified Access, the default limit is 16 KB (16,384 bytes), and you can increase the limit for each resource type in the web ACL `AssociationConfig` , for additional processing fees.
	//
	// The options for oversize handling are the following:
	//
	// - `CONTINUE` - Inspect the available body contents normally, according to the rule inspection criteria.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	//
	// You can combine the `MATCH` or `NO_MATCH` settings for oversize handling with your rule and web ACL action settings, so that you block any request whose body is over the limit.
	//
	// Default: `CONTINUE`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-jsonbody.html#cfn-wafv2-rulegroup-jsonbody-oversizehandling
	//
	OversizeHandling *string `field:"optional" json:"oversizeHandling" yaml:"oversizeHandling"`
}

