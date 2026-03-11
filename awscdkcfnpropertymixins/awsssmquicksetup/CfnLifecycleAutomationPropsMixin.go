package awsssmquicksetup

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsssmquicksetup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a lifecycle automation resource that executes SSM Automation documents during CloudFormation stack operations.
//
// This resource replaces inline AWS Lambda custom resources and provides a managed way to handle lifecycle events in Quick Setup configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnLifecycleAutomationPropsMixin := awscdkcfnpropertymixins.Aws_ssmquicksetup.NewCfnLifecycleAutomationPropsMixin(&CfnLifecycleAutomationMixinProps{
//   	AutomationDocument: jsii.String("automationDocument"),
//   	AutomationParameters: map[string][]*string{
//   		"automationParametersKey": []*string{
//   			jsii.String("automationParameters"),
//   		},
//   	},
//   	ResourceKey: jsii.String("resourceKey"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmquicksetup-lifecycleautomation.html
//
type CfnLifecycleAutomationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLifecycleAutomationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLifecycleAutomationPropsMixin
type jsiiProxy_CfnLifecycleAutomationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLifecycleAutomationPropsMixin) Props() *CfnLifecycleAutomationMixinProps {
	var returns *CfnLifecycleAutomationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleAutomationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSMQuickSetup::LifecycleAutomation`.
func NewCfnLifecycleAutomationPropsMixin(props *CfnLifecycleAutomationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLifecycleAutomationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLifecycleAutomationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLifecycleAutomationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ssmquicksetup.CfnLifecycleAutomationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSMQuickSetup::LifecycleAutomation`.
func NewCfnLifecycleAutomationPropsMixin_Override(c CfnLifecycleAutomationPropsMixin, props *CfnLifecycleAutomationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ssmquicksetup.CfnLifecycleAutomationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLifecycleAutomationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLifecycleAutomationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ssmquicksetup.CfnLifecycleAutomationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLifecycleAutomationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ssmquicksetup.CfnLifecycleAutomationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLifecycleAutomationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLifecycleAutomationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

