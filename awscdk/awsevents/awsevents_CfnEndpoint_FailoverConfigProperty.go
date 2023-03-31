package awsevents


// The failover configuration for an endpoint.
//
// This includes what triggers failover and what happens when it's triggered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failoverConfigProperty := &failoverConfigProperty{
//   	primary: &primaryProperty{
//   		healthCheck: jsii.String("healthCheck"),
//   	},
//   	secondary: &secondaryProperty{
//   		route: jsii.String("route"),
//   	},
//   }
//
type CfnEndpoint_FailoverConfigProperty struct {
	// The main Region of the endpoint.
	Primary interface{} `field:"required" json:"primary" yaml:"primary"`
	// The Region that events are routed to when failover is triggered or event replication is enabled.
	Secondary interface{} `field:"required" json:"secondary" yaml:"secondary"`
}

