package previewawsworkspacesthinclientmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsworkspacesthinclient/previewawsworkspacesthinclientmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes an environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEnvironmentPropsMixin := awscdkmixinspreview.Mixins.NewCfnEnvironmentPropsMixin(&CfnEnvironmentMixinProps{
//   	DesiredSoftwareSetId: jsii.String("desiredSoftwareSetId"),
//   	DesktopArn: jsii.String("desktopArn"),
//   	DesktopEndpoint: jsii.String("desktopEndpoint"),
//   	DeviceCreationTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	MaintenanceWindow: &MaintenanceWindowProperty{
//   		ApplyTimeOf: jsii.String("applyTimeOf"),
//   		DaysOfTheWeek: []*string{
//   			jsii.String("daysOfTheWeek"),
//   		},
//   		EndTimeHour: jsii.Number(123),
//   		EndTimeMinute: jsii.Number(123),
//   		StartTimeHour: jsii.Number(123),
//   		StartTimeMinute: jsii.Number(123),
//   		Type: jsii.String("type"),
//   	},
//   	Name: jsii.String("name"),
//   	SoftwareSetUpdateMode: jsii.String("softwareSetUpdateMode"),
//   	SoftwareSetUpdateSchedule: jsii.String("softwareSetUpdateSchedule"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html
//
type CfnEnvironmentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEnvironmentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEnvironmentPropsMixin
type jsiiProxy_CfnEnvironmentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEnvironmentPropsMixin) Props() *CfnEnvironmentMixinProps {
	var returns *CfnEnvironmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WorkSpacesThinClient::Environment`.
func NewCfnEnvironmentPropsMixin(props *CfnEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEnvironmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEnvironmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnvironmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspacesthinclient.mixins.CfnEnvironmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WorkSpacesThinClient::Environment`.
func NewCfnEnvironmentPropsMixin_Override(c CfnEnvironmentPropsMixin, props *CfnEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspacesthinclient.mixins.CfnEnvironmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEnvironmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_workspacesthinclient.mixins.CfnEnvironmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnvironmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_workspacesthinclient.mixins.CfnEnvironmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnvironmentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEnvironmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

