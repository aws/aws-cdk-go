package awsappmesh


// An object that represents the action to take if a match is determined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tcpRouteActionProperty := &TcpRouteActionProperty{
//   	WeightedTargets: []interface{}{
//   		&WeightedTargetProperty{
//   			VirtualNode: jsii.String("virtualNode"),
//   			Weight: jsii.Number(123),
//
//   			// the properties below are optional
//   			Port: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-tcprouteaction.html
//
type CfnRoute_TcpRouteActionProperty struct {
	// An object that represents the targets that traffic is routed to when a request matches the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-tcprouteaction.html#cfn-appmesh-route-tcprouteaction-weightedtargets
	//
	WeightedTargets interface{} `field:"required" json:"weightedTargets" yaml:"weightedTargets"`
}

