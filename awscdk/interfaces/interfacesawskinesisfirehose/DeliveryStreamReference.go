package interfacesawskinesisfirehose


// A reference to a DeliveryStream resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliveryStreamReference := &DeliveryStreamReference{
//   	DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   	DeliveryStreamName: jsii.String("deliveryStreamName"),
//   }
//
type DeliveryStreamReference struct {
	// The ARN of the DeliveryStream resource.
	DeliveryStreamArn *string `field:"required" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// The DeliveryStreamName of the DeliveryStream resource.
	DeliveryStreamName *string `field:"required" json:"deliveryStreamName" yaml:"deliveryStreamName"`
}

