package awsappmesh


// An object that represents the type of virtual gateway connection pool.
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
//   virtualGatewayConnectionPoolProperty := &VirtualGatewayConnectionPoolProperty{
//   	Grpc: &VirtualGatewayGrpcConnectionPoolProperty{
//   		MaxRequests: jsii.Number(123),
//   	},
//   	Http: &VirtualGatewayHttpConnectionPoolProperty{
//   		MaxConnections: jsii.Number(123),
//
//   		// the properties below are optional
//   		MaxPendingRequests: jsii.Number(123),
//   	},
//   	Http2: &VirtualGatewayHttp2ConnectionPoolProperty{
//   		MaxRequests: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayconnectionpool.html
//
type CfnVirtualGateway_VirtualGatewayConnectionPoolProperty struct {
	// An object that represents a type of connection pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayconnectionpool.html#cfn-appmesh-virtualgateway-virtualgatewayconnectionpool-grpc
	//
	Grpc interface{} `field:"optional" json:"grpc" yaml:"grpc"`
	// An object that represents a type of connection pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayconnectionpool.html#cfn-appmesh-virtualgateway-virtualgatewayconnectionpool-http
	//
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// An object that represents a type of connection pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayconnectionpool.html#cfn-appmesh-virtualgateway-virtualgatewayconnectionpool-http2
	//
	Http2 interface{} `field:"optional" json:"http2" yaml:"http2"`
}

