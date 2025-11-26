package previewawsappmeshmixins


// An object that represents a gateway route target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gatewayRouteTargetProperty := &GatewayRouteTargetProperty{
//   	Port: jsii.Number(123),
//   	VirtualService: &GatewayRouteVirtualServiceProperty{
//   		VirtualServiceName: jsii.String("virtualServiceName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutetarget.html
//
type CfnGatewayRoutePropsMixin_GatewayRouteTargetProperty struct {
	// The port number of the gateway route target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutetarget.html#cfn-appmesh-gatewayroute-gatewayroutetarget-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// An object that represents a virtual service gateway route target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutetarget.html#cfn-appmesh-gatewayroute-gatewayroutetarget-virtualservice
	//
	VirtualService interface{} `field:"optional" json:"virtualService" yaml:"virtualService"`
}

