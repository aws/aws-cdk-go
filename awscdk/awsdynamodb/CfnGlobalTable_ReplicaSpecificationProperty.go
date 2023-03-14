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
type CfnGlobalTable_ReplicaSpecificationProperty struct {
	// The region in which this replica exists.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The settings used to enable or disable CloudWatch Contributor Insights for the specified replica.
	//
	// When not specified, defaults to contributor insights disabled for the replica.
	ContributorInsightsSpecification interface{} `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Defines additional settings for the global secondary indexes of this replica.
	GlobalSecondaryIndexes interface{} `field:"optional" json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
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

