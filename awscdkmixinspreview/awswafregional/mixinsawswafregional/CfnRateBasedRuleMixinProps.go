package mixinsawswafregional


// Properties for CfnRateBasedRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRateBasedRuleMixinProps := &CfnRateBasedRuleMixinProps{
//   	MatchPredicates: []interface{}{
//   		&PredicateProperty{
//   			DataId: jsii.String("dataId"),
//   			Negated: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	MetricName: jsii.String("metricName"),
//   	Name: jsii.String("name"),
//   	RateKey: jsii.String("rateKey"),
//   	RateLimit: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html
//
type CfnRateBasedRuleMixinProps struct {
	// The `Predicates` object contains one `Predicate` element for each `ByteMatchSet` , `IPSet` , or `SqlInjectionMatchSet>` object that you want to include in a `RateBasedRule` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-matchpredicates
	//
	MatchPredicates interface{} `field:"optional" json:"matchPredicates" yaml:"matchPredicates"`
	// A name for the metrics for a `RateBasedRule` .
	//
	// The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one. It can't contain whitespace or metric names reserved for AWS WAF , including "All" and "Default_Action." You can't change the name of the metric after you create the `RateBasedRule` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// A friendly name or description for a `RateBasedRule` .
	//
	// You can't change the name of a `RateBasedRule` after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The field that AWS WAF uses to determine if requests are likely arriving from single source and thus subject to rate monitoring.
	//
	// The only valid value for `RateKey` is `IP` . `IP` indicates that requests arriving from the same IP address are subject to the `RateLimit` that is specified in the `RateBasedRule` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-ratekey
	//
	RateKey *string `field:"optional" json:"rateKey" yaml:"rateKey"`
	// The maximum number of requests, which have an identical value in the field specified by the `RateKey` , allowed in a five-minute period.
	//
	// If the number of requests exceeds the `RateLimit` and the other predicates specified in the rule are also met, AWS WAF triggers the action that is specified for this rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-ratelimit
	//
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

