package awss3


// Specifies whether Amazon S3 replicates delete markers.
//
// If you specify a `Filter` in your replication configuration, you must also include a `DeleteMarkerReplication` element. If your `Filter` includes a `Tag` element, the `DeleteMarkerReplication` `Status` must be set to Disabled, because Amazon S3 does not support replicating delete markers for tag-based rules. For an example configuration, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-config-min-rule-config) .
//
// For more information about delete marker replication, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/delete-marker-replication.html) .
//
// > If you are using an earlier version of the replication configuration, Amazon S3 handles replication of delete markers differently. For more information, see [Backward Compatibility](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deleteMarkerReplicationProperty := &deleteMarkerReplicationProperty{
//   	status: jsii.String("status"),
//   }
//
type CfnBucket_DeleteMarkerReplicationProperty struct {
	// Indicates whether to replicate delete markers.
	//
	// Disabled by default.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

