package mixinsawsappmesh


// An object that represents the type of virtual node connection pool.
//
// Only one protocol is used at a time and should be the same protocol as the one chosen under port mapping.
//
// If not present the default value for `maxPendingRequests` is `2147483647` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualNodeConnectionPoolProperty := &VirtualNodeConnectionPoolProperty{
//   	Grpc: &VirtualNodeGrpcConnectionPoolProperty{
//   		MaxRequests: jsii.Number(123),
//   	},
//   	Http: &VirtualNodeHttpConnectionPoolProperty{
//   		MaxConnections: jsii.Number(123),
//   		MaxPendingRequests: jsii.Number(123),
//   	},
//   	Http2: &VirtualNodeHttp2ConnectionPoolProperty{
//   		MaxRequests: jsii.Number(123),
//   	},
//   	Tcp: &VirtualNodeTcpConnectionPoolProperty{
//   		MaxConnections: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodeconnectionpool.html
//
type CfnVirtualNodePropsMixin_VirtualNodeConnectionPoolProperty struct {
	// An object that represents a type of connection pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodeconnectionpool.html#cfn-appmesh-virtualnode-virtualnodeconnectionpool-grpc
	//
	Grpc interface{} `field:"optional" json:"grpc" yaml:"grpc"`
	// An object that represents a type of connection pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodeconnectionpool.html#cfn-appmesh-virtualnode-virtualnodeconnectionpool-http
	//
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// An object that represents a type of connection pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodeconnectionpool.html#cfn-appmesh-virtualnode-virtualnodeconnectionpool-http2
	//
	Http2 interface{} `field:"optional" json:"http2" yaml:"http2"`
	// An object that represents a type of connection pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-virtualnodeconnectionpool.html#cfn-appmesh-virtualnode-virtualnodeconnectionpool-tcp
	//
	Tcp interface{} `field:"optional" json:"tcp" yaml:"tcp"`
}

