package awselasticache


// The configuration details of the Kinesis Data Firehose destination.
//
// Note that this field is marked as required but only if Kinesis Data Firehose was chosen as the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisFirehoseDestinationDetailsProperty := &kinesisFirehoseDestinationDetailsProperty{
//   	deliveryStream: jsii.String("deliveryStream"),
//   }
//
type CfnCacheCluster_KinesisFirehoseDestinationDetailsProperty struct {
	// The name of the Kinesis Data Firehose delivery stream.
	DeliveryStream *string `field:"required" json:"deliveryStream" yaml:"deliveryStream"`
}

