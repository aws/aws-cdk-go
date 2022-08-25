package awss3


// Specifies which Amazon S3 objects to replicate and where to store the replicas.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationRuleProperty := &replicationRuleProperty{
//   	destination: &replicationDestinationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		accessControlTranslation: &accessControlTranslationProperty{
//   			owner: jsii.String("owner"),
//   		},
//   		account: jsii.String("account"),
//   		encryptionConfiguration: &encryptionConfigurationProperty{
//   			replicaKmsKeyId: jsii.String("replicaKmsKeyId"),
//   		},
//   		metrics: &metricsProperty{
//   			status: jsii.String("status"),
//
//   			// the properties below are optional
//   			eventThreshold: &replicationTimeValueProperty{
//   				minutes: jsii.Number(123),
//   			},
//   		},
//   		replicationTime: &replicationTimeProperty{
//   			status: jsii.String("status"),
//   			time: &replicationTimeValueProperty{
//   				minutes: jsii.Number(123),
//   			},
//   		},
//   		storageClass: jsii.String("storageClass"),
//   	},
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	deleteMarkerReplication: &deleteMarkerReplicationProperty{
//   		status: jsii.String("status"),
//   	},
//   	filter: &replicationRuleFilterProperty{
//   		and: &replicationRuleAndOperatorProperty{
//   			prefix: jsii.String("prefix"),
//   			tagFilters: []interface{}{
//   				&tagFilterProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		prefix: jsii.String("prefix"),
//   		tagFilter: &tagFilterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	id: jsii.String("id"),
//   	prefix: jsii.String("prefix"),
//   	priority: jsii.Number(123),
//   	sourceSelectionCriteria: &sourceSelectionCriteriaProperty{
//   		replicaModifications: &replicaModificationsProperty{
//   			status: jsii.String("status"),
//   		},
//   		sseKmsEncryptedObjects: &sseKmsEncryptedObjectsProperty{
//   			status: jsii.String("status"),
//   		},
//   	},
//   }
//
type CfnBucket_ReplicationRuleProperty struct {
	// A container for information about the replication destination and its configurations including enabling the S3 Replication Time Control (S3 RTC).
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Specifies whether the rule is enabled.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Specifies whether Amazon S3 replicates delete markers.
	//
	// If you specify a `Filter` in your replication configuration, you must also include a `DeleteMarkerReplication` element. If your `Filter` includes a `Tag` element, the `DeleteMarkerReplication` `Status` must be set to Disabled, because Amazon S3 does not support replicating delete markers for tag-based rules. For an example configuration, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-config-min-rule-config) .
	//
	// For more information about delete marker replication, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/delete-marker-replication.html) .
	//
	// > If you are using an earlier version of the replication configuration, Amazon S3 handles replication of delete markers differently. For more information, see [Backward Compatibility](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations) .
	DeleteMarkerReplication interface{} `field:"optional" json:"deleteMarkerReplication" yaml:"deleteMarkerReplication"`
	// A filter that identifies the subset of objects to which the replication rule applies.
	//
	// A `Filter` must specify exactly one `Prefix` , `TagFilter` , or an `And` child element. The use of the filter field indicates this is a V2 replication configuration. V1 does not have this field.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// A unique identifier for the rule.
	//
	// The maximum value is 255 characters. If you don't specify a value, AWS CloudFormation generates a random ID. When using a V2 replication configuration this property is capitalized as "ID".
	Id *string `field:"optional" json:"id" yaml:"id"`
	// An object key name prefix that identifies the object or objects to which the rule applies.
	//
	// The maximum prefix length is 1,024 characters. To include all objects in a bucket, specify an empty string.
	//
	// > Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints) .
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The priority indicates which rule has precedence whenever two or more replication rules conflict.
	//
	// Amazon S3 will attempt to replicate objects according to all replication rules. However, if there are two or more rules with the same destination bucket, then objects will be replicated according to the rule with the highest priority. The higher the number, the higher the priority.
	//
	// For more information, see [Replication](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html) in the *Amazon S3 User Guide* .
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A container that describes additional filters for identifying the source objects that you want to replicate.
	//
	// You can choose to enable or disable the replication of these objects.
	SourceSelectionCriteria interface{} `field:"optional" json:"sourceSelectionCriteria" yaml:"sourceSelectionCriteria"`
}

