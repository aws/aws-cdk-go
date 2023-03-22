package awsappmesh


// An object representing the gateway route host name to match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteHostnameMatchProperty := &gatewayRouteHostnameMatchProperty{
//   	exact: jsii.String("exact"),
//   	suffix: jsii.String("suffix"),
//   }
//
type CfnGatewayRoute_GatewayRouteHostnameMatchProperty struct {
	// The exact host name to match on.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// The specified ending characters of the host name to match on.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

