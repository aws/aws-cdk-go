package awseventsv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awseventsv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource type definition for AWS::EventsV2::ResourcePolicy, the resource policy of an Amazon EventBridge event bus.
//
// This manages only the resource policy named "default", which the bus owner writes. It does not manage the policy named "AWS_RAM", which AWS Resource Access Manager owns on behalf of the bus owner. If the bus already has a default policy, creating this resource fails. Deleting this resource removes all permissions granted by it. An explicit Deny in this policy takes precedence over an Allow in the "AWS_RAM" policy, so deleting this resource can widen access. Set DeletionPolicy: Retain if the policy carries a Deny that you rely on. Required permissions: events:PutResourcePolicy, events:GetResourcePolicy, and events:DeleteResourcePolicy. Listing resources of this type also requires events:ListEventBuses and events:ListResourcePolicies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var policyDocument interface{}
//
//   cfnResourcePolicyPropsMixin := awscdkcfnpropertymixins.Aws_eventsv2.NewCfnResourcePolicyPropsMixin(&CfnResourcePolicyMixinProps{
//   	EventBusArn: jsii.String("eventBusArn"),
//   	PolicyDocument: policyDocument,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-resourcepolicy.html
//
type CfnResourcePolicyPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnResourcePolicyMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResourcePolicyPropsMixin
type jsiiProxy_CfnResourcePolicyPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnResourcePolicyPropsMixin) Props() *CfnResourcePolicyMixinProps {
	var returns *CfnResourcePolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicyPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EventsV2::ResourcePolicy`.
func NewCfnResourcePolicyPropsMixin(props *CfnResourcePolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnResourcePolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResourcePolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResourcePolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnResourcePolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EventsV2::ResourcePolicy`.
func NewCfnResourcePolicyPropsMixin_Override(c CfnResourcePolicyPropsMixin, props *CfnResourcePolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnResourcePolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnResourcePolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResourcePolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnResourcePolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourcePolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnResourcePolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourcePolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnResourcePolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

