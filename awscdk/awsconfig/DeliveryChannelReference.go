package awsconfig


// A reference to a DeliveryChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliveryChannelReference := &DeliveryChannelReference{
//   	DeliveryChannelId: jsii.String("deliveryChannelId"),
//   }
//
type DeliveryChannelReference struct {
	// The Id of the DeliveryChannel resource.
	DeliveryChannelId *string `field:"required" json:"deliveryChannelId" yaml:"deliveryChannelId"`
}

