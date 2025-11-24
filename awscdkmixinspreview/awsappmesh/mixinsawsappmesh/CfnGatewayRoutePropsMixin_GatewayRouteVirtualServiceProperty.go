package mixinsawsappmesh


// An object that represents the virtual service that traffic is routed to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gatewayRouteVirtualServiceProperty := &GatewayRouteVirtualServiceProperty{
//   	VirtualServiceName: jsii.String("virtualServiceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutevirtualservice.html
//
type CfnGatewayRoutePropsMixin_GatewayRouteVirtualServiceProperty struct {
	// The name of the virtual service that traffic is routed to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutevirtualservice.html#cfn-appmesh-gatewayroute-gatewayroutevirtualservice-virtualservicename
	//
	VirtualServiceName *string `field:"optional" json:"virtualServiceName" yaml:"virtualServiceName"`
}

