package awscassandra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTableProps := &CfnTableProps{
//   	KeyspaceName: jsii.String("keyspaceName"),
//   	PartitionKeyColumns: []interface{}{
//   		&ColumnProperty{
//   			ColumnName: jsii.String("columnName"),
//   			ColumnType: jsii.String("columnType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AutoScalingSpecifications: &AutoScalingSpecificationProperty{
//   		ReadCapacityAutoScaling: &AutoScalingSettingProperty{
//   			AutoScalingDisabled: jsii.Boolean(false),
//   			MaximumUnits: jsii.Number(123),
//   			MinimumUnits: jsii.Number(123),
//   			ScalingPolicy: &ScalingPolicyProperty{
//   				TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   					TargetValue: jsii.Number(123),
//
//   					// the properties below are optional
//   					DisableScaleIn: jsii.Boolean(false),
//   					ScaleInCooldown: jsii.Number(123),
//   					ScaleOutCooldown: jsii.Number(123),
//   				},
//   			},
//   		},
//   		WriteCapacityAutoScaling: &AutoScalingSettingProperty{
//   			AutoScalingDisabled: jsii.Boolean(false),
//   			MaximumUnits: jsii.Number(123),
//   			MinimumUnits: jsii.Number(123),
//   			ScalingPolicy: &ScalingPolicyProperty{
//   				TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   					TargetValue: jsii.Number(123),
//
//   					// the properties below are optional
//   					DisableScaleIn: jsii.Boolean(false),
//   					ScaleInCooldown: jsii.Number(123),
//   					ScaleOutCooldown: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	BillingMode: &BillingModeProperty{
//   		Mode: jsii.String("mode"),
//
//   		// the properties below are optional
//   		ProvisionedThroughput: &ProvisionedThroughputProperty{
//   			ReadCapacityUnits: jsii.Number(123),
//   			WriteCapacityUnits: jsii.Number(123),
//   		},
//   	},
//   	CdcSpecification: &CdcSpecificationProperty{
//   		Status: jsii.String("status"),
//
//   		// the properties below are optional
//   		ViewType: jsii.String("viewType"),
//   	},
//   	ClientSideTimestampsEnabled: jsii.Boolean(false),
//   	ClusteringKeyColumns: []interface{}{
//   		&ClusteringKeyColumnProperty{
//   			Column: &ColumnProperty{
//   				ColumnName: jsii.String("columnName"),
//   				ColumnType: jsii.String("columnType"),
//   			},
//
//   			// the properties below are optional
//   			OrderBy: jsii.String("orderBy"),
//   		},
//   	},
//   	DefaultTimeToLive: jsii.Number(123),
//   	EncryptionSpecification: &EncryptionSpecificationProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//
//   		// the properties below are optional
//   		KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   	},
//   	PointInTimeRecoveryEnabled: jsii.Boolean(false),
//   	RegularColumns: []interface{}{
//   		&ColumnProperty{
//   			ColumnName: jsii.String("columnName"),
//   			ColumnType: jsii.String("columnType"),
//   		},
//   	},
//   	ReplicaSpecifications: []interface{}{
//   		&ReplicaSpecificationProperty{
//   			Region: jsii.String("region"),
//
//   			// the properties below are optional
//   			ReadCapacityAutoScaling: &AutoScalingSettingProperty{
//   				AutoScalingDisabled: jsii.Boolean(false),
//   				MaximumUnits: jsii.Number(123),
//   				MinimumUnits: jsii.Number(123),
//   				ScalingPolicy: &ScalingPolicyProperty{
//   					TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   						TargetValue: jsii.Number(123),
//
//   						// the properties below are optional
//   						DisableScaleIn: jsii.Boolean(false),
//   						ScaleInCooldown: jsii.Number(123),
//   						ScaleOutCooldown: jsii.Number(123),
//   					},
//   				},
//   			},
//   			ReadCapacityUnits: jsii.Number(123),
//   		},
//   	},
//   	TableName: jsii.String("tableName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html
//
type CfnTableProps struct {
	// The name of the keyspace to create the table in.
	//
	// The keyspace must already exist.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-keyspacename
	//
	KeyspaceName *string `field:"required" json:"keyspaceName" yaml:"keyspaceName"`
	// One or more columns that uniquely identify every row in the table.
	//
	// Every table must have a partition key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-partitionkeycolumns
	//
	PartitionKeyColumns interface{} `field:"required" json:"partitionKeyColumns" yaml:"partitionKeyColumns"`
	// The optional auto scaling capacity settings for a table in provisioned capacity mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-autoscalingspecifications
	//
	AutoScalingSpecifications interface{} `field:"optional" json:"autoScalingSpecifications" yaml:"autoScalingSpecifications"`
	// The billing mode for the table, which determines how you'll be charged for reads and writes:.
	//
	// - *On-demand mode* (default) - You pay based on the actual reads and writes your application performs.
	// - *Provisioned mode* - Lets you specify the number of reads and writes per second that you need for your application.
	//
	// If you don't specify a value for this property, then the table will use on-demand mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-billingmode
	//
	BillingMode interface{} `field:"optional" json:"billingMode" yaml:"billingMode"`
	// The settings for the CDC stream of a table.
	//
	// For more information about CDC streams, see [Working with change data capture (CDC) streams in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/cdc.html) in the *Amazon Keyspaces Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-cdcspecification
	//
	CdcSpecification interface{} `field:"optional" json:"cdcSpecification" yaml:"cdcSpecification"`
	// Enables client-side timestamps for the table.
	//
	// By default, the setting is disabled. You can enable client-side timestamps with the following option:
	//
	// - `status: "enabled"`
	//
	// After client-side timestamps are enabled for a table, you can't disable this setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-clientsidetimestampsenabled
	//
	ClientSideTimestampsEnabled interface{} `field:"optional" json:"clientSideTimestampsEnabled" yaml:"clientSideTimestampsEnabled"`
	// One or more columns that determine how the table data is sorted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-clusteringkeycolumns
	//
	ClusteringKeyColumns interface{} `field:"optional" json:"clusteringKeyColumns" yaml:"clusteringKeyColumns"`
	// The default Time To Live (TTL) value for all rows in a table in seconds.
	//
	// The maximum configurable value is 630,720,000 seconds, which is the equivalent of 20 years. By default, the TTL value for a table is 0, which means data does not expire.
	//
	// For more information, see [Setting the default TTL value for a table](https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_default_ttl) in the *Amazon Keyspaces Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-defaulttimetolive
	//
	DefaultTimeToLive *float64 `field:"optional" json:"defaultTimeToLive" yaml:"defaultTimeToLive"`
	// The encryption at rest options for the table.
	//
	// - *AWS owned key* (default) - The key is owned by Amazon Keyspaces .
	// - *Customer managed key* - The key is stored in your account and is created, owned, and managed by you.
	//
	// > If you choose encryption with a customer managed key, you must specify a valid customer managed KMS key with permissions granted to Amazon Keyspaces.
	//
	// For more information, see [Encryption at rest in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html) in the *Amazon Keyspaces Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-encryptionspecification
	//
	EncryptionSpecification interface{} `field:"optional" json:"encryptionSpecification" yaml:"encryptionSpecification"`
	// Specifies if point-in-time recovery is enabled or disabled for the table.
	//
	// The options are `PointInTimeRecoveryEnabled=true` and `PointInTimeRecoveryEnabled=false` . If not specified, the default is `PointInTimeRecoveryEnabled=false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-pointintimerecoveryenabled
	//
	PointInTimeRecoveryEnabled interface{} `field:"optional" json:"pointInTimeRecoveryEnabled" yaml:"pointInTimeRecoveryEnabled"`
	// One or more columns that are not part of the primary key - that is, columns that are *not* defined as partition key columns or clustering key columns.
	//
	// You can add regular columns to existing tables by adding them to the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-regularcolumns
	//
	RegularColumns interface{} `field:"optional" json:"regularColumns" yaml:"regularColumns"`
	// The AWS Region specific settings of a multi-Region table.
	//
	// For a multi-Region table, you can configure the table's read capacity differently per AWS Region. You can do this by configuring the following parameters.
	//
	// - `region` : The Region where these settings are applied. (Required)
	// - `readCapacityUnits` : The provisioned read capacity units. (Optional)
	// - `readCapacityAutoScaling` : The read capacity auto scaling settings for the table. (Optional)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-replicaspecifications
	//
	ReplicaSpecifications interface{} `field:"optional" json:"replicaSpecifications" yaml:"replicaSpecifications"`
	// The name of the table to be created.
	//
	// The table name is case sensitive. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the table name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you can't perform updates that require replacing this resource. You can perform updates that require no interruption or some interruption. If you must replace the resource, specify a new name.
	//
	// *Length constraints:* Minimum length of 3. Maximum length of 255.
	//
	// *Pattern:* `^[a-zA-Z0-9][a-zA-Z0-9_]{1,47}$`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

