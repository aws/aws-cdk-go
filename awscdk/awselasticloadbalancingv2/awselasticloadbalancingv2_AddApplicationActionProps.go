package awselasticloadbalancingv2


// Properties for adding a new action to a listener.
//
// Example:
//   var listener applicationListener
//
//
//   listener.addAction(jsii.String("Fixed"), &addApplicationActionProps{
//   	priority: jsii.Number(10),
//   	conditions: []listenerCondition{
//   		elbv2.*listenerCondition.pathPatterns([]*string{
//   			jsii.String("/ok"),
//   		}),
//   	},
//   	action: elbv2.listenerAction.fixedResponse(jsii.Number(200), &fixedResponseOptions{
//   		contentType: elbv2.contentType_TEXT_PLAIN,
//   		messageBody: jsii.String("OK"),
//   	}),
//   })
//
// Experimental.
type AddApplicationActionProps struct {
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	// Experimental.
	Conditions *[]ListenerCondition `field:"optional" json:"conditions" yaml:"conditions"`
	// Rule applies if the requested host matches the indicated host.
	//
	// May contain up to three '*' wildcards.
	//
	// Requires that priority is set.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#host-conditions
	//
	// Deprecated: Use `conditions` instead.
	HostHeader *string `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// Rule applies if the requested path matches the given path pattern.
	//
	// May contain up to three '*' wildcards.
	//
	// Requires that priority is set.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#path-conditions
	//
	// Deprecated: Use `conditions` instead.
	PathPattern *string `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// Rule applies if the requested path matches any of the given patterns.
	//
	// May contain up to three '*' wildcards.
	//
	// Requires that priority is set.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#path-conditions
	//
	// Deprecated: Use `conditions` instead.
	PathPatterns *[]*string `field:"optional" json:"pathPatterns" yaml:"pathPatterns"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Action to perform.
	// Experimental.
	Action ListenerAction `field:"required" json:"action" yaml:"action"`
}

