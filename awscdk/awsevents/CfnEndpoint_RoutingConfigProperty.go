package awsevents


// The routing configuration of the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingConfigProperty := &RoutingConfigProperty{
//   	FailoverConfig: &FailoverConfigProperty{
//   		Primary: &PrimaryProperty{
//   			HealthCheck: jsii.String("healthCheck"),
//   		},
//   		Secondary: &SecondaryProperty{
//   			Route: jsii.String("route"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-endpoint-routingconfig.html
//
type CfnEndpoint_RoutingConfigProperty struct {
	// The failover configuration for an endpoint.
	//
	// This includes what triggers failover and what happens when it's triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-endpoint-routingconfig.html#cfn-events-endpoint-routingconfig-failoverconfig
	//
	FailoverConfig interface{} `field:"required" json:"failoverConfig" yaml:"failoverConfig"`
}

