package awsappmesh


// An object that represents the action to take if a match is determined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tcpRouteActionProperty := &tcpRouteActionProperty{
//   	weightedTargets: []interface{}{
//   		&weightedTargetProperty{
//   			virtualNode: jsii.String("virtualNode"),
//   			weight: jsii.Number(123),
//
//   			// the properties below are optional
//   			port: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnRoute_TcpRouteActionProperty struct {
	// An object that represents the targets that traffic is routed to when a request matches the route.
	WeightedTargets interface{} `field:"required" json:"weightedTargets" yaml:"weightedTargets"`
}

