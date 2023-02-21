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
//   replicationConfigurationProperty := &replicationConfigurationProperty{
//   	role: jsii.String("role"),
//   	rules: []interface{}{
//   		&replicationRuleProperty{
//   			destination: &replicationDestinationProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				accessControlTranslation: &accessControlTranslationProperty{
//   					owner: jsii.String("owner"),
//   				},
//   				account: jsii.String("account"),
//   				encryptionConfiguration: &encryptionConfigurationProperty{
//   					replicaKmsKeyId: jsii.String("replicaKmsKeyId"),
//   				},
//   				metrics: &metricsProperty{
//   					status: jsii.String("status"),
//
//   					// the properties below are optional
//   					eventThreshold: &replicationTimeValueProperty{
//   						minutes: jsii.Number(123),
//   					},
//   				},
//   				replicationTime: &replicationTimeProperty{
//   					status: jsii.String("status"),
//   					time: &replicationTimeValueProperty{
//   						minutes: jsii.Number(123),
//   					},
//   				},
//   				storageClass: jsii.String("storageClass"),
//   			},
//   			status: jsii.String("status"),
//
//   			// the properties below are optional
//   			deleteMarkerReplication: &deleteMarkerReplicationProperty{
//   				status: jsii.String("status"),
//   			},
//   			filter: &replicationRuleFilterProperty{
//   				and: &replicationRuleAndOperatorProperty{
//   					prefix: jsii.String("prefix"),
//   					tagFilters: []interface{}{
//   						&tagFilterProperty{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				prefix: jsii.String("prefix"),
//   				tagFilter: &tagFilterProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			id: jsii.String("id"),
//   			prefix: jsii.String("prefix"),
//   			priority: jsii.Number(123),
//   			sourceSelectionCriteria: &sourceSelectionCriteriaProperty{
//   				replicaModifications: &replicaModificationsProperty{
//   					status: jsii.String("status"),
//   				},
//   				sseKmsEncryptedObjects: &sseKmsEncryptedObjectsProperty{
//   					status: jsii.String("status"),
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

