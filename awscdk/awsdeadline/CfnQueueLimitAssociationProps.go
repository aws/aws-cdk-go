package awsdeadline


// Properties for defining a `CfnQueueLimitAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnQueueLimitAssociationProps := &CfnQueueLimitAssociationProps{
//   	FarmId: jsii.String("farmId"),
//   	LimitId: jsii.String("limitId"),
//   	QueueId: jsii.String("queueId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuelimitassociation.html
//
type CfnQueueLimitAssociationProps struct {
	// The unique identifier of the farm that contains the queue-limit association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuelimitassociation.html#cfn-deadline-queuelimitassociation-farmid
	//
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// The unique identifier of the limit in the association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuelimitassociation.html#cfn-deadline-queuelimitassociation-limitid
	//
	LimitId *string `field:"required" json:"limitId" yaml:"limitId"`
	// The unique identifier of the queue in the association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuelimitassociation.html#cfn-deadline-queuelimitassociation-queueid
	//
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

