package awsappintegrations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsappintegrations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an event integration.
//
// You provide a name, description, and a reference to an Amazon EventBridge bus in your account and a partner event source that will push events to that bus. No objects are created in your account, only metadata that is persisted on the EventIntegration control plane.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnEventIntegrationPropsMixin := awscdkcfnpropertymixins.Aws_appintegrations.NewCfnEventIntegrationPropsMixin(&CfnEventIntegrationMixinProps{
//   	Description: jsii.String("description"),
//   	EventBridgeBus: jsii.String("eventBridgeBus"),
//   	EventFilter: &EventFilterProperty{
//   		Source: jsii.String("source"),
//   	},
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-eventintegration.html
//
type CfnEventIntegrationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnEventIntegrationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEventIntegrationPropsMixin
type jsiiProxy_CfnEventIntegrationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEventIntegrationPropsMixin) Props() *CfnEventIntegrationMixinProps {
	var returns *CfnEventIntegrationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventIntegrationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppIntegrations::EventIntegration`.
func NewCfnEventIntegrationPropsMixin(props *CfnEventIntegrationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnEventIntegrationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEventIntegrationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEventIntegrationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnEventIntegrationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppIntegrations::EventIntegration`.
func NewCfnEventIntegrationPropsMixin_Override(c CfnEventIntegrationPropsMixin, props *CfnEventIntegrationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnEventIntegrationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEventIntegrationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventIntegrationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnEventIntegrationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventIntegrationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_appintegrations.CfnEventIntegrationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventIntegrationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEventIntegrationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

