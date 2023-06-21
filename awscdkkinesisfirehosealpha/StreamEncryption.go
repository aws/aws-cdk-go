package awscdkkinesisfirehosealpha


// Options for server-side encryption of a delivery stream.
//
// Example:
//   var destination iDestination
//   // SSE with an customer-managed key that is explicitly specified
//   var key key
//
//
//   // SSE with an AWS-owned key
//   // SSE with an AWS-owned key
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream AWS Owned"), &DeliveryStreamProps{
//   	Encryption: firehose.StreamEncryption_AWS_OWNED,
//   	Destinations: []*iDestination{
//   		destination,
//   	},
//   })
//   // SSE with an customer-managed key that is created automatically by the CDK
//   // SSE with an customer-managed key that is created automatically by the CDK
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Implicit Customer Managed"), &DeliveryStreamProps{
//   	Encryption: firehose.StreamEncryption_CUSTOMER_MANAGED,
//   	Destinations: []*iDestination{
//   		destination,
//   	},
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Explicit Customer Managed"), &DeliveryStreamProps{
//   	EncryptionKey: key,
//   	Destinations: []*iDestination{
//   		destination,
//   	},
//   })
//
// Experimental.
type StreamEncryption string

const (
	// Data in the stream is stored unencrypted.
	// Experimental.
	StreamEncryption_UNENCRYPTED StreamEncryption = "UNENCRYPTED"
	// Data in the stream is stored encrypted by a KMS key managed by the customer.
	// Experimental.
	StreamEncryption_CUSTOMER_MANAGED StreamEncryption = "CUSTOMER_MANAGED"
	// Data in the stream is stored encrypted by a KMS key owned by AWS and managed for use in multiple AWS accounts.
	// Experimental.
	StreamEncryption_AWS_OWNED StreamEncryption = "AWS_OWNED"
)

