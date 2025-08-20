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
//   failoverConfigProperty := &FailoverConfigProperty{
//   	Primary: &PrimaryProperty{
//   		HealthCheck: jsii.String("healthCheck"),
//   	},
//   	Secondary: &SecondaryProperty{
//   		Route: jsii.String("route"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-endpoint-failoverconfig.html
//
type CfnEndpoint_FailoverConfigProperty struct {
	// The main Region of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-endpoint-failoverconfig.html#cfn-events-endpoint-failoverconfig-primary
	//
	Primary interface{} `field:"required" json:"primary" yaml:"primary"`
	// The Region that events are routed to when failover is triggered or event replication is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-endpoint-failoverconfig.html#cfn-events-endpoint-failoverconfig-secondary
	//
	Secondary interface{} `field:"required" json:"secondary" yaml:"secondary"`
}

