package awsmpa

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmpa/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new identity source.
//
// For more information, see [Identity Source](https://docs.aws.amazon.com/mpa/latest/userguide/mpa-concepts.html) in the *Multi-party approval User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnIdentitySourcePropsMixin := awscdkcfnpropertymixins.Aws_mpa.NewCfnIdentitySourcePropsMixin(&CfnIdentitySourceMixinProps{
//   	IdentitySourceParameters: &IdentitySourceParametersProperty{
//   		IamIdentityCenter: &IamIdentityCenterProperty{
//   			ApprovalPortalUrl: jsii.String("approvalPortalUrl"),
//   			InstanceArn: jsii.String("instanceArn"),
//   			Region: jsii.String("region"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-identitysource.html
//
type CfnIdentitySourcePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnIdentitySourceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIdentitySourcePropsMixin
type jsiiProxy_CfnIdentitySourcePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnIdentitySourcePropsMixin) Props() *CfnIdentitySourceMixinProps {
	var returns *CfnIdentitySourceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentitySourcePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MPA::IdentitySource`.
func NewCfnIdentitySourcePropsMixin(props *CfnIdentitySourceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnIdentitySourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIdentitySourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIdentitySourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mpa.CfnIdentitySourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MPA::IdentitySource`.
func NewCfnIdentitySourcePropsMixin_Override(c CfnIdentitySourcePropsMixin, props *CfnIdentitySourceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mpa.CfnIdentitySourcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnIdentitySourcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIdentitySourcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mpa.CfnIdentitySourcePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdentitySourcePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mpa.CfnIdentitySourcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdentitySourcePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIdentitySourcePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

