package awsdynamodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdynamodb/internal"
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
//   // TableV2 uses a Box internally for replicas.
//   // The mixin defers the merge and appends the new replica at synthesis time.
//   table := dynamodb.NewTableV2(scope, jsii.String("Table"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	// L2 prop: pointInTimeRecovery is a boolean
//   	Replicas: []ReplicaTableProps{
//   		&ReplicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   			PointInTimeRecovery: jsii.Boolean(true),
//   		},
//   	},
//   })
//
//   // Mixins always use L1 (CloudFormation) property names and shapes,
//   // regardless of what the L2 API looks like.
//   table.With(awscdkcfnpropertymixins.NewCfnGlobalTablePropsMixin(&CfnGlobalTableMixinProps{
//   	Replicas: []interface{}{
//   		&ReplicaSpecificationProperty{
//   			Region: jsii.String("eu-west-1"),
//   			// L1 prop: pointInTimeRecoverySpecification is an object
//   			PointInTimeRecoverySpecification: &PointInTimeRecoverySpecificationProperty{
//   				PointInTimeRecoveryEnabled: jsii.Boolean(true),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdk.PropertyMergeStrategy_Combine(&CombineStrategyOptions{
//   		Arrays: awscdk.ArrayMergeStrategy_Append(),
//   	}),
//   }))
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html
//
type CfnGlobalTablePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnGlobalTableMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGlobalTablePropsMixin
type jsiiProxy_CfnGlobalTablePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnGlobalTablePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DynamoDB::GlobalTable`.
func NewCfnGlobalTablePropsMixin(props *CfnGlobalTableMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnGlobalTablePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGlobalTablePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGlobalTablePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dynamodb.CfnGlobalTablePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DynamoDB::GlobalTable`.
func NewCfnGlobalTablePropsMixin_Override(c CfnGlobalTablePropsMixin, props *CfnGlobalTableMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_dynamodb.CfnGlobalTablePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnGlobalTablePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGlobalTablePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_dynamodb.CfnGlobalTablePropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_dynamodb.CfnGlobalTablePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGlobalTablePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

