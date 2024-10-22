package awscdkkinesisfirehosealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
)

// A Kinesis Data Firehose delivery stream source.
//
// Example:
//   var destination iDestination
//
//   sourceStream := kinesis.NewStream(this, jsii.String("Source Stream"))
//
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Source: firehose.NewKinesisStreamSource(sourceStream),
//   	Destination: destination,
//   })
//
// Experimental.
type KinesisStreamSource interface {
	ISource
	// Grant read permissions for this source resource and its contents to an IAM principal (the delivery stream).
	//
	// If an encryption key is used, permission to use the key to decrypt the
	// contents of the stream will also be granted.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for KinesisStreamSource
type jsiiProxy_KinesisStreamSource struct {
	jsiiProxy_ISource
}

// Creates a new KinesisStreamSource.
// Experimental.
func NewKinesisStreamSource(stream awskinesis.IStream) KinesisStreamSource {
	_init_.Initialize()

	if err := validateNewKinesisStreamSourceParameters(stream); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisStreamSource{}

	_jsii_.Create(
		"@aws-cdk/aws-kinesisfirehose-alpha.KinesisStreamSource",
		[]interface{}{stream},
		&j,
	)

	return &j
}

// Creates a new KinesisStreamSource.
// Experimental.
func NewKinesisStreamSource_Override(k KinesisStreamSource, stream awskinesis.IStream) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-kinesisfirehose-alpha.KinesisStreamSource",
		[]interface{}{stream},
		k,
	)
}

func (k *jsiiProxy_KinesisStreamSource) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

