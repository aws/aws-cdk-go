package awsevents


// Properties for defining a `CfnEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointProps := &CfnEndpointProps{
//   	EventBuses: []interface{}{
//   		&EndpointEventBusProperty{
//   			EventBusArn: jsii.String("eventBusArn"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RoutingConfig: &RoutingConfigProperty{
//   		FailoverConfig: &FailoverConfigProperty{
//   			Primary: &PrimaryProperty{
//   				HealthCheck: jsii.String("healthCheck"),
//   			},
//   			Secondary: &SecondaryProperty{
//   				Route: jsii.String("route"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ReplicationConfig: &ReplicationConfigProperty{
//   		State: jsii.String("state"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
type CfnEndpointProps struct {
	// The event buses being used by the endpoint.
	//
	// *Exactly* : `2`.
	EventBuses interface{} `field:"required" json:"eventBuses" yaml:"eventBuses"`
	// The name of the endpoint.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The routing configuration of the endpoint.
	RoutingConfig interface{} `field:"required" json:"routingConfig" yaml:"routingConfig"`
	// A description for the endpoint.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether event replication was enabled or disabled for this endpoint.
	//
	// The default state is `ENABLED` which means you must supply a `RoleArn` . If you don't have a `RoleArn` or you don't want event replication enabled, set the state to `DISABLED` .
	ReplicationConfig interface{} `field:"optional" json:"replicationConfig" yaml:"replicationConfig"`
	// The ARN of the role used by event replication for the endpoint.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

