package previewawscodestarconnectionsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscodestarconnections/previewawscodestarconnectionsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Information, such as repository, branch, provider, and resource names for a specific sync configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSyncConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnSyncConfigurationPropsMixin(&CfnSyncConfigurationMixinProps{
//   	Branch: jsii.String("branch"),
//   	ConfigFile: jsii.String("configFile"),
//   	PublishDeploymentStatus: jsii.String("publishDeploymentStatus"),
//   	RepositoryLinkId: jsii.String("repositoryLinkId"),
//   	ResourceName: jsii.String("resourceName"),
//   	RoleArn: jsii.String("roleArn"),
//   	SyncType: jsii.String("syncType"),
//   	TriggerResourceUpdateOn: jsii.String("triggerResourceUpdateOn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-syncconfiguration.html
//
type CfnSyncConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSyncConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSyncConfigurationPropsMixin
type jsiiProxy_CfnSyncConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSyncConfigurationPropsMixin) Props() *CfnSyncConfigurationMixinProps {
	var returns *CfnSyncConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSyncConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodeStarConnections::SyncConfiguration`.
func NewCfnSyncConfigurationPropsMixin(props *CfnSyncConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSyncConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSyncConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSyncConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codestarconnections.mixins.CfnSyncConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodeStarConnections::SyncConfiguration`.
func NewCfnSyncConfigurationPropsMixin_Override(c CfnSyncConfigurationPropsMixin, props *CfnSyncConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codestarconnections.mixins.CfnSyncConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSyncConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSyncConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codestarconnections.mixins.CfnSyncConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSyncConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codestarconnections.mixins.CfnSyncConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSyncConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSyncConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

