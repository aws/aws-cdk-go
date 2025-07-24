package awss3


// The date source for the partitioned prefix.
//
// Example:
//   accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"))
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
//   	ServerAccessLogsBucket: accessLogsBucket,
//   	ServerAccessLogsPrefix: jsii.String("logs"),
//   	TargetObjectKeyFormat: s3.TargetObjectKeyFormat_PartitionedPrefix(s3.PartitionDateSource_EVENT_TIME),
//   })
//
type PartitionDateSource string

const (
	// The year, month, and day will be based on the timestamp of the S3 event in the file that's been delivered.
	PartitionDateSource_EVENT_TIME PartitionDateSource = "EVENT_TIME"
	// The year, month, and day will be based on the time when the log file was delivered to S3.
	PartitionDateSource_DELIVERY_TIME PartitionDateSource = "DELIVERY_TIME"
)

