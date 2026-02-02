package previewawsnimblestudiomixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsnimblestudio/previewawsnimblestudiomixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLaunchProfilePropsMixin := awscdkmixinspreview.Mixins.NewCfnLaunchProfilePropsMixin(&CfnLaunchProfileMixinProps{
//   	Description: jsii.String("description"),
//   	Ec2SubnetIds: []*string{
//   		jsii.String("ec2SubnetIds"),
//   	},
//   	LaunchProfileProtocolVersions: []*string{
//   		jsii.String("launchProfileProtocolVersions"),
//   	},
//   	Name: jsii.String("name"),
//   	StreamConfiguration: &StreamConfigurationProperty{
//   		AutomaticTerminationMode: jsii.String("automaticTerminationMode"),
//   		ClipboardMode: jsii.String("clipboardMode"),
//   		Ec2InstanceTypes: []*string{
//   			jsii.String("ec2InstanceTypes"),
//   		},
//   		MaxSessionLengthInMinutes: jsii.Number(123),
//   		MaxStoppedSessionLengthInMinutes: jsii.Number(123),
//   		SessionBackup: &StreamConfigurationSessionBackupProperty{
//   			MaxBackupsToRetain: jsii.Number(123),
//   			Mode: jsii.String("mode"),
//   		},
//   		SessionPersistenceMode: jsii.String("sessionPersistenceMode"),
//   		SessionStorage: &StreamConfigurationSessionStorageProperty{
//   			Mode: []*string{
//   				jsii.String("mode"),
//   			},
//   			Root: &StreamingSessionStorageRootProperty{
//   				Linux: jsii.String("linux"),
//   				Windows: jsii.String("windows"),
//   			},
//   		},
//   		StreamingImageIds: []*string{
//   			jsii.String("streamingImageIds"),
//   		},
//   		VolumeConfiguration: &VolumeConfigurationProperty{
//   			Iops: jsii.Number(123),
//   			Size: jsii.Number(123),
//   			Throughput: jsii.Number(123),
//   		},
//   	},
//   	StudioComponentIds: []*string{
//   		jsii.String("studioComponentIds"),
//   	},
//   	StudioId: jsii.String("studioId"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html
//
type CfnLaunchProfilePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLaunchProfileMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLaunchProfilePropsMixin
type jsiiProxy_CfnLaunchProfilePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLaunchProfilePropsMixin) Props() *CfnLaunchProfileMixinProps {
	var returns *CfnLaunchProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunchProfilePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NimbleStudio::LaunchProfile`.
func NewCfnLaunchProfilePropsMixin(props *CfnLaunchProfileMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLaunchProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLaunchProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLaunchProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_nimblestudio.mixins.CfnLaunchProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NimbleStudio::LaunchProfile`.
func NewCfnLaunchProfilePropsMixin_Override(c CfnLaunchProfilePropsMixin, props *CfnLaunchProfileMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_nimblestudio.mixins.CfnLaunchProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLaunchProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLaunchProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_nimblestudio.mixins.CfnLaunchProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLaunchProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_nimblestudio.mixins.CfnLaunchProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLaunchProfilePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLaunchProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

