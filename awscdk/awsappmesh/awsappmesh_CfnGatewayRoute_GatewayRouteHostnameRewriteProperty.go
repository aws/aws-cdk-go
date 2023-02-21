package awsappmesh


// An object representing the gateway route host name to rewrite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteHostnameRewriteProperty := &gatewayRouteHostnameRewriteProperty{
//   	defaultTargetHostname: jsii.String("defaultTargetHostname"),
//   }
//
type CfnGatewayRoute_GatewayRouteHostnameRewriteProperty struct {
	// The default target host name to write to.
	DefaultTargetHostname *string `field:"optional" json:"defaultTargetHostname" yaml:"defaultTargetHostname"`
}

