package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Represents timeouts for GRPC protocols.
//
// Example:
//   var router VirtualRouter
//   var node VirtualNode
//
//
//   router.addRoute(jsii.String("route-http"), &RouteBaseProps{
//   	RouteSpec: appmesh.RouteSpec_Grpc(&GrpcRouteSpecOptions{
//   		WeightedTargets: []WeightedTarget{
//   			&WeightedTarget{
//   				VirtualNode: node,
//   			},
//   		},
//   		Match: &GrpcRouteMatch{
//   			ServiceName: jsii.String("my-service.default.svc.cluster.local"),
//   		},
//   		Timeout: &GrpcTimeout{
//   			Idle: awscdk.Duration_Seconds(jsii.Number(2)),
//   			PerRequest: awscdk.Duration_*Seconds(jsii.Number(1)),
//   		},
//   	}),
//   })
//
type GrpcTimeout struct {
	// Represents an idle timeout.
	//
	// The amount of time that a connection may be idle.
	// Default: - none.
	//
	Idle awscdk.Duration `field:"optional" json:"idle" yaml:"idle"`
	// Represents per request timeout.
	// Default: - 15 s.
	//
	PerRequest awscdk.Duration `field:"optional" json:"perRequest" yaml:"perRequest"`
}

