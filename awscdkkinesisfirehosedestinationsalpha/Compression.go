package awscdkkinesisfirehosedestinationsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Possible compression options Kinesis Data Firehose can use to compress data on delivery.
//
// Example:
//   // Compress data delivered to S3 using Snappy
//   var bucket bucket
//
//   s3Destination := destinations.NewS3Bucket(bucket, &S3BucketProps{
//   	Compression: destinations.Compression_SNAPPY(),
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destination: s3Destination,
//   })
//
// Deprecated.
type Compression interface {
	// the string value of the Compression.
	// Deprecated.
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
// Deprecated.
func Compression_Of(value *string) Compression {
	_init_.Initialize()

	if err := validateCompression_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Compression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.Compression",
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
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.Compression",
		"GZIP",
		&returns,
	)
	return returns
}

func Compression_HADOOP_SNAPPY() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.Compression",
		"HADOOP_SNAPPY",
		&returns,
	)
	return returns
}

func Compression_SNAPPY() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.Compression",
		"SNAPPY",
		&returns,
	)
	return returns
}

func Compression_UNCOMPRESSED() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.Compression",
		"UNCOMPRESSED",
		&returns,
	)
	return returns
}

func Compression_ZIP() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.Compression",
		"ZIP",
		&returns,
	)
	return returns
}

