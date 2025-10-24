package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Represents server-side encryption for an Amazon Firehose Delivery Stream.
//
// Example:
//   var destination IDestination
//   // SSE with an customer-managed key that is explicitly specified
//   var key Key
//
//
//   // SSE with an AWS-owned key
//   // SSE with an AWS-owned key
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream with AWS Owned Key"), &DeliveryStreamProps{
//   	Encryption: firehose.StreamEncryption_AwsOwnedKey(),
//   	Destination: destination,
//   })
//   // SSE with an customer-managed key that is created automatically by the CDK
//   // SSE with an customer-managed key that is created automatically by the CDK
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream with Customer Managed Key"), &DeliveryStreamProps{
//   	Encryption: firehose.StreamEncryption_CustomerManagedKey(),
//   	Destination: destination,
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream with Customer Managed and Provided Key"), &DeliveryStreamProps{
//   	Encryption: firehose.StreamEncryption_*CustomerManagedKey(key),
//   	Destination: destination,
//   })
//
type StreamEncryption interface {
	// Optional KMS key used for customer managed encryption.
	EncryptionKey() awskms.IKey
	// The type of server-side encryption for the Amazon Firehose delivery stream.
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
func StreamEncryption_AwsOwnedKey() StreamEncryption {
	_init_.Initialize()

	var returns StreamEncryption

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.StreamEncryption",
		"awsOwnedKey",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Configure server-side encryption using customer managed keys.
func StreamEncryption_CustomerManagedKey(encryptionKey awskms.IKey) StreamEncryption {
	_init_.Initialize()

	var returns StreamEncryption

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.StreamEncryption",
		"customerManagedKey",
		[]interface{}{encryptionKey},
		&returns,
	)

	return returns
}

// No server-side encryption is configured.
func StreamEncryption_Unencrypted() StreamEncryption {
	_init_.Initialize()

	var returns StreamEncryption

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.StreamEncryption",
		"unencrypted",
		nil, // no parameters
		&returns,
	)

	return returns
}

