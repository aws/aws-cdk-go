// CDK Destinations Constructs for AWS Kinesis Firehose
package awscdkkinesisfirehosedestinationsalpha


// Options for S3 record backup of a delivery stream.
//
// Example:
//   // Enable backup of all source records (to an S3 bucket created by CDK).
//   var bucket bucket
//   // Explicitly provide an S3 bucket to which all source records will be backed up.
//   var backupBucket bucket
//
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket, &s3BucketProps{
//   			s3Backup: &destinationS3BackupProps{
//   				mode: destinations.backupMode_ALL,
//   			},
//   		}),
//   	},
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All Explicit Bucket"), &deliveryStreamProps{
//   	destinations: []*iDestination{
//   		destinations.NewS3Bucket(bucket, &s3BucketProps{
//   			s3Backup: &destinationS3BackupProps{
//   				bucket: backupBucket,
//   			},
//   		}),
//   	},
//   })
//   // Explicitly provide an S3 prefix under which all source records will be backed up.
//   // Explicitly provide an S3 prefix under which all source records will be backed up.
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream Backup All Explicit Prefix"), &deliveryStreamProps{
//   	destinations: []*iDestination{
//   		destinations.NewS3Bucket(bucket, &s3BucketProps{
//   			s3Backup: &destinationS3BackupProps{
//   				mode: destinations.*backupMode_ALL,
//   				dataOutputPrefix: jsii.String("mybackup"),
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type BackupMode string

const (
	// All records are backed up.
	// Experimental.
	BackupMode_ALL BackupMode = "ALL"
	// Only records that failed to deliver or transform are backed up.
	// Experimental.
	BackupMode_FAILED BackupMode = "FAILED"
)

