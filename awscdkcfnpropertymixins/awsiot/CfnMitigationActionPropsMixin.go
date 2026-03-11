package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines an action that can be applied to audit findings by using StartAuditMitigationActionsTask.
//
// For API reference, see [CreateMitigationAction](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateMitigationAction.html) and for general information, see [Mitigation actions](https://docs.aws.amazon.com/iot/latest/developerguide/dd-mitigation-actions.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMitigationActionPropsMixin := awscdkcfnpropertymixins.Aws_iot.NewCfnMitigationActionPropsMixin(&CfnMitigationActionMixinProps{
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
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-mitigationaction.html
//
type CfnMitigationActionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMitigationActionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMitigationActionPropsMixin
type jsiiProxy_CfnMitigationActionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnMitigationActionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::MitigationAction`.
func NewCfnMitigationActionPropsMixin(props *CfnMitigationActionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMitigationActionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMitigationActionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMitigationActionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnMitigationActionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::MitigationAction`.
func NewCfnMitigationActionPropsMixin_Override(c CfnMitigationActionPropsMixin, props *CfnMitigationActionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnMitigationActionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMitigationActionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMitigationActionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnMitigationActionPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnMitigationActionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMitigationActionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

