package previewawsgluemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsglue/previewawsgluemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::Glue::MLTransform is an AWS Glue resource type that manages machine learning transforms.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   cfnMLTransformPropsMixin := awscdkmixinspreview.Mixins.NewCfnMLTransformPropsMixin(&CfnMLTransformMixinProps{
//   	Description: jsii.String("description"),
//   	GlueVersion: jsii.String("glueVersion"),
//   	InputRecordTables: &InputRecordTablesProperty{
//   		GlueTables: []interface{}{
//   			&GlueTablesProperty{
//   				CatalogId: jsii.String("catalogId"),
//   				ConnectionName: jsii.String("connectionName"),
//   				DatabaseName: jsii.String("databaseName"),
//   				TableName: jsii.String("tableName"),
//   			},
//   		},
//   	},
//   	MaxCapacity: jsii.Number(123),
//   	MaxRetries: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	NumberOfWorkers: jsii.Number(123),
//   	Role: jsii.String("role"),
//   	Tags: tags,
//   	Timeout: jsii.Number(123),
//   	TransformEncryption: &TransformEncryptionProperty{
//   		MlUserDataEncryption: &MLUserDataEncryptionProperty{
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			MlUserDataEncryptionMode: jsii.String("mlUserDataEncryptionMode"),
//   		},
//   		TaskRunSecurityConfigurationName: jsii.String("taskRunSecurityConfigurationName"),
//   	},
//   	TransformParameters: &TransformParametersProperty{
//   		FindMatchesParameters: &FindMatchesParametersProperty{
//   			AccuracyCostTradeoff: jsii.Number(123),
//   			EnforceProvidedLabels: jsii.Boolean(false),
//   			PrecisionRecallTradeoff: jsii.Number(123),
//   			PrimaryKeyColumnName: jsii.String("primaryKeyColumnName"),
//   		},
//   		TransformType: jsii.String("transformType"),
//   	},
//   	WorkerType: jsii.String("workerType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html
//
type CfnMLTransformPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMLTransformMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMLTransformPropsMixin
type jsiiProxy_CfnMLTransformPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMLTransformPropsMixin) Props() *CfnMLTransformMixinProps {
	var returns *CfnMLTransformMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransformPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::MLTransform`.
func NewCfnMLTransformPropsMixin(props *CfnMLTransformMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMLTransformPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMLTransformPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMLTransformPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnMLTransformPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::MLTransform`.
func NewCfnMLTransformPropsMixin_Override(c CfnMLTransformPropsMixin, props *CfnMLTransformMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnMLTransformPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMLTransformPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMLTransformPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnMLTransformPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMLTransformPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnMLTransformPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMLTransformPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMLTransformPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

