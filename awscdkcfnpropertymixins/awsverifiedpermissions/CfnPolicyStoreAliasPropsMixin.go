package awsverifiedpermissions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsverifiedpermissions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::VerifiedPermissions::PolicyStoreAlias Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPolicyStoreAliasPropsMixin := awscdkcfnpropertymixins.Aws_verifiedpermissions.NewCfnPolicyStoreAliasPropsMixin(&CfnPolicyStoreAliasMixinProps{
//   	AliasName: jsii.String("aliasName"),
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policystorealias.html
//
type CfnPolicyStoreAliasPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPolicyStoreAliasMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPolicyStoreAliasPropsMixin
type jsiiProxy_CfnPolicyStoreAliasPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPolicyStoreAliasPropsMixin) Props() *CfnPolicyStoreAliasMixinProps {
	var returns *CfnPolicyStoreAliasMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyStoreAliasPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::VerifiedPermissions::PolicyStoreAlias`.
func NewCfnPolicyStoreAliasPropsMixin(props *CfnPolicyStoreAliasMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPolicyStoreAliasPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPolicyStoreAliasPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPolicyStoreAliasPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyStoreAliasPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::VerifiedPermissions::PolicyStoreAlias`.
func NewCfnPolicyStoreAliasPropsMixin_Override(c CfnPolicyStoreAliasPropsMixin, props *CfnPolicyStoreAliasMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyStoreAliasPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPolicyStoreAliasPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPolicyStoreAliasPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyStoreAliasPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPolicyStoreAliasPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_verifiedpermissions.CfnPolicyStoreAliasPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPolicyStoreAliasPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPolicyStoreAliasPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

