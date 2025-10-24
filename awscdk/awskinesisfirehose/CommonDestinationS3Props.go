package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Common properties for defining a backup, intermediary, or final S3 destination for a Amazon Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var compression Compression
//   var key Key
//   var size Size
//
//   commonDestinationS3Props := &CommonDestinationS3Props{
//   	BufferingInterval: cdk.Duration_Minutes(jsii.Number(30)),
//   	BufferingSize: size,
//   	Compression: compression,
//   	DataOutputPrefix: jsii.String("dataOutputPrefix"),
//   	EncryptionKey: key,
//   	ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   }
//
type CommonDestinationS3Props struct {
	// The length of time that Firehose buffers incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Duration.seconds(0)
	// Maximum: Duration.seconds(900)
	// Default: Duration.seconds(300)
	//
	BufferingInterval awscdk.Duration `field:"optional" json:"bufferingInterval" yaml:"bufferingInterval"`
	// The size of the buffer that Amazon Data Firehose uses for incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Size.mebibytes(1) when record data format conversion is disabled, Size.mebibytes(64) when it is enabled
	// Maximum: Size.mebibytes(128)
	// Default: Size.mebibytes(5) when record data format conversion is disabled, Size.mebibytes(128) when it is enabled
	//
	BufferingSize awscdk.Size `field:"optional" json:"bufferingSize" yaml:"bufferingSize"`
	// The type of compression that Amazon Data Firehose uses to compress the data that it delivers to the Amazon S3 bucket.
	//
	// The compression formats SNAPPY or ZIP cannot be specified for Amazon Redshift
	// destinations because they are not supported by the Amazon Redshift COPY operation
	// that reads from the S3 bucket.
	// Default: - UNCOMPRESSED.
	//
	Compression Compression `field:"optional" json:"compression" yaml:"compression"`
	// A prefix that Amazon Data Firehose evaluates and adds to records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Default: "YYYY/MM/DD/HH".
	//
	DataOutputPrefix *string `field:"optional" json:"dataOutputPrefix" yaml:"dataOutputPrefix"`
	// The AWS KMS key used to encrypt the data that it delivers to your Amazon S3 bucket.
	// Default: - Data is not encrypted.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// A prefix that Amazon Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Default: "YYYY/MM/DD/HH".
	//
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
}

