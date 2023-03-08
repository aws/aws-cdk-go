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
//   gatewayRouteSpecProperty := &GatewayRouteSpecProperty{
//   	GrpcRoute: &GrpcGatewayRouteProperty{
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
//   	Http2Route: &HttpGatewayRouteProperty{
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
//   	HttpRoute: &HttpGatewayRouteProperty{
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

