package awss3


// A container for replication rules.
//
// You can add up to 1,000 rules. The maximum size of a replication configuration is 2 MB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationConfigurationProperty := &ReplicationConfigurationProperty{
//   	Role: jsii.String("role"),
//   	Rules: []interface{}{
//   		&ReplicationRuleProperty{
//   			Destination: &ReplicationDestinationProperty{
//   				Bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				AccessControlTranslation: &AccessControlTranslationProperty{
//   					Owner: jsii.String("owner"),
//   				},
//   				Account: jsii.String("account"),
//   				EncryptionConfiguration: &EncryptionConfigurationProperty{
//   					ReplicaKmsKeyId: jsii.String("replicaKmsKeyId"),
//   				},
//   				Metrics: &MetricsProperty{
//   					Status: jsii.String("status"),
//
//   					// the properties below are optional
//   					EventThreshold: &ReplicationTimeValueProperty{
//   						Minutes: jsii.Number(123),
//   					},
//   				},
//   				ReplicationTime: &ReplicationTimeProperty{
//   					Status: jsii.String("status"),
//   					Time: &ReplicationTimeValueProperty{
//   						Minutes: jsii.Number(123),
//   					},
//   				},
//   				StorageClass: jsii.String("storageClass"),
//   			},
//   			Status: jsii.String("status"),
//
//   			// the properties below are optional
//   			DeleteMarkerReplication: &DeleteMarkerReplicationProperty{
//   				Status: jsii.String("status"),
//   			},
//   			Filter: &ReplicationRuleFilterProperty{
//   				And: &ReplicationRuleAndOperatorProperty{
//   					Prefix: jsii.String("prefix"),
//   					TagFilters: []interface{}{
//   						&TagFilterProperty{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				Prefix: jsii.String("prefix"),
//   				TagFilter: &TagFilterProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Id: jsii.String("id"),
//   			Prefix: jsii.String("prefix"),
//   			Priority: jsii.Number(123),
//   			SourceSelectionCriteria: &SourceSelectionCriteriaProperty{
//   				ReplicaModifications: &ReplicaModificationsProperty{
//   					Status: jsii.String("status"),
//   				},
//   				SseKmsEncryptedObjects: &SseKmsEncryptedObjectsProperty{
//   					Status: jsii.String("status"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_ReplicationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that Amazon S3 assumes when replicating objects.
	//
	// For more information, see [How to Set Up Replication](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-how-setup.html) in the *Amazon S3 User Guide* .
	Role *string `field:"required" json:"role" yaml:"role"`
	// A container for one or more replication rules.
	//
	// A replication configuration must have at least one rule and can contain a maximum of 1,000 rules.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

