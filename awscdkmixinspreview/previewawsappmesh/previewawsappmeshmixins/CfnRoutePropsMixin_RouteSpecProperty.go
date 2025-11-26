package previewawsappmeshmixins


// An object that represents a route specification.
//
// Specify one route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   routeSpecProperty := &RouteSpecProperty{
//   	GrpcRoute: &GrpcRouteProperty{
//   		Action: &GrpcRouteActionProperty{
//   			WeightedTargets: []interface{}{
//   				&WeightedTargetProperty{
//   					Port: jsii.Number(123),
//   					VirtualNode: jsii.String("virtualNode"),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Match: &GrpcRouteMatchProperty{
//   			Metadata: []interface{}{
//   				&GrpcRouteMetadataProperty{
//   					Invert: jsii.Boolean(false),
//   					Match: &GrpcRouteMetadataMatchMethodProperty{
//   						Exact: jsii.String("exact"),
//   						Prefix: jsii.String("prefix"),
//   						Range: &MatchRangeProperty{
//   							End: jsii.Number(123),
//   							Start: jsii.Number(123),
//   						},
//   						Regex: jsii.String("regex"),
//   						Suffix: jsii.String("suffix"),
//   					},
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			MethodName: jsii.String("methodName"),
//   			Port: jsii.Number(123),
//   			ServiceName: jsii.String("serviceName"),
//   		},
//   		RetryPolicy: &GrpcRetryPolicyProperty{
//   			GrpcRetryEvents: []*string{
//   				jsii.String("grpcRetryEvents"),
//   			},
//   			HttpRetryEvents: []*string{
//   				jsii.String("httpRetryEvents"),
//   			},
//   			MaxRetries: jsii.Number(123),
//   			PerRetryTimeout: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			TcpRetryEvents: []*string{
//   				jsii.String("tcpRetryEvents"),
//   			},
//   		},
//   		Timeout: &GrpcTimeoutProperty{
//   			Idle: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			PerRequest: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Http2Route: &HttpRouteProperty{
//   		Action: &HttpRouteActionProperty{
//   			WeightedTargets: []interface{}{
//   				&WeightedTargetProperty{
//   					Port: jsii.Number(123),
//   					VirtualNode: jsii.String("virtualNode"),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Match: &HttpRouteMatchProperty{
//   			Headers: []interface{}{
//   				&HttpRouteHeaderProperty{
//   					Invert: jsii.Boolean(false),
//   					Match: &HeaderMatchMethodProperty{
//   						Exact: jsii.String("exact"),
//   						Prefix: jsii.String("prefix"),
//   						Range: &MatchRangeProperty{
//   							End: jsii.Number(123),
//   							Start: jsii.Number(123),
//   						},
//   						Regex: jsii.String("regex"),
//   						Suffix: jsii.String("suffix"),
//   					},
//   					Name: jsii.String("name"),
//   				},
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
//   			Scheme: jsii.String("scheme"),
//   		},
//   		RetryPolicy: &HttpRetryPolicyProperty{
//   			HttpRetryEvents: []*string{
//   				jsii.String("httpRetryEvents"),
//   			},
//   			MaxRetries: jsii.Number(123),
//   			PerRetryTimeout: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			TcpRetryEvents: []*string{
//   				jsii.String("tcpRetryEvents"),
//   			},
//   		},
//   		Timeout: &HttpTimeoutProperty{
//   			Idle: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			PerRequest: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	HttpRoute: &HttpRouteProperty{
//   		Action: &HttpRouteActionProperty{
//   			WeightedTargets: []interface{}{
//   				&WeightedTargetProperty{
//   					Port: jsii.Number(123),
//   					VirtualNode: jsii.String("virtualNode"),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Match: &HttpRouteMatchProperty{
//   			Headers: []interface{}{
//   				&HttpRouteHeaderProperty{
//   					Invert: jsii.Boolean(false),
//   					Match: &HeaderMatchMethodProperty{
//   						Exact: jsii.String("exact"),
//   						Prefix: jsii.String("prefix"),
//   						Range: &MatchRangeProperty{
//   							End: jsii.Number(123),
//   							Start: jsii.Number(123),
//   						},
//   						Regex: jsii.String("regex"),
//   						Suffix: jsii.String("suffix"),
//   					},
//   					Name: jsii.String("name"),
//   				},
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
//   			Scheme: jsii.String("scheme"),
//   		},
//   		RetryPolicy: &HttpRetryPolicyProperty{
//   			HttpRetryEvents: []*string{
//   				jsii.String("httpRetryEvents"),
//   			},
//   			MaxRetries: jsii.Number(123),
//   			PerRetryTimeout: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			TcpRetryEvents: []*string{
//   				jsii.String("tcpRetryEvents"),
//   			},
//   		},
//   		Timeout: &HttpTimeoutProperty{
//   			Idle: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			PerRequest: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Priority: jsii.Number(123),
//   	TcpRoute: &TcpRouteProperty{
//   		Action: &TcpRouteActionProperty{
//   			WeightedTargets: []interface{}{
//   				&WeightedTargetProperty{
//   					Port: jsii.Number(123),
//   					VirtualNode: jsii.String("virtualNode"),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Match: &TcpRouteMatchProperty{
//   			Port: jsii.Number(123),
//   		},
//   		Timeout: &TcpTimeoutProperty{
//   			Idle: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-routespec.html
//
type CfnRoutePropsMixin_RouteSpecProperty struct {
	// An object that represents the specification of a gRPC route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-routespec.html#cfn-appmesh-route-routespec-grpcroute
	//
	GrpcRoute interface{} `field:"optional" json:"grpcRoute" yaml:"grpcRoute"`
	// An object that represents the specification of an HTTP/2 route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-routespec.html#cfn-appmesh-route-routespec-http2route
	//
	Http2Route interface{} `field:"optional" json:"http2Route" yaml:"http2Route"`
	// An object that represents the specification of an HTTP route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-routespec.html#cfn-appmesh-route-routespec-httproute
	//
	HttpRoute interface{} `field:"optional" json:"httpRoute" yaml:"httpRoute"`
	// The priority for the route.
	//
	// Routes are matched based on the specified value, where 0 is the highest priority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-routespec.html#cfn-appmesh-route-routespec-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// An object that represents the specification of a TCP route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-routespec.html#cfn-appmesh-route-routespec-tcproute
	//
	TcpRoute interface{} `field:"optional" json:"tcpRoute" yaml:"tcpRoute"`
}

