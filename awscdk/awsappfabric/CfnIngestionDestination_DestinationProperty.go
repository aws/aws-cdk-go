package awsappfabric


// Contains information about an audit log destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationProperty := &DestinationProperty{
//   	FirehoseStream: &FirehoseStreamProperty{
//   		StreamName: jsii.String("streamName"),
//   	},
//   	S3Bucket: &S3BucketProperty{
//   		BucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-destination.html
//
type CfnIngestionDestination_DestinationProperty struct {
	// Contains information about an Amazon Kinesis Data Firehose delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-destination.html#cfn-appfabric-ingestiondestination-destination-firehosestream
	//
	FirehoseStream interface{} `field:"optional" json:"firehoseStream" yaml:"firehoseStream"`
	// Contains information about an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-destination.html#cfn-appfabric-ingestiondestination-destination-s3bucket
	//
	S3Bucket interface{} `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

