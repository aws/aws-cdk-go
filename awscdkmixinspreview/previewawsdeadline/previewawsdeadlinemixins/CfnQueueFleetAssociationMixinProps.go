package previewawsdeadlinemixins


// Properties for CfnQueueFleetAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnQueueFleetAssociationMixinProps := &CfnQueueFleetAssociationMixinProps{
//   	FarmId: jsii.String("farmId"),
//   	FleetId: jsii.String("fleetId"),
//   	QueueId: jsii.String("queueId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuefleetassociation.html
//
type CfnQueueFleetAssociationMixinProps struct {
	// The identifier of the farm that contains the queue and the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuefleetassociation.html#cfn-deadline-queuefleetassociation-farmid
	//
	FarmId *string `field:"optional" json:"farmId" yaml:"farmId"`
	// The fleet ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuefleetassociation.html#cfn-deadline-queuefleetassociation-fleetid
	//
	FleetId *string `field:"optional" json:"fleetId" yaml:"fleetId"`
	// The queue ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuefleetassociation.html#cfn-deadline-queuefleetassociation-queueid
	//
	QueueId *string `field:"optional" json:"queueId" yaml:"queueId"`
}

