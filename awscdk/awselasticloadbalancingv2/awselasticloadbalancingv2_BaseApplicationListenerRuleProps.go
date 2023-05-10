package awselasticloadbalancingv2


// Basic properties for defining a rule on a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationTargetGroup applicationTargetGroup
//   var listenerAction listenerAction
//   var listenerCondition listenerCondition
//
//   baseApplicationListenerRuleProps := &BaseApplicationListenerRuleProps{
//   	Priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	Action: listenerAction,
//   	Conditions: []*listenerCondition{
//   		listenerCondition,
//   	},
//   	FixedResponse: &FixedResponse{
//   		StatusCode: jsii.String("statusCode"),
//
//   		// the properties below are optional
//   		ContentType: awscdk.Aws_elasticloadbalancingv2.ContentType_TEXT_PLAIN,
//   		MessageBody: jsii.String("messageBody"),
//   	},
//   	HostHeader: jsii.String("hostHeader"),
//   	PathPattern: jsii.String("pathPattern"),
//   	PathPatterns: []*string{
//   		jsii.String("pathPatterns"),
//   	},
//   	RedirectResponse: &RedirectResponse{
//   		StatusCode: jsii.String("statusCode"),
//
//   		// the properties below are optional
//   		Host: jsii.String("host"),
//   		Path: jsii.String("path"),
//   		Port: jsii.String("port"),
//   		Protocol: jsii.String("protocol"),
//   		Query: jsii.String("query"),
//   	},
//   	TargetGroups: []iApplicationTargetGroup{
//   		applicationTargetGroup,
//   	},
//   }
//
// Experimental.
type BaseApplicationListenerRuleProps struct {
	// Priority of the rule.
	//
	// The rule with the lowest priority will be used for every request.
	//
	// Priorities must be unique.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Action to perform when requests are received.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	// Experimental.
	Action ListenerAction `field:"optional" json:"action" yaml:"action"`
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	// Experimental.
	Conditions *[]ListenerCondition `field:"optional" json:"conditions" yaml:"conditions"`
	// Fixed response to return.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	// Deprecated: Use `action` instead.
	FixedResponse *FixedResponse `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// Rule applies if the requested host matches the indicated host.
	//
	// May contain up to three '*' wildcards.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#host-conditions
	//
	// Deprecated: Use `conditions` instead.
	HostHeader *string `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// Rule applies if the requested path matches the given path pattern.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#path-conditions
	//
	// Deprecated: Use `conditions` instead.
	PathPattern *string `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// Rule applies if the requested path matches any of the given patterns.
	//
	// Paths may contain up to three '*' wildcards.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#path-conditions
	//
	// Deprecated: Use `conditions` instead.
	PathPatterns *[]*string `field:"optional" json:"pathPatterns" yaml:"pathPatterns"`
	// Redirect response to return.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	// Deprecated: Use `action` instead.
	RedirectResponse *RedirectResponse `field:"optional" json:"redirectResponse" yaml:"redirectResponse"`
	// Target groups to forward requests to.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	//
	// Implies a `forward` action.
	// Experimental.
	TargetGroups *[]IApplicationTargetGroup `field:"optional" json:"targetGroups" yaml:"targetGroups"`
}

