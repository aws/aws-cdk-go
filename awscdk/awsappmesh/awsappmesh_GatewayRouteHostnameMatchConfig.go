package awsappmesh


// Configuration for gateway route host name match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteHostnameMatchConfig := &gatewayRouteHostnameMatchConfig{
//   	hostnameMatch: &gatewayRouteHostnameMatchProperty{
//   		exact: jsii.String("exact"),
//   		suffix: jsii.String("suffix"),
//   	},
//   }
//
type GatewayRouteHostnameMatchConfig struct {
	// GatewayRoute CFN configuration for host name match.
	HostnameMatch *CfnGatewayRoute_GatewayRouteHostnameMatchProperty `field:"required" json:"hostnameMatch" yaml:"hostnameMatch"`
}

