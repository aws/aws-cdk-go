package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new instance profile. For information about instance profiles, see [Using instance profiles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html) .
//
// For information about the number of instance profiles you can create, see [IAM object quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnInstanceProfilePropsMixin := awscdkcfnpropertymixins.Aws_iam.NewCfnInstanceProfilePropsMixin(&CfnInstanceProfileMixinProps{
//   	InstanceProfileName: jsii.String("instanceProfileName"),
//   	Path: jsii.String("path"),
//   	Roles: []interface{}{
//   		jsii.String("roles"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
//
type CfnInstanceProfilePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnInstanceProfileMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInstanceProfilePropsMixin
type jsiiProxy_CfnInstanceProfilePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnInstanceProfilePropsMixin) Props() *CfnInstanceProfileMixinProps {
	var returns *CfnInstanceProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceProfilePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IAM::InstanceProfile`.
func NewCfnInstanceProfilePropsMixin(props *CfnInstanceProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnInstanceProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInstanceProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstanceProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iam.CfnInstanceProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IAM::InstanceProfile`.
func NewCfnInstanceProfilePropsMixin_Override(c CfnInstanceProfilePropsMixin, props *CfnInstanceProfileMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iam.CfnInstanceProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnInstanceProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstanceProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iam.CfnInstanceProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstanceProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iam.CfnInstanceProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstanceProfilePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnInstanceProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

