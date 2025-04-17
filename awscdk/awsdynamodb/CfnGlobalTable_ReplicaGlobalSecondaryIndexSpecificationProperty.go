package awsdynamodb


// Represents the properties of a global secondary index that can be set on a per-replica basis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicaGlobalSecondaryIndexSpecificationProperty := &ReplicaGlobalSecondaryIndexSpecificationProperty{
//   	IndexName: jsii.String("indexName"),
//
//   	// the properties below are optional
//   	ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	ReadOnDemandThroughputSettings: &ReadOnDemandThroughputSettingsProperty{
//   		MaxReadRequestUnits: jsii.Number(123),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaglobalsecondaryindexspecification.html
//
type CfnGlobalTable_ReplicaGlobalSecondaryIndexSpecificationProperty struct {
	// The name of the global secondary index.
	//
	// The name must be unique among all other indexes on this table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaglobalsecondaryindexspecification.html#cfn-dynamodb-globaltable-replicaglobalsecondaryindexspecification-indexname
	//
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Updates the status for contributor insights for a specific table or index.
	//
	// CloudWatch Contributor Insights for DynamoDB graphs display the partition key and (if applicable) sort key of frequently accessed items and frequently throttled items in plaintext. If you require the use of AWS Key Management Service (KMS) to encrypt this table’s partition key and sort key data with an AWS managed key or customer managed key, you should not enable CloudWatch Contributor Insights for DynamoDB for this table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaglobalsecondaryindexspecification.html#cfn-dynamodb-globaltable-replicaglobalsecondaryindexspecification-contributorinsightsspecification
	//
	ContributorInsightsSpecification interface{} `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Sets the read request settings for a replica global secondary index.
	//
	// You can only specify this setting if your resource uses the `PAY_PER_REQUEST` `BillingMode` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaglobalsecondaryindexspecification.html#cfn-dynamodb-globaltable-replicaglobalsecondaryindexspecification-readondemandthroughputsettings
	//
	ReadOnDemandThroughputSettings interface{} `field:"optional" json:"readOnDemandThroughputSettings" yaml:"readOnDemandThroughputSettings"`
	// Allows you to specify the read capacity settings for a replica global secondary index when the `BillingMode` is set to `PROVISIONED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaglobalsecondaryindexspecification.html#cfn-dynamodb-globaltable-replicaglobalsecondaryindexspecification-readprovisionedthroughputsettings
	//
	ReadProvisionedThroughputSettings interface{} `field:"optional" json:"readProvisionedThroughputSettings" yaml:"readProvisionedThroughputSettings"`
}

