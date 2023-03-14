package awsappmesh


// An object that represents an HTTP or HTTP/2 route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpRouteProperty := &HttpRouteProperty{
//   	Action: &HttpRouteActionProperty{
//   		WeightedTargets: []interface{}{
//   			&WeightedTargetProperty{
//   				VirtualNode: jsii.String("virtualNode"),
//   				Weight: jsii.Number(123),
//
//   				// the properties below are optional
//   				Port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Match: &HttpRouteMatchProperty{
//   		Headers: []interface{}{
//   			&HttpRouteHeaderProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Invert: jsii.Boolean(false),
//   				Match: &HeaderMatchMethodProperty{
//   					Exact: jsii.String("exact"),
//   					Prefix: jsii.String("prefix"),
//   					Range: &MatchRangeProperty{
//   						End: jsii.Number(123),
//   						Start: jsii.Number(123),
//   					},
//   					Regex: jsii.String("regex"),
//   					Suffix: jsii.String("suffix"),
//   				},
//   			},
//   		},
//   		Method: jsii.String("method"),
//   		Path: &HttpPathMatchProperty{
//   			Exact: jsii.String("exact"),
//   			Regex: jsii.String("regex"),
//   		},
//   		Port: jsii.Number(123),
//   		Prefix: jsii.String("prefix"),
//   		QueryParameters: []interface{}{
//   			&QueryParameterProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Match: &HttpQueryParameterMatchProperty{
//   					Exact: jsii.String("exact"),
//   				},
//   			},
//   		},
//   		Scheme: jsii.String("scheme"),
//   	},
//
//   	// the properties below are optional
//   	RetryPolicy: &HttpRetryPolicyProperty{
//   		MaxRetries: jsii.Number(123),
//   		PerRetryTimeout: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		HttpRetryEvents: []*string{
//   			jsii.String("httpRetryEvents"),
//   		},
//   		TcpRetryEvents: []*string{
//   			jsii.String("tcpRetryEvents"),
//   		},
//   	},
//   	Timeout: &HttpTimeoutProperty{
//   		Idle: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		PerRequest: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
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

