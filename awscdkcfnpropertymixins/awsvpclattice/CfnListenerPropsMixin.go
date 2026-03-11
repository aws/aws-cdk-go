package awsvpclattice

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a listener for a service.
//
// Before you start using your Amazon VPC Lattice service, you must add one or more listeners. A listener is a process that checks for connection requests to your services. For more information, see [Listeners](https://docs.aws.amazon.com/vpc-lattice/latest/ug/listeners.html) in the *Amazon VPC Lattice User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnListenerPropsMixin := awscdkcfnpropertymixins.Aws_vpclattice.NewCfnListenerPropsMixin(&CfnListenerMixinProps{
//   	DefaultAction: &DefaultActionProperty{
//   		FixedResponse: &FixedResponseProperty{
//   			StatusCode: jsii.Number(123),
//   		},
//   		Forward: &ForwardProperty{
//   			TargetGroups: []interface{}{
//   				&WeightedTargetGroupProperty{
//   					TargetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	ServiceIdentifier: jsii.String("serviceIdentifier"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-listener.html
//
type CfnListenerPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnListenerMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnListenerPropsMixin
type jsiiProxy_CfnListenerPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnListenerPropsMixin) Props() *CfnListenerMixinProps {
	var returns *CfnListenerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::VpcLattice::Listener`.
func NewCfnListenerPropsMixin(props *CfnListenerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnListenerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnListenerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnListenerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_vpclattice.CfnListenerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::VpcLattice::Listener`.
func NewCfnListenerPropsMixin_Override(c CfnListenerPropsMixin, props *CfnListenerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_vpclattice.CfnListenerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnListenerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnListenerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_vpclattice.CfnListenerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnListenerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_vpclattice.CfnListenerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnListenerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnListenerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

