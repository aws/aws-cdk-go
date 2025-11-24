package mixinsawsdynamodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsdynamodb/mixinsawsdynamodb/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
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
// - `application-autoscaling:DeleteScalingPolicy`
// - `application-autoscaling:DeleteScheduledAction`
// - `application-autoscaling:DeregisterScalableTarget`
// - `application-autoscaling:DescribeScalableTargets`
// - `application-autoscaling:DescribeScalingPolicies`
// - `application-autoscaling:PutScalingPolicy`
// - `application-autoscaling:PutScheduledAction`
// - `application-autoscaling:RegisterScalableTarget`
// - `dynamodb:BatchWriteItem`
// - `dynamodb:CreateGlobalTableWitness`
// - `dynamodb:CreateTable`
// - `dynamodb:CreateTableReplica`
// - `dynamodb:DeleteGlobalTableWitness`
// - `dynamodb:DeleteItem`
// - `dynamodb:DeleteTable`
// - `dynamodb:DeleteTableReplica`
// - `dynamodb:DescribeContinuousBackups`
// - `dynamodb:DescribeContributorInsights`
// - `dynamodb:DescribeTable`
// - `dynamodb:DescribeTableReplicaAutoScaling`
// - `dynamodb:DescribeTimeToLive`
// - `dynamodb:DisableKinesisStreamingDestination`
// - `dynamodb:EnableKinesisStreamingDestination`
// - `dynamodb:GetItem`
// - `dynamodb:ListTables`
// - `dynamodb:ListTagsOfResource`
// - `dynamodb:PutItem`
// - `dynamodb:Query`
// - `dynamodb:Scan`
// - `dynamodb:TagResource`
// - `dynamodb:UntagResource`
// - `dynamodb:UpdateContinuousBackups`
// - `dynamodb:UpdateContributorInsights`
// - `dynamodb:UpdateItem`
// - `dynamodb:UpdateTable`
// - `dynamodb:UpdateTableReplicaAutoScaling`
// - `dynamodb:UpdateTimeToLive`
// - `iam:CreateServiceLinkedRole`
// - `kms:CreateGrant`
// - `kms:DescribeKey`
// - When using provisioned billing mode, CloudFormation will create an auto scaling policy on each of your replicas to control their write capacities. You must configure this policy using the `WriteProvisionedThroughputSettings` property. CloudFormation will ensure that all replicas have the same write capacity auto scaling property. You cannot directly specify a value for write capacity for a global table.
// - If your table uses provisioned capacity, you must configure auto scaling directly in the `AWS::DynamoDB::GlobalTable` resource. You should not configure additional auto scaling policies on any of the table replicas or global secondary indexes, either via API or via `AWS::ApplicationAutoScaling::ScalableTarget` or `AWS::ApplicationAutoScaling::ScalingPolicy` . Doing so might result in unexpected behavior and is unsupported.
// - In AWS CloudFormation , each global table is controlled by a single stack, in a single region, regardless of the number of replicas. When you deploy your template, CloudFormation will create/update all replicas as part of a single stack operation. You should not deploy the same `AWS::DynamoDB::GlobalTable` resource in multiple regions. Doing so will result in errors, and is unsupported. If you deploy your application template in multiple regions, you can use conditions to only create the resource in a single region. Alternatively, you can choose to define your `AWS::DynamoDB::GlobalTable` resources in a stack separate from your application stack, and make sure it is only deployed to a single region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policyDocument interface{}
//
//   cfnGlobalTablePropsMixin := awscdkmixinspreview.Mixins.NewCfnGlobalTablePropsMixin(&CfnGlobalTableMixinProps{
//   	AttributeDefinitions: []interface{}{
//   		&AttributeDefinitionProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			AttributeType: jsii.String("attributeType"),
//   		},
//   	},
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
//   					SeedCapacity: jsii.Number(123),
//   					TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   						DisableScaleIn: jsii.Boolean(false),
//   						ScaleInCooldown: jsii.Number(123),
//   						ScaleOutCooldown: jsii.Number(123),
//   						TargetValue: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	GlobalTableWitnesses: []interface{}{
//   		&GlobalTableWitnessProperty{
//   			Region: jsii.String("region"),
//   		},
//   	},
//   	KeySchema: []interface{}{
//   		&KeySchemaProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			KeyType: jsii.String("keyType"),
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
//   	MultiRegionConsistency: jsii.String("multiRegionConsistency"),
//   	Replicas: []interface{}{
//   		&ReplicaSpecificationProperty{
//   			ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   				Enabled: jsii.Boolean(false),
//   				Mode: jsii.String("mode"),
//   			},
//   			DeletionProtectionEnabled: jsii.Boolean(false),
//   			GlobalSecondaryIndexes: []interface{}{
//   				&ReplicaGlobalSecondaryIndexSpecificationProperty{
//   					ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   						Enabled: jsii.Boolean(false),
//   						Mode: jsii.String("mode"),
//   					},
//   					IndexName: jsii.String("indexName"),
//   					ReadOnDemandThroughputSettings: &ReadOnDemandThroughputSettingsProperty{
//   						MaxReadRequestUnits: jsii.Number(123),
//   					},
//   					ReadProvisionedThroughputSettings: &ReadProvisionedThroughputSettingsProperty{
//   						ReadCapacityAutoScalingSettings: &CapacityAutoScalingSettingsProperty{
//   							MaxCapacity: jsii.Number(123),
//   							MinCapacity: jsii.Number(123),
//   							SeedCapacity: jsii.Number(123),
//   							TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   								DisableScaleIn: jsii.Boolean(false),
//   								ScaleInCooldown: jsii.Number(123),
//   								ScaleOutCooldown: jsii.Number(123),
//   								TargetValue: jsii.Number(123),
//   							},
//   						},
//   						ReadCapacityUnits: jsii.Number(123),
//   					},
//   				},
//   			},
//   			KinesisStreamSpecification: &KinesisStreamSpecificationProperty{
//   				ApproximateCreationDateTimePrecision: jsii.String("approximateCreationDateTimePrecision"),
//   				StreamArn: jsii.String("streamArn"),
//   			},
//   			PointInTimeRecoverySpecification: &PointInTimeRecoverySpecificationProperty{
//   				PointInTimeRecoveryEnabled: jsii.Boolean(false),
//   				RecoveryPeriodInDays: jsii.Number(123),
//   			},
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
//   			Region: jsii.String("region"),
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
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	SseSpecification: &SSESpecificationProperty{
//   		SseEnabled: jsii.Boolean(false),
//   		SseType: jsii.String("sseType"),
//   	},
//   	StreamSpecification: &StreamSpecificationProperty{
//   		StreamViewType: jsii.String("streamViewType"),
//   	},
//   	TableName: jsii.String("tableName"),
//   	TimeToLiveSpecification: &TimeToLiveSpecificationProperty{
//   		AttributeName: jsii.String("attributeName"),
//   		Enabled: jsii.Boolean(false),
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
//   			SeedCapacity: jsii.Number(123),
//   			TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   				DisableScaleIn: jsii.Boolean(false),
//   				ScaleInCooldown: jsii.Number(123),
//   				ScaleOutCooldown: jsii.Number(123),
//   				TargetValue: jsii.Number(123),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html
//
type CfnGlobalTablePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGlobalTableMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGlobalTablePropsMixin
type jsiiProxy_CfnGlobalTablePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGlobalTablePropsMixin) Props() *CfnGlobalTableMixinProps {
	var returns *CfnGlobalTableMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTablePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DynamoDB::GlobalTable`.
func NewCfnGlobalTablePropsMixin(props *CfnGlobalTableMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGlobalTablePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGlobalTablePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGlobalTablePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dynamodb.mixins.CfnGlobalTablePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DynamoDB::GlobalTable`.
func NewCfnGlobalTablePropsMixin_Override(c CfnGlobalTablePropsMixin, props *CfnGlobalTableMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dynamodb.mixins.CfnGlobalTablePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGlobalTablePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGlobalTablePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_dynamodb.mixins.CfnGlobalTablePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGlobalTablePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_dynamodb.mixins.CfnGlobalTablePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGlobalTablePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGlobalTablePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

