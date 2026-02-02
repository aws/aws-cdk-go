package previewawscodepipelinemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscodepipeline/previewawscodepipelinemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CodePipeline::CustomActionType` resource creates a custom action for activities that aren't included in the CodePipeline default actions, such as running an internally developed build process or a test suite.
//
// You can use these custom actions in the stage of a pipeline. For more information, see [Create and Add a Custom Action in AWS CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html) in the *AWS CodePipeline User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCustomActionTypePropsMixin := awscdkmixinspreview.Mixins.NewCfnCustomActionTypePropsMixin(&CfnCustomActionTypeMixinProps{
//   	Category: jsii.String("category"),
//   	ConfigurationProperties: []interface{}{
//   		&ConfigurationPropertiesProperty{
//   			Description: jsii.String("description"),
//   			Key: jsii.Boolean(false),
//   			Name: jsii.String("name"),
//   			Queryable: jsii.Boolean(false),
//   			Required: jsii.Boolean(false),
//   			Secret: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	InputArtifactDetails: &ArtifactDetailsProperty{
//   		MaximumCount: jsii.Number(123),
//   		MinimumCount: jsii.Number(123),
//   	},
//   	OutputArtifactDetails: &ArtifactDetailsProperty{
//   		MaximumCount: jsii.Number(123),
//   		MinimumCount: jsii.Number(123),
//   	},
//   	Provider: jsii.String("provider"),
//   	Settings: &SettingsProperty{
//   		EntityUrlTemplate: jsii.String("entityUrlTemplate"),
//   		ExecutionUrlTemplate: jsii.String("executionUrlTemplate"),
//   		RevisionUrlTemplate: jsii.String("revisionUrlTemplate"),
//   		ThirdPartyConfigurationUrl: jsii.String("thirdPartyConfigurationUrl"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Version: jsii.String("version"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html
//
type CfnCustomActionTypePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCustomActionTypeMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCustomActionTypePropsMixin
type jsiiProxy_CfnCustomActionTypePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCustomActionTypePropsMixin) Props() *CfnCustomActionTypeMixinProps {
	var returns *CfnCustomActionTypeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomActionTypePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodePipeline::CustomActionType`.
func NewCfnCustomActionTypePropsMixin(props *CfnCustomActionTypeMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCustomActionTypePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCustomActionTypePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCustomActionTypePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnCustomActionTypePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodePipeline::CustomActionType`.
func NewCfnCustomActionTypePropsMixin_Override(c CfnCustomActionTypePropsMixin, props *CfnCustomActionTypeMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnCustomActionTypePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCustomActionTypePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomActionTypePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnCustomActionTypePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCustomActionTypePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnCustomActionTypePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomActionTypePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCustomActionTypePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

