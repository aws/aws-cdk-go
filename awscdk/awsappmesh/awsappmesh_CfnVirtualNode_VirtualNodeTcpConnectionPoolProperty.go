package awsappmesh


// An object that represents a type of connection pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualNodeTcpConnectionPoolProperty := &virtualNodeTcpConnectionPoolProperty{
//   	maxConnections: jsii.Number(123),
//   }
//
type CfnVirtualNode_VirtualNodeTcpConnectionPoolProperty struct {
	// Maximum number of outbound TCP connections Envoy can establish concurrently with all hosts in upstream cluster.
	MaxConnections *float64 `field:"required" json:"maxConnections" yaml:"maxConnections"`
}

