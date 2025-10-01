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
//   	Name: jsii.String("name"),
//   	ReplicationConfig: &ReplicationConfigProperty{
//   		State: jsii.String("state"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-endpoint.html
//
type CfnEndpointProps struct {
	// The event buses being used by the endpoint.
	//
	// *Exactly* : `2`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-endpoint.html#cfn-events-endpoint-eventbuses
	//
	EventBuses interface{} `field:"required" json:"eventBuses" yaml:"eventBuses"`
	// The routing configuration of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-endpoint.html#cfn-events-endpoint-routingconfig
	//
	RoutingConfig interface{} `field:"required" json:"routingConfig" yaml:"routingConfig"`
	// A description for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-endpoint.html#cfn-events-endpoint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-endpoint.html#cfn-events-endpoint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether event replication was enabled or disabled for this endpoint.
	//
	// The default state is `ENABLED` which means you must supply a `RoleArn` . If you don't have a `RoleArn` or you don't want event replication enabled, set the state to `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-endpoint.html#cfn-events-endpoint-replicationconfig
	//
	ReplicationConfig interface{} `field:"optional" json:"replicationConfig" yaml:"replicationConfig"`
	// The ARN of the role used by event replication for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-endpoint.html#cfn-events-endpoint-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

