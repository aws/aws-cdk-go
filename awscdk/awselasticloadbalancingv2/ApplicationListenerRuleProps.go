package awselasticloadbalancingv2


// Properties for defining a listener rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationListener applicationListener
//   var applicationTargetGroup applicationTargetGroup
//   var listenerAction listenerAction
//   var listenerCondition listenerCondition
//
//   applicationListenerRuleProps := &ApplicationListenerRuleProps{
//   	Listener: applicationListener,
//   	Priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	Action: listenerAction,
//   	Conditions: []*listenerCondition{
//   		listenerCondition,
//   	},
//   	TargetGroups: []iApplicationTargetGroup{
//   		applicationTargetGroup,
//   	},
//   }
//
type ApplicationListenerRuleProps struct {
	// Priority of the rule.
	//
	// The rule with the lowest priority will be used for every request.
	//
	// Priorities must be unique.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Action to perform when requests are received.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	// Default: - No action.
	//
	Action ListenerAction `field:"optional" json:"action" yaml:"action"`
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	// Default: - No conditions.
	//
	Conditions *[]ListenerCondition `field:"optional" json:"conditions" yaml:"conditions"`
	// Target groups to forward requests to.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	//
	// Implies a `forward` action.
	// Default: - No target groups.
	//
	TargetGroups *[]IApplicationTargetGroup `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// The listener to attach the rule to.
	Listener IApplicationListener `field:"required" json:"listener" yaml:"listener"`
}

