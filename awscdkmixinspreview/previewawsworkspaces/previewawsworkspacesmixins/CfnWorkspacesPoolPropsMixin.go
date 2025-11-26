package previewawsworkspacesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsworkspaces/previewawsworkspacesmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes a pool of WorkSpaces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkspacesPoolPropsMixin := awscdkmixinspreview.Mixins.NewCfnWorkspacesPoolPropsMixin(&CfnWorkspacesPoolMixinProps{
//   	ApplicationSettings: &ApplicationSettingsProperty{
//   		SettingsGroup: jsii.String("settingsGroup"),
//   		Status: jsii.String("status"),
//   	},
//   	BundleId: jsii.String("bundleId"),
//   	Capacity: &CapacityProperty{
//   		DesiredUserSessions: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	DirectoryId: jsii.String("directoryId"),
//   	PoolName: jsii.String("poolName"),
//   	RunningMode: jsii.String("runningMode"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeoutSettings: &TimeoutSettingsProperty{
//   		DisconnectTimeoutInSeconds: jsii.Number(123),
//   		IdleDisconnectTimeoutInSeconds: jsii.Number(123),
//   		MaxUserDurationInSeconds: jsii.Number(123),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html
//
type CfnWorkspacesPoolPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWorkspacesPoolMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkspacesPoolPropsMixin
type jsiiProxy_CfnWorkspacesPoolPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWorkspacesPoolPropsMixin) Props() *CfnWorkspacesPoolMixinProps {
	var returns *CfnWorkspacesPoolMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspacesPoolPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WorkSpaces::WorkspacesPool`.
func NewCfnWorkspacesPoolPropsMixin(props *CfnWorkspacesPoolMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWorkspacesPoolPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkspacesPoolPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkspacesPoolPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspaces.mixins.CfnWorkspacesPoolPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WorkSpaces::WorkspacesPool`.
func NewCfnWorkspacesPoolPropsMixin_Override(c CfnWorkspacesPoolPropsMixin, props *CfnWorkspacesPoolMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspaces.mixins.CfnWorkspacesPoolPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWorkspacesPoolPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspacesPoolPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_workspaces.mixins.CfnWorkspacesPoolPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkspacesPoolPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_workspaces.mixins.CfnWorkspacesPoolPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkspacesPoolPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnWorkspacesPoolPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

