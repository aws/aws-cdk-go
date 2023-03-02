package awsappmesh


// An object that represents a gateway route specification.
//
// Specify one gateway route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteSpecProperty := &gatewayRouteSpecProperty{
//   	grpcRoute: &grpcGatewayRouteProperty{
//   		action: &grpcGatewayRouteActionProperty{
//   			target: &gatewayRouteTargetProperty{
//   				virtualService: &gatewayRouteVirtualServiceProperty{
//   					virtualServiceName: jsii.String("virtualServiceName"),
//   				},
//
//   				// the properties below are optional
//   				port: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			rewrite: &grpcGatewayRouteRewriteProperty{
//   				hostname: &gatewayRouteHostnameRewriteProperty{
//   					defaultTargetHostname: jsii.String("defaultTargetHostname"),
//   				},
//   			},
//   		},
//   		match: &grpcGatewayRouteMatchProperty{
//   			hostname: &gatewayRouteHostnameMatchProperty{
//   				exact: jsii.String("exact"),
//   				suffix: jsii.String("suffix"),
//   			},
//   			metadata: []interface{}{
//   				&grpcGatewayRouteMetadataProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					invert: jsii.Boolean(false),
//   					match: &gatewayRouteMetadataMatchProperty{
//   						exact: jsii.String("exact"),
//   						prefix: jsii.String("prefix"),
//   						range: &gatewayRouteRangeMatchProperty{
//   							end: jsii.Number(123),
//   							start: jsii.Number(123),
//   						},
//   						regex: jsii.String("regex"),
//   						suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			port: jsii.Number(123),
//   			serviceName: jsii.String("serviceName"),
//   		},
//   	},
//   	http2Route: &httpGatewayRouteProperty{
//   		action: &httpGatewayRouteActionProperty{
//   			target: &gatewayRouteTargetProperty{
//   				virtualService: &gatewayRouteVirtualServiceProperty{
//   					virtualServiceName: jsii.String("virtualServiceName"),
//   				},
//
//   				// the properties below are optional
//   				port: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			rewrite: &httpGatewayRouteRewriteProperty{
//   				hostname: &gatewayRouteHostnameRewriteProperty{
//   					defaultTargetHostname: jsii.String("defaultTargetHostname"),
//   				},
//   				path: &httpGatewayRoutePathRewriteProperty{
//   					exact: jsii.String("exact"),
//   				},
//   				prefix: &httpGatewayRoutePrefixRewriteProperty{
//   					defaultPrefix: jsii.String("defaultPrefix"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		match: &httpGatewayRouteMatchProperty{
//   			headers: []interface{}{
//   				&httpGatewayRouteHeaderProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					invert: jsii.Boolean(false),
//   					match: &httpGatewayRouteHeaderMatchProperty{
//   						exact: jsii.String("exact"),
//   						prefix: jsii.String("prefix"),
//   						range: &gatewayRouteRangeMatchProperty{
//   							end: jsii.Number(123),
//   							start: jsii.Number(123),
//   						},
//   						regex: jsii.String("regex"),
//   						suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			hostname: &gatewayRouteHostnameMatchProperty{
//   				exact: jsii.String("exact"),
//   				suffix: jsii.String("suffix"),
//   			},
//   			method: jsii.String("method"),
//   			path: &httpPathMatchProperty{
//   				exact: jsii.String("exact"),
//   				regex: jsii.String("regex"),
//   			},
//   			port: jsii.Number(123),
//   			prefix: jsii.String("prefix"),
//   			queryParameters: []interface{}{
//   				&queryParameterProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					match: &httpQueryParameterMatchProperty{
//   						exact: jsii.String("exact"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	httpRoute: &httpGatewayRouteProperty{
//   		action: &httpGatewayRouteActionProperty{
//   			target: &gatewayRouteTargetProperty{
//   				virtualService: &gatewayRouteVirtualServiceProperty{
//   					virtualServiceName: jsii.String("virtualServiceName"),
//   				},
//
//   				// the properties below are optional
//   				port: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			rewrite: &httpGatewayRouteRewriteProperty{
//   				hostname: &gatewayRouteHostnameRewriteProperty{
//   					defaultTargetHostname: jsii.String("defaultTargetHostname"),
//   				},
//   				path: &httpGatewayRoutePathRewriteProperty{
//   					exact: jsii.String("exact"),
//   				},
//   				prefix: &httpGatewayRoutePrefixRewriteProperty{
//   					defaultPrefix: jsii.String("defaultPrefix"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		match: &httpGatewayRouteMatchProperty{
//   			headers: []interface{}{
//   				&httpGatewayRouteHeaderProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					invert: jsii.Boolean(false),
//   					match: &httpGatewayRouteHeaderMatchProperty{
//   						exact: jsii.String("exact"),
//   						prefix: jsii.String("prefix"),
//   						range: &gatewayRouteRangeMatchProperty{
//   							end: jsii.Number(123),
//   							start: jsii.Number(123),
//   						},
//   						regex: jsii.String("regex"),
//   						suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			hostname: &gatewayRouteHostnameMatchProperty{
//   				exact: jsii.String("exact"),
//   				suffix: jsii.String("suffix"),
//   			},
//   			method: jsii.String("method"),
//   			path: &httpPathMatchProperty{
//   				exact: jsii.String("exact"),
//   				regex: jsii.String("regex"),
//   			},
//   			port: jsii.Number(123),
//   			prefix: jsii.String("prefix"),
//   			queryParameters: []interface{}{
//   				&queryParameterProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					match: &httpQueryParameterMatchProperty{
//   						exact: jsii.String("exact"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	priority: jsii.Number(123),
//   }
//
type CfnGatewayRoute_GatewayRouteSpecProperty struct {
	// An object that represents the specification of a gRPC gateway route.
	GrpcRoute interface{} `field:"optional" json:"grpcRoute" yaml:"grpcRoute"`
	// An object that represents the specification of an HTTP/2 gateway route.
	Http2Route interface{} `field:"optional" json:"http2Route" yaml:"http2Route"`
	// An object that represents the specification of an HTTP gateway route.
	HttpRoute interface{} `field:"optional" json:"httpRoute" yaml:"httpRoute"`
	// The ordering of the gateway routes spec.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

