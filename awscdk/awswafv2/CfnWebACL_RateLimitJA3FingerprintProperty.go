package awswafv2


// Specifies the request's JA3 fingerprint as an aggregate key for a rate-based rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rateLimitJA3FingerprintProperty := &RateLimitJA3FingerprintProperty{
//   	FallbackBehavior: jsii.String("fallbackBehavior"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratelimitja3fingerprint.html
//
type CfnWebACL_RateLimitJA3FingerprintProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratelimitja3fingerprint.html#cfn-wafv2-webacl-ratelimitja3fingerprint-fallbackbehavior
	//
	FallbackBehavior *string `field:"required" json:"fallbackBehavior" yaml:"fallbackBehavior"`
}

