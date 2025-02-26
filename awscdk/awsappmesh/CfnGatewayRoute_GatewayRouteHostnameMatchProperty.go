package awsappmesh


// An object representing the gateway route host name to match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteHostnameMatchProperty := &GatewayRouteHostnameMatchProperty{
//   	Exact: jsii.String("exact"),
//   	Suffix: jsii.String("suffix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutehostnamematch.html
//
type CfnGatewayRoute_GatewayRouteHostnameMatchProperty struct {
	// The exact host name to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutehostnamematch.html#cfn-appmesh-gatewayroute-gatewayroutehostnamematch-exact
	//
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// The specified ending characters of the host name to match on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-gatewayroute-gatewayroutehostnamematch.html#cfn-appmesh-gatewayroute-gatewayroutehostnamematch-suffix
	//
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

