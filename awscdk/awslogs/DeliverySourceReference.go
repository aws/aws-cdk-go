package awslogs


// A reference to a DeliverySource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliverySourceReference := &DeliverySourceReference{
//   	DeliverySourceArn: jsii.String("deliverySourceArn"),
//   	DeliverySourceName: jsii.String("deliverySourceName"),
//   }
//
type DeliverySourceReference struct {
	// The ARN of the DeliverySource resource.
	DeliverySourceArn *string `field:"required" json:"deliverySourceArn" yaml:"deliverySourceArn"`
	// The Name of the DeliverySource resource.
	DeliverySourceName *string `field:"required" json:"deliverySourceName" yaml:"deliverySourceName"`
}

