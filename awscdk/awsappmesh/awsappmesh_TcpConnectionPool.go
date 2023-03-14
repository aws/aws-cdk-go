package awsappmesh


// Connection pool properties for TCP listeners.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tcpConnectionPool := &tcpConnectionPool{
//   	maxConnections: jsii.Number(123),
//   }
//
// Experimental.
type TcpConnectionPool struct {
	// The maximum connections in the pool.
	// Experimental.
	MaxConnections *float64 `field:"required" json:"maxConnections" yaml:"maxConnections"`
}

