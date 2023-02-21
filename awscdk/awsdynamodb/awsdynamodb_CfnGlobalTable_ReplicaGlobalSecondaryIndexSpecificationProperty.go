package awsdynamodb


// Represents the properties of a global secondary index that can be set on a per-replica basis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicaGlobalSecondaryIndexSpecificationProperty := &replicaGlobalSecondaryIndexSpecificationProperty{
//   	indexName: jsii.String("indexName"),
//
//   	// the properties below are optional
//   	contributorInsightsSpecification: &contributorInsightsSpecificationProperty{
//   		enabled: jsii.Boolean(false),
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
//   }
//
type CfnGlobalTable_ReplicaGlobalSecondaryIndexSpecificationProperty struct {
	// The name of the global secondary index.
	//
	// The name must be unique among all other indexes on this table.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Updates the status for contributor insights for a specific table or index.
	//
	// CloudWatch Contributor Insights for DynamoDB graphs display the partition key and (if applicable) sort key of frequently accessed items and frequently throttled items in plaintext. If you require the use of AWS Key Management Service (KMS) to encrypt this table’s partition key and sort key data with an AWS managed key or customer managed key, you should not enable CloudWatch Contributor Insights for DynamoDB for this table.
	ContributorInsightsSpecification interface{} `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Allows you to specify the read capacity settings for a replica global secondary index when the `BillingMode` is set to `PROVISIONED` .
	ReadProvisionedThroughputSettings interface{} `field:"optional" json:"readProvisionedThroughputSettings" yaml:"readProvisionedThroughputSettings"`
}

