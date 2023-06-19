package awselasticloadbalancingv2


// Properties for adding a new target group to a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationTargetGroup applicationTargetGroup
//   var listenerCondition listenerCondition
//
//   addApplicationTargetGroupsProps := &AddApplicationTargetGroupsProps{
//   	TargetGroups: []iApplicationTargetGroup{
//   		applicationTargetGroup,
//   	},
//
//   	// the properties below are optional
//   	Conditions: []*listenerCondition{
//   		listenerCondition,
//   	},
//   	HostHeader: jsii.String("hostHeader"),
//   	PathPattern: jsii.String("pathPattern"),
//   	PathPatterns: []*string{
//   		jsii.String("pathPatterns"),
//   	},
//   	Priority: jsii.Number(123),
//   }
//
// Experimental.
type AddApplicationTargetGroupsProps struct {
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
	// Target groups to forward requests to.
	// Experimental.
	TargetGroups *[]IApplicationTargetGroup `field:"required" json:"targetGroups" yaml:"targetGroups"`
}

