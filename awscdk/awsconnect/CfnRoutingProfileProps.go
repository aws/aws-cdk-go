package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRoutingProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRoutingProfileProps := &CfnRoutingProfileProps{
//   	DefaultOutboundQueueArn: jsii.String("defaultOutboundQueueArn"),
//   	Description: jsii.String("description"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	MediaConcurrencies: []interface{}{
//   		&MediaConcurrencyProperty{
//   			Channel: jsii.String("channel"),
//   			Concurrency: jsii.Number(123),
//
//   			// the properties below are optional
//   			CrossChannelBehavior: &CrossChannelBehaviorProperty{
//   				BehaviorType: jsii.String("behaviorType"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AgentAvailabilityTimer: jsii.String("agentAvailabilityTimer"),
//   	ManualAssignmentQueueConfigs: []interface{}{
//   		&RoutingProfileManualAssignmentQueueConfigProperty{
//   			QueueReference: &RoutingProfileQueueReferenceProperty{
//   				Channel: jsii.String("channel"),
//   				QueueArn: jsii.String("queueArn"),
//   			},
//   		},
//   	},
//   	QueueConfigs: []interface{}{
//   		&RoutingProfileQueueConfigProperty{
//   			Delay: jsii.Number(123),
//   			Priority: jsii.Number(123),
//   			QueueReference: &RoutingProfileQueueReferenceProperty{
//   				Channel: jsii.String("channel"),
//   				QueueArn: jsii.String("queueArn"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-routingprofile.html
//
type CfnRoutingProfileProps struct {
	// The Amazon Resource Name (ARN) of the default outbound queue for the routing profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-routingprofile.html#cfn-connect-routingprofile-defaultoutboundqueuearn
	//
	DefaultOutboundQueueArn *string `field:"required" json:"defaultOutboundQueueArn" yaml:"defaultOutboundQueueArn"`
	// The description of the routing profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-routingprofile.html#cfn-connect-routingprofile-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-routingprofile.html#cfn-connect-routingprofile-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The channels agents can handle in the Contact Control Panel (CCP) for this routing profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-routingprofile.html#cfn-connect-routingprofile-mediaconcurrencies
	//
	MediaConcurrencies interface{} `field:"required" json:"mediaConcurrencies" yaml:"mediaConcurrencies"`
	// The name of the routing profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-routingprofile.html#cfn-connect-routingprofile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether agents with this routing profile will have their routing order calculated based on *time since their last inbound contact* or *longest idle time* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-routingprofile.html#cfn-connect-routingprofile-agentavailabilitytimer
	//
	AgentAvailabilityTimer *string `field:"optional" json:"agentAvailabilityTimer" yaml:"agentAvailabilityTimer"`
	// Contains information about the queue and channel for manual assignment behaviour can be enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-routingprofile.html#cfn-connect-routingprofile-manualassignmentqueueconfigs
	//
	ManualAssignmentQueueConfigs interface{} `field:"optional" json:"manualAssignmentQueueConfigs" yaml:"manualAssignmentQueueConfigs"`
	// The inbound queues associated with the routing profile.
	//
	// If no queue is added, the agent can make only outbound calls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-routingprofile.html#cfn-connect-routingprofile-queueconfigs
	//
	QueueConfigs interface{} `field:"optional" json:"queueConfigs" yaml:"queueConfigs"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "Tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-routingprofile.html#cfn-connect-routingprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

