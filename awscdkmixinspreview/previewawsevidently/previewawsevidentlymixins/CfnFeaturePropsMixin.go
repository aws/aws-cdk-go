package previewawsevidentlymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsevidently/previewawsevidentlymixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates an Evidently *feature* that you want to launch or test.
//
// You can define up to five variations of a feature, and use these variations in your launches and experiments. A feature must be created in a project. For information about creating a project, see [CreateProject](https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_CreateProject.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFeaturePropsMixin := awscdkmixinspreview.Mixins.NewCfnFeaturePropsMixin(&CfnFeatureMixinProps{
//   	DefaultVariation: jsii.String("defaultVariation"),
//   	Description: jsii.String("description"),
//   	EntityOverrides: []interface{}{
//   		&EntityOverrideProperty{
//   			EntityId: jsii.String("entityId"),
//   			Variation: jsii.String("variation"),
//   		},
//   	},
//   	EvaluationStrategy: jsii.String("evaluationStrategy"),
//   	Name: jsii.String("name"),
//   	Project: jsii.String("project"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Variations: []interface{}{
//   		&VariationObjectProperty{
//   			BooleanValue: jsii.Boolean(false),
//   			DoubleValue: jsii.Number(123),
//   			LongValue: jsii.Number(123),
//   			StringValue: jsii.String("stringValue"),
//   			VariationName: jsii.String("variationName"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-feature.html
//
type CfnFeaturePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFeatureMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFeaturePropsMixin
type jsiiProxy_CfnFeaturePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFeaturePropsMixin) Props() *CfnFeatureMixinProps {
	var returns *CfnFeatureMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeaturePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Evidently::Feature`.
func NewCfnFeaturePropsMixin(props *CfnFeatureMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFeaturePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFeaturePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFeaturePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnFeaturePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Evidently::Feature`.
func NewCfnFeaturePropsMixin_Override(c CfnFeaturePropsMixin, props *CfnFeatureMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnFeaturePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFeaturePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFeaturePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnFeaturePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFeaturePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_evidently.mixins.CfnFeaturePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFeaturePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFeaturePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

