package interfacesawsdeadline


// A reference to a QueueLimitAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueLimitAssociationReference := &QueueLimitAssociationReference{
//   	FarmId: jsii.String("farmId"),
//   	LimitId: jsii.String("limitId"),
//   	QueueId: jsii.String("queueId"),
//   }
//
type QueueLimitAssociationReference struct {
	// The FarmId of the QueueLimitAssociation resource.
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// The LimitId of the QueueLimitAssociation resource.
	LimitId *string `field:"required" json:"limitId" yaml:"limitId"`
	// The QueueId of the QueueLimitAssociation resource.
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

