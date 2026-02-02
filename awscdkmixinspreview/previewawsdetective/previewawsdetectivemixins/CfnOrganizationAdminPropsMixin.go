package previewawsdetectivemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdetective/previewawsdetectivemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Detective::OrganizationAdmin` resource is an Amazon Detective resource type that designates the Detective administrator account for the organization in the current region.
//
// If the account does not have Detective enabled, then this resource enables Detective for that account and creates a new behavior graph.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOrganizationAdminPropsMixin := awscdkmixinspreview.Mixins.NewCfnOrganizationAdminPropsMixin(&CfnOrganizationAdminMixinProps{
//   	AccountId: jsii.String("accountId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-organizationadmin.html
//
type CfnOrganizationAdminPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnOrganizationAdminMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOrganizationAdminPropsMixin
type jsiiProxy_CfnOrganizationAdminPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnOrganizationAdminPropsMixin) Props() *CfnOrganizationAdminMixinProps {
	var returns *CfnOrganizationAdminMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationAdminPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Detective::OrganizationAdmin`.
func NewCfnOrganizationAdminPropsMixin(props *CfnOrganizationAdminMixinProps, options *mixins.CfnPropertyMixinOptions) CfnOrganizationAdminPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOrganizationAdminPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOrganizationAdminPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_detective.mixins.CfnOrganizationAdminPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Detective::OrganizationAdmin`.
func NewCfnOrganizationAdminPropsMixin_Override(c CfnOrganizationAdminPropsMixin, props *CfnOrganizationAdminMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_detective.mixins.CfnOrganizationAdminPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnOrganizationAdminPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOrganizationAdminPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_detective.mixins.CfnOrganizationAdminPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOrganizationAdminPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_detective.mixins.CfnOrganizationAdminPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOrganizationAdminPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnOrganizationAdminPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

