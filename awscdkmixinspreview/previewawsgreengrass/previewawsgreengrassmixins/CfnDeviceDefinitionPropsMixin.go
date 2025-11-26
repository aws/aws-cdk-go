package previewawsgreengrassmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgreengrass/previewawsgreengrassmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Greengrass::DeviceDefinition` resource represents a device definition for AWS IoT Greengrass .
//
// Device definitions are used to organize your device definition versions.
//
// Device definitions can reference multiple device definition versions. All device definition versions must be associated with a device definition. Each device definition version can contain one or more devices.
//
// > When you create a device definition, you can optionally include an initial device definition version. To associate a device definition version later, create an [`AWS::Greengrass::DeviceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html) resource and specify the ID of this device definition.
// >
// > After you create the device definition version that contains the devices you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   cfnDeviceDefinitionPropsMixin := awscdkmixinspreview.Mixins.NewCfnDeviceDefinitionPropsMixin(&CfnDeviceDefinitionMixinProps{
//   	InitialVersion: &DeviceDefinitionVersionProperty{
//   		Devices: []interface{}{
//   			&DeviceProperty{
//   				CertificateArn: jsii.String("certificateArn"),
//   				Id: jsii.String("id"),
//   				SyncShadow: jsii.Boolean(false),
//   				ThingArn: jsii.String("thingArn"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: tags,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html
//
type CfnDeviceDefinitionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDeviceDefinitionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDeviceDefinitionPropsMixin
type jsiiProxy_CfnDeviceDefinitionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDeviceDefinitionPropsMixin) Props() *CfnDeviceDefinitionMixinProps {
	var returns *CfnDeviceDefinitionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeviceDefinitionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Greengrass::DeviceDefinition`.
func NewCfnDeviceDefinitionPropsMixin(props *CfnDeviceDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDeviceDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDeviceDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDeviceDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnDeviceDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Greengrass::DeviceDefinition`.
func NewCfnDeviceDefinitionPropsMixin_Override(c CfnDeviceDefinitionPropsMixin, props *CfnDeviceDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnDeviceDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDeviceDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDeviceDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnDeviceDefinitionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeviceDefinitionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnDeviceDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeviceDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDeviceDefinitionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

