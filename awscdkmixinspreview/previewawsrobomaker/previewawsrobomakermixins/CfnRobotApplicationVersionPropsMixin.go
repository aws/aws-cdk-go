package previewawsrobomakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrobomaker/previewawsrobomakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::RoboMaker::RobotApplicationVersion` resource creates an AWS RoboMaker robot version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRobotApplicationVersionPropsMixin := awscdkmixinspreview.Mixins.NewCfnRobotApplicationVersionPropsMixin(&CfnRobotApplicationVersionMixinProps{
//   	Application: jsii.String("application"),
//   	CurrentRevisionId: jsii.String("currentRevisionId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-robomaker-robotapplicationversion.html
//
type CfnRobotApplicationVersionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRobotApplicationVersionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRobotApplicationVersionPropsMixin
type jsiiProxy_CfnRobotApplicationVersionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRobotApplicationVersionPropsMixin) Props() *CfnRobotApplicationVersionMixinProps {
	var returns *CfnRobotApplicationVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRobotApplicationVersionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RoboMaker::RobotApplicationVersion`.
func NewCfnRobotApplicationVersionPropsMixin(props *CfnRobotApplicationVersionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRobotApplicationVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRobotApplicationVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRobotApplicationVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_robomaker.mixins.CfnRobotApplicationVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RoboMaker::RobotApplicationVersion`.
func NewCfnRobotApplicationVersionPropsMixin_Override(c CfnRobotApplicationVersionPropsMixin, props *CfnRobotApplicationVersionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_robomaker.mixins.CfnRobotApplicationVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRobotApplicationVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRobotApplicationVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_robomaker.mixins.CfnRobotApplicationVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRobotApplicationVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_robomaker.mixins.CfnRobotApplicationVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRobotApplicationVersionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRobotApplicationVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

