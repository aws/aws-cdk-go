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
//   router.addRoute(jsii.String("route-http"), &routeBaseProps{
//   	routeSpec: appmesh.routeSpec.grpc(&grpcRouteSpecOptions{
//   		weightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				virtualNode: node,
//   			},
//   		},
//   		match: &grpcRouteMatch{
//   			serviceName: jsii.String("my-service.default.svc.cluster.local"),
//   		},
//   		timeout: &grpcTimeout{
//   			idle: cdk.duration.seconds(jsii.Number(2)),
//   			perRequest: cdk.*duration.seconds(jsii.Number(1)),
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

