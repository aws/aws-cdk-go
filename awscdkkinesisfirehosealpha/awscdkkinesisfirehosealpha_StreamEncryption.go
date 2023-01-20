// The CDK Construct Library for AWS::KinesisFirehose
package awscdkkinesisfirehosealpha


// Options for server-side encryption of a delivery stream.
//
// Example:
//   var destination iDestination
//   // SSE with an customer-managed CMK that is explicitly specified
//   var key key
//
//
//   // SSE with an AWS-owned CMK
//   // SSE with an AWS-owned CMK
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream AWS Owned"), &deliveryStreamProps{
//   	encryption: firehose.streamEncryption_AWS_OWNED,
//   	destinations: []*iDestination{
//   		destination,
//   	},
//   })
//   // SSE with an customer-managed CMK that is created automatically by the CDK
//   // SSE with an customer-managed CMK that is created automatically by the CDK
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Implicit Customer Managed"), &deliveryStreamProps{
//   	encryption: firehose.*streamEncryption_CUSTOMER_MANAGED,
//   	destinations: []*iDestination{
//   		destination,
//   	},
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Explicit Customer Managed"), &deliveryStreamProps{
//   	encryptionKey: key,
//   	destinations: []*iDestination{
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

