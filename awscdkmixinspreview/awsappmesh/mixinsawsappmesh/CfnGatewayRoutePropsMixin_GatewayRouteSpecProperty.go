package mixinsawsappmesh


// An object that represents a gateway route specification.
//
// Specify one gateway route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gatewayRouteSpecProperty := &GatewayRouteSpecProperty{
//   	GrpcRoute: &GrpcGatewayRouteProperty{
//   		Action: &GrpcGatewayRouteActionProperty{
//   			Rewrite: &GrpcGatewayRouteRewriteProperty{
//   				Hostname: &GatewayRouteHostnameRewriteProperty{
//   					DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   				},
//   			},
//   			Target: &GatewayRouteTargetProperty{
//   				Port: jsii.Number(123),
//   				VirtualService: &GatewayRouteVirtualServiceProperty{
//   					VirtualServiceName: jsii.String("virtualServiceName"),
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
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Port: jsii.Number(123),
//   			ServiceName: jsii.String("serviceName"),
//   		},
//   	},
//   	Http2Route: &HttpGatewayRouteProperty{
//   		Action: &HttpGatewayRouteActionProperty{
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
//   			Target: &GatewayRouteTargetProperty{
//   				Port: jsii.Number(123),
//   				VirtualService: &GatewayRouteVirtualServiceProperty{
//   					VirtualServiceName: jsii.String("virtualServiceName"),
//   				},
//   			},
//   		},
//   		Match: &HttpGatewayRouteMatchProperty{
//   			Headers: []interface{}{
//   				&HttpGatewayRouteHeaderProperty{
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
//   					Name: jsii.String("name"),
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
//   					Match: &HttpQueryParameterMatchProperty{
//   						Exact: jsii.String("exact"),
//   					},
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	HttpRoute: &HttpGatewayRouteProperty{
//   		Action: &HttpGatewayRouteActionProperty{
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
//   			Target: &GatewayRouteTargetProperty{
//   				Port: jsii.Number(123),
//   				VirtualService: &GatewayRouteVirtualServiceProperty{
//   					VirtualServiceName: jsii.String("virtualServiceName"),
//   				},
//   			},
//   		},
//   		Match: &HttpGatewayRouteMatchProperty{
//   			Headers: []interface{}{
//   				&HttpGatewayRouteHeaderProperty{
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
//   					Name: jsii.String("name"),
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
//   					Match: &HttpQueryParameterMatchProperty{
//   						Exact: jsii.String("exact"),
//   					},
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	Priority: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutespec.html
//
type CfnGatewayRoutePropsMixin_GatewayRouteSpecProperty struct {
	// An object that represents the specification of a gRPC gateway route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutespec.html#cfn-appmesh-gatewayroute-gatewayroutespec-grpcroute
	//
	GrpcRoute interface{} `field:"optional" json:"grpcRoute" yaml:"grpcRoute"`
	// An object that represents the specification of an HTTP/2 gateway route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutespec.html#cfn-appmesh-gatewayroute-gatewayroutespec-http2route
	//
	Http2Route interface{} `field:"optional" json:"http2Route" yaml:"http2Route"`
	// An object that represents the specification of an HTTP gateway route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutespec.html#cfn-appmesh-gatewayroute-gatewayroutespec-httproute
	//
	HttpRoute interface{} `field:"optional" json:"httpRoute" yaml:"httpRoute"`
	// The ordering of the gateway routes spec.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutespec.html#cfn-appmesh-gatewayroute-gatewayroutespec-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

