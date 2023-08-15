package awselasticloadbalancingv2


// Properties for adding a new action to a listener.
//
// Example:
//   var listener applicationListener
//
//
//   listener.AddAction(jsii.String("Fixed"), &AddApplicationActionProps{
//   	Priority: jsii.Number(10),
//   	Conditions: []listenerCondition{
//   		elbv2.*listenerCondition_PathPatterns([]*string{
//   			jsii.String("/ok"),
//   		}),
//   	},
//   	Action: elbv2.ListenerAction_FixedResponse(jsii.Number(200), &FixedResponseOptions{
//   		ContentType: jsii.String("text/plain"),
//   		MessageBody: jsii.String("OK"),
//   	}),
//   })
//
type AddApplicationActionProps struct {
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	// Default: - No conditions.
	//
	Conditions *[]ListenerCondition `field:"optional" json:"conditions" yaml:"conditions"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	// Default: Target groups are used as defaults.
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Action to perform.
	Action ListenerAction `field:"required" json:"action" yaml:"action"`
}

