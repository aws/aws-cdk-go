package awskinesisfirehose


// Options for CloudWatchLogProcessor.
//
// Example:
//   var bucket Bucket
//
//   s3Destination := firehose.NewS3Bucket(bucket, &S3BucketProps{
//   	Processors: []IDataProcessor{
//   		firehose.NewDecompressionProcessor(),
//   		firehose.NewCloudWatchLogProcessor(&CloudWatchLogProcessorOptions{
//   			DataMessageExtraction: jsii.Boolean(true),
//   		}),
//   	},
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destination: s3Destination,
//   })
//
type CloudWatchLogProcessorOptions struct {
	// Extract message from CloudWatch logs.
	//
	// This must be true.
	DataMessageExtraction *bool `field:"required" json:"dataMessageExtraction" yaml:"dataMessageExtraction"`
}

