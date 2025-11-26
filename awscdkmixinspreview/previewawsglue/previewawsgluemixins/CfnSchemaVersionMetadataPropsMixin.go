package previewawsgluemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsglue/previewawsgluemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Glue::SchemaVersionMetadata` is an AWS Glue resource type that defines the metadata key-value pairs for a schema version in AWS Glue Schema Registry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSchemaVersionMetadataPropsMixin := awscdkmixinspreview.Mixins.NewCfnSchemaVersionMetadataPropsMixin(&CfnSchemaVersionMetadataMixinProps{
//   	Key: jsii.String("key"),
//   	SchemaVersionId: jsii.String("schemaVersionId"),
//   	Value: jsii.String("value"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schemaversionmetadata.html
//
type CfnSchemaVersionMetadataPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSchemaVersionMetadataMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSchemaVersionMetadataPropsMixin
type jsiiProxy_CfnSchemaVersionMetadataPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSchemaVersionMetadataPropsMixin) Props() *CfnSchemaVersionMetadataMixinProps {
	var returns *CfnSchemaVersionMetadataMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersionMetadataPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::SchemaVersionMetadata`.
func NewCfnSchemaVersionMetadataPropsMixin(props *CfnSchemaVersionMetadataMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSchemaVersionMetadataPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSchemaVersionMetadataPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSchemaVersionMetadataPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSchemaVersionMetadataPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::SchemaVersionMetadata`.
func NewCfnSchemaVersionMetadataPropsMixin_Override(c CfnSchemaVersionMetadataPropsMixin, props *CfnSchemaVersionMetadataMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSchemaVersionMetadataPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSchemaVersionMetadataPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSchemaVersionMetadataPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSchemaVersionMetadataPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSchemaVersionMetadataPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnSchemaVersionMetadataPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSchemaVersionMetadataPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSchemaVersionMetadataPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

