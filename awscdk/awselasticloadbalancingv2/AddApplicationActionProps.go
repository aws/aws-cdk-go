package awselasticloadbalancingv2


// Properties for adding a new action to a listener.
//
// Example:
//   var listener ApplicationListener
//
//
//   listener.AddAction(jsii.String("Fixed"), &AddApplicationActionProps{
//   	Priority: jsii.Number(10),
//   	Conditions: []ListenerCondition{
//   		elbv2.ListenerCondition_PathPatterns([]*string{
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
	// `ListenerRule`s have a `Rule` suffix on their logicalId by default. This allows you to remove that suffix.
	//
	// Legacy behavior of the `addTargetGroups()` convenience method did not include the `Rule` suffix on the logicalId of the generated `ListenerRule`.
	// At some point, increasing complexity of requirements can require users to switch from the `addTargetGroups()` method
	// to the `addAction()` method.
	// When migrating `ListenerRule`s deployed by a legacy version of `addTargetGroups()`,
	// you will need to enable this flag to avoid changing the logicalId of your resource.
	// Otherwise Cfn will attempt to replace the `ListenerRule` and fail.
	// Default: - use standard logicalId with the `Rule` suffix.
	//
	RemoveSuffix *bool `field:"optional" json:"removeSuffix" yaml:"removeSuffix"`
}

