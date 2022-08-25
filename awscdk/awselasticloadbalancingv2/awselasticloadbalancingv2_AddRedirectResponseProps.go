package awselasticloadbalancingv2


// Properties for adding a redirect response to a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var listenerCondition listenerCondition
//
//   addRedirectResponseProps := &addRedirectResponseProps{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	conditions: []*listenerCondition{
//   		listenerCondition,
//   	},
//   	host: jsii.String("host"),
//   	hostHeader: jsii.String("hostHeader"),
//   	path: jsii.String("path"),
//   	pathPattern: jsii.String("pathPattern"),
//   	pathPatterns: []*string{
//   		jsii.String("pathPatterns"),
//   	},
//   	port: jsii.String("port"),
//   	priority: jsii.Number(123),
//   	protocol: jsii.String("protocol"),
//   	query: jsii.String("query"),
//   }
//
// Deprecated: Use `ApplicationListener.addAction` instead.
type AddRedirectResponseProps struct {
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
	// The HTTP redirect code (HTTP_301 or HTTP_302).
	// Deprecated: Use `ApplicationListener.addAction` instead.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The hostname.
	//
	// This component is not percent-encoded. The hostname can contain #{host}.
	// Deprecated: Use `ApplicationListener.addAction` instead.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The absolute path, starting with the leading "/".
	//
	// This component is not percent-encoded.
	// The path can contain #{host}, #{path}, and #{port}.
	// Deprecated: Use `ApplicationListener.addAction` instead.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The port.
	//
	// You can specify a value from 1 to 65535 or #{port}.
	// Deprecated: Use `ApplicationListener.addAction` instead.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocol.
	//
	// You can specify HTTP, HTTPS, or #{protocol}. You can redirect HTTP to HTTP,
	// HTTP to HTTPS, and HTTPS to HTTPS. You cannot redirect HTTPS to HTTP.
	// Deprecated: Use `ApplicationListener.addAction` instead.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The query parameters, URL-encoded when necessary, but not percent-encoded.
	//
	// Do not include the leading "?", as it is automatically added.
	// You can specify any of the reserved keywords.
	// Deprecated: Use `ApplicationListener.addAction` instead.
	Query *string `field:"optional" json:"query" yaml:"query"`
}

