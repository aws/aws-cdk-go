package awsappmesh


// An object that represents a type of connection pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualNodeHttp2ConnectionPoolProperty := &virtualNodeHttp2ConnectionPoolProperty{
//   	maxRequests: jsii.Number(123),
//   }
//
type CfnVirtualNode_VirtualNodeHttp2ConnectionPoolProperty struct {
	// Maximum number of inflight requests Envoy can concurrently support across hosts in upstream cluster.
	MaxRequests *float64 `field:"required" json:"maxRequests" yaml:"maxRequests"`
}

