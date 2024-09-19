package awscdkkinesisfirehosealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Represents server-side encryption for a Kinesis Firehose Delivery Stream.
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
//   	Encryption: firehose.StreamEncryption_AwsOwnedKey(),
//   	Destinations: []*iDestination{
//   		destination,
//   	},
//   })
//   // SSE with an customer-managed key that is created automatically by the CDK
//   // SSE with an customer-managed key that is created automatically by the CDK
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Implicit Customer Managed"), &DeliveryStreamProps{
//   	Encryption: firehose.StreamEncryption_CustomerManagedKey(),
//   	Destinations: []*iDestination{
//   		destination,
//   	},
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Explicit Customer Managed"), &DeliveryStreamProps{
//   	Encryption: firehose.StreamEncryption_*CustomerManagedKey(key),
//   	Destinations: []*iDestination{
//   		destination,
//   	},
//   })
//
// Experimental.
type StreamEncryption interface {
	// Optional KMS key used for customer managed encryption.
	// Experimental.
	EncryptionKey() awskms.IKey
	// The type of server-side encryption for the Kinesis Firehose delivery stream.
	// Experimental.
	Type() StreamEncryptionType
}

// The jsii proxy struct for StreamEncryption
type jsiiProxy_StreamEncryption struct {
	_ byte // padding
}

func (j *jsiiProxy_StreamEncryption) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamEncryption) Type() StreamEncryptionType {
	var returns StreamEncryptionType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Configure server-side encryption using an AWS owned key.
// Experimental.
func StreamEncryption_AwsOwnedKey() StreamEncryption {
	_init_.Initialize()

	var returns StreamEncryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisfirehose-alpha.StreamEncryption",
		"awsOwnedKey",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Configure server-side encryption using customer managed keys.
// Experimental.
func StreamEncryption_CustomerManagedKey(encryptionKey awskms.IKey) StreamEncryption {
	_init_.Initialize()

	var returns StreamEncryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisfirehose-alpha.StreamEncryption",
		"customerManagedKey",
		[]interface{}{encryptionKey},
		&returns,
	)

	return returns
}

// No server-side encryption is configured.
// Experimental.
func StreamEncryption_Unencrypted() StreamEncryption {
	_init_.Initialize()

	var returns StreamEncryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisfirehose-alpha.StreamEncryption",
		"unencrypted",
		nil, // no parameters
		&returns,
	)

	return returns
}

