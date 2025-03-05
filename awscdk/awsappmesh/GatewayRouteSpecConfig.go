package awsappmesh


// All Properties for GatewayRoute Specs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteSpecConfig := &GatewayRouteSpecConfig{
//   	GrpcSpecConfig: &GrpcGatewayRouteProperty{
//   		Action: &GrpcGatewayRouteActionProperty{
//   			Target: &GatewayRouteTargetProperty{
//   				VirtualService: &GatewayRouteVirtualServiceProperty{
//   					VirtualServiceName: jsii.String("virtualServiceName"),
//   				},
//
//   				// the properties below are optional
//   				Port: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			Rewrite: &GrpcGatewayRouteRewriteProperty{
//   				Hostname: &GatewayRouteHostnameRewriteProperty{
//   					DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   				},
//   			},
//   		},
//   		Match: &GrpcGatewayRouteMatchProperty{
//   			Hostname: &GatewayRouteHostnameMatchProperty{
//   				Exact: jsii.String("exact"),
//   				Suffix: jsii.String("suffix"),
//   			},
//   			Metadata: []interface{}{
//   				&GrpcGatewayRouteMetadataProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Invert: jsii.Boolean(false),
//   					Match: &GatewayRouteMetadataMatchProperty{
//   						Exact: jsii.String("exact"),
//   						Prefix: jsii.String("prefix"),
//   						Range: &GatewayRouteRangeMatchProperty{
//   							End: jsii.Number(123),
//   							Start: jsii.Number(123),
//   						},
//   						Regex: jsii.String("regex"),
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			Port: jsii.Number(123),
//   			ServiceName: jsii.String("serviceName"),
//   		},
//   	},
//   	Http2SpecConfig: &HttpGatewayRouteProperty{
//   		Action: &HttpGatewayRouteActionProperty{
//   			Target: &GatewayRouteTargetProperty{
//   				VirtualService: &GatewayRouteVirtualServiceProperty{
//   					VirtualServiceName: jsii.String("virtualServiceName"),
//   				},
//
//   				// the properties below are optional
//   				Port: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			Rewrite: &HttpGatewayRouteRewriteProperty{
//   				Hostname: &GatewayRouteHostnameRewriteProperty{
//   					DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   				},
//   				Path: &HttpGatewayRoutePathRewriteProperty{
//   					Exact: jsii.String("exact"),
//   				},
//   				Prefix: &HttpGatewayRoutePrefixRewriteProperty{
//   					DefaultPrefix: jsii.String("defaultPrefix"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		Match: &HttpGatewayRouteMatchProperty{
//   			Headers: []interface{}{
//   				&HttpGatewayRouteHeaderProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Invert: jsii.Boolean(false),
//   					Match: &HttpGatewayRouteHeaderMatchProperty{
//   						Exact: jsii.String("exact"),
//   						Prefix: jsii.String("prefix"),
//   						Range: &GatewayRouteRangeMatchProperty{
//   							End: jsii.Number(123),
//   							Start: jsii.Number(123),
//   						},
//   						Regex: jsii.String("regex"),
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			Hostname: &GatewayRouteHostnameMatchProperty{
//   				Exact: jsii.String("exact"),
//   				Suffix: jsii.String("suffix"),
//   			},
//   			Method: jsii.String("method"),
//   			Path: &HttpPathMatchProperty{
//   				Exact: jsii.String("exact"),
//   				Regex: jsii.String("regex"),
//   			},
//   			Port: jsii.Number(123),
//   			Prefix: jsii.String("prefix"),
//   			QueryParameters: []interface{}{
//   				&QueryParameterProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Match: &HttpQueryParameterMatchProperty{
//   						Exact: jsii.String("exact"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	HttpSpecConfig: &HttpGatewayRouteProperty{
//   		Action: &HttpGatewayRouteActionProperty{
//   			Target: &GatewayRouteTargetProperty{
//   				VirtualService: &GatewayRouteVirtualServiceProperty{
//   					VirtualServiceName: jsii.String("virtualServiceName"),
//   				},
//
//   				// the properties below are optional
//   				Port: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			Rewrite: &HttpGatewayRouteRewriteProperty{
//   				Hostname: &GatewayRouteHostnameRewriteProperty{
//   					DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   				},
//   				Path: &HttpGatewayRoutePathRewriteProperty{
//   					Exact: jsii.String("exact"),
//   				},
//   				Prefix: &HttpGatewayRoutePrefixRewriteProperty{
//   					DefaultPrefix: jsii.String("defaultPrefix"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		Match: &HttpGatewayRouteMatchProperty{
//   			Headers: []interface{}{
//   				&HttpGatewayRouteHeaderProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Invert: jsii.Boolean(false),
//   					Match: &HttpGatewayRouteHeaderMatchProperty{
//   						Exact: jsii.String("exact"),
//   						Prefix: jsii.String("prefix"),
//   						Range: &GatewayRouteRangeMatchProperty{
//   							End: jsii.Number(123),
//   							Start: jsii.Number(123),
//   						},
//   						Regex: jsii.String("regex"),
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			Hostname: &GatewayRouteHostnameMatchProperty{
//   				Exact: jsii.String("exact"),
//   				Suffix: jsii.String("suffix"),
//   			},
//   			Method: jsii.String("method"),
//   			Path: &HttpPathMatchProperty{
//   				Exact: jsii.String("exact"),
//   				Regex: jsii.String("regex"),
//   			},
//   			Port: jsii.Number(123),
//   			Prefix: jsii.String("prefix"),
//   			QueryParameters: []interface{}{
//   				&QueryParameterProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Match: &HttpQueryParameterMatchProperty{
//   						Exact: jsii.String("exact"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Priority: jsii.Number(123),
//   }
//
type GatewayRouteSpecConfig struct {
	// The spec for a grpc gateway route.
	// Default: - no grpc spec.
	//
	GrpcSpecConfig *CfnGatewayRoute_GrpcGatewayRouteProperty `field:"optional" json:"grpcSpecConfig" yaml:"grpcSpecConfig"`
	// The spec for an http2 gateway route.
	// Default: - no http2 spec.
	//
	Http2SpecConfig *CfnGatewayRoute_HttpGatewayRouteProperty `field:"optional" json:"http2SpecConfig" yaml:"http2SpecConfig"`
	// The spec for an http gateway route.
	// Default: - no http spec.
	//
	HttpSpecConfig *CfnGatewayRoute_HttpGatewayRouteProperty `field:"optional" json:"httpSpecConfig" yaml:"httpSpecConfig"`
	// The priority for the gateway route.
	//
	// When a Virtual Gateway has multiple gateway routes, gateway route match
	// is performed in the order of specified value, where 0 is the highest priority,
	// and first matched gateway route is selected.
	// Default: - no particular priority.
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

