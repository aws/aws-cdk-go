package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Properties to define an application target group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationTargetProps := &ApplicationTargetProps{
//   	ContainerPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	HostHeader: jsii.String("hostHeader"),
//   	Listener: jsii.String("listener"),
//   	PathPattern: jsii.String("pathPattern"),
//   	Priority: jsii.Number(123),
//   	Protocol: awscdk.Aws_ecs.Protocol_TCP,
//   }
//
type ApplicationTargetProps struct {
	// The port number of the container.
	//
	// Only applicable when using application/network load balancers.
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// Rule applies if the requested host matches the indicated host.
	//
	// May contain up to three '*' wildcards.
	//
	// Requires that priority is set.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#host-conditions
	//
	// Default: No host condition.
	//
	HostHeader *string `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// Name of the listener the target group attached to.
	// Default: - default listener (first added listener).
	//
	Listener *string `field:"optional" json:"listener" yaml:"listener"`
	// Rule applies if the requested path matches the given path pattern.
	//
	// May contain up to three '*' wildcards.
	//
	// Requires that priority is set.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#path-conditions
	//
	// Default: No path condition.
	//
	PathPattern *string `field:"optional" json:"pathPattern" yaml:"pathPattern"`
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
	// The protocol used for the port mapping.
	//
	// Only applicable when using application load balancers.
	// Default: ecs.Protocol.TCP
	//
	Protocol awsecs.Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

