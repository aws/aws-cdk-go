package awsconnect


// Contains information about the queue and channel for which priority and delay can be set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingProfileQueueConfigProperty := &RoutingProfileQueueConfigProperty{
//   	Delay: jsii.Number(123),
//   	Priority: jsii.Number(123),
//   	QueueReference: &RoutingProfileQueueReferenceProperty{
//   		Channel: jsii.String("channel"),
//   		QueueArn: jsii.String("queueArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-routingprofilequeueconfig.html
//
type CfnRoutingProfile_RoutingProfileQueueConfigProperty struct {
	// The delay, in seconds, a contact should wait in the queue before they are routed to an available agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-routingprofilequeueconfig.html#cfn-connect-routingprofile-routingprofilequeueconfig-delay
	//
	Delay *float64 `field:"required" json:"delay" yaml:"delay"`
	// The order in which contacts are to be handled for the queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-routingprofilequeueconfig.html#cfn-connect-routingprofile-routingprofilequeueconfig-priority
	//
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Contains the channel and queue identifier for a routing profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-routingprofilequeueconfig.html#cfn-connect-routingprofile-routingprofilequeueconfig-queuereference
	//
	QueueReference interface{} `field:"required" json:"queueReference" yaml:"queueReference"`
}

