package awsappmesh


// All Properties for Route Specs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeSpecConfig := &RouteSpecConfig{
//   	GrpcRouteSpec: &GrpcRouteProperty{
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
//   	Http2RouteSpec: &HttpRouteProperty{
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
//   	HttpRouteSpec: &HttpRouteProperty{
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
//   	TcpRouteSpec: &TcpRouteProperty{
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
type RouteSpecConfig struct {
	// The spec for a grpc route.
	GrpcRouteSpec *CfnRoute_GrpcRouteProperty `field:"optional" json:"grpcRouteSpec" yaml:"grpcRouteSpec"`
	// The spec for an http2 route.
	Http2RouteSpec *CfnRoute_HttpRouteProperty `field:"optional" json:"http2RouteSpec" yaml:"http2RouteSpec"`
	// The spec for an http route.
	HttpRouteSpec *CfnRoute_HttpRouteProperty `field:"optional" json:"httpRouteSpec" yaml:"httpRouteSpec"`
	// The priority for the route.
	//
	// When a Virtual Router has multiple routes, route match is performed in the
	// order of specified value, where 0 is the highest priority, and first matched route is selected.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The spec for a tcp route.
	TcpRouteSpec *CfnRoute_TcpRouteProperty `field:"optional" json:"tcpRouteSpec" yaml:"tcpRouteSpec"`
}

