package interfacesawsconfig


// A reference to a DeliveryChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliveryChannelReference := &DeliveryChannelReference{
//   	DeliveryChannelName: jsii.String("deliveryChannelName"),
//   }
//
type DeliveryChannelReference struct {
	// The Name of the DeliveryChannel resource.
	DeliveryChannelName *string `field:"required" json:"deliveryChannelName" yaml:"deliveryChannelName"`
}

