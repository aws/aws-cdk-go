package awsdeadline


// Properties for defining a `CfnQueueFleetAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnQueueFleetAssociationProps := &CfnQueueFleetAssociationProps{
//   	FarmId: jsii.String("farmId"),
//   	FleetId: jsii.String("fleetId"),
//   	QueueId: jsii.String("queueId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuefleetassociation.html
//
type CfnQueueFleetAssociationProps struct {
	// The identifier of the farm that contains the queue and the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuefleetassociation.html#cfn-deadline-queuefleetassociation-farmid
	//
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// The fleet ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuefleetassociation.html#cfn-deadline-queuefleetassociation-fleetid
	//
	FleetId *string `field:"required" json:"fleetId" yaml:"fleetId"`
	// The queue ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuefleetassociation.html#cfn-deadline-queuefleetassociation-queueid
	//
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

