package awsappmesh


// An object that represents a route specification.
//
// Specify one route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeSpecProperty := &RouteSpecProperty{
//   	GrpcRoute: &GrpcRouteProperty{
//   		Action: &GrpcRouteActionProperty{
//   			WeightedTargets: []interface{}{
//   				&WeightedTargetProperty{
//   					VirtualNode: jsii.String("virtualNode"),
//   					Weight: jsii.Number(123),
//
//   					// the properties below are optional
//   					Port: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Match: &GrpcRouteMatchProperty{
//   			Metadata: []interface{}{
//   				&GrpcRouteMetadataProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
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
//   				},
//   			},
//   			MethodName: jsii.String("methodName"),
//   			Port: jsii.Number(123),
//   			ServiceName: jsii.String("serviceName"),
//   		},
//
//   		// the properties below are optional
//   		RetryPolicy: &GrpcRetryPolicyProperty{
//   			MaxRetries: jsii.Number(123),
//   			PerRetryTimeout: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			GrpcRetryEvents: []*string{
//   				jsii.String("grpcRetryEvents"),
//   			},
//   			HttpRetryEvents: []*string{
//   				jsii.String("httpRetryEvents"),
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
//   					VirtualNode: jsii.String("virtualNode"),
//   					Weight: jsii.Number(123),
//
//   					// the properties below are optional
//   					Port: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Match: &HttpRouteMatchProperty{
//   			Headers: []interface{}{
//   				&HttpRouteHeaderProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
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
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Match: &HttpQueryParameterMatchProperty{
//   						Exact: jsii.String("exact"),
//   					},
//   				},
//   			},
//   			Scheme: jsii.String("scheme"),
//   		},
//
//   		// the properties below are optional
//   		RetryPolicy: &HttpRetryPolicyProperty{
//   			MaxRetries: jsii.Number(123),
//   			PerRetryTimeout: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			HttpRetryEvents: []*string{
//   				jsii.String("httpRetryEvents"),
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
//   					VirtualNode: jsii.String("virtualNode"),
//   					Weight: jsii.Number(123),
//
//   					// the properties below are optional
//   					Port: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Match: &HttpRouteMatchProperty{
//   			Headers: []interface{}{
//   				&HttpRouteHeaderProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
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
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Match: &HttpQueryParameterMatchProperty{
//   						Exact: jsii.String("exact"),
//   					},
//   				},
//   			},
//   			Scheme: jsii.String("scheme"),
//   		},
//
//   		// the properties below are optional
//   		RetryPolicy: &HttpRetryPolicyProperty{
//   			MaxRetries: jsii.Number(123),
//   			PerRetryTimeout: &DurationProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			HttpRetryEvents: []*string{
//   				jsii.String("httpRetryEvents"),
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
//   					VirtualNode: jsii.String("virtualNode"),
//   					Weight: jsii.Number(123),
//
//   					// the properties below are optional
//   					Port: jsii.Number(123),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
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
type CfnRoute_RouteSpecProperty struct {
	// An object that represents the specification of a gRPC route.
	GrpcRoute interface{} `field:"optional" json:"grpcRoute" yaml:"grpcRoute"`
	// An object that represents the specification of an HTTP/2 route.
	Http2Route interface{} `field:"optional" json:"http2Route" yaml:"http2Route"`
	// An object that represents the specification of an HTTP route.
	HttpRoute interface{} `field:"optional" json:"httpRoute" yaml:"httpRoute"`
	// The priority for the route.
	//
	// Routes are matched based on the specified value, where 0 is the highest priority.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// An object that represents the specification of a TCP route.
	TcpRoute interface{} `field:"optional" json:"tcpRoute" yaml:"tcpRoute"`
}

