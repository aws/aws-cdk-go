package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Represents timeouts for GRPC protocols.
//
// Example:
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
// Experimental.
type GrpcTimeout struct {
	// Represents an idle timeout.
	//
	// The amount of time that a connection may be idle.
	// Experimental.
	Idle awscdk.Duration `field:"optional" json:"idle" yaml:"idle"`
	// Represents per request timeout.
	// Experimental.
	PerRequest awscdk.Duration `field:"optional" json:"perRequest" yaml:"perRequest"`
}

