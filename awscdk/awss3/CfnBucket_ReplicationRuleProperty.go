package awss3


// Specifies which Amazon S3 objects to replicate and where to store the replicas.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationRuleProperty := &ReplicationRuleProperty{
//   	Destination: &ReplicationDestinationProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		AccessControlTranslation: &AccessControlTranslationProperty{
//   			Owner: jsii.String("owner"),
//   		},
//   		Account: jsii.String("account"),
//   		EncryptionConfiguration: &EncryptionConfigurationProperty{
//   			ReplicaKmsKeyId: jsii.String("replicaKmsKeyId"),
//   		},
//   		Metrics: &MetricsProperty{
//   			Status: jsii.String("status"),
//
//   			// the properties below are optional
//   			EventThreshold: &ReplicationTimeValueProperty{
//   				Minutes: jsii.Number(123),
//   			},
//   		},
//   		ReplicationTime: &ReplicationTimeProperty{
//   			Status: jsii.String("status"),
//   			Time: &ReplicationTimeValueProperty{
//   				Minutes: jsii.Number(123),
//   			},
//   		},
//   		StorageClass: jsii.String("storageClass"),
//   	},
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	DeleteMarkerReplication: &DeleteMarkerReplicationProperty{
//   		Status: jsii.String("status"),
//   	},
//   	Filter: &ReplicationRuleFilterProperty{
//   		And: &ReplicationRuleAndOperatorProperty{
//   			Prefix: jsii.String("prefix"),
//   			TagFilters: []interface{}{
//   				&TagFilterProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		Prefix: jsii.String("prefix"),
//   		TagFilter: &TagFilterProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Id: jsii.String("id"),
//   	Prefix: jsii.String("prefix"),
//   	Priority: jsii.Number(123),
//   	SourceSelectionCriteria: &SourceSelectionCriteriaProperty{
//   		ReplicaModifications: &ReplicaModificationsProperty{
//   			Status: jsii.String("status"),
//   		},
//   		SseKmsEncryptedObjects: &SseKmsEncryptedObjectsProperty{
//   			Status: jsii.String("status"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationrule.html
//
type CfnBucket_ReplicationRuleProperty struct {
	// A container for information about the replication destination and its configurations including enabling the S3 Replication Time Control (S3 RTC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationrule.html#cfn-s3-bucket-replicationrule-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Specifies whether the rule is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationrule.html#cfn-s3-bucket-replicationrule-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// Specifies whether Amazon S3 replicates delete markers.
	//
	// If you specify a `Filter` in your replication configuration, you must also include a `DeleteMarkerReplication` element. If your `Filter` includes a `Tag` element, the `DeleteMarkerReplication` `Status` must be set to Disabled, because Amazon S3 does not support replicating delete markers for tag-based rules. For an example configuration, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-config-min-rule-config) .
	//
	// For more information about delete marker replication, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/delete-marker-replication.html) .
	//
	// > If you are using an earlier version of the replication configuration, Amazon S3 handles replication of delete markers differently. For more information, see [Backward Compatibility](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationrule.html#cfn-s3-bucket-replicationrule-deletemarkerreplication
	//
	DeleteMarkerReplication interface{} `field:"optional" json:"deleteMarkerReplication" yaml:"deleteMarkerReplication"`
	// A filter that identifies the subset of objects to which the replication rule applies.
	//
	// A `Filter` must specify exactly one `Prefix` , `TagFilter` , or an `And` child element. The use of the filter field indicates that this is a V2 replication configuration. This field isn't supported in a V1 replication configuration.
	//
	// > V1 replication configuration only supports filtering by key prefix. To filter using a V1 replication configuration, add the `Prefix` directly as a child element of the `Rule` element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationrule.html#cfn-s3-bucket-replicationrule-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// A unique identifier for the rule.
	//
	// The maximum value is 255 characters. If you don't specify a value, AWS CloudFormation generates a random ID. When using a V2 replication configuration this property is capitalized as "ID".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationrule.html#cfn-s3-bucket-replicationrule-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// An object key name prefix that identifies the object or objects to which the rule applies.
	//
	// The maximum prefix length is 1,024 characters. To include all objects in a bucket, specify an empty string. To filter using a V1 replication configuration, add the `Prefix` directly as a child element of the `Rule` element.
	//
	// > Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationrule.html#cfn-s3-bucket-replicationrule-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The priority indicates which rule has precedence whenever two or more replication rules conflict.
	//
	// Amazon S3 will attempt to replicate objects according to all replication rules. However, if there are two or more rules with the same destination bucket, then objects will be replicated according to the rule with the highest priority. The higher the number, the higher the priority.
	//
	// For more information, see [Replication](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationrule.html#cfn-s3-bucket-replicationrule-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A container that describes additional filters for identifying the source objects that you want to replicate.
	//
	// You can choose to enable or disable the replication of these objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationrule.html#cfn-s3-bucket-replicationrule-sourceselectioncriteria
	//
	SourceSelectionCriteria interface{} `field:"optional" json:"sourceSelectionCriteria" yaml:"sourceSelectionCriteria"`
}

