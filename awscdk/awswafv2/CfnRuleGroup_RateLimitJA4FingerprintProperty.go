package awswafv2


// Use the request's JA4 fingerprint derived from the TLS Client Hello of an incoming request as an aggregate key.
//
// If you use a single JA4 fingerprint as your custom key, then each value fully defines an aggregation instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rateLimitJA4FingerprintProperty := &RateLimitJA4FingerprintProperty{
//   	FallbackBehavior: jsii.String("fallbackBehavior"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratelimitja4fingerprint.html
//
type CfnRuleGroup_RateLimitJA4FingerprintProperty struct {
	// The match status to assign to the web request if there is insufficient TSL Client Hello information to compute the JA4 fingerprint.
	//
	// You can specify the following fallback behaviors:
	//
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratelimitja4fingerprint.html#cfn-wafv2-rulegroup-ratelimitja4fingerprint-fallbackbehavior
	//
	FallbackBehavior *string `field:"required" json:"fallbackBehavior" yaml:"fallbackBehavior"`
}

