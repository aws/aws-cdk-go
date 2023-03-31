package awswafregional


// Properties for defining a `CfnRateBasedRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRateBasedRuleProps := &cfnRateBasedRuleProps{
//   	metricName: jsii.String("metricName"),
//   	name: jsii.String("name"),
//   	rateKey: jsii.String("rateKey"),
//   	rateLimit: jsii.Number(123),
//
//   	// the properties below are optional
//   	matchPredicates: []interface{}{
//   		&predicateProperty{
//   			dataId: jsii.String("dataId"),
//   			negated: jsii.Boolean(false),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnRateBasedRuleProps struct {
	// A name for the metrics for a `RateBasedRule` .
	//
	// The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one. It can't contain whitespace or metric names reserved for AWS WAF , including "All" and "Default_Action." You can't change the name of the metric after you create the `RateBasedRule` .
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// A friendly name or description for a `RateBasedRule` .
	//
	// You can't change the name of a `RateBasedRule` after you create it.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The field that AWS WAF uses to determine if requests are likely arriving from single source and thus subject to rate monitoring.
	//
	// The only valid value for `RateKey` is `IP` . `IP` indicates that requests arriving from the same IP address are subject to the `RateLimit` that is specified in the `RateBasedRule` .
	RateKey *string `field:"required" json:"rateKey" yaml:"rateKey"`
	// The maximum number of requests, which have an identical value in the field specified by the `RateKey` , allowed in a five-minute period.
	//
	// If the number of requests exceeds the `RateLimit` and the other predicates specified in the rule are also met, AWS WAF triggers the action that is specified for this rule.
	RateLimit *float64 `field:"required" json:"rateLimit" yaml:"rateLimit"`
	// The `Predicates` object contains one `Predicate` element for each `ByteMatchSet` , `IPSet` , or `SqlInjectionMatchSet>` object that you want to include in a `RateBasedRule` .
	MatchPredicates interface{} `field:"optional" json:"matchPredicates" yaml:"matchPredicates"`
}

