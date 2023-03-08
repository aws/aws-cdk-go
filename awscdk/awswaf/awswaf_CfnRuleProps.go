package awswaf


// Properties for defining a `CfnRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRuleProps := &cfnRuleProps{
//   	metricName: jsii.String("metricName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	predicates: []interface{}{
//   		&predicateProperty{
//   			dataId: jsii.String("dataId"),
//   			negated: jsii.Boolean(false),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnRuleProps struct {
	// The name of the metrics for this `Rule` .
	//
	// The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one. It can't contain whitespace or metric names reserved for AWS WAF , including "All" and "Default_Action." You can't change `MetricName` after you create the `Rule` .
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The friendly name or description for the `Rule` .
	//
	// You can't change the name of a `Rule` after you create it.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The `Predicates` object contains one `Predicate` element for each `ByteMatchSet` , `IPSet` , or `SqlInjectionMatchSet` object that you want to include in a `Rule` .
	Predicates interface{} `field:"optional" json:"predicates" yaml:"predicates"`
}

