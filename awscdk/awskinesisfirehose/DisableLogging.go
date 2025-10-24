package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Disables logging for error logs.
//
// When this class is used, logging is disabled (`logging: false`)
// and no CloudWatch log group can be specified.
//
// Example:
//   var bucket Bucket
//
//   destination := firehose.NewS3Bucket(bucket, &S3BucketProps{
//   	LoggingConfig: firehose.NewDisableLogging(),
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destination: destination,
//   })
//
type DisableLogging interface {
	ILoggingConfig
	// If true, log errors when data transformation or data delivery fails.
	//
	// `true` when using `EnableLogging`, `false` when using `DisableLogging`.
	Logging() *bool
}

// The jsii proxy struct for DisableLogging
type jsiiProxy_DisableLogging struct {
	jsiiProxy_ILoggingConfig
}

func (j *jsiiProxy_DisableLogging) Logging() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}


func NewDisableLogging() DisableLogging {
	_init_.Initialize()

	j := jsiiProxy_DisableLogging{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.DisableLogging",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewDisableLogging_Override(d DisableLogging) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.DisableLogging",
		nil, // no parameters
		d,
	)
}

