package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Possible compression options Amazon Data Firehose can use to compress data on delivery.
//
// Example:
//   // Compress data delivered to S3 using Snappy
//   var bucket Bucket
//
//   s3Destination := firehose.NewS3Bucket(bucket, &S3BucketProps{
//   	Compression: firehose.Compression_SNAPPY(),
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destination: s3Destination,
//   })
//
type Compression interface {
	// the string value of the Compression.
	Value() *string
}

// The jsii proxy struct for Compression
type jsiiProxy_Compression struct {
	_ byte // padding
}

func (j *jsiiProxy_Compression) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Creates a new Compression instance with a custom value.
func Compression_Of(value *string) Compression {
	_init_.Initialize()

	if err := validateCompression_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Compression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.Compression",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Compression_GZIP() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.Compression",
		"GZIP",
		&returns,
	)
	return returns
}

func Compression_HADOOP_SNAPPY() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.Compression",
		"HADOOP_SNAPPY",
		&returns,
	)
	return returns
}

func Compression_SNAPPY() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.Compression",
		"SNAPPY",
		&returns,
	)
	return returns
}

func Compression_UNCOMPRESSED() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.Compression",
		"UNCOMPRESSED",
		&returns,
	)
	return returns
}

func Compression_ZIP() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.Compression",
		"ZIP",
		&returns,
	)
	return returns
}

