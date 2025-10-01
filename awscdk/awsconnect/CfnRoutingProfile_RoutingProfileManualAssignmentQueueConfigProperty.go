package awsconnect


// Contains information about the queue and channel for manual assignment behaviour can be enabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingProfileManualAssignmentQueueConfigProperty := &RoutingProfileManualAssignmentQueueConfigProperty{
//   	QueueReference: &RoutingProfileQueueReferenceProperty{
//   		Channel: jsii.String("channel"),
//   		QueueArn: jsii.String("queueArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-routingprofilemanualassignmentqueueconfig.html
//
type CfnRoutingProfile_RoutingProfileManualAssignmentQueueConfigProperty struct {
	// Contains information about a queue resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-routingprofile-routingprofilemanualassignmentqueueconfig.html#cfn-connect-routingprofile-routingprofilemanualassignmentqueueconfig-queuereference
	//
	QueueReference interface{} `field:"required" json:"queueReference" yaml:"queueReference"`
}

