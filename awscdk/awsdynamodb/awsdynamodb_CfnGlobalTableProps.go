package awsdynamodb


// Properties for defining a `CfnGlobalTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGlobalTableProps := &cfnGlobalTableProps{
//   	attributeDefinitions: []interface{}{
//   		&attributeDefinitionProperty{
//   			attributeName: jsii.String("attributeName"),
//   			attributeType: jsii.String("attributeType"),
//   		},
//   	},
//   	keySchema: []interface{}{
//   		&keySchemaProperty{
//   			attributeName: jsii.String("attributeName"),
//   			keyType: jsii.String("keyType"),
//   		},
//   	},
//   	replicas: []interface{}{
//   		&replicaSpecificationProperty{
//   			region: jsii.String("region"),
//
//   			// the properties below are optional
//   			contributorInsightsSpecification: &contributorInsightsSpecificationProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   			globalSecondaryIndexes: []interface{}{
//   				&replicaGlobalSecondaryIndexSpecificationProperty{
//   					indexName: jsii.String("indexName"),
//
//   					// the properties below are optional
//   					contributorInsightsSpecification: &contributorInsightsSpecificationProperty{
//   						enabled: jsii.Boolean(false),
//   					},
//   					readProvisionedThroughputSettings: &readProvisionedThroughputSettingsProperty{
//   						readCapacityAutoScalingSettings: &capacityAutoScalingSettingsProperty{
//   							maxCapacity: jsii.Number(123),
//   							minCapacity: jsii.Number(123),
//   							targetTrackingScalingPolicyConfiguration: &targetTrackingScalingPolicyConfigurationProperty{
//   								targetValue: jsii.Number(123),
//
//   								// the properties below are optional
//   								disableScaleIn: jsii.Boolean(false),
//   								scaleInCooldown: jsii.Number(123),
//   								scaleOutCooldown: jsii.Number(123),
//   							},
//
//   							// the properties below are optional
//   							seedCapacity: jsii.Number(123),
//   						},
//   						readCapacityUnits: jsii.Number(123),
//   					},
//   				},
//   			},
//   			pointInTimeRecoverySpecification: &pointInTimeRecoverySpecificationProperty{
//   				pointInTimeRecoveryEnabled: jsii.Boolean(false),
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
//   			sseSpecification: &replicaSSESpecificationProperty{
//   				kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   			},
//   			tableClass: jsii.String("tableClass"),
//   			tags: []cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	billingMode: jsii.String("billingMode"),
//   	globalSecondaryIndexes: []interface{}{
//   		&globalSecondaryIndexProperty{
//   			indexName: jsii.String("indexName"),
//   			keySchema: []interface{}{
//   				&keySchemaProperty{
//   					attributeName: jsii.String("attributeName"),
//   					keyType: jsii.String("keyType"),
//   				},
//   			},
//   			projection: &projectionProperty{
//   				nonKeyAttributes: []*string{
//   					jsii.String("nonKeyAttributes"),
//   				},
//   				projectionType: jsii.String("projectionType"),
//   			},
//
//   			// the properties below are optional
//   			writeProvisionedThroughputSettings: &writeProvisionedThroughputSettingsProperty{
//   				writeCapacityAutoScalingSettings: &capacityAutoScalingSettingsProperty{
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
//   			},
//   		},
//   	},
//   	localSecondaryIndexes: []interface{}{
//   		&localSecondaryIndexProperty{
//   			indexName: jsii.String("indexName"),
//   			keySchema: []interface{}{
//   				&keySchemaProperty{
//   					attributeName: jsii.String("attributeName"),
//   					keyType: jsii.String("keyType"),
//   				},
//   			},
//   			projection: &projectionProperty{
//   				nonKeyAttributes: []*string{
//   					jsii.String("nonKeyAttributes"),
//   				},
//   				projectionType: jsii.String("projectionType"),
//   			},
//   		},
//   	},
//   	sseSpecification: &sSESpecificationProperty{
//   		sseEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		sseType: jsii.String("sseType"),
//   	},
//   	streamSpecification: &streamSpecificationProperty{
//   		streamViewType: jsii.String("streamViewType"),
//   	},
//   	tableName: jsii.String("tableName"),
//   	timeToLiveSpecification: &timeToLiveSpecificationProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		attributeName: jsii.String("attributeName"),
//   	},
//   	writeProvisionedThroughputSettings: &writeProvisionedThroughputSettingsProperty{
//   		writeCapacityAutoScalingSettings: &capacityAutoScalingSettingsProperty{
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
//   	},
//   }
//
type CfnGlobalTableProps struct {
	// A list of attributes that describe the key schema for the global table and indexes.
	AttributeDefinitions interface{} `field:"required" json:"attributeDefinitions" yaml:"attributeDefinitions"`
	// Specifies the attributes that make up the primary key for the table.
	//
	// The attributes in the `KeySchema` property must also be defined in the `AttributeDefinitions` property.
	KeySchema interface{} `field:"required" json:"keySchema" yaml:"keySchema"`
	// Specifies the list of replicas for your global table.
	//
	// The list must contain at least one element, the region where the stack defining the global table is deployed. For example, if you define your table in a stack deployed to us-east-1, you must have an entry in `Replicas` with the region us-east-1. You cannot remove the replica in the stack region.
	//
	// > Adding a replica might take a few minutes for an empty table, or up to several hours for large tables. If you want to add or remove a replica, we recommend submitting an `UpdateStack` operation containing only that change.
	// >
	// > If you add or delete a replica during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new replica, you might need to manually delete the replica.
	//
	// You can create a new global table with as many replicas as needed. You can add or remove replicas after table creation, but you can only add or remove a single replica in each update.
	Replicas interface{} `field:"required" json:"replicas" yaml:"replicas"`
	// Specifies how you are charged for read and write throughput and how you manage capacity. Valid values are:.
	//
	// - `PAY_PER_REQUEST`
	// - `PROVISIONED`
	//
	// All replicas in your global table will have the same billing mode. If you use `PROVISIONED` billing mode, you must provide an auto scaling configuration via the `WriteProvisionedThroughputSettings` property. The default value of this property is `PROVISIONED` .
	BillingMode *string `field:"optional" json:"billingMode" yaml:"billingMode"`
	// Global secondary indexes to be created on the global table.
	//
	// You can create up to 20 global secondary indexes. Each replica in your global table will have the same global secondary index settings. You can only create or delete one global secondary index in a single stack operation.
	//
	// Since the backfilling of an index could take a long time, CloudFormation does not wait for the index to become active. If a stack operation rolls back, CloudFormation might not delete an index that has been added. In that case, you will need to delete the index manually.
	GlobalSecondaryIndexes interface{} `field:"optional" json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// Local secondary indexes to be created on the table.
	//
	// You can create up to five local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes. Each replica in your global table will have the same local secondary index settings.
	LocalSecondaryIndexes interface{} `field:"optional" json:"localSecondaryIndexes" yaml:"localSecondaryIndexes"`
	// Specifies the settings to enable server-side encryption.
	//
	// These settings will be applied to all replicas. If you plan to use customer-managed KMS keys, you must provide a key for each replica using the `ReplicaSpecification.ReplicaSSESpecification` property.
	SseSpecification interface{} `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// Specifies the streams settings on your global table.
	//
	// You must provide a value for this property if your global table contains more than one replica. You can only change the streams settings if your global table has only one replica.
	StreamSpecification interface{} `field:"optional" json:"streamSpecification" yaml:"streamSpecification"`
	// A name for the global table.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID as the table name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Specifies the time to live (TTL) settings for the table.
	//
	// This setting will be applied to all replicas.
	TimeToLiveSpecification interface{} `field:"optional" json:"timeToLiveSpecification" yaml:"timeToLiveSpecification"`
	// Specifies an auto scaling policy for write capacity.
	//
	// This policy will be applied to all replicas. This setting must be specified if `BillingMode` is set to `PROVISIONED` .
	WriteProvisionedThroughputSettings interface{} `field:"optional" json:"writeProvisionedThroughputSettings" yaml:"writeProvisionedThroughputSettings"`
}

