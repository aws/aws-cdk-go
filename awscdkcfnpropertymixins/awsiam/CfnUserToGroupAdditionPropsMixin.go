package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds the specified user to the specified group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnUserToGroupAdditionPropsMixin := awscdkcfnpropertymixins.Aws_iam.NewCfnUserToGroupAdditionPropsMixin(&CfnUserToGroupAdditionMixinProps{
//   	GroupName: jsii.String("groupName"),
//   	Users: []*string{
//   		jsii.String("users"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-usertogroupaddition.html
//
type CfnUserToGroupAdditionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnUserToGroupAdditionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUserToGroupAdditionPropsMixin
type jsiiProxy_CfnUserToGroupAdditionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnUserToGroupAdditionPropsMixin) Props() *CfnUserToGroupAdditionMixinProps {
	var returns *CfnUserToGroupAdditionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserToGroupAdditionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IAM::UserToGroupAddition`.
func NewCfnUserToGroupAdditionPropsMixin(props *CfnUserToGroupAdditionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnUserToGroupAdditionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserToGroupAdditionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserToGroupAdditionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iam.CfnUserToGroupAdditionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IAM::UserToGroupAddition`.
func NewCfnUserToGroupAdditionPropsMixin_Override(c CfnUserToGroupAdditionPropsMixin, props *CfnUserToGroupAdditionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iam.CfnUserToGroupAdditionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnUserToGroupAdditionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserToGroupAdditionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iam.CfnUserToGroupAdditionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserToGroupAdditionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iam.CfnUserToGroupAdditionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserToGroupAdditionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnUserToGroupAdditionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

