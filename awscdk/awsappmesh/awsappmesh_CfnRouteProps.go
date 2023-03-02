package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteProps := &cfnRouteProps{
//   	meshName: jsii.String("meshName"),
//   	spec: &routeSpecProperty{
//   		grpcRoute: &grpcRouteProperty{
//   			action: &grpcRouteActionProperty{
//   				weightedTargets: []interface{}{
//   					&weightedTargetProperty{
//   						virtualNode: jsii.String("virtualNode"),
//   						weight: jsii.Number(123),
//
//   						// the properties below are optional
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			match: &grpcRouteMatchProperty{
//   				metadata: []interface{}{
//   					&grpcRouteMetadataProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						invert: jsii.Boolean(false),
//   						match: &grpcRouteMetadataMatchMethodProperty{
//   							exact: jsii.String("exact"),
//   							prefix: jsii.String("prefix"),
//   							range: &matchRangeProperty{
//   								end: jsii.Number(123),
//   								start: jsii.Number(123),
//   							},
//   							regex: jsii.String("regex"),
//   							suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				methodName: jsii.String("methodName"),
//   				port: jsii.Number(123),
//   				serviceName: jsii.String("serviceName"),
//   			},
//
//   			// the properties below are optional
//   			retryPolicy: &grpcRetryPolicyProperty{
//   				maxRetries: jsii.Number(123),
//   				perRetryTimeout: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				grpcRetryEvents: []*string{
//   					jsii.String("grpcRetryEvents"),
//   				},
//   				httpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				tcpRetryEvents: []*string{
//   					jsii.String("tcpRetryEvents"),
//   				},
//   			},
//   			timeout: &grpcTimeoutProperty{
//   				idle: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   				perRequest: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   			},
//   		},
//   		http2Route: &httpRouteProperty{
//   			action: &httpRouteActionProperty{
//   				weightedTargets: []interface{}{
//   					&weightedTargetProperty{
//   						virtualNode: jsii.String("virtualNode"),
//   						weight: jsii.Number(123),
//
//   						// the properties below are optional
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			match: &httpRouteMatchProperty{
//   				headers: []interface{}{
//   					&httpRouteHeaderProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						invert: jsii.Boolean(false),
//   						match: &headerMatchMethodProperty{
//   							exact: jsii.String("exact"),
//   							prefix: jsii.String("prefix"),
//   							range: &matchRangeProperty{
//   								end: jsii.Number(123),
//   								start: jsii.Number(123),
//   							},
//   							regex: jsii.String("regex"),
//   							suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				method: jsii.String("method"),
//   				path: &httpPathMatchProperty{
//   					exact: jsii.String("exact"),
//   					regex: jsii.String("regex"),
//   				},
//   				port: jsii.Number(123),
//   				prefix: jsii.String("prefix"),
//   				queryParameters: []interface{}{
//   					&queryParameterProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						match: &httpQueryParameterMatchProperty{
//   							exact: jsii.String("exact"),
//   						},
//   					},
//   				},
//   				scheme: jsii.String("scheme"),
//   			},
//
//   			// the properties below are optional
//   			retryPolicy: &httpRetryPolicyProperty{
//   				maxRetries: jsii.Number(123),
//   				perRetryTimeout: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				httpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				tcpRetryEvents: []*string{
//   					jsii.String("tcpRetryEvents"),
//   				},
//   			},
//   			timeout: &httpTimeoutProperty{
//   				idle: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   				perRequest: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   			},
//   		},
//   		httpRoute: &httpRouteProperty{
//   			action: &httpRouteActionProperty{
//   				weightedTargets: []interface{}{
//   					&weightedTargetProperty{
//   						virtualNode: jsii.String("virtualNode"),
//   						weight: jsii.Number(123),
//
//   						// the properties below are optional
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			match: &httpRouteMatchProperty{
//   				headers: []interface{}{
//   					&httpRouteHeaderProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						invert: jsii.Boolean(false),
//   						match: &headerMatchMethodProperty{
//   							exact: jsii.String("exact"),
//   							prefix: jsii.String("prefix"),
//   							range: &matchRangeProperty{
//   								end: jsii.Number(123),
//   								start: jsii.Number(123),
//   							},
//   							regex: jsii.String("regex"),
//   							suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				method: jsii.String("method"),
//   				path: &httpPathMatchProperty{
//   					exact: jsii.String("exact"),
//   					regex: jsii.String("regex"),
//   				},
//   				port: jsii.Number(123),
//   				prefix: jsii.String("prefix"),
//   				queryParameters: []interface{}{
//   					&queryParameterProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						match: &httpQueryParameterMatchProperty{
//   							exact: jsii.String("exact"),
//   						},
//   					},
//   				},
//   				scheme: jsii.String("scheme"),
//   			},
//
//   			// the properties below are optional
//   			retryPolicy: &httpRetryPolicyProperty{
//   				maxRetries: jsii.Number(123),
//   				perRetryTimeout: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				httpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				tcpRetryEvents: []*string{
//   					jsii.String("tcpRetryEvents"),
//   				},
//   			},
//   			timeout: &httpTimeoutProperty{
//   				idle: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   				perRequest: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   			},
//   		},
//   		priority: jsii.Number(123),
//   		tcpRoute: &tcpRouteProperty{
//   			action: &tcpRouteActionProperty{
//   				weightedTargets: []interface{}{
//   					&weightedTargetProperty{
//   						virtualNode: jsii.String("virtualNode"),
//   						weight: jsii.Number(123),
//
//   						// the properties below are optional
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			match: &tcpRouteMatchProperty{
//   				port: jsii.Number(123),
//   			},
//   			timeout: &tcpTimeoutProperty{
//   				idle: &durationProperty{
//   					unit: jsii.String("unit"),
//   					value: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	virtualRouterName: jsii.String("virtualRouterName"),
//
//   	// the properties below are optional
//   	meshOwner: jsii.String("meshOwner"),
//   	routeName: jsii.String("routeName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRouteProps struct {
	// The name of the service mesh to create the route in.
	MeshName *string `field:"required" json:"meshName" yaml:"meshName"`
	// The route specification to apply.
	Spec interface{} `field:"required" json:"spec" yaml:"spec"`
	// The name of the virtual router in which to create the route.
	//
	// If the virtual router is in a shared mesh, then you must be the owner of the virtual router resource.
	VirtualRouterName *string `field:"required" json:"virtualRouterName" yaml:"virtualRouterName"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then the account that you specify must share the mesh with your account before you can create the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	MeshOwner *string `field:"optional" json:"meshOwner" yaml:"meshOwner"`
	// The name to use for the route.
	RouteName *string `field:"optional" json:"routeName" yaml:"routeName"`
	// Optional metadata that you can apply to the route to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

