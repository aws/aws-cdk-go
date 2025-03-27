package awskinesisfirehose


// Options for S3 record backup of a delivery stream.
//
// Example:
//   // Enable backup of all source records (to an S3 bucket created by CDK).
//   var bucket bucket
//   // Explicitly provide an S3 bucket to which all source records will be backed up.
//   var backupBucket bucket
//
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All"), &DeliveryStreamProps{
//   	Destination:
//   	firehose.NewS3Bucket(bucket, &S3BucketProps{
//   		S3Backup: &DestinationS3BackupProps{
//   			Mode: firehose.BackupMode_ALL,
//   		},
//   	}),
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All Explicit Bucket"), &DeliveryStreamProps{
//   	Destination:
//   	firehose.NewS3Bucket(bucket, &S3BucketProps{
//   		S3Backup: &DestinationS3BackupProps{
//   			Bucket: backupBucket,
//   		},
//   	}),
//   })
//   // Explicitly provide an S3 prefix under which all source records will be backed up.
//   // Explicitly provide an S3 prefix under which all source records will be backed up.
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All Explicit Prefix"), &DeliveryStreamProps{
//   	Destination:
//   	firehose.NewS3Bucket(bucket, &S3BucketProps{
//   		S3Backup: &DestinationS3BackupProps{
//   			Mode: firehose.BackupMode_ALL,
//   			DataOutputPrefix: jsii.String("mybackup"),
//   		},
//   	}),
//   })
//
type BackupMode string

const (
	// All records are backed up.
	BackupMode_ALL BackupMode = "ALL"
	// Only records that failed to deliver or transform are backed up.
	BackupMode_FAILED BackupMode = "FAILED"
)

