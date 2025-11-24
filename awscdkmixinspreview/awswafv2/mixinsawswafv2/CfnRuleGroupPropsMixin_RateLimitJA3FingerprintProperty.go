package mixinsawswafv2


// Use the request's JA3 fingerprint derived from the TLS Client Hello of an incoming request as an aggregate key.
//
// If you use a single JA3 fingerprint as your custom key, then each value fully defines an aggregation instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rateLimitJA3FingerprintProperty := &RateLimitJA3FingerprintProperty{
//   	FallbackBehavior: jsii.String("fallbackBehavior"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratelimitja3fingerprint.html
//
type CfnRuleGroupPropsMixin_RateLimitJA3FingerprintProperty struct {
	// The match status to assign to the web request if there is insufficient TSL Client Hello information to compute the JA3 fingerprint.
	//
	// You can specify the following fallback behaviors:
	//
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratelimitja3fingerprint.html#cfn-wafv2-rulegroup-ratelimitja3fingerprint-fallbackbehavior
	//
	FallbackBehavior *string `field:"optional" json:"fallbackBehavior" yaml:"fallbackBehavior"`
}

