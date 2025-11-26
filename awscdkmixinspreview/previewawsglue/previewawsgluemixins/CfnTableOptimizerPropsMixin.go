package previewawsgluemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsglue/previewawsgluemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AWS Glue resource for enabling table optimizers to improve read performance for open table formats.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTableOptimizerPropsMixin := awscdkmixinspreview.Mixins.NewCfnTableOptimizerPropsMixin(&CfnTableOptimizerMixinProps{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	TableName: jsii.String("tableName"),
//   	TableOptimizerConfiguration: &TableOptimizerConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		OrphanFileDeletionConfiguration: &OrphanFileDeletionConfigurationProperty{
//   			IcebergConfiguration: &IcebergConfigurationProperty{
//   				Location: jsii.String("location"),
//   				OrphanFileRetentionPeriodInDays: jsii.Number(123),
//   			},
//   		},
//   		RetentionConfiguration: &RetentionConfigurationProperty{
//   			IcebergConfiguration: &IcebergRetentionConfigurationProperty{
//   				CleanExpiredFiles: jsii.Boolean(false),
//   				NumberOfSnapshotsToRetain: jsii.Number(123),
//   				SnapshotRetentionPeriodInDays: jsii.Number(123),
//   			},
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			GlueConnectionName: jsii.String("glueConnectionName"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-tableoptimizer.html
//
type CfnTableOptimizerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTableOptimizerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTableOptimizerPropsMixin
type jsiiProxy_CfnTableOptimizerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTableOptimizerPropsMixin) Props() *CfnTableOptimizerMixinProps {
	var returns *CfnTableOptimizerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTableOptimizerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::TableOptimizer`.
func NewCfnTableOptimizerPropsMixin(props *CfnTableOptimizerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTableOptimizerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTableOptimizerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTableOptimizerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTableOptimizerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::TableOptimizer`.
func NewCfnTableOptimizerPropsMixin_Override(c CfnTableOptimizerPropsMixin, props *CfnTableOptimizerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTableOptimizerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTableOptimizerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTableOptimizerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTableOptimizerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTableOptimizerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTableOptimizerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTableOptimizerPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTableOptimizerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

