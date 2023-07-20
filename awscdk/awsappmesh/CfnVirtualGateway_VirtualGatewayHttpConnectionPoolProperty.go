package awsappmesh


// An object that represents a type of connection pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayHttpConnectionPoolProperty := &VirtualGatewayHttpConnectionPoolProperty{
//   	MaxConnections: jsii.Number(123),
//
//   	// the properties below are optional
//   	MaxPendingRequests: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhttpconnectionpool.html
//
type CfnVirtualGateway_VirtualGatewayHttpConnectionPoolProperty struct {
	// Maximum number of outbound TCP connections Envoy can establish concurrently with all hosts in upstream cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhttpconnectionpool.html#cfn-appmesh-virtualgateway-virtualgatewayhttpconnectionpool-maxconnections
	//
	MaxConnections *float64 `field:"required" json:"maxConnections" yaml:"maxConnections"`
	// Number of overflowing requests after `max_connections` Envoy will queue to upstream cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhttpconnectionpool.html#cfn-appmesh-virtualgateway-virtualgatewayhttpconnectionpool-maxpendingrequests
	//
	MaxPendingRequests *float64 `field:"optional" json:"maxPendingRequests" yaml:"maxPendingRequests"`
}

