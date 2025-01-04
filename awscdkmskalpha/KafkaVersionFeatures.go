package awscdkmskalpha


// Available features for a given Kafka version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import msk_alpha "github.com/aws/aws-cdk-go/awscdkmskalpha"
//
//   kafkaVersionFeatures := &KafkaVersionFeatures{
//   	TieredStorage: jsii.Boolean(false),
//   }
//
// Experimental.
type KafkaVersionFeatures struct {
	// Whether the Kafka version supports tiered storage mode.
	// See: https://docs.aws.amazon.com/msk/latest/developerguide/msk-tiered-storage.html#msk-tiered-storage-requirements
	//
	// Default: false.
	//
	// Experimental.
	TieredStorage *bool `field:"optional" json:"tieredStorage" yaml:"tieredStorage"`
}

