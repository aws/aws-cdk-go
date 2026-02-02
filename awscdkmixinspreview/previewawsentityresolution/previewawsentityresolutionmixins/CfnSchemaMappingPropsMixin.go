package previewawsentityresolutionmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsentityresolution/previewawsentityresolutionmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a schema mapping, which defines the schema of the input customer records table.
//
// The `SchemaMapping` also provides AWS Entity Resolution with some metadata about the table, such as the attribute types of the columns and which columns to match on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSchemaMappingPropsMixin := awscdkmixinspreview.Mixins.NewCfnSchemaMappingPropsMixin(&CfnSchemaMappingMixinProps{
//   	Description: jsii.String("description"),
//   	MappedInputFields: []interface{}{
//   		&SchemaInputAttributeProperty{
//   			FieldName: jsii.String("fieldName"),
//   			GroupName: jsii.String("groupName"),
//   			Hashed: jsii.Boolean(false),
//   			MatchKey: jsii.String("matchKey"),
//   			SubType: jsii.String("subType"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SchemaName: jsii.String("schemaName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-schemamapping.html
//
type CfnSchemaMappingPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSchemaMappingMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSchemaMappingPropsMixin
type jsiiProxy_CfnSchemaMappingPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSchemaMappingPropsMixin) Props() *CfnSchemaMappingMixinProps {
	var returns *CfnSchemaMappingMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaMappingPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EntityResolution::SchemaMapping`.
func NewCfnSchemaMappingPropsMixin(props *CfnSchemaMappingMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSchemaMappingPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSchemaMappingPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSchemaMappingPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnSchemaMappingPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EntityResolution::SchemaMapping`.
func NewCfnSchemaMappingPropsMixin_Override(c CfnSchemaMappingPropsMixin, props *CfnSchemaMappingMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnSchemaMappingPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSchemaMappingPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSchemaMappingPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnSchemaMappingPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSchemaMappingPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnSchemaMappingPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSchemaMappingPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSchemaMappingPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

