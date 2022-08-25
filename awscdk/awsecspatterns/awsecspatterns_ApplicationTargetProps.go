package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
)

// Properties to define an application target group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationTargetProps := &applicationTargetProps{
//   	containerPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	hostHeader: jsii.String("hostHeader"),
//   	listener: jsii.String("listener"),
//   	pathPattern: jsii.String("pathPattern"),
//   	priority: jsii.Number(123),
//   	protocol: awscdk.Aws_ecs.protocol_TCP,
//   }
//
// Experimental.
type ApplicationTargetProps struct {
	// The port number of the container.
	//
	// Only applicable when using application/network load balancers.
	// Experimental.
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// Rule applies if the requested host matches the indicated host.
	//
	// May contain up to three '*' wildcards.
	//
	// Requires that priority is set.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#host-conditions
	//
	// Experimental.
	HostHeader *string `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// Name of the listener the target group attached to.
	// Experimental.
	Listener *string `field:"optional" json:"listener" yaml:"listener"`
	// Rule applies if the requested path matches the given path pattern.
	//
	// May contain up to three '*' wildcards.
	//
	// Requires that priority is set.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#path-conditions
	//
	// Experimental.
	PathPattern *string `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The protocol used for the port mapping.
	//
	// Only applicable when using application load balancers.
	// Experimental.
	Protocol awsecs.Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

