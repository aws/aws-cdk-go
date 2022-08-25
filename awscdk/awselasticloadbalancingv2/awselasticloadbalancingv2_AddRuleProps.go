package awselasticloadbalancingv2


// Properties for adding a conditional load balancing rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var listenerCondition listenerCondition
//
//   addRuleProps := &addRuleProps{
//   	conditions: []*listenerCondition{
//   		listenerCondition,
//   	},
//   	priority: jsii.Number(123),
//   }
//
type AddRuleProps struct {
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	Conditions *[]ListenerCondition `field:"optional" json:"conditions" yaml:"conditions"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

