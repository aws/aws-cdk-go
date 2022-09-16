package awsappmesh


// All Properties for GatewayRoute Specs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteSpecConfig := &gatewayRouteSpecConfig{
//   	grpcSpecConfig: &grpcGatewayRouteProperty{
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
//   	http2SpecConfig: &httpGatewayRouteProperty{
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
//   	httpSpecConfig: &httpGatewayRouteProperty{
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
type GatewayRouteSpecConfig struct {
	// The spec for a grpc gateway route.
	GrpcSpecConfig *CfnGatewayRoute_GrpcGatewayRouteProperty `field:"optional" json:"grpcSpecConfig" yaml:"grpcSpecConfig"`
	// The spec for an http2 gateway route.
	Http2SpecConfig *CfnGatewayRoute_HttpGatewayRouteProperty `field:"optional" json:"http2SpecConfig" yaml:"http2SpecConfig"`
	// The spec for an http gateway route.
	HttpSpecConfig *CfnGatewayRoute_HttpGatewayRouteProperty `field:"optional" json:"httpSpecConfig" yaml:"httpSpecConfig"`
	// The priority for the gateway route.
	//
	// When a Virtual Gateway has multiple gateway routes, gateway route match
	// is performed in the order of specified value, where 0 is the highest priority,
	// and first matched gateway route is selected.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

