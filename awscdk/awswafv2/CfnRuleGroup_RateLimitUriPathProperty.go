package awswafv2


// Specifies the request's URI path as an aggregate key for a rate-based rule.
//
// Each distinct URI path contributes to the aggregation instance. If you use just the URI path as your custom key, then each URI path fully defines an aggregation instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rateLimitUriPathProperty := &RateLimitUriPathProperty{
//   	TextTransformations: []interface{}{
//   		&TextTransformationProperty{
//   			Priority: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratelimituripath.html
//
type CfnRuleGroup_RateLimitUriPathProperty struct {
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// Text transformations are used in rule match statements, to transform the `FieldToMatch` request component before inspecting it, and they're used in rate-based rule statements, to transform request components before using them as custom aggregation keys. If you specify one or more transformations to apply, AWS WAF performs all transformations on the specified content, starting from the lowest priority setting, and then uses the component contents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratelimituripath.html#cfn-wafv2-rulegroup-ratelimituripath-texttransformations
	//
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

