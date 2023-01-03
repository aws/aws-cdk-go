package awselasticloadbalancingv2


// Properties for adding a fixed response to a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var listenerCondition listenerCondition
//
//   addFixedResponseProps := &addFixedResponseProps{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	conditions: []*listenerCondition{
//   		listenerCondition,
//   	},
//   	contentType: awscdk.Aws_elasticloadbalancingv2.contentType_TEXT_PLAIN,
//   	hostHeader: jsii.String("hostHeader"),
//   	messageBody: jsii.String("messageBody"),
//   	pathPattern: jsii.String("pathPattern"),
//   	pathPatterns: []*string{
//   		jsii.String("pathPatterns"),
//   	},
//   	priority: jsii.Number(123),
//   }
//
// Deprecated: Use `ApplicationListener.addAction` instead.
type AddFixedResponseProps struct {
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	// Deprecated: Use `ApplicationListener.addAction` instead.
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
	// Deprecated: Use `ApplicationListener.addAction` instead.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The HTTP response code (2XX, 4XX or 5XX).
	// Deprecated: Use `ApplicationListener.addAction` instead.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The content type.
	// Deprecated: Use `ApplicationListener.addAction` instead.
	ContentType ContentType `field:"optional" json:"contentType" yaml:"contentType"`
	// The message.
	// Deprecated: Use `ApplicationListener.addAction` instead.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
}

