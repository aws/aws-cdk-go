package awsimagebuilder

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Returns the list of workflow build versions for a specified workflow version ARN.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAllWorkflowBuildVersionsPropsMixin := awscdkcfnpropertymixins.Aws_imagebuilder.NewCfnAllWorkflowBuildVersionsPropsMixin(&CfnAllWorkflowBuildVersionsMixinProps{
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-allworkflowbuildversions.html
//
type CfnAllWorkflowBuildVersionsPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAllWorkflowBuildVersionsMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAllWorkflowBuildVersionsPropsMixin
type jsiiProxy_CfnAllWorkflowBuildVersionsPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAllWorkflowBuildVersionsPropsMixin) Props() *CfnAllWorkflowBuildVersionsMixinProps {
	var returns *CfnAllWorkflowBuildVersionsMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAllWorkflowBuildVersionsPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ImageBuilder::AllWorkflowBuildVersions`.
func NewCfnAllWorkflowBuildVersionsPropsMixin(props *CfnAllWorkflowBuildVersionsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAllWorkflowBuildVersionsPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAllWorkflowBuildVersionsPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAllWorkflowBuildVersionsPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnAllWorkflowBuildVersionsPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ImageBuilder::AllWorkflowBuildVersions`.
func NewCfnAllWorkflowBuildVersionsPropsMixin_Override(c CfnAllWorkflowBuildVersionsPropsMixin, props *CfnAllWorkflowBuildVersionsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnAllWorkflowBuildVersionsPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAllWorkflowBuildVersionsPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAllWorkflowBuildVersionsPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnAllWorkflowBuildVersionsPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAllWorkflowBuildVersionsPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnAllWorkflowBuildVersionsPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAllWorkflowBuildVersionsPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAllWorkflowBuildVersionsPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

