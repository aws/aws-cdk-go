package awsdynamodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DynamoDB::GlobalTable` resource enables you to create and manage a Version 2019.11.21 global table. This resource cannot be used to create or manage a Version 2017.11.29 global table. For more information, see [Global tables](https://docs.aws.amazon.com//amazondynamodb/latest/developerguide/GlobalTables.html) .
//
// > You cannot convert a resource of type `AWS::DynamoDB::Table` into a resource of type `AWS::DynamoDB::GlobalTable` by changing its type in your template. *Doing so might result in the deletion of your DynamoDB table.*
// >
// > You can instead use the GlobalTable resource to create a new table in a single Region. This will be billed the same as a single Region table. If you later update the stack to add other Regions then Global Tables pricing will apply.
//
// You should be aware of the following behaviors when working with DynamoDB global tables.
//
// - The IAM Principal executing the stack operation must have the permissions listed below in all regions where you plan to have a global table replica. The IAM Principal's permissions should not have restrictions based on IP source address. Some global tables operations (for example, adding a replica) are asynchronous, and require that the IAM Principal is valid until they complete. You should not delete the Principal (user or IAM role) until CloudFormation has finished updating your stack.
//
// - `dynamodb:CreateTable`
// - `dynamodb:UpdateTable`
// - `dynamodb:DeleteTable`
// - `dynamodb:DescribeContinuousBackups`
// - `dynamodb:DescribeContributorInsights`
// - `dynamodb:DescribeTable`
// - `dynamodb:DescribeTableReplicaAutoScaling`
// - `dynamodb:DescribeTimeToLive`
// - `dynamodb:ListTables`
// - `dynamodb:UpdateTimeToLive`
// - `dynamodb:UpdateContributorInsights`
// - `dynamodb:UpdateContinuousBackups`
// - `dynamodb:ListTagsOfResource`
// - `dynamodb:TagResource`
// - `dynamodb:UntagResource`
// - `dynamodb:BatchWriteItem`
// - `dynamodb:CreateTableReplica`
// - `dynamodb:DeleteItem`
// - `dynamodb:DeleteTableReplica`
// - `dynamodb:DisableKinesisStreamingDestination`
// - `dynamodb:EnableKinesisStreamingDestination`
// - `dynamodb:GetItem`
// - `dynamodb:PutItem`
// - `dynamodb:Query`
// - `dynamodb:Scan`
// - `dynamodb:UpdateItem`
// - `dynamodb:DescribeTableReplicaAutoScaling`
// - `dynamodb:UpdateTableReplicaAutoScaling`
// - `iam:CreateServiceLinkedRole`
// - `kms:CreateGrant`
// - `kms:DescribeKey`
// - `application-autoscaling:DeleteScalingPolicy`
// - `application-autoscaling:DeleteScheduledAction`
// - `application-autoscaling:DeregisterScalableTarget`
// - `application-autoscaling:DescribeScalingPolicies`
// - `application-autoscaling:DescribeScalableTargets`
// - `application-autoscaling:PutScalingPolicy`
// - `application-autoscaling:PutScheduledAction`
// - `application-autoscaling:RegisterScalableTarget`
// - When using provisioned billing mode, CloudFormation will create an auto scaling policy on each of your replicas to control their write capacities. You must configure this policy using the `WriteProvisionedThroughputSettings` property. CloudFormation will ensure that all replicas have the same write capacity auto scaling property. You cannot directly specify a value for write capacity for a global table.
// - If your table uses provisioned capacity, you must configure auto scaling directly in the `AWS::DynamoDB::GlobalTable` resource. You should not configure additional auto scaling policies on any of the table replicas or global secondary indexes, either via API or via `AWS::ApplicationAutoScaling::ScalableTarget` or `AWS::ApplicationAutoScaling::ScalingPolicy` . Doing so might result in unexpected behavior and is unsupported.
// - In AWS CloudFormation , each global table is controlled by a single stack, in a single region, regardless of the number of replicas. When you deploy your template, CloudFormation will create/update all replicas as part of a single stack operation. You should not deploy the same `AWS::DynamoDB::GlobalTable` resource in multiple regions. Doing so will result in errors, and is unsupported. If you deploy your application template in multiple regions, you can use conditions to only create the resource in a single region. Alternatively, you can choose to define your `AWS::DynamoDB::GlobalTable` resources in a stack separate from your application stack, and make sure it is only deployed to a single region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnGlobalTable := awscdk.Aws_dynamodb.NewCfnGlobalTable(this, jsii.String("MyCfnGlobalTable"), &CfnGlobalTableProps{
//   	AttributeDefinitions: []interface{}{
//   		&AttributeDefinitionProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			AttributeType: jsii.String("attributeType"),
//   		},
//   	},
//   	KeySchema: []interface{}{
//   		&KeySchemaProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			KeyType: jsii.String("keyType"),
//   		},
//   	},
//   	Replicas: []interface{}{
//   		&ReplicaSpecificationProperty{
//   			Region: jsii.String("region"),
//
//   			// the properties below are optional
//   			ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   			DeletionProtectionEnabled: jsii.Boolean(false),
//   			GlobalSecondaryIndexes: []interface{}{
//   				&ReplicaGlobalSecondaryIndexSpecificationProperty{
//   					IndexName: jsii.String("indexName"),
//
//   					// the properties below are optional
//   					ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   						Enabled: jsii.Boolean(false),
//   					},
//   					ReadOnDemandThroughputSettings: &ReadOnDemandThroughputSettingsProperty{
//   						MaxReadRequestUnits: jsii.Number(123),
//   					},
//   					ReadProvisionedThroughputSettings: &ReadProvisionedThroughputSettingsProperty{
//   						ReadCapacityAutoScalingSettings: &CapacityAutoScalingSettingsProperty{
//   							MaxCapacity: jsii.Number(123),
//   							MinCapacity: jsii.Number(123),
//   							TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   								TargetValue: jsii.Number(123),
//
//   								// the properties below are optional
//   								DisableScaleIn: jsii.Boolean(false),
//   								ScaleInCooldown: jsii.Number(123),
//   								ScaleOutCooldown: jsii.Number(123),
//   							},
//
//   							// the properties below are optional
//   							SeedCapacity: jsii.Number(123),
//   						},
//   						ReadCapacityUnits: jsii.Number(123),
//   					},
//   				},
//   			},
//   			KinesisStreamSpecification: &KinesisStreamSpecificationProperty{
//   				StreamArn: jsii.String("streamArn"),
//
//   				// the properties below are optional
//   				ApproximateCreationDateTimePrecision: jsii.String("approximateCreationDateTimePrecision"),
//   			},
//   			PointInTimeRecoverySpecification: &PointInTimeRecoverySpecificationProperty{
//   				PointInTimeRecoveryEnabled: jsii.Boolean(false),
//   			},
//   			ReadOnDemandThroughputSettings: &ReadOnDemandThroughputSettingsProperty{
//   				MaxReadRequestUnits: jsii.Number(123),
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
//   			ReplicaStreamSpecification: &ReplicaStreamSpecificationProperty{
//   				ResourcePolicy: &ResourcePolicyProperty{
//   					PolicyDocument: policyDocument,
//   				},
//   			},
//   			ResourcePolicy: &ResourcePolicyProperty{
//   				PolicyDocument: policyDocument,
//   			},
//   			SseSpecification: &ReplicaSSESpecificationProperty{
//   				KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   			},
//   			TableClass: jsii.String("tableClass"),
//   			Tags: []cfnTag{
//   				&cfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	BillingMode: jsii.String("billingMode"),
//   	GlobalSecondaryIndexes: []interface{}{
//   		&GlobalSecondaryIndexProperty{
//   			IndexName: jsii.String("indexName"),
//   			KeySchema: []interface{}{
//   				&KeySchemaProperty{
//   					AttributeName: jsii.String("attributeName"),
//   					KeyType: jsii.String("keyType"),
//   				},
//   			},
//   			Projection: &ProjectionProperty{
//   				NonKeyAttributes: []*string{
//   					jsii.String("nonKeyAttributes"),
//   				},
//   				ProjectionType: jsii.String("projectionType"),
//   			},
//
//   			// the properties below are optional
//   			WarmThroughput: &WarmThroughputProperty{
//   				ReadUnitsPerSecond: jsii.Number(123),
//   				WriteUnitsPerSecond: jsii.Number(123),
//   			},
//   			WriteOnDemandThroughputSettings: &WriteOnDemandThroughputSettingsProperty{
//   				MaxWriteRequestUnits: jsii.Number(123),
//   			},
//   			WriteProvisionedThroughputSettings: &WriteProvisionedThroughputSettingsProperty{
//   				WriteCapacityAutoScalingSettings: &CapacityAutoScalingSettingsProperty{
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
//   			},
//   		},
//   	},
//   	LocalSecondaryIndexes: []interface{}{
//   		&LocalSecondaryIndexProperty{
//   			IndexName: jsii.String("indexName"),
//   			KeySchema: []interface{}{
//   				&KeySchemaProperty{
//   					AttributeName: jsii.String("attributeName"),
//   					KeyType: jsii.String("keyType"),
//   				},
//   			},
//   			Projection: &ProjectionProperty{
//   				NonKeyAttributes: []*string{
//   					jsii.String("nonKeyAttributes"),
//   				},
//   				ProjectionType: jsii.String("projectionType"),
//   			},
//   		},
//   	},
//   	SseSpecification: &SSESpecificationProperty{
//   		SseEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		SseType: jsii.String("sseType"),
//   	},
//   	StreamSpecification: &StreamSpecificationProperty{
//   		StreamViewType: jsii.String("streamViewType"),
//   	},
//   	TableName: jsii.String("tableName"),
//   	TimeToLiveSpecification: &TimeToLiveSpecificationProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		AttributeName: jsii.String("attributeName"),
//   	},
//   	WarmThroughput: &WarmThroughputProperty{
//   		ReadUnitsPerSecond: jsii.Number(123),
//   		WriteUnitsPerSecond: jsii.Number(123),
//   	},
//   	WriteOnDemandThroughputSettings: &WriteOnDemandThroughputSettingsProperty{
//   		MaxWriteRequestUnits: jsii.Number(123),
//   	},
//   	WriteProvisionedThroughputSettings: &WriteProvisionedThroughputSettingsProperty{
//   		WriteCapacityAutoScalingSettings: &CapacityAutoScalingSettingsProperty{
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
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html
//
type CfnGlobalTable interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the DynamoDB table, such as `arn:aws:dynamodb:us-east-2:123456789012:table/myDynamoDBTable` .
	//
	// The ARN returned is that of the replica in the region the stack is deployed to.
	AttrArn() *string
	// A list of attributes that describe the key schema for the global table and indexes.
	AttributeDefinitions() interface{}
	SetAttributeDefinitions(val interface{})
	// The ARN of the DynamoDB stream, such as `arn:aws:dynamodb:us-east-1:123456789012:table/testddbstack-myDynamoDBTable-012A1SL7SMP5Q/stream/2015-11-30T20:10:00.000` . The `StreamArn` returned is that of the replica in the region the stack is deployed to.
	//
	// > You must specify the `StreamSpecification` property to use this attribute.
	AttrStreamArn() *string
	// Unique identifier for the table, such as `a123b456-01ab-23cd-123a-111222aaabbb` .
	//
	// The `TableId` returned is that of the replica in the region the stack is deployed to.
	AttrTableId() *string
	// Specifies how you are charged for read and write throughput and how you manage capacity.
	//
	// Valid values are:.
	BillingMode() *string
	SetBillingMode(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Global secondary indexes to be created on the global table.
	GlobalSecondaryIndexes() interface{}
	SetGlobalSecondaryIndexes(val interface{})
	// Specifies the attributes that make up the primary key for the table.
	KeySchema() interface{}
	SetKeySchema(val interface{})
	// Local secondary indexes to be created on the table.
	LocalSecondaryIndexes() interface{}
	SetLocalSecondaryIndexes(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Specifies the list of replicas for your global table.
	Replicas() interface{}
	SetReplicas(val interface{})
	// Specifies the settings to enable server-side encryption.
	SseSpecification() interface{}
	SetSseSpecification(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Specifies the streams settings on your global table.
	StreamSpecification() interface{}
	SetStreamSpecification(val interface{})
	// A name for the global table.
	TableName() *string
	SetTableName(val *string)
	// Specifies the time to live (TTL) settings for the table.
	TimeToLiveSpecification() interface{}
	SetTimeToLiveSpecification(val interface{})
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Provides visibility into the number of read and write operations your table or secondary index can instantaneously support.
	WarmThroughput() interface{}
	SetWarmThroughput(val interface{})
	// Sets the write request settings for a global table or a global secondary index.
	WriteOnDemandThroughputSettings() interface{}
	SetWriteOnDemandThroughputSettings(val interface{})
	// Specifies an auto scaling policy for write capacity.
	WriteProvisionedThroughputSettings() interface{}
	SetWriteProvisionedThroughputSettings(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnGlobalTable
type jsiiProxy_CfnGlobalTable struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGlobalTable) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) AttributeDefinitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) AttrStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) AttrTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) BillingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) GlobalSecondaryIndexes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) KeySchema() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) LocalSecondaryIndexes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) Replicas() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) SseSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) StreamSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) TimeToLiveSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeToLiveSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) WarmThroughput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) WriteOnDemandThroughputSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeOnDemandThroughputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) WriteProvisionedThroughputSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeProvisionedThroughputSettings",
		&returns,
	)
	return returns
}


func NewCfnGlobalTable(scope constructs.Construct, id *string, props *CfnGlobalTableProps) CfnGlobalTable {
	_init_.Initialize()

	if err := validateNewCfnGlobalTableParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGlobalTable{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnGlobalTable_Override(c CfnGlobalTable, scope constructs.Construct, id *string, props *CfnGlobalTableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGlobalTable)SetAttributeDefinitions(val interface{}) {
	if err := j.validateSetAttributeDefinitionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeDefinitions",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable)SetBillingMode(val *string) {
	_jsii_.Set(
		j,
		"billingMode",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable)SetGlobalSecondaryIndexes(val interface{}) {
	if err := j.validateSetGlobalSecondaryIndexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalSecondaryIndexes",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable)SetKeySchema(val interface{}) {
	if err := j.validateSetKeySchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keySchema",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable)SetLocalSecondaryIndexes(val interface{}) {
	if err := j.validateSetLocalSecondaryIndexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSecondaryIndexes",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable)SetReplicas(val interface{}) {
	if err := j.validateSetReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicas",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable)SetSseSpecification(val interface{}) {
	if err := j.validateSetSseSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sseSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable)SetStreamSpecification(val interface{}) {
	if err := j.validateSetStreamSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable)SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable)SetTimeToLiveSpecification(val interface{}) {
	if err := j.validateSetTimeToLiveSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeToLiveSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable)SetWarmThroughput(val interface{}) {
	if err := j.validateSetWarmThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warmThroughput",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable)SetWriteOnDemandThroughputSettings(val interface{}) {
	if err := j.validateSetWriteOnDemandThroughputSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeOnDemandThroughputSettings",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable)SetWriteProvisionedThroughputSettings(val interface{}) {
	if err := j.validateSetWriteProvisionedThroughputSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeProvisionedThroughputSettings",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnGlobalTable_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGlobalTable_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnGlobalTable_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGlobalTable_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnGlobalTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGlobalTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGlobalTable_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_dynamodb.CfnGlobalTable",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGlobalTable) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnGlobalTable) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGlobalTable) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGlobalTable) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnGlobalTable) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnGlobalTable) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnGlobalTable) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnGlobalTable) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnGlobalTable) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGlobalTable) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGlobalTable) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnGlobalTable) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGlobalTable) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGlobalTable) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGlobalTable) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGlobalTable) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGlobalTable) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnGlobalTable) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGlobalTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGlobalTable) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

