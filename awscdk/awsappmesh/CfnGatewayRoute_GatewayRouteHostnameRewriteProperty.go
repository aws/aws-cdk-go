package awsappmesh


// An object representing the gateway route host name to rewrite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteHostnameRewriteProperty := &GatewayRouteHostnameRewriteProperty{
//   	DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutehostnamerewrite.html
//
type CfnGatewayRoute_GatewayRouteHostnameRewriteProperty struct {
	// The default target host name to write to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutehostnamerewrite.html#cfn-appmesh-gatewayroute-gatewayroutehostnamerewrite-defaulttargethostname
	//
	DefaultTargetHostname *string `field:"optional" json:"defaultTargetHostname" yaml:"defaultTargetHostname"`
}

