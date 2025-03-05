package awswafv2


// Specifies the request's JA4 fingerprint as an aggregate key for a rate-based rule.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratelimitja4fingerprint.html#cfn-wafv2-rulegroup-ratelimitja4fingerprint-fallbackbehavior
	//
	FallbackBehavior *string `field:"required" json:"fallbackBehavior" yaml:"fallbackBehavior"`
}

