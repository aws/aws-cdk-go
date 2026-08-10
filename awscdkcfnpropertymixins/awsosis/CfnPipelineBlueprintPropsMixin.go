package awsosis

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsosis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An OpenSearch Ingestion blueprint, which is a preconfigured template for creating an ingestion pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPipelineBlueprintPropsMixin := awscdkcfnpropertymixins.Aws_osis.NewCfnPipelineBlueprintPropsMixin(&CfnPipelineBlueprintMixinProps{
//   	BlueprintName: jsii.String("blueprintName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipelineblueprint.html
//
type CfnPipelineBlueprintPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPipelineBlueprintMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPipelineBlueprintPropsMixin
type jsiiProxy_CfnPipelineBlueprintPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPipelineBlueprintPropsMixin) Props() *CfnPipelineBlueprintMixinProps {
	var returns *CfnPipelineBlueprintMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipelineBlueprintPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::OSIS::PipelineBlueprint`.
func NewCfnPipelineBlueprintPropsMixin(props *CfnPipelineBlueprintMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPipelineBlueprintPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPipelineBlueprintPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPipelineBlueprintPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_osis.CfnPipelineBlueprintPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::OSIS::PipelineBlueprint`.
func NewCfnPipelineBlueprintPropsMixin_Override(c CfnPipelineBlueprintPropsMixin, props *CfnPipelineBlueprintMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_osis.CfnPipelineBlueprintPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPipelineBlueprintPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPipelineBlueprintPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_osis.CfnPipelineBlueprintPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPipelineBlueprintPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_osis.CfnPipelineBlueprintPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPipelineBlueprintPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPipelineBlueprintPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

