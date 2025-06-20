package awsappmesh


// An object that represents a retry policy.
//
// Specify at least one value for at least one of the types of `RetryEvents` , a value for `maxRetries` , and a value for `perRetryTimeout` . Both `server-error` and `gateway-error` under `httpRetryEvents` include the Envoy `reset` policy. For more information on the `reset` policy, see the [Envoy documentation](https://docs.aws.amazon.com/https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-on) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpRetryPolicyProperty := &HttpRetryPolicyProperty{
//   	MaxRetries: jsii.Number(123),
//   	PerRetryTimeout: &DurationProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	HttpRetryEvents: []*string{
//   		jsii.String("httpRetryEvents"),
//   	},
//   	TcpRetryEvents: []*string{
//   		jsii.String("tcpRetryEvents"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html
//
type CfnRoute_HttpRetryPolicyProperty struct {
	// The maximum number of retry attempts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html#cfn-appmesh-route-httpretrypolicy-maxretries
	//
	MaxRetries *float64 `field:"required" json:"maxRetries" yaml:"maxRetries"`
	// The timeout for each retry attempt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html#cfn-appmesh-route-httpretrypolicy-perretrytimeout
	//
	PerRetryTimeout interface{} `field:"required" json:"perRetryTimeout" yaml:"perRetryTimeout"`
	// Specify at least one of the following values.
	//
	// - *server-error* – HTTP status codes 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, and 511
	// - *gateway-error* – HTTP status codes 502, 503, and 504
	// - *client-error* – HTTP status code 409
	// - *stream-error* – Retry on refused stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html#cfn-appmesh-route-httpretrypolicy-httpretryevents
	//
	HttpRetryEvents *[]*string `field:"optional" json:"httpRetryEvents" yaml:"httpRetryEvents"`
	// Specify a valid value.
	//
	// The event occurs before any processing of a request has started and is encountered when the upstream is temporarily or permanently unavailable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-httpretrypolicy.html#cfn-appmesh-route-httpretrypolicy-tcpretryevents
	//
	TcpRetryEvents *[]*string `field:"optional" json:"tcpRetryEvents" yaml:"tcpRetryEvents"`
}

