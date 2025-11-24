package mixinsawss3tables

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awss3tables/mixinsawss3tables/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new table associated with the given namespace in a table bucket.
//
// For more information, see [Creating an Amazon S3 table](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-create.html) in the *Amazon Simple Storage Service User Guide* .
//
// - **Permissions** - - You must have the `s3tables:CreateTable` permission to use this operation.
// - If you use this operation with the optional `metadata` request parameter you must have the `s3tables:PutTableData` permission.
// - If you use this operation with the optional `encryptionConfiguration` request parameter you must have the `s3tables:PutTableEncryption` permission.
//
// > Additionally, If you choose SSE-KMS encryption you must grant the S3 Tables maintenance principal access to your KMS key. For more information, see [Permissions requirements for S3 Tables SSE-KMS encryption](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-kms-permissions.html) .
// - **Cloud Development Kit** - To use S3 Tables AWS CDK constructs, add the `@aws-cdk/aws-s3tables-alpha` dependency with one of the following options:
//
// - NPM: `npm i.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTablePropsMixin := awscdkmixinspreview.Mixins.NewCfnTablePropsMixin(&CfnTableMixinProps{
//   	Compaction: &CompactionProperty{
//   		Status: jsii.String("status"),
//   		TargetFileSizeMb: jsii.Number(123),
//   	},
//   	IcebergMetadata: &IcebergMetadataProperty{
//   		IcebergSchema: &IcebergSchemaProperty{
//   			SchemaFieldList: []interface{}{
//   				&SchemaFieldProperty{
//   					Name: jsii.String("name"),
//   					Required: jsii.Boolean(false),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   	Namespace: jsii.String("namespace"),
//   	OpenTableFormat: jsii.String("openTableFormat"),
//   	SnapshotManagement: &SnapshotManagementProperty{
//   		MaxSnapshotAgeHours: jsii.Number(123),
//   		MinSnapshotsToKeep: jsii.Number(123),
//   		Status: jsii.String("status"),
//   	},
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   	TableName: jsii.String("tableName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WithoutMetadata: jsii.String("withoutMetadata"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3tables-table.html
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


// Create a mixin to apply properties to `AWS::S3Tables::Table`.
func NewCfnTablePropsMixin(props *CfnTableMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTablePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTablePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTablePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3tables.mixins.CfnTablePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3Tables::Table`.
func NewCfnTablePropsMixin_Override(c CfnTablePropsMixin, props *CfnTableMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3tables.mixins.CfnTablePropsMixin",
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
		"@aws-cdk/mixins-preview.aws_s3tables.mixins.CfnTablePropsMixin",
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
		"@aws-cdk/mixins-preview.aws_s3tables.mixins.CfnTablePropsMixin",
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

