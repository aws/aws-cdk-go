package interfacesawslogs


// A reference to a Delivery resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliveryReference := &DeliveryReference{
//   	DeliveryArn: jsii.String("deliveryArn"),
//   	DeliveryId: jsii.String("deliveryId"),
//   }
//
type DeliveryReference struct {
	// The ARN of the Delivery resource.
	DeliveryArn *string `field:"required" json:"deliveryArn" yaml:"deliveryArn"`
	// The DeliveryId of the Delivery resource.
	DeliveryId *string `field:"required" json:"deliveryId" yaml:"deliveryId"`
}

