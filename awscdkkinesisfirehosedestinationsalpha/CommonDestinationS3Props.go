// CDK Destinations Constructs for AWS Kinesis Firehose
package awscdkkinesisfirehosedestinationsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Common properties for defining a backup, intermediary, or final S3 destination for a Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import kinesisfirehose_destinations_alpha "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var compression compression
//   var key key
//   var size size
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
// Experimental.
type CommonDestinationS3Props struct {
	// The length of time that Firehose buffers incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Duration.seconds(60)
	// Maximum: Duration.seconds(900)
	// Experimental.
	BufferingInterval awscdk.Duration `field:"optional" json:"bufferingInterval" yaml:"bufferingInterval"`
	// The size of the buffer that Kinesis Data Firehose uses for incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Size.mebibytes(1)
	// Maximum: Size.mebibytes(128)
	// Experimental.
	BufferingSize awscdk.Size `field:"optional" json:"bufferingSize" yaml:"bufferingSize"`
	// The type of compression that Kinesis Data Firehose uses to compress the data that it delivers to the Amazon S3 bucket.
	//
	// The compression formats SNAPPY or ZIP cannot be specified for Amazon Redshift
	// destinations because they are not supported by the Amazon Redshift COPY operation
	// that reads from the S3 bucket.
	// Experimental.
	Compression Compression `field:"optional" json:"compression" yaml:"compression"`
	// A prefix that Kinesis Data Firehose evaluates and adds to records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	DataOutputPrefix *string `field:"optional" json:"dataOutputPrefix" yaml:"dataOutputPrefix"`
	// The AWS KMS key used to encrypt the data that it delivers to your Amazon S3 bucket.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
}

