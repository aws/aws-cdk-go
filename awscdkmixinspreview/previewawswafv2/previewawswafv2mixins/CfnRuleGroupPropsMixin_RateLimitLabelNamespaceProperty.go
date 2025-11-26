package previewawswafv2mixins


// Specifies a label namespace to use as an aggregate key for a rate-based rule.
//
// Each distinct fully qualified label name that has the specified label namespace contributes to the aggregation instance. If you use just one label namespace as your custom key, then each label name fully defines an aggregation instance.
//
// This uses only labels that have been added to the request by rules that are evaluated before this rate-based rule in the web ACL.
//
// For information about label namespaces and names, see [Label syntax and naming requirements](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-label-requirements.html) in the *AWS WAF Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rateLimitLabelNamespaceProperty := &RateLimitLabelNamespaceProperty{
//   	Namespace: jsii.String("namespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratelimitlabelnamespace.html
//
type CfnRuleGroupPropsMixin_RateLimitLabelNamespaceProperty struct {
	// The namespace to use for aggregation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratelimitlabelnamespace.html#cfn-wafv2-rulegroup-ratelimitlabelnamespace-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

