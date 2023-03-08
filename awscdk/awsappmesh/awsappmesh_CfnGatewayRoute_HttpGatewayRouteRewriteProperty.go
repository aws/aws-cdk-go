package awsappmesh


// An object representing the gateway route to rewrite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpGatewayRouteRewriteProperty := &httpGatewayRouteRewriteProperty{
//   	hostname: &gatewayRouteHostnameRewriteProperty{
//   		defaultTargetHostname: jsii.String("defaultTargetHostname"),
//   	},
//   	path: &httpGatewayRoutePathRewriteProperty{
//   		exact: jsii.String("exact"),
//   	},
//   	prefix: &httpGatewayRoutePrefixRewriteProperty{
//   		defaultPrefix: jsii.String("defaultPrefix"),
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnGatewayRoute_HttpGatewayRouteRewriteProperty struct {
	// The host name to rewrite.
	Hostname interface{} `field:"optional" json:"hostname" yaml:"hostname"`
	// The path to rewrite.
	Path interface{} `field:"optional" json:"path" yaml:"path"`
	// The specified beginning characters to rewrite.
	Prefix interface{} `field:"optional" json:"prefix" yaml:"prefix"`
}

