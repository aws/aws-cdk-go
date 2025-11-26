package previewawsdynamodbmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdynamodb/previewawsdynamodbmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DynamoDB::Table` resource creates a DynamoDB table. For more information, see [CreateTable](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html) in the *Amazon DynamoDB API Reference* .
//
// You should be aware of the following behaviors when working with DynamoDB tables:
//
// - AWS CloudFormation typically creates DynamoDB tables in parallel. However, if your template includes multiple DynamoDB tables with indexes, you must declare dependencies so that the tables are created sequentially. Amazon DynamoDB limits the number of tables with secondary indexes that are in the creating state. If you create multiple tables with indexes at the same time, DynamoDB returns an error and the stack operation fails. For an example, see [DynamoDB Table with a DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#aws-resource-dynamodb-table--examples--DynamoDB_Table_with_a_DependsOn_Attribute) .
//
// > Our guidance is to use the latest schema documented for your AWS CloudFormation templates. This schema supports the provisioning of all table settings below. When using this schema in your AWS CloudFormation templates, please ensure that your Identity and Access Management ( IAM ) policies are updated with appropriate permissions to allow for the authorization of these setting changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policyDocument interface{}
//
//   cfnTablePropsMixin := awscdkmixinspreview.Mixins.NewCfnTablePropsMixin(&CfnTableMixinProps{
//   	AttributeDefinitions: []interface{}{
//   		&AttributeDefinitionProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			AttributeType: jsii.String("attributeType"),
//   		},
//   	},
//   	BillingMode: jsii.String("billingMode"),
//   	ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   		Enabled: jsii.Boolean(false),
//   		Mode: jsii.String("mode"),
//   	},
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	GlobalSecondaryIndexes: []interface{}{
//   		&GlobalSecondaryIndexProperty{
//   			ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   				Enabled: jsii.Boolean(false),
//   				Mode: jsii.String("mode"),
//   			},
//   			IndexName: jsii.String("indexName"),
//   			KeySchema: []interface{}{
//   				&KeySchemaProperty{
//   					AttributeName: jsii.String("attributeName"),
//   					KeyType: jsii.String("keyType"),
//   				},
//   			},
//   			OnDemandThroughput: &OnDemandThroughputProperty{
//   				MaxReadRequestUnits: jsii.Number(123),
//   				MaxWriteRequestUnits: jsii.Number(123),
//   			},
//   			Projection: &ProjectionProperty{
//   				NonKeyAttributes: []*string{
//   					jsii.String("nonKeyAttributes"),
//   				},
//   				ProjectionType: jsii.String("projectionType"),
//   			},
//   			ProvisionedThroughput: &ProvisionedThroughputProperty{
//   				ReadCapacityUnits: jsii.Number(123),
//   				WriteCapacityUnits: jsii.Number(123),
//   			},
//   			WarmThroughput: &WarmThroughputProperty{
//   				ReadUnitsPerSecond: jsii.Number(123),
//   				WriteUnitsPerSecond: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ImportSourceSpecification: &ImportSourceSpecificationProperty{
//   		InputCompressionType: jsii.String("inputCompressionType"),
//   		InputFormat: jsii.String("inputFormat"),
//   		InputFormatOptions: &InputFormatOptionsProperty{
//   			Csv: &CsvProperty{
//   				Delimiter: jsii.String("delimiter"),
//   				HeaderList: []*string{
//   					jsii.String("headerList"),
//   				},
//   			},
//   		},
//   		S3BucketSource: &S3BucketSourceProperty{
//   			S3Bucket: jsii.String("s3Bucket"),
//   			S3BucketOwner: jsii.String("s3BucketOwner"),
//   			S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//   	},
//   	KeySchema: []interface{}{
//   		&KeySchemaProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			KeyType: jsii.String("keyType"),
//   		},
//   	},
//   	KinesisStreamSpecification: &KinesisStreamSpecificationProperty{
//   		ApproximateCreationDateTimePrecision: jsii.String("approximateCreationDateTimePrecision"),
//   		StreamArn: jsii.String("streamArn"),
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
//   	OnDemandThroughput: &OnDemandThroughputProperty{
//   		MaxReadRequestUnits: jsii.Number(123),
//   		MaxWriteRequestUnits: jsii.Number(123),
//   	},
//   	PointInTimeRecoverySpecification: &PointInTimeRecoverySpecificationProperty{
//   		PointInTimeRecoveryEnabled: jsii.Boolean(false),
//   		RecoveryPeriodInDays: jsii.Number(123),
//   	},
//   	ProvisionedThroughput: &ProvisionedThroughputProperty{
//   		ReadCapacityUnits: jsii.Number(123),
//   		WriteCapacityUnits: jsii.Number(123),
//   	},
//   	ResourcePolicy: &ResourcePolicyProperty{
//   		PolicyDocument: policyDocument,
//   	},
//   	SseSpecification: &SSESpecificationProperty{
//   		KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   		SseEnabled: jsii.Boolean(false),
//   		SseType: jsii.String("sseType"),
//   	},
//   	StreamSpecification: &StreamSpecificationProperty{
//   		ResourcePolicy: &ResourcePolicyProperty{
//   			PolicyDocument: policyDocument,
//   		},
//   		StreamViewType: jsii.String("streamViewType"),
//   	},
//   	TableClass: jsii.String("tableClass"),
//   	TableName: jsii.String("tableName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeToLiveSpecification: &TimeToLiveSpecificationProperty{
//   		AttributeName: jsii.String("attributeName"),
//   		Enabled: jsii.Boolean(false),
//   	},
//   	WarmThroughput: &WarmThroughputProperty{
//   		ReadUnitsPerSecond: jsii.Number(123),
//   		WriteUnitsPerSecond: jsii.Number(123),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html
//
type CfnTablePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTableMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTablePropsMixin
type jsiiProxy_CfnTablePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTablePropsMixin) Props() *CfnTableMixinProps {
	var returns *CfnTableMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTablePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DynamoDB::Table`.
func NewCfnTablePropsMixin(props *CfnTableMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTablePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTablePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTablePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dynamodb.mixins.CfnTablePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DynamoDB::Table`.
func NewCfnTablePropsMixin_Override(c CfnTablePropsMixin, props *CfnTableMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dynamodb.mixins.CfnTablePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTablePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTablePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_dynamodb.mixins.CfnTablePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTablePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_dynamodb.mixins.CfnTablePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTablePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTablePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

