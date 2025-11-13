package interfacesawslogs


// A reference to a DeliveryDestination resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliveryDestinationReference := &DeliveryDestinationReference{
//   	DeliveryDestinationArn: jsii.String("deliveryDestinationArn"),
//   	DeliveryDestinationName: jsii.String("deliveryDestinationName"),
//   }
//
type DeliveryDestinationReference struct {
	// The ARN of the DeliveryDestination resource.
	DeliveryDestinationArn *string `field:"required" json:"deliveryDestinationArn" yaml:"deliveryDestinationArn"`
	// The Name of the DeliveryDestination resource.
	DeliveryDestinationName *string `field:"required" json:"deliveryDestinationName" yaml:"deliveryDestinationName"`
}

