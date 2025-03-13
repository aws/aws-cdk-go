package awsappmesh


// Properties specific for a TCP Based Routes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var virtualNode virtualNode
//
//   tcpRouteSpecOptions := &TcpRouteSpecOptions{
//   	WeightedTargets: []weightedTarget{
//   		&weightedTarget{
//   			VirtualNode: virtualNode,
//
//   			// the properties below are optional
//   			Port: jsii.Number(123),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Priority: jsii.Number(123),
//   	Timeout: &TcpTimeout{
//   		Idle: cdk.Duration_Minutes(jsii.Number(30)),
//   	},
//   }
//
type TcpRouteSpecOptions struct {
	// The priority for the route.
	//
	// When a Virtual Router has multiple routes, route match is performed in the
	// order of specified value, where 0 is the highest priority, and first matched route is selected.
	// Default: - no particular priority.
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// List of targets that traffic is routed to when a request matches the route.
	WeightedTargets *[]*WeightedTarget `field:"required" json:"weightedTargets" yaml:"weightedTargets"`
	// An object that represents a tcp timeout.
	// Default: - None.
	//
	Timeout *TcpTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}

