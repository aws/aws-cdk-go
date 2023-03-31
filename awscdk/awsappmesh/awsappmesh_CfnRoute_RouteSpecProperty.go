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
//   routeSpecProperty := &routeSpecProperty{
//   	grpcRoute: &grpcRouteProperty{
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
//   	http2Route: &httpRouteProperty{
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
//   	httpRoute: &httpRouteProperty{
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
//   	tcpRoute: &tcpRouteProperty{
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

