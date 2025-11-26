package previewawsprotonmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsproton/previewawsprotonmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Detailed data of an AWS Proton environment account connection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEnvironmentAccountConnectionPropsMixin := awscdkmixinspreview.Mixins.NewCfnEnvironmentAccountConnectionPropsMixin(&CfnEnvironmentAccountConnectionMixinProps{
//   	CodebuildRoleArn: jsii.String("codebuildRoleArn"),
//   	ComponentRoleArn: jsii.String("componentRoleArn"),
//   	EnvironmentAccountId: jsii.String("environmentAccountId"),
//   	EnvironmentName: jsii.String("environmentName"),
//   	ManagementAccountId: jsii.String("managementAccountId"),
//   	RoleArn: jsii.String("roleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-proton-environmentaccountconnection.html
//
type CfnEnvironmentAccountConnectionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEnvironmentAccountConnectionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEnvironmentAccountConnectionPropsMixin
type jsiiProxy_CfnEnvironmentAccountConnectionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEnvironmentAccountConnectionPropsMixin) Props() *CfnEnvironmentAccountConnectionMixinProps {
	var returns *CfnEnvironmentAccountConnectionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentAccountConnectionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Proton::EnvironmentAccountConnection`.
func NewCfnEnvironmentAccountConnectionPropsMixin(props *CfnEnvironmentAccountConnectionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEnvironmentAccountConnectionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEnvironmentAccountConnectionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnvironmentAccountConnectionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_proton.mixins.CfnEnvironmentAccountConnectionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Proton::EnvironmentAccountConnection`.
func NewCfnEnvironmentAccountConnectionPropsMixin_Override(c CfnEnvironmentAccountConnectionPropsMixin, props *CfnEnvironmentAccountConnectionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_proton.mixins.CfnEnvironmentAccountConnectionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEnvironmentAccountConnectionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironmentAccountConnectionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_proton.mixins.CfnEnvironmentAccountConnectionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnvironmentAccountConnectionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_proton.mixins.CfnEnvironmentAccountConnectionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnvironmentAccountConnectionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEnvironmentAccountConnectionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

