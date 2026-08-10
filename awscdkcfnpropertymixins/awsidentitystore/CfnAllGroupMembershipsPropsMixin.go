package awsidentitystore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsidentitystore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Retrieves membership metadata and attributes for a group membership in an identity store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAllGroupMembershipsPropsMixin := awscdkcfnpropertymixins.Aws_identitystore.NewCfnAllGroupMembershipsPropsMixin(&CfnAllGroupMembershipsMixinProps{
//   	GroupId: jsii.String("groupId"),
//   	IdentityStoreId: jsii.String("identityStoreId"),
//   	MemberId: &MemberIdProperty{
//   		UserId: jsii.String("userId"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-identitystore-allgroupmemberships.html
//
type CfnAllGroupMembershipsPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAllGroupMembershipsMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAllGroupMembershipsPropsMixin
type jsiiProxy_CfnAllGroupMembershipsPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAllGroupMembershipsPropsMixin) Props() *CfnAllGroupMembershipsMixinProps {
	var returns *CfnAllGroupMembershipsMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAllGroupMembershipsPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IdentityStore::AllGroupMemberships`.
func NewCfnAllGroupMembershipsPropsMixin(props *CfnAllGroupMembershipsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAllGroupMembershipsPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAllGroupMembershipsPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAllGroupMembershipsPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_identitystore.CfnAllGroupMembershipsPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IdentityStore::AllGroupMemberships`.
func NewCfnAllGroupMembershipsPropsMixin_Override(c CfnAllGroupMembershipsPropsMixin, props *CfnAllGroupMembershipsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_identitystore.CfnAllGroupMembershipsPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAllGroupMembershipsPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAllGroupMembershipsPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_identitystore.CfnAllGroupMembershipsPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAllGroupMembershipsPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_identitystore.CfnAllGroupMembershipsPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAllGroupMembershipsPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAllGroupMembershipsPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

