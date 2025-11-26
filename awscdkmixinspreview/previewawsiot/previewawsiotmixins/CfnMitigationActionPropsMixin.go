package previewawsiotmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiot/previewawsiotmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines an action that can be applied to audit findings by using StartAuditMitigationActionsTask.
//
// For API reference, see [CreateMitigationAction](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateMitigationAction.html) and for general information, see [Mitigation actions](https://docs.aws.amazon.com/iot/latest/developerguide/dd-mitigation-actions.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMitigationActionPropsMixin := awscdkmixinspreview.Mixins.NewCfnMitigationActionPropsMixin(&CfnMitigationActionMixinProps{
//   	ActionName: jsii.String("actionName"),
//   	ActionParams: &ActionParamsProperty{
//   		AddThingsToThingGroupParams: &AddThingsToThingGroupParamsProperty{
//   			OverrideDynamicGroups: jsii.Boolean(false),
//   			ThingGroupNames: []*string{
//   				jsii.String("thingGroupNames"),
//   			},
//   		},
//   		EnableIoTLoggingParams: &EnableIoTLoggingParamsProperty{
//   			LogLevel: jsii.String("logLevel"),
//   			RoleArnForLogging: jsii.String("roleArnForLogging"),
//   		},
//   		PublishFindingToSnsParams: &PublishFindingToSnsParamsProperty{
//   			TopicArn: jsii.String("topicArn"),
//   		},
//   		ReplaceDefaultPolicyVersionParams: &ReplaceDefaultPolicyVersionParamsProperty{
//   			TemplateName: jsii.String("templateName"),
//   		},
//   		UpdateCaCertificateParams: &UpdateCACertificateParamsProperty{
//   			Action: jsii.String("action"),
//   		},
//   		UpdateDeviceCertificateParams: &UpdateDeviceCertificateParamsProperty{
//   			Action: jsii.String("action"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-mitigationaction.html
//
type CfnMitigationActionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMitigationActionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMitigationActionPropsMixin
type jsiiProxy_CfnMitigationActionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMitigationActionPropsMixin) Props() *CfnMitigationActionMixinProps {
	var returns *CfnMitigationActionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMitigationActionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::MitigationAction`.
func NewCfnMitigationActionPropsMixin(props *CfnMitigationActionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMitigationActionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMitigationActionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMitigationActionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnMitigationActionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::MitigationAction`.
func NewCfnMitigationActionPropsMixin_Override(c CfnMitigationActionPropsMixin, props *CfnMitigationActionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnMitigationActionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMitigationActionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMitigationActionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnMitigationActionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMitigationActionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnMitigationActionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMitigationActionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMitigationActionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

