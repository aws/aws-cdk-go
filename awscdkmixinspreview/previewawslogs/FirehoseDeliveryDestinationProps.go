package previewawslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
)

// Properties for Firehose delivery destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var deliveryStreamRef IDeliveryStreamRef
//
//   firehoseDeliveryDestinationProps := &FirehoseDeliveryDestinationProps{
//   	DeliveryStream: deliveryStreamRef,
//
//   	// the properties below are optional
//   	OutputFormat: jsii.String("outputFormat"),
//   	SourceAccountId: jsii.String("sourceAccountId"),
//   }
//
// Experimental.
type FirehoseDeliveryDestinationProps struct {
	// Delivery stream to delivery logs to.
	// Experimental.
	DeliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef `field:"required" json:"deliveryStream" yaml:"deliveryStream"`
	// Format of the logs that are sent to this delivery destination.
	// Experimental.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Optional acount id for account the delivery source is in for cross account Vended Logs.
	// Experimental.
	SourceAccountId *string `field:"optional" json:"sourceAccountId" yaml:"sourceAccountId"`
}

