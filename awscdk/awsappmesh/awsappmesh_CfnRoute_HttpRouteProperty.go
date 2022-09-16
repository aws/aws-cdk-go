package awsappmesh


// An object that represents an HTTP or HTTP/2 route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpRouteProperty := &httpRouteProperty{
//   	action: &httpRouteActionProperty{
//   		weightedTargets: []interface{}{
//   			&weightedTargetProperty{
//   				virtualNode: jsii.String("virtualNode"),
//   				weight: jsii.Number(123),
//
//   				// the properties below are optional
//   				port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	match: &httpRouteMatchProperty{
//   		headers: []interface{}{
//   			&httpRouteHeaderProperty{
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				invert: jsii.Boolean(false),
//   				match: &headerMatchMethodProperty{
//   					exact: jsii.String("exact"),
//   					prefix: jsii.String("prefix"),
//   					range: &matchRangeProperty{
//   						end: jsii.Number(123),
//   						start: jsii.Number(123),
//   					},
//   					regex: jsii.String("regex"),
//   					suffix: jsii.String("suffix"),
//   				},
//   			},
//   		},
//   		method: jsii.String("method"),
//   		path: &httpPathMatchProperty{
//   			exact: jsii.String("exact"),
//   			regex: jsii.String("regex"),
//   		},
//   		port: jsii.Number(123),
//   		prefix: jsii.String("prefix"),
//   		queryParameters: []interface{}{
//   			&queryParameterProperty{
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				match: &httpQueryParameterMatchProperty{
//   					exact: jsii.String("exact"),
//   				},
//   			},
//   		},
//   		scheme: jsii.String("scheme"),
//   	},
//
//   	// the properties below are optional
//   	retryPolicy: &httpRetryPolicyProperty{
//   		maxRetries: jsii.Number(123),
//   		perRetryTimeout: &durationProperty{
//   			unit: jsii.String("unit"),
//   			value: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		httpRetryEvents: []*string{
//   			jsii.String("httpRetryEvents"),
//   		},
//   		tcpRetryEvents: []*string{
//   			jsii.String("tcpRetryEvents"),
//   		},
//   	},
//   	timeout: &httpTimeoutProperty{
//   		idle: &durationProperty{
//   			unit: jsii.String("unit"),
//   			value: jsii.Number(123),
//   		},
//   		perRequest: &durationProperty{
//   			unit: jsii.String("unit"),
//   			value: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnRoute_HttpRouteProperty struct {
	// An object that represents the action to take if a match is determined.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	Match interface{} `field:"required" json:"match" yaml:"match"`
	// An object that represents a retry policy.
	RetryPolicy interface{} `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// An object that represents types of timeouts.
	Timeout interface{} `field:"optional" json:"timeout" yaml:"timeout"`
}

