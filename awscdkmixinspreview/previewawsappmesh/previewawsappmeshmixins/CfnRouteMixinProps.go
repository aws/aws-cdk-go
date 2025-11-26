package previewawsappmeshmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRoutePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRouteMixinProps := &CfnRouteMixinProps{
//   	MeshName: jsii.String("meshName"),
//   	MeshOwner: jsii.String("meshOwner"),
//   	RouteName: jsii.String("routeName"),
//   	Spec: &RouteSpecProperty{
//   		GrpcRoute: &GrpcRouteProperty{
//   			Action: &GrpcRouteActionProperty{
//   				WeightedTargets: []interface{}{
//   					&WeightedTargetProperty{
//   						Port: jsii.Number(123),
//   						VirtualNode: jsii.String("virtualNode"),
//   						Weight: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Match: &GrpcRouteMatchProperty{
//   				Metadata: []interface{}{
//   					&GrpcRouteMetadataProperty{
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
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				MethodName: jsii.String("methodName"),
//   				Port: jsii.Number(123),
//   				ServiceName: jsii.String("serviceName"),
//   			},
//   			RetryPolicy: &GrpcRetryPolicyProperty{
//   				GrpcRetryEvents: []*string{
//   					jsii.String("grpcRetryEvents"),
//   				},
//   				HttpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				MaxRetries: jsii.Number(123),
//   				PerRetryTimeout: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
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
//   						Port: jsii.Number(123),
//   						VirtualNode: jsii.String("virtualNode"),
//   						Weight: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Match: &HttpRouteMatchProperty{
//   				Headers: []interface{}{
//   					&HttpRouteHeaderProperty{
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
//   						Name: jsii.String("name"),
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
//   						Match: &HttpQueryParameterMatchProperty{
//   							Exact: jsii.String("exact"),
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Scheme: jsii.String("scheme"),
//   			},
//   			RetryPolicy: &HttpRetryPolicyProperty{
//   				HttpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				MaxRetries: jsii.Number(123),
//   				PerRetryTimeout: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
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
//   						Port: jsii.Number(123),
//   						VirtualNode: jsii.String("virtualNode"),
//   						Weight: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Match: &HttpRouteMatchProperty{
//   				Headers: []interface{}{
//   					&HttpRouteHeaderProperty{
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
//   						Name: jsii.String("name"),
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
//   						Match: &HttpQueryParameterMatchProperty{
//   							Exact: jsii.String("exact"),
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Scheme: jsii.String("scheme"),
//   			},
//   			RetryPolicy: &HttpRetryPolicyProperty{
//   				HttpRetryEvents: []*string{
//   					jsii.String("httpRetryEvents"),
//   				},
//   				MaxRetries: jsii.Number(123),
//   				PerRetryTimeout: &DurationProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
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
//   						Port: jsii.Number(123),
//   						VirtualNode: jsii.String("virtualNode"),
//   						Weight: jsii.Number(123),
//   					},
//   				},
//   			},
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VirtualRouterName: jsii.String("virtualRouterName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html
//
type CfnRouteMixinProps struct {
	// The name of the service mesh to create the route in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html#cfn-appmesh-route-meshname
	//
	MeshName *string `field:"optional" json:"meshName" yaml:"meshName"`
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
	// The route specification to apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html#cfn-appmesh-route-spec
	//
	Spec interface{} `field:"optional" json:"spec" yaml:"spec"`
	// Optional metadata that you can apply to the route to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html#cfn-appmesh-route-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the virtual router in which to create the route.
	//
	// If the virtual router is in a shared mesh, then you must be the owner of the virtual router resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html#cfn-appmesh-route-virtualroutername
	//
	VirtualRouterName *string `field:"optional" json:"virtualRouterName" yaml:"virtualRouterName"`
}

