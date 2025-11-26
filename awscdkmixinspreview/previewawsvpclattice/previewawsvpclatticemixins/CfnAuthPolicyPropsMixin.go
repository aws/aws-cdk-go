package previewawsvpclatticemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsvpclattice/previewawsvpclatticemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates the auth policy. The policy string in JSON must not contain newlines or blank lines.
//
// For more information, see [Auth policies](https://docs.aws.amazon.com/vpc-lattice/latest/ug/auth-policies.html) in the *Amazon VPC Lattice User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policy interface{}
//
//   cfnAuthPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnAuthPolicyPropsMixin(&CfnAuthPolicyMixinProps{
//   	Policy: policy,
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-authpolicy.html
//
type CfnAuthPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAuthPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAuthPolicyPropsMixin
type jsiiProxy_CfnAuthPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAuthPolicyPropsMixin) Props() *CfnAuthPolicyMixinProps {
	var returns *CfnAuthPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::VpcLattice::AuthPolicy`.
func NewCfnAuthPolicyPropsMixin(props *CfnAuthPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAuthPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAuthPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAuthPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnAuthPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::VpcLattice::AuthPolicy`.
func NewCfnAuthPolicyPropsMixin_Override(c CfnAuthPolicyPropsMixin, props *CfnAuthPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnAuthPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAuthPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAuthPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnAuthPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAuthPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnAuthPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAuthPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAuthPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

