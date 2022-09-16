package awsappmesh


// All Properties for Route Specs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeSpecConfig := &routeSpecConfig{
//   	grpcRouteSpec: &grpcRouteProperty{
//   		action: &grpcRouteActionProperty{
//   			weightedTargets: []interface{}{
//   				&weightedTargetProperty{
//   					virtualNode: jsii.String("virtualNode"),
//   					weight: jsii.Number(123),
//
//   					// the properties below are optional
//   					port: jsii.Number(123),
//   				},
//   			},
//   		},
//   		match: &grpcRouteMatchProperty{
//   			metadata: []interface{}{
//   				&grpcRouteMetadataProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					invert: jsii.Boolean(false),
//   					match: &grpcRouteMetadataMatchMethodProperty{
//   						exact: jsii.String("exact"),
//   						prefix: jsii.String("prefix"),
//   						range: &matchRangeProperty{
//   							end: jsii.Number(123),
//   							start: jsii.Number(123),
//   						},
//   						regex: jsii.String("regex"),
//   						suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			methodName: jsii.String("methodName"),
//   			port: jsii.Number(123),
//   			serviceName: jsii.String("serviceName"),
//   		},
//
//   		// the properties below are optional
//   		retryPolicy: &grpcRetryPolicyProperty{
//   			maxRetries: jsii.Number(123),
//   			perRetryTimeout: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			grpcRetryEvents: []*string{
//   				jsii.String("grpcRetryEvents"),
//   			},
//   			httpRetryEvents: []*string{
//   				jsii.String("httpRetryEvents"),
//   			},
//   			tcpRetryEvents: []*string{
//   				jsii.String("tcpRetryEvents"),
//   			},
//   		},
//   		timeout: &grpcTimeoutProperty{
//   			idle: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   			perRequest: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	http2RouteSpec: &httpRouteProperty{
//   		action: &httpRouteActionProperty{
//   			weightedTargets: []interface{}{
//   				&weightedTargetProperty{
//   					virtualNode: jsii.String("virtualNode"),
//   					weight: jsii.Number(123),
//
//   					// the properties below are optional
//   					port: jsii.Number(123),
//   				},
//   			},
//   		},
//   		match: &httpRouteMatchProperty{
//   			headers: []interface{}{
//   				&httpRouteHeaderProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					invert: jsii.Boolean(false),
//   					match: &headerMatchMethodProperty{
//   						exact: jsii.String("exact"),
//   						prefix: jsii.String("prefix"),
//   						range: &matchRangeProperty{
//   							end: jsii.Number(123),
//   							start: jsii.Number(123),
//   						},
//   						regex: jsii.String("regex"),
//   						suffix: jsii.String("suffix"),
//   					},
//   				},
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
//   			scheme: jsii.String("scheme"),
//   		},
//
//   		// the properties below are optional
//   		retryPolicy: &httpRetryPolicyProperty{
//   			maxRetries: jsii.Number(123),
//   			perRetryTimeout: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			httpRetryEvents: []*string{
//   				jsii.String("httpRetryEvents"),
//   			},
//   			tcpRetryEvents: []*string{
//   				jsii.String("tcpRetryEvents"),
//   			},
//   		},
//   		timeout: &httpTimeoutProperty{
//   			idle: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   			perRequest: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	httpRouteSpec: &httpRouteProperty{
//   		action: &httpRouteActionProperty{
//   			weightedTargets: []interface{}{
//   				&weightedTargetProperty{
//   					virtualNode: jsii.String("virtualNode"),
//   					weight: jsii.Number(123),
//
//   					// the properties below are optional
//   					port: jsii.Number(123),
//   				},
//   			},
//   		},
//   		match: &httpRouteMatchProperty{
//   			headers: []interface{}{
//   				&httpRouteHeaderProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					invert: jsii.Boolean(false),
//   					match: &headerMatchMethodProperty{
//   						exact: jsii.String("exact"),
//   						prefix: jsii.String("prefix"),
//   						range: &matchRangeProperty{
//   							end: jsii.Number(123),
//   							start: jsii.Number(123),
//   						},
//   						regex: jsii.String("regex"),
//   						suffix: jsii.String("suffix"),
//   					},
//   				},
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
//   			scheme: jsii.String("scheme"),
//   		},
//
//   		// the properties below are optional
//   		retryPolicy: &httpRetryPolicyProperty{
//   			maxRetries: jsii.Number(123),
//   			perRetryTimeout: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			httpRetryEvents: []*string{
//   				jsii.String("httpRetryEvents"),
//   			},
//   			tcpRetryEvents: []*string{
//   				jsii.String("tcpRetryEvents"),
//   			},
//   		},
//   		timeout: &httpTimeoutProperty{
//   			idle: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   			perRequest: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	priority: jsii.Number(123),
//   	tcpRouteSpec: &tcpRouteProperty{
//   		action: &tcpRouteActionProperty{
//   			weightedTargets: []interface{}{
//   				&weightedTargetProperty{
//   					virtualNode: jsii.String("virtualNode"),
//   					weight: jsii.Number(123),
//
//   					// the properties below are optional
//   					port: jsii.Number(123),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		match: &tcpRouteMatchProperty{
//   			port: jsii.Number(123),
//   		},
//   		timeout: &tcpTimeoutProperty{
//   			idle: &durationProperty{
//   				unit: jsii.String("unit"),
//   				value: jsii.Number(123),
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

