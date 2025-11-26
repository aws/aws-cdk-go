package previewawsdynamodbmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Defines settings specific to a single replica of a global table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policyDocument interface{}
//
//   replicaSpecificationProperty := &ReplicaSpecificationProperty{
//   	ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   		Enabled: jsii.Boolean(false),
//   		Mode: jsii.String("mode"),
//   	},
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	GlobalSecondaryIndexes: []interface{}{
//   		&ReplicaGlobalSecondaryIndexSpecificationProperty{
//   			ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   				Enabled: jsii.Boolean(false),
//   				Mode: jsii.String("mode"),
//   			},
//   			IndexName: jsii.String("indexName"),
//   			ReadOnDemandThroughputSettings: &ReadOnDemandThroughputSettingsProperty{
//   				MaxReadRequestUnits: jsii.Number(123),
//   			},
//   			ReadProvisionedThroughputSettings: &ReadProvisionedThroughputSettingsProperty{
//   				ReadCapacityAutoScalingSettings: &CapacityAutoScalingSettingsProperty{
//   					MaxCapacity: jsii.Number(123),
//   					MinCapacity: jsii.Number(123),
//   					SeedCapacity: jsii.Number(123),
//   					TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   						DisableScaleIn: jsii.Boolean(false),
//   						ScaleInCooldown: jsii.Number(123),
//   						ScaleOutCooldown: jsii.Number(123),
//   						TargetValue: jsii.Number(123),
//   					},
//   				},
//   				ReadCapacityUnits: jsii.Number(123),
//   			},
//   		},
//   	},
//   	KinesisStreamSpecification: &KinesisStreamSpecificationProperty{
//   		ApproximateCreationDateTimePrecision: jsii.String("approximateCreationDateTimePrecision"),
//   		StreamArn: jsii.String("streamArn"),
//   	},
//   	PointInTimeRecoverySpecification: &PointInTimeRecoverySpecificationProperty{
//   		PointInTimeRecoveryEnabled: jsii.Boolean(false),
//   		RecoveryPeriodInDays: jsii.Number(123),
//   	},
//   	ReadOnDemandThroughputSettings: &ReadOnDemandThroughputSettingsProperty{
//   		MaxReadRequestUnits: jsii.Number(123),
//   	},
//   	ReadProvisionedThroughputSettings: &ReadProvisionedThroughputSettingsProperty{
//   		ReadCapacityAutoScalingSettings: &CapacityAutoScalingSettingsProperty{
//   			MaxCapacity: jsii.Number(123),
//   			MinCapacity: jsii.Number(123),
//   			SeedCapacity: jsii.Number(123),
//   			TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   				DisableScaleIn: jsii.Boolean(false),
//   				ScaleInCooldown: jsii.Number(123),
//   				ScaleOutCooldown: jsii.Number(123),
//   				TargetValue: jsii.Number(123),
//   			},
//   		},
//   		ReadCapacityUnits: jsii.Number(123),
//   	},
//   	Region: jsii.String("region"),
//   	ReplicaStreamSpecification: &ReplicaStreamSpecificationProperty{
//   		ResourcePolicy: &ResourcePolicyProperty{
//   			PolicyDocument: policyDocument,
//   		},
//   	},
//   	ResourcePolicy: &ResourcePolicyProperty{
//   		PolicyDocument: policyDocument,
//   	},
//   	SseSpecification: &ReplicaSSESpecificationProperty{
//   		KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	},
//   	TableClass: jsii.String("tableClass"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html
//
type CfnGlobalTablePropsMixin_ReplicaSpecificationProperty struct {
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
	// Sets read request settings for the replica table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-readondemandthroughputsettings
	//
	ReadOnDemandThroughputSettings interface{} `field:"optional" json:"readOnDemandThroughputSettings" yaml:"readOnDemandThroughputSettings"`
	// Defines read capacity settings for the replica table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-readprovisionedthroughputsettings
	//
	ReadProvisionedThroughputSettings interface{} `field:"optional" json:"readProvisionedThroughputSettings" yaml:"readProvisionedThroughputSettings"`
	// The region in which this replica exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Represents the DynamoDB Streams configuration for a global table replica.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-replicastreamspecification
	//
	ReplicaStreamSpecification interface{} `field:"optional" json:"replicaStreamSpecification" yaml:"replicaStreamSpecification"`
	// A resource-based policy document that contains permissions to add to the specified replica of a DynamoDB global table.
	//
	// Resource-based policies let you define access permissions by specifying who has access to each resource, and the actions they are allowed to perform on each resource.
	//
	// In a CloudFormation template, you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to DynamoDB . For more information about resource-based policies, see [Using resource-based policies for DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/access-control-resource-based.html) and [Resource-based policy examples](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-examples.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-resourcepolicy
	//
	ResourcePolicy interface{} `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
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

