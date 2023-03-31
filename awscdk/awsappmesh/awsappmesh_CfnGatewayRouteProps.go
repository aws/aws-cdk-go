package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnGatewayRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayRouteProps := &cfnGatewayRouteProps{
//   	meshName: jsii.String("meshName"),
//   	spec: &gatewayRouteSpecProperty{
//   		grpcRoute: &grpcGatewayRouteProperty{
//   			action: &grpcGatewayRouteActionProperty{
//   				target: &gatewayRouteTargetProperty{
//   					virtualService: &gatewayRouteVirtualServiceProperty{
//   						virtualServiceName: jsii.String("virtualServiceName"),
//   					},
//
//   					// the properties below are optional
//   					port: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				rewrite: &grpcGatewayRouteRewriteProperty{
//   					hostname: &gatewayRouteHostnameRewriteProperty{
//   						defaultTargetHostname: jsii.String("defaultTargetHostname"),
//   					},
//   				},
//   			},
//   			match: &grpcGatewayRouteMatchProperty{
//   				hostname: &gatewayRouteHostnameMatchProperty{
//   					exact: jsii.String("exact"),
//   					suffix: jsii.String("suffix"),
//   				},
//   				metadata: []interface{}{
//   					&grpcGatewayRouteMetadataProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						invert: jsii.Boolean(false),
//   						match: &gatewayRouteMetadataMatchProperty{
//   							exact: jsii.String("exact"),
//   							prefix: jsii.String("prefix"),
//   							range: &gatewayRouteRangeMatchProperty{
//   								end: jsii.Number(123),
//   								start: jsii.Number(123),
//   							},
//   							regex: jsii.String("regex"),
//   							suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				port: jsii.Number(123),
//   				serviceName: jsii.String("serviceName"),
//   			},
//   		},
//   		http2Route: &httpGatewayRouteProperty{
//   			action: &httpGatewayRouteActionProperty{
//   				target: &gatewayRouteTargetProperty{
//   					virtualService: &gatewayRouteVirtualServiceProperty{
//   						virtualServiceName: jsii.String("virtualServiceName"),
//   					},
//
//   					// the properties below are optional
//   					port: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				rewrite: &httpGatewayRouteRewriteProperty{
//   					hostname: &gatewayRouteHostnameRewriteProperty{
//   						defaultTargetHostname: jsii.String("defaultTargetHostname"),
//   					},
//   					path: &httpGatewayRoutePathRewriteProperty{
//   						exact: jsii.String("exact"),
//   					},
//   					prefix: &httpGatewayRoutePrefixRewriteProperty{
//   						defaultPrefix: jsii.String("defaultPrefix"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			match: &httpGatewayRouteMatchProperty{
//   				headers: []interface{}{
//   					&httpGatewayRouteHeaderProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						invert: jsii.Boolean(false),
//   						match: &httpGatewayRouteHeaderMatchProperty{
//   							exact: jsii.String("exact"),
//   							prefix: jsii.String("prefix"),
//   							range: &gatewayRouteRangeMatchProperty{
//   								end: jsii.Number(123),
//   								start: jsii.Number(123),
//   							},
//   							regex: jsii.String("regex"),
//   							suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				hostname: &gatewayRouteHostnameMatchProperty{
//   					exact: jsii.String("exact"),
//   					suffix: jsii.String("suffix"),
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
//   			},
//   		},
//   		httpRoute: &httpGatewayRouteProperty{
//   			action: &httpGatewayRouteActionProperty{
//   				target: &gatewayRouteTargetProperty{
//   					virtualService: &gatewayRouteVirtualServiceProperty{
//   						virtualServiceName: jsii.String("virtualServiceName"),
//   					},
//
//   					// the properties below are optional
//   					port: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				rewrite: &httpGatewayRouteRewriteProperty{
//   					hostname: &gatewayRouteHostnameRewriteProperty{
//   						defaultTargetHostname: jsii.String("defaultTargetHostname"),
//   					},
//   					path: &httpGatewayRoutePathRewriteProperty{
//   						exact: jsii.String("exact"),
//   					},
//   					prefix: &httpGatewayRoutePrefixRewriteProperty{
//   						defaultPrefix: jsii.String("defaultPrefix"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			match: &httpGatewayRouteMatchProperty{
//   				headers: []interface{}{
//   					&httpGatewayRouteHeaderProperty{
//   						name: jsii.String("name"),
//
//   						// the properties below are optional
//   						invert: jsii.Boolean(false),
//   						match: &httpGatewayRouteHeaderMatchProperty{
//   							exact: jsii.String("exact"),
//   							prefix: jsii.String("prefix"),
//   							range: &gatewayRouteRangeMatchProperty{
//   								end: jsii.Number(123),
//   								start: jsii.Number(123),
//   							},
//   							regex: jsii.String("regex"),
//   							suffix: jsii.String("suffix"),
//   						},
//   					},
//   				},
//   				hostname: &gatewayRouteHostnameMatchProperty{
//   					exact: jsii.String("exact"),
//   					suffix: jsii.String("suffix"),
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
//   			},
//   		},
//   		priority: jsii.Number(123),
//   	},
//   	virtualGatewayName: jsii.String("virtualGatewayName"),
//
//   	// the properties below are optional
//   	gatewayRouteName: jsii.String("gatewayRouteName"),
//   	meshOwner: jsii.String("meshOwner"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnGatewayRouteProps struct {
	// The name of the service mesh that the resource resides in.
	MeshName *string `field:"required" json:"meshName" yaml:"meshName"`
	// The specifications of the gateway route.
	Spec interface{} `field:"required" json:"spec" yaml:"spec"`
	// The virtual gateway that the gateway route is associated with.
	VirtualGatewayName *string `field:"required" json:"virtualGatewayName" yaml:"virtualGatewayName"`
	// The name of the gateway route.
	GatewayRouteName *string `field:"optional" json:"gatewayRouteName" yaml:"gatewayRouteName"`
	// The AWS IAM account ID of the service mesh owner.
	//
	// If the account ID is not your own, then it's the ID of the account that shared the mesh with your account. For more information about mesh sharing, see [Working with shared meshes](https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html) .
	MeshOwner *string `field:"optional" json:"meshOwner" yaml:"meshOwner"`
	// Optional metadata that you can apply to the gateway route to assist with categorization and organization.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

