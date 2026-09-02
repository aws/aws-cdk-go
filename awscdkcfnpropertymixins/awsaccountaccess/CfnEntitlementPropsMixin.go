package awsaccountaccess

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsaccountaccess/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::AccountAccess::Entitlement specifying an entitlement for account access.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnEntitlementPropsMixin := awscdkcfnpropertymixins.Aws_accountaccess.NewCfnEntitlementPropsMixin(&CfnEntitlementMixinProps{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	Entitlement: &EntitlementProperty{
//   		PrincipalRole: &PrincipalRoleEntitlementProperty{
//   			Account: jsii.String("account"),
//   			Principal: &PrincipalProperty{
//   				IdentityCenter: &IdentityCenterPrincipalProperty{
//   					GroupId: jsii.String("groupId"),
//   					UserId: jsii.String("userId"),
//   				},
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accountaccess-entitlement.html
//
type CfnEntitlementPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnEntitlementMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEntitlementPropsMixin
type jsiiProxy_CfnEntitlementPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEntitlementPropsMixin) Props() *CfnEntitlementMixinProps {
	var returns *CfnEntitlementMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEntitlementPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AccountAccess::Entitlement`.
func NewCfnEntitlementPropsMixin(props *CfnEntitlementMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnEntitlementPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEntitlementPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEntitlementPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_accountaccess.CfnEntitlementPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AccountAccess::Entitlement`.
func NewCfnEntitlementPropsMixin_Override(c CfnEntitlementPropsMixin, props *CfnEntitlementMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_accountaccess.CfnEntitlementPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEntitlementPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEntitlementPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_accountaccess.CfnEntitlementPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEntitlementPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_accountaccess.CfnEntitlementPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEntitlementPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEntitlementPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

