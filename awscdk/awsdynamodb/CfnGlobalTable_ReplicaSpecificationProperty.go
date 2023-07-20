package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Defines settings specific to a single replica of a global table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicaSpecificationProperty := &ReplicaSpecificationProperty{
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	GlobalSecondaryIndexes: []interface{}{
//   		&ReplicaGlobalSecondaryIndexSpecificationProperty{
//   			IndexName: jsii.String("indexName"),
//
//   			// the properties below are optional
//   			ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   			ReadProvisionedThroughputSettings: &ReadProvisionedThroughputSettingsProperty{
//   				ReadCapacityAutoScalingSettings: &CapacityAutoScalingSettingsProperty{
//   					MaxCapacity: jsii.Number(123),
//   					MinCapacity: jsii.Number(123),
//   					TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   						TargetValue: jsii.Number(123),
//
//   						// the properties below are optional
//   						DisableScaleIn: jsii.Boolean(false),
//   						ScaleInCooldown: jsii.Number(123),
//   						ScaleOutCooldown: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					SeedCapacity: jsii.Number(123),
//   				},
//   				ReadCapacityUnits: jsii.Number(123),
//   			},
//   		},
//   	},
//   	KinesisStreamSpecification: &KinesisStreamSpecificationProperty{
//   		StreamArn: jsii.String("streamArn"),
//   	},
//   	PointInTimeRecoverySpecification: &PointInTimeRecoverySpecificationProperty{
//   		PointInTimeRecoveryEnabled: jsii.Boolean(false),
//   	},
//   	ReadProvisionedThroughputSettings: &ReadProvisionedThroughputSettingsProperty{
//   		ReadCapacityAutoScalingSettings: &CapacityAutoScalingSettingsProperty{
//   			MaxCapacity: jsii.Number(123),
//   			MinCapacity: jsii.Number(123),
//   			TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   				TargetValue: jsii.Number(123),
//
//   				// the properties below are optional
//   				DisableScaleIn: jsii.Boolean(false),
//   				ScaleInCooldown: jsii.Number(123),
//   				ScaleOutCooldown: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			SeedCapacity: jsii.Number(123),
//   		},
//   		ReadCapacityUnits: jsii.Number(123),
//   	},
//   	SseSpecification: &ReplicaSSESpecificationProperty{
//   		KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	},
//   	TableClass: jsii.String("tableClass"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html
//
type CfnGlobalTable_ReplicaSpecificationProperty struct {
	// The region in which this replica exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
	// The settings used to enable or disable CloudWatch Contributor Insights for the specified replica.
	//
	// When not specified, defaults to contributor insights disabled for the replica.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-contributorinsightsspecification
	//
	ContributorInsightsSpecification interface{} `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Determines if a replica is protected from deletion.
	//
	// When enabled, the table cannot be deleted by any user or process. This setting is disabled by default. For more information, see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the *Amazon DynamoDB Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-deletionprotectionenabled
	//
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// Defines additional settings for the global secondary indexes of this replica.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-globalsecondaryindexes
	//
	GlobalSecondaryIndexes interface{} `field:"optional" json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// Defines the Kinesis Data Streams configuration for the specified replica.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-kinesisstreamspecification
	//
	KinesisStreamSpecification interface{} `field:"optional" json:"kinesisStreamSpecification" yaml:"kinesisStreamSpecification"`
	// The settings used to enable point in time recovery.
	//
	// When not specified, defaults to point in time recovery disabled for the replica.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-pointintimerecoveryspecification
	//
	PointInTimeRecoverySpecification interface{} `field:"optional" json:"pointInTimeRecoverySpecification" yaml:"pointInTimeRecoverySpecification"`
	// Defines read capacity settings for the replica table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-readprovisionedthroughputsettings
	//
	ReadProvisionedThroughputSettings interface{} `field:"optional" json:"readProvisionedThroughputSettings" yaml:"readProvisionedThroughputSettings"`
	// Allows you to specify a customer-managed key for the replica.
	//
	// When using customer-managed keys for server-side encryption, this property must have a value in all replicas.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-ssespecification
	//
	SseSpecification interface{} `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// The table class of the specified table.
	//
	// Valid values are `STANDARD` and `STANDARD_INFREQUENT_ACCESS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-tableclass
	//
	TableClass *string `field:"optional" json:"tableClass" yaml:"tableClass"`
	// An array of key-value pairs to apply to this replica.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

