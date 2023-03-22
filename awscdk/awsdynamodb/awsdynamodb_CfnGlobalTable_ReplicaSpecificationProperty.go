package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Defines settings specific to a single replica of a global table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicaSpecificationProperty := &replicaSpecificationProperty{
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	contributorInsightsSpecification: &contributorInsightsSpecificationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	deletionProtectionEnabled: jsii.Boolean(false),
//   	globalSecondaryIndexes: []interface{}{
//   		&replicaGlobalSecondaryIndexSpecificationProperty{
//   			indexName: jsii.String("indexName"),
//
//   			// the properties below are optional
//   			contributorInsightsSpecification: &contributorInsightsSpecificationProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   			readProvisionedThroughputSettings: &readProvisionedThroughputSettingsProperty{
//   				readCapacityAutoScalingSettings: &capacityAutoScalingSettingsProperty{
//   					maxCapacity: jsii.Number(123),
//   					minCapacity: jsii.Number(123),
//   					targetTrackingScalingPolicyConfiguration: &targetTrackingScalingPolicyConfigurationProperty{
//   						targetValue: jsii.Number(123),
//
//   						// the properties below are optional
//   						disableScaleIn: jsii.Boolean(false),
//   						scaleInCooldown: jsii.Number(123),
//   						scaleOutCooldown: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					seedCapacity: jsii.Number(123),
//   				},
//   				readCapacityUnits: jsii.Number(123),
//   			},
//   		},
//   	},
//   	kinesisStreamSpecification: &kinesisStreamSpecificationProperty{
//   		streamArn: jsii.String("streamArn"),
//   	},
//   	pointInTimeRecoverySpecification: &pointInTimeRecoverySpecificationProperty{
//   		pointInTimeRecoveryEnabled: jsii.Boolean(false),
//   	},
//   	readProvisionedThroughputSettings: &readProvisionedThroughputSettingsProperty{
//   		readCapacityAutoScalingSettings: &capacityAutoScalingSettingsProperty{
//   			maxCapacity: jsii.Number(123),
//   			minCapacity: jsii.Number(123),
//   			targetTrackingScalingPolicyConfiguration: &targetTrackingScalingPolicyConfigurationProperty{
//   				targetValue: jsii.Number(123),
//
//   				// the properties below are optional
//   				disableScaleIn: jsii.Boolean(false),
//   				scaleInCooldown: jsii.Number(123),
//   				scaleOutCooldown: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			seedCapacity: jsii.Number(123),
//   		},
//   		readCapacityUnits: jsii.Number(123),
//   	},
//   	sseSpecification: &replicaSSESpecificationProperty{
//   		kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	},
//   	tableClass: jsii.String("tableClass"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnGlobalTable_ReplicaSpecificationProperty struct {
	// The region in which this replica exists.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The settings used to enable or disable CloudWatch Contributor Insights for the specified replica.
	//
	// When not specified, defaults to contributor insights disabled for the replica.
	ContributorInsightsSpecification interface{} `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Determines if a replica is protected from deletion.
	//
	// When enabled, the table cannot be deleted by any user or process. This setting is disabled by default. For more information, see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the *Amazon DynamoDB Developer Guide* .
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// Defines additional settings for the global secondary indexes of this replica.
	GlobalSecondaryIndexes interface{} `field:"optional" json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// Defines the Kinesis Data Streams configuration for the specified replica.
	KinesisStreamSpecification interface{} `field:"optional" json:"kinesisStreamSpecification" yaml:"kinesisStreamSpecification"`
	// The settings used to enable point in time recovery.
	//
	// When not specified, defaults to point in time recovery disabled for the replica.
	PointInTimeRecoverySpecification interface{} `field:"optional" json:"pointInTimeRecoverySpecification" yaml:"pointInTimeRecoverySpecification"`
	// Defines read capacity settings for the replica table.
	ReadProvisionedThroughputSettings interface{} `field:"optional" json:"readProvisionedThroughputSettings" yaml:"readProvisionedThroughputSettings"`
	// Allows you to specify a customer-managed key for the replica.
	//
	// When using customer-managed keys for server-side encryption, this property must have a value in all replicas.
	SseSpecification interface{} `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// The table class of the specified table.
	//
	// Valid values are `STANDARD` and `STANDARD_INFREQUENT_ACCESS` .
	TableClass *string `field:"optional" json:"tableClass" yaml:"tableClass"`
	// An array of key-value pairs to apply to this replica.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

