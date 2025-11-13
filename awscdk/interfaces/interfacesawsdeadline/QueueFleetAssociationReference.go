package interfacesawsdeadline


// A reference to a QueueFleetAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueFleetAssociationReference := &QueueFleetAssociationReference{
//   	FarmId: jsii.String("farmId"),
//   	FleetId: jsii.String("fleetId"),
//   	QueueId: jsii.String("queueId"),
//   }
//
type QueueFleetAssociationReference struct {
	// The FarmId of the QueueFleetAssociation resource.
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// The FleetId of the QueueFleetAssociation resource.
	FleetId *string `field:"required" json:"fleetId" yaml:"fleetId"`
	// The QueueId of the QueueFleetAssociation resource.
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

