package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteProps := &CfnRouteProps{
//   	MeshName: jsii.String("meshName"),
//   	Spec: &RouteSpecProperty{
//   		GrpcRoute: &GrpcRouteProperty{
//   			Action: &GrpcRouteActionProperty{
//   				WeightedTargets: []interface{}{
//   					&WeightedTargetProperty{
//   						VirtualNode: jsii.String("virtualNode"),
//   						Weight: jsii.Number(123),
//
//   						// the properties below are optional
//   						Port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Match: &GrpcRouteMatchProperty{
//   				Metadata: []interface{}{
//   					&GrpcRouteMetadataProperty{
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Invert: jsii.Boolean(false),
//   						Match: &GrpcRouteMetadataMatchMethodProperty{
//   							Exact: jsii.String("exact"),
//   							Prefix: jsii.String("prefix"),
//   							Range: &MatchRangeProperty{
//   								End: jsii.Number(123),
//   								Start: jsii.Number(123),
//   							},
//   							Regex: jsii.String("regex"),
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				MethodName: jsii.String("methodName"),
//   				Port: jsii.Number(123),
//   				ServiceName: jsii.String("serviceName"),
//   			},
//
//   			// the properties below are optional
//   			RetryPolicy: &GrpcRetryPolicyProperty{
//   				MaxRetries: jsii.Number(123),
//   				PerRetryTimeout: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				GrpcRetryEvents: []*string{
//   					jsii.String("grpcRetryEvents"),
//   				},
//   				HttpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				TcpRetryEvents: []*string{
//   					jsii.String("tcpRetryEvents"),
//   				},
//   			},
//   			Timeout: &GrpcTimeoutProperty{
//   				Idle: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				PerRequest: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Http2Route: &HttpRouteProperty{
//   			Action: &HttpRouteActionProperty{
//   				WeightedTargets: []interface{}{
//   					&WeightedTargetProperty{
//   						VirtualNode: jsii.String("virtualNode"),
//   						Weight: jsii.Number(123),
//
//   						// the properties below are optional
//   						Port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Match: &HttpRouteMatchProperty{
//   				Headers: []interface{}{
//   					&HttpRouteHeaderProperty{
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Invert: jsii.Boolean(false),
//   						Match: &HeaderMatchMethodProperty{
//   							Exact: jsii.String("exact"),
//   							Prefix: jsii.String("prefix"),
//   							Range: &MatchRangeProperty{
//   								End: jsii.Number(123),
//   								Start: jsii.Number(123),
//   							},
//   							Regex: jsii.String("regex"),
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				Method: jsii.String("method"),
//   				Path: &HttpPathMatchProperty{
//   					Exact: jsii.String("exact"),
//   					Regex: jsii.String("regex"),
//   				},
//   				Port: jsii.Number(123),
//   				Prefix: jsii.String("prefix"),
//   				QueryParameters: []interface{}{
//   					&QueryParameterProperty{
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Match: &HttpQueryParameterMatchProperty{
//   							Exact: jsii.String("exact"),
//   						},
//   					},
//   				},
//   				Scheme: jsii.String("scheme"),
//   			},
//
//   			// the properties below are optional
//   			RetryPolicy: &HttpRetryPolicyProperty{
//   				MaxRetries: jsii.Number(123),
//   				PerRetryTimeout: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				HttpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				TcpRetryEvents: []*string{
//   					jsii.String("tcpRetryEvents"),
//   				},
//   			},
//   			Timeout: &HttpTimeoutProperty{
//   				Idle: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				PerRequest: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   		},
//   		HttpRoute: &HttpRouteProperty{
//   			Action: &HttpRouteActionProperty{
//   				WeightedTargets: []interface{}{
//   					&WeightedTargetProperty{
//   						VirtualNode: jsii.String("virtualNode"),
//   						Weight: jsii.Number(123),
//
//   						// the properties below are optional
//   						Port: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Match: &HttpRouteMatchProperty{
//   				Headers: []interface{}{
//   					&HttpRouteHeaderProperty{
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Invert: jsii.Boolean(false),
//   						Match: &HeaderMatchMethodProperty{
//   							Exact: jsii.String("exact"),
//   							Prefix: jsii.String("prefix"),
//   							Range: &MatchRangeProperty{
//   								End: jsii.Number(123),
//   								Start: jsii.Number(123),
//   							},
//   							Regex: jsii.String("regex"),
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				Method: jsii.String("method"),
//   				Path: &HttpPathMatchProperty{
//   					Exact: jsii.String("exact"),
//   					Regex: jsii.String("regex"),
//   				},
//   				Port: jsii.Number(123),
//   				Prefix: jsii.String("prefix"),
//   				QueryParameters: []interface{}{
//   					&QueryParameterProperty{
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Match: &HttpQueryParameterMatchProperty{
//   							Exact: jsii.String("exact"),
//   						},
//   					},
//   				},
//   				Scheme: jsii.String("scheme"),
//   			},
//
//   			// the properties below are optional
//   			RetryPolicy: &HttpRetryPolicyProperty{
//   				MaxRetries: jsii.Number(123),
//   				PerRetryTimeout: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				HttpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				TcpRetryEvents: []*string{
//   					jsii.String("tcpRetryEvents"),
//   				},
//   			},
//   			Timeout: &HttpTimeoutProperty{
//   				Idle: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				PerRequest: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Priority: jsii.Number(123),
//   		TcpRoute: &TcpRouteProperty{
//   			Action: &TcpRouteActionProperty{
//   				WeightedTargets: []interface{}{
//   					&WeightedTargetProperty{
//   						VirtualNode: jsii.String("virtualNode"),
//   						Weight: jsii.Number(123),
//
//   						// the properties below are optional
//   						Port: jsii.Number(123),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			Match: &TcpRouteMatchProperty{
//   				Port: jsii.Number(123),
//   			},
//   			Timeout: &TcpTimeoutProperty{
//   				Idle: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	VirtualRouterName: jsii.String("virtualRouterName"),
//
//   	// the properties below are optional
//   	MeshOwner: jsii.String("meshOwner"),
//   	RouteName: jsii.String("routeName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html
//
type CfnRouteProps struct {
	// The name of the service mesh to create the route in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html#cfn-appmesh-route-meshname
	//
	MeshName *string `field:"required" json:"meshName" yaml:"meshName"`
	// The route specification to apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html#cfn-appmesh-route-spec
	//
	Spec interface{} `field:"required" json:"spec" yaml:"spec"`
	// The name of the virtual router in which to create the route.
	//
	// If the virtual router is in a shared mesh, then you must be the owner of the virtual router resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html#cfn-appmesh-route-virtualroutername
	//
	VirtualRouterName *string `field:"required" json:"virtualRouterName" yaml:"virtualRouterName"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then the account that you specify must share the mesh with your account before you can create the resource in the service mesh. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html#cfn-appmesh-route-meshowner
	//
	MeshOwner *string `field:"optional" json:"meshOwner" yaml:"meshOwner"`
	// The name to use for the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html#cfn-appmesh-route-routename
	//
	RouteName *string `field:"optional" json:"routeName" yaml:"routeName"`
	// Optional metadata that you can apply to the route to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html#cfn-appmesh-route-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

