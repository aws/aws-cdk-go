package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGatewayRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayRouteProps := &CfnGatewayRouteProps{
//   	MeshName: jsii.String("meshName"),
//   	Spec: &GatewayRouteSpecProperty{
//   		GrpcRoute: &GrpcGatewayRouteProperty{
//   			Action: &GrpcGatewayRouteActionProperty{
//   				Target: &GatewayRouteTargetProperty{
//   					VirtualService: &GatewayRouteVirtualServiceProperty{
//   						VirtualServiceName: jsii.String("virtualServiceName"),
//   					},
//
//   					// the properties below are optional
//   					Port: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				Rewrite: &GrpcGatewayRouteRewriteProperty{
//   					Hostname: &GatewayRouteHostnameRewriteProperty{
//   						DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   					},
//   				},
//   			},
//   			Match: &GrpcGatewayRouteMatchProperty{
//   				Hostname: &GatewayRouteHostnameMatchProperty{
//   					Exact: jsii.String("exact"),
//   					Suffix: jsii.String("suffix"),
//   				},
//   				Metadata: []interface{}{
//   					&GrpcGatewayRouteMetadataProperty{
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Invert: jsii.Boolean(false),
//   						Match: &GatewayRouteMetadataMatchProperty{
//   							Exact: jsii.String("exact"),
//   							Prefix: jsii.String("prefix"),
//   							Range: &GatewayRouteRangeMatchProperty{
//   								End: jsii.Number(123),
//   								Start: jsii.Number(123),
//   							},
//   							Regex: jsii.String("regex"),
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				Port: jsii.Number(123),
//   				ServiceName: jsii.String("serviceName"),
//   			},
//   		},
//   		Http2Route: &HttpGatewayRouteProperty{
//   			Action: &HttpGatewayRouteActionProperty{
//   				Target: &GatewayRouteTargetProperty{
//   					VirtualService: &GatewayRouteVirtualServiceProperty{
//   						VirtualServiceName: jsii.String("virtualServiceName"),
//   					},
//
//   					// the properties below are optional
//   					Port: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				Rewrite: &HttpGatewayRouteRewriteProperty{
//   					Hostname: &GatewayRouteHostnameRewriteProperty{
//   						DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   					},
//   					Path: &HttpGatewayRoutePathRewriteProperty{
//   						Exact: jsii.String("exact"),
//   					},
//   					Prefix: &HttpGatewayRoutePrefixRewriteProperty{
//   						DefaultPrefix: jsii.String("defaultPrefix"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			Match: &HttpGatewayRouteMatchProperty{
//   				Headers: []interface{}{
//   					&HttpGatewayRouteHeaderProperty{
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Invert: jsii.Boolean(false),
//   						Match: &HttpGatewayRouteHeaderMatchProperty{
//   							Exact: jsii.String("exact"),
//   							Prefix: jsii.String("prefix"),
//   							Range: &GatewayRouteRangeMatchProperty{
//   								End: jsii.Number(123),
//   								Start: jsii.Number(123),
//   							},
//   							Regex: jsii.String("regex"),
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				Hostname: &GatewayRouteHostnameMatchProperty{
//   					Exact: jsii.String("exact"),
//   					Suffix: jsii.String("suffix"),
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
//   			},
//   		},
//   		HttpRoute: &HttpGatewayRouteProperty{
//   			Action: &HttpGatewayRouteActionProperty{
//   				Target: &GatewayRouteTargetProperty{
//   					VirtualService: &GatewayRouteVirtualServiceProperty{
//   						VirtualServiceName: jsii.String("virtualServiceName"),
//   					},
//
//   					// the properties below are optional
//   					Port: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				Rewrite: &HttpGatewayRouteRewriteProperty{
//   					Hostname: &GatewayRouteHostnameRewriteProperty{
//   						DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   					},
//   					Path: &HttpGatewayRoutePathRewriteProperty{
//   						Exact: jsii.String("exact"),
//   					},
//   					Prefix: &HttpGatewayRoutePrefixRewriteProperty{
//   						DefaultPrefix: jsii.String("defaultPrefix"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			Match: &HttpGatewayRouteMatchProperty{
//   				Headers: []interface{}{
//   					&HttpGatewayRouteHeaderProperty{
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Invert: jsii.Boolean(false),
//   						Match: &HttpGatewayRouteHeaderMatchProperty{
//   							Exact: jsii.String("exact"),
//   							Prefix: jsii.String("prefix"),
//   							Range: &GatewayRouteRangeMatchProperty{
//   								End: jsii.Number(123),
//   								Start: jsii.Number(123),
//   							},
//   							Regex: jsii.String("regex"),
//   							Suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				Hostname: &GatewayRouteHostnameMatchProperty{
//   					Exact: jsii.String("exact"),
//   					Suffix: jsii.String("suffix"),
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
//   			},
//   		},
//   		Priority: jsii.Number(123),
//   	},
//   	VirtualGatewayName: jsii.String("virtualGatewayName"),
//
//   	// the properties below are optional
//   	GatewayRouteName: jsii.String("gatewayRouteName"),
//   	MeshOwner: jsii.String("meshOwner"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-gatewayroute.html
//
type CfnGatewayRouteProps struct {
	// The name of the service mesh that the resource resides in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-gatewayroute.html#cfn-appmesh-gatewayroute-meshname
	//
	MeshName *string `field:"required" json:"meshName" yaml:"meshName"`
	// The specifications of the gateway route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-gatewayroute.html#cfn-appmesh-gatewayroute-spec
	//
	Spec interface{} `field:"required" json:"spec" yaml:"spec"`
	// The virtual gateway that the gateway route is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-gatewayroute.html#cfn-appmesh-gatewayroute-virtualgatewayname
	//
	VirtualGatewayName *string `field:"required" json:"virtualGatewayName" yaml:"virtualGatewayName"`
	// The name of the gateway route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-gatewayroute.html#cfn-appmesh-gatewayroute-gatewayroutename
	//
	GatewayRouteName *string `field:"optional" json:"gatewayRouteName" yaml:"gatewayRouteName"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then it's the ID of the account that shared the mesh with your account. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-gatewayroute.html#cfn-appmesh-gatewayroute-meshowner
	//
	MeshOwner *string `field:"optional" json:"meshOwner" yaml:"meshOwner"`
	// Optional metadata that you can apply to the gateway route to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-gatewayroute.html#cfn-appmesh-gatewayroute-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

