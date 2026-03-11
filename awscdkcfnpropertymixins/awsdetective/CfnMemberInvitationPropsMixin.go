package awsdetective

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdetective/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Detective::MemberInvitation` resource is an Amazon Detective resource type that creates an invitation to join a Detective behavior graph.
//
// The administrator account can choose whether to send an email notification of the invitation to the root user email address of the AWS account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMemberInvitationPropsMixin := awscdkcfnpropertymixins.Aws_detective.NewCfnMemberInvitationPropsMixin(&CfnMemberInvitationMixinProps{
//   	DisableEmailNotification: jsii.Boolean(false),
//   	GraphArn: jsii.String("graphArn"),
//   	MemberEmailAddress: jsii.String("memberEmailAddress"),
//   	MemberId: jsii.String("memberId"),
//   	Message: jsii.String("message"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html
//
type CfnMemberInvitationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMemberInvitationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMemberInvitationPropsMixin
type jsiiProxy_CfnMemberInvitationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMemberInvitationPropsMixin) Props() *CfnMemberInvitationMixinProps {
	var returns *CfnMemberInvitationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMemberInvitationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Detective::MemberInvitation`.
func NewCfnMemberInvitationPropsMixin(props *CfnMemberInvitationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMemberInvitationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMemberInvitationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMemberInvitationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_detective.CfnMemberInvitationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Detective::MemberInvitation`.
func NewCfnMemberInvitationPropsMixin_Override(c CfnMemberInvitationPropsMixin, props *CfnMemberInvitationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_detective.CfnMemberInvitationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMemberInvitationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMemberInvitationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_detective.CfnMemberInvitationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMemberInvitationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_detective.CfnMemberInvitationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMemberInvitationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMemberInvitationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

