package mixinsawsappmesh


// An object that represents the action to take if a match is determined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   grpcRouteActionProperty := &GrpcRouteActionProperty{
//   	WeightedTargets: []interface{}{
//   		&WeightedTargetProperty{
//   			Port: jsii.Number(123),
//   			VirtualNode: jsii.String("virtualNode"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcrouteaction.html
//
type CfnRoutePropsMixin_GrpcRouteActionProperty struct {
	// An object that represents the targets that traffic is routed to when a request matches the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-grpcrouteaction.html#cfn-appmesh-route-grpcrouteaction-weightedtargets
	//
	WeightedTargets interface{} `field:"optional" json:"weightedTargets" yaml:"weightedTargets"`
}

