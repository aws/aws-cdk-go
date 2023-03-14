package awsappmesh


// An object that represents the type of virtual node connection pool.
//
// Only one protocol is used at a time and should be the same protocol as the one chosen under port mapping.
//
// If not present the default value for `maxPendingRequests` is `2147483647` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualNodeConnectionPoolProperty := &virtualNodeConnectionPoolProperty{
//   	grpc: &virtualNodeGrpcConnectionPoolProperty{
//   		maxRequests: jsii.Number(123),
//   	},
//   	http: &virtualNodeHttpConnectionPoolProperty{
//   		maxConnections: jsii.Number(123),
//
//   		// the properties below are optional
//   		maxPendingRequests: jsii.Number(123),
//   	},
//   	http2: &virtualNodeHttp2ConnectionPoolProperty{
//   		maxRequests: jsii.Number(123),
//   	},
//   	tcp: &virtualNodeTcpConnectionPoolProperty{
//   		maxConnections: jsii.Number(123),
//   	},
//   }
//
type CfnVirtualNode_VirtualNodeConnectionPoolProperty struct {
	// An object that represents a type of connection pool.
	Grpc interface{} `field:"optional" json:"grpc" yaml:"grpc"`
	// An object that represents a type of connection pool.
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// An object that represents a type of connection pool.
	Http2 interface{} `field:"optional" json:"http2" yaml:"http2"`
	// An object that represents a type of connection pool.
	Tcp interface{} `field:"optional" json:"tcp" yaml:"tcp"`
}

