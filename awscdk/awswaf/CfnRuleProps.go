package awswaf


// Properties for defining a `CfnRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRuleProps := &CfnRuleProps{
//   	MetricName: jsii.String("metricName"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Predicates: []interface{}{
//   		&PredicateProperty{
//   			DataId: jsii.String("dataId"),
//   			Negated: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html
//
type CfnRuleProps struct {
	// The name of the metrics for this `Rule` .
	//
	// The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one. It can't contain whitespace or metric names reserved for AWS WAF , including "All" and "Default_Action." You can't change `MetricName` after you create the `Rule` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html#cfn-waf-rule-metricname
	//
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The friendly name or description for the `Rule` .
	//
	// You can't change the name of a `Rule` after you create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html#cfn-waf-rule-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The `Predicates` object contains one `Predicate` element for each `ByteMatchSet` , `IPSet` , or `SqlInjectionMatchSet` object that you want to include in a `Rule` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html#cfn-waf-rule-predicates
	//
	Predicates interface{} `field:"optional" json:"predicates" yaml:"predicates"`
}

