package previewawsemrmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsemr/previewawsemrmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::EMR::StudioSessionMapping` resource is an Amazon EMR resource type that maps a user or group to the Amazon EMR Studio specified by `StudioId` , and applies a session policy that defines Studio permissions for that user or group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStudioSessionMappingPropsMixin := awscdkmixinspreview.Mixins.NewCfnStudioSessionMappingPropsMixin(&CfnStudioSessionMappingMixinProps{
//   	IdentityName: jsii.String("identityName"),
//   	IdentityType: jsii.String("identityType"),
//   	SessionPolicyArn: jsii.String("sessionPolicyArn"),
//   	StudioId: jsii.String("studioId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studiosessionmapping.html
//
type CfnStudioSessionMappingPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnStudioSessionMappingMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStudioSessionMappingPropsMixin
type jsiiProxy_CfnStudioSessionMappingPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStudioSessionMappingPropsMixin) Props() *CfnStudioSessionMappingMixinProps {
	var returns *CfnStudioSessionMappingMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMappingPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EMR::StudioSessionMapping`.
func NewCfnStudioSessionMappingPropsMixin(props *CfnStudioSessionMappingMixinProps, options *mixins.CfnPropertyMixinOptions) CfnStudioSessionMappingPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStudioSessionMappingPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStudioSessionMappingPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnStudioSessionMappingPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EMR::StudioSessionMapping`.
func NewCfnStudioSessionMappingPropsMixin_Override(c CfnStudioSessionMappingPropsMixin, props *CfnStudioSessionMappingMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnStudioSessionMappingPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStudioSessionMappingPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStudioSessionMappingPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnStudioSessionMappingPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStudioSessionMappingPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnStudioSessionMappingPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStudioSessionMappingPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnStudioSessionMappingPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

