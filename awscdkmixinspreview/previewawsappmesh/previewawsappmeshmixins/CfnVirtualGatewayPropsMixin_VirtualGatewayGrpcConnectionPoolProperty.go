package previewawsappmeshmixins


// An object that represents a type of connection pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualGatewayGrpcConnectionPoolProperty := &VirtualGatewayGrpcConnectionPoolProperty{
//   	MaxRequests: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaygrpcconnectionpool.html
//
type CfnVirtualGatewayPropsMixin_VirtualGatewayGrpcConnectionPoolProperty struct {
	// Maximum number of inflight requests Envoy can concurrently support across hosts in upstream cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaygrpcconnectionpool.html#cfn-appmesh-virtualgateway-virtualgatewaygrpcconnectionpool-maxrequests
	//
	MaxRequests *float64 `field:"optional" json:"maxRequests" yaml:"maxRequests"`
}

