package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Represents timeouts for GRPC protocols.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http"), &RouteBaseProps{
//   	RouteSpec: appmesh.RouteSpec_Grpc(&GrpcRouteSpecOptions{
//   		WeightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				VirtualNode: node,
//   			},
//   		},
//   		Match: &GrpcRouteMatch{
//   			ServiceName: jsii.String("my-service.default.svc.cluster.local"),
//   		},
//   		Timeout: &GrpcTimeout{
//   			Idle: cdk.Duration_Seconds(jsii.Number(2)),
//   			PerRequest: cdk.Duration_*Seconds(jsii.Number(1)),
//   		},
//   	}),
//   })
//
type GrpcTimeout struct {
	// Represents an idle timeout.
	//
	// The amount of time that a connection may be idle.
	Idle awscdk.Duration `field:"optional" json:"idle" yaml:"idle"`
	// Represents per request timeout.
	PerRequest awscdk.Duration `field:"optional" json:"perRequest" yaml:"perRequest"`
}

