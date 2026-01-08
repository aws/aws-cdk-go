package previewawsecrmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsecr/previewawsecrmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The ARN of the IAM principal to remove from the pull time update exclusion list.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPullTimeUpdateExclusionPropsMixin := awscdkmixinspreview.Mixins.NewCfnPullTimeUpdateExclusionPropsMixin(&CfnPullTimeUpdateExclusionMixinProps{
//   	PrincipalArn: jsii.String("principalArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pulltimeupdateexclusion.html
//
type CfnPullTimeUpdateExclusionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPullTimeUpdateExclusionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPullTimeUpdateExclusionPropsMixin
type jsiiProxy_CfnPullTimeUpdateExclusionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPullTimeUpdateExclusionPropsMixin) Props() *CfnPullTimeUpdateExclusionMixinProps {
	var returns *CfnPullTimeUpdateExclusionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPullTimeUpdateExclusionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ECR::PullTimeUpdateExclusion`.
func NewCfnPullTimeUpdateExclusionPropsMixin(props *CfnPullTimeUpdateExclusionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPullTimeUpdateExclusionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPullTimeUpdateExclusionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPullTimeUpdateExclusionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.mixins.CfnPullTimeUpdateExclusionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ECR::PullTimeUpdateExclusion`.
func NewCfnPullTimeUpdateExclusionPropsMixin_Override(c CfnPullTimeUpdateExclusionPropsMixin, props *CfnPullTimeUpdateExclusionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.mixins.CfnPullTimeUpdateExclusionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPullTimeUpdateExclusionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPullTimeUpdateExclusionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecr.mixins.CfnPullTimeUpdateExclusionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPullTimeUpdateExclusionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ecr.mixins.CfnPullTimeUpdateExclusionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPullTimeUpdateExclusionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPullTimeUpdateExclusionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

