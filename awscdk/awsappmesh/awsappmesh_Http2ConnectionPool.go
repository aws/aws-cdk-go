package awsappmesh


// Connection pool properties for HTTP2 listeners.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   http2ConnectionPool := &http2ConnectionPool{
//   	maxRequests: jsii.Number(123),
//   }
//
type Http2ConnectionPool struct {
	// The maximum requests in the pool.
	MaxRequests *float64 `field:"required" json:"maxRequests" yaml:"maxRequests"`
}

