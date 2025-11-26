package previewawsnimblestudiomixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsnimblestudio/previewawsnimblestudiomixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studiocomponent.html.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStudioComponentPropsMixin := awscdkmixinspreview.Mixins.NewCfnStudioComponentPropsMixin(&CfnStudioComponentMixinProps{
//   	Configuration: &StudioComponentConfigurationProperty{
//   		ActiveDirectoryConfiguration: &ActiveDirectoryConfigurationProperty{
//   			ComputerAttributes: []interface{}{
//   				&ActiveDirectoryComputerAttributeProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			DirectoryId: jsii.String("directoryId"),
//   			OrganizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   		},
//   		ComputeFarmConfiguration: &ComputeFarmConfigurationProperty{
//   			ActiveDirectoryUser: jsii.String("activeDirectoryUser"),
//   			Endpoint: jsii.String("endpoint"),
//   		},
//   		LicenseServiceConfiguration: &LicenseServiceConfigurationProperty{
//   			Endpoint: jsii.String("endpoint"),
//   		},
//   		SharedFileSystemConfiguration: &SharedFileSystemConfigurationProperty{
//   			Endpoint: jsii.String("endpoint"),
//   			FileSystemId: jsii.String("fileSystemId"),
//   			LinuxMountPoint: jsii.String("linuxMountPoint"),
//   			ShareName: jsii.String("shareName"),
//   			WindowsMountDrive: jsii.String("windowsMountDrive"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Ec2SecurityGroupIds: []*string{
//   		jsii.String("ec2SecurityGroupIds"),
//   	},
//   	InitializationScripts: []interface{}{
//   		&StudioComponentInitializationScriptProperty{
//   			LaunchProfileProtocolVersion: jsii.String("launchProfileProtocolVersion"),
//   			Platform: jsii.String("platform"),
//   			RunContext: jsii.String("runContext"),
//   			Script: jsii.String("script"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	ScriptParameters: []interface{}{
//   		&ScriptParameterKeyValueProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	StudioId: jsii.String("studioId"),
//   	Subtype: jsii.String("subtype"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studiocomponent.html
//
type CfnStudioComponentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnStudioComponentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStudioComponentPropsMixin
type jsiiProxy_CfnStudioComponentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStudioComponentPropsMixin) Props() *CfnStudioComponentMixinProps {
	var returns *CfnStudioComponentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioComponentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NimbleStudio::StudioComponent`.
func NewCfnStudioComponentPropsMixin(props *CfnStudioComponentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnStudioComponentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStudioComponentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStudioComponentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_nimblestudio.mixins.CfnStudioComponentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NimbleStudio::StudioComponent`.
func NewCfnStudioComponentPropsMixin_Override(c CfnStudioComponentPropsMixin, props *CfnStudioComponentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_nimblestudio.mixins.CfnStudioComponentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStudioComponentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStudioComponentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_nimblestudio.mixins.CfnStudioComponentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStudioComponentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_nimblestudio.mixins.CfnStudioComponentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStudioComponentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnStudioComponentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

