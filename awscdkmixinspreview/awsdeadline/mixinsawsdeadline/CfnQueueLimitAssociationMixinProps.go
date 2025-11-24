package mixinsawsdeadline


// Properties for CfnQueueLimitAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnQueueLimitAssociationMixinProps := &CfnQueueLimitAssociationMixinProps{
//   	FarmId: jsii.String("farmId"),
//   	LimitId: jsii.String("limitId"),
//   	QueueId: jsii.String("queueId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuelimitassociation.html
//
type CfnQueueLimitAssociationMixinProps struct {
	// The unique identifier of the farm that contains the queue-limit association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuelimitassociation.html#cfn-deadline-queuelimitassociation-farmid
	//
	FarmId *string `field:"optional" json:"farmId" yaml:"farmId"`
	// The unique identifier of the limit in the association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuelimitassociation.html#cfn-deadline-queuelimitassociation-limitid
	//
	LimitId *string `field:"optional" json:"limitId" yaml:"limitId"`
	// The unique identifier of the queue in the association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queuelimitassociation.html#cfn-deadline-queuelimitassociation-queueid
	//
	QueueId *string `field:"optional" json:"queueId" yaml:"queueId"`
}

