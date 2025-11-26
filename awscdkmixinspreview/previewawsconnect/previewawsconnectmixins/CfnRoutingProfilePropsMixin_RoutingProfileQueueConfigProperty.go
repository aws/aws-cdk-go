package previewawsconnectmixins


// Contains information about the queue and channel for which priority and delay can be set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnRoutingProfilePropsMixin_RoutingProfileQueueConfigProperty struct {
	// The delay, in seconds, a contact should be in the queue before they are routed to an available agent.
	//
	// For more information, see [Queues: priority and delay](https://docs.aws.amazon.com/connect/latest/adminguide/concepts-routing-profiles-priority.html) in the *Amazon Connect Administrator Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-routingprofilequeueconfig.html#cfn-connect-routingprofile-routingprofilequeueconfig-delay
	//
	Delay *float64 `field:"optional" json:"delay" yaml:"delay"`
	// The order in which contacts are to be handled for the queue.
	//
	// For more information, see [Queues: priority and delay](https://docs.aws.amazon.com/connect/latest/adminguide/concepts-routing-profiles-priority.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-routingprofilequeueconfig.html#cfn-connect-routingprofile-routingprofilequeueconfig-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Contains information about a queue resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-routingprofilequeueconfig.html#cfn-connect-routingprofile-routingprofilequeueconfig-queuereference
	//
	QueueReference interface{} `field:"optional" json:"queueReference" yaml:"queueReference"`
}

