package awsevents


// The routing configuration of the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingConfigProperty := &routingConfigProperty{
//   	failoverConfig: &failoverConfigProperty{
//   		primary: &primaryProperty{
//   			healthCheck: jsii.String("healthCheck"),
//   		},
//   		secondary: &secondaryProperty{
//   			route: jsii.String("route"),
//   		},
//   	},
//   }
//
type CfnEndpoint_RoutingConfigProperty struct {
	// The failover configuration for an endpoint.
	//
	// This includes what triggers failover and what happens when it's triggered.
	FailoverConfig interface{} `field:"required" json:"failoverConfig" yaml:"failoverConfig"`
}

