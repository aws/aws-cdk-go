package awsappmesh


// Configuration for gateway route host name match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteHostnameMatchConfig := &GatewayRouteHostnameMatchConfig{
//   	HostnameMatch: &GatewayRouteHostnameMatchProperty{
//   		Exact: jsii.String("exact"),
//   		Suffix: jsii.String("suffix"),
//   	},
//   }
//
// Experimental.
type GatewayRouteHostnameMatchConfig struct {
	// GatewayRoute CFN configuration for host name match.
	// Experimental.
	HostnameMatch *CfnGatewayRoute_GatewayRouteHostnameMatchProperty `field:"required" json:"hostnameMatch" yaml:"hostnameMatch"`
}

