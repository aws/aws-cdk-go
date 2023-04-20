package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Properties for a new delivery stream.
//
// Example:
//   var bucket bucket
//   // Provide a Lambda function that will transform records before delivery, with custom
//   // buffering and retry configuration
//   lambdaFunction := lambda.NewFunction(this, jsii.String("Processor"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_14_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("process-records"))),
//   })
//   lambdaProcessor := firehose.NewLambdaFunctionProcessor(lambdaFunction, &DataProcessorProps{
//   	BufferInterval: awscdk.Duration_Minutes(jsii.Number(5)),
//   	BufferSize: awscdk.Size_Mebibytes(jsii.Number(5)),
//   	Retries: jsii.Number(5),
//   })
//   s3Destination := destinations.NewS3Bucket(bucket, &S3BucketProps{
//   	Processor: lambdaProcessor,
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destinations: []iDestination{
//   		s3Destination,
//   	},
//   })
//
// Experimental.
type DeliveryStreamProps struct {
	// The destinations that this delivery stream will deliver data to.
	//
	// Only a singleton array is supported at this time.
	// Experimental.
	Destinations *[]IDestination `field:"required" json:"destinations" yaml:"destinations"`
	// A name for the delivery stream.
	// Experimental.
	DeliveryStreamName *string `field:"optional" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// Indicates the type of customer master key (CMK) to use for server-side encryption, if any.
	// Experimental.
	Encryption StreamEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Customer managed key to server-side encrypt data in the stream.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The IAM role associated with this delivery stream.
	//
	// Assumed by Kinesis Data Firehose to read from sources and encrypt data server-side.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The Kinesis data stream to use as a source for this delivery stream.
	// Experimental.
	SourceStream awskinesis.IStream `field:"optional" json:"sourceStream" yaml:"sourceStream"`
}

