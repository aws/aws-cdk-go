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
//   grpcRetryPolicyProperty := &grpcRetryPolicyProperty{
//   	maxRetries: jsii.Number(123),
//   	perRetryTimeout: &durationProperty{
//   		unit: jsii.String("unit"),
//   		value: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	grpcRetryEvents: []*string{
//   		jsii.String("grpcRetryEvents"),
//   	},
//   	httpRetryEvents: []*string{
//   		jsii.String("httpRetryEvents"),
//   	},
//   	tcpRetryEvents: []*string{
//   		jsii.String("tcpRetryEvents"),
//   	},
//   }
//
type CfnRoute_GrpcRetryPolicyProperty struct {
	// The maximum number of retry attempts.
	MaxRetries *float64 `field:"required" json:"maxRetries" yaml:"maxRetries"`
	// The timeout for each retry attempt.
	PerRetryTimeout interface{} `field:"required" json:"perRetryTimeout" yaml:"perRetryTimeout"`
	// Specify at least one of the valid values.
	GrpcRetryEvents *[]*string `field:"optional" json:"grpcRetryEvents" yaml:"grpcRetryEvents"`
	// Specify at least one of the following values.
	//
	// - *server-error* – HTTP status codes 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, and 511
	// - *gateway-error* – HTTP status codes 502, 503, and 504
	// - *client-error* – HTTP status code 409
	// - *stream-error* – Retry on refused stream.
	HttpRetryEvents *[]*string `field:"optional" json:"httpRetryEvents" yaml:"httpRetryEvents"`
	// Specify a valid value.
	//
	// The event occurs before any processing of a request has started and is encountered when the upstream is temporarily or permanently unavailable.
	TcpRetryEvents *[]*string `field:"optional" json:"tcpRetryEvents" yaml:"tcpRetryEvents"`
}

