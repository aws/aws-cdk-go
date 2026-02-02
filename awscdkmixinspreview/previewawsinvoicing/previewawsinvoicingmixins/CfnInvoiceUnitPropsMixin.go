package previewawsinvoicingmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsinvoicing/previewawsinvoicingmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An invoice unit is a set of mutually exclusive account that correspond to your business entity.
//
// Invoice units allow you separate AWS account costs and configures your invoice for each business entity going forward.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInvoiceUnitPropsMixin := awscdkmixinspreview.Mixins.NewCfnInvoiceUnitPropsMixin(&CfnInvoiceUnitMixinProps{
//   	Description: jsii.String("description"),
//   	InvoiceReceiver: jsii.String("invoiceReceiver"),
//   	Name: jsii.String("name"),
//   	ResourceTags: []ResourceTagProperty{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Rule: &RuleProperty{
//   		LinkedAccounts: []*string{
//   			jsii.String("linkedAccounts"),
//   		},
//   	},
//   	TaxInheritanceDisabled: jsii.Boolean(false),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html
//
type CfnInvoiceUnitPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInvoiceUnitMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInvoiceUnitPropsMixin
type jsiiProxy_CfnInvoiceUnitPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInvoiceUnitPropsMixin) Props() *CfnInvoiceUnitMixinProps {
	var returns *CfnInvoiceUnitMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInvoiceUnitPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Invoicing::InvoiceUnit`.
func NewCfnInvoiceUnitPropsMixin(props *CfnInvoiceUnitMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInvoiceUnitPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInvoiceUnitPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInvoiceUnitPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_invoicing.mixins.CfnInvoiceUnitPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Invoicing::InvoiceUnit`.
func NewCfnInvoiceUnitPropsMixin_Override(c CfnInvoiceUnitPropsMixin, props *CfnInvoiceUnitMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_invoicing.mixins.CfnInvoiceUnitPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInvoiceUnitPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInvoiceUnitPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_invoicing.mixins.CfnInvoiceUnitPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInvoiceUnitPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_invoicing.mixins.CfnInvoiceUnitPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInvoiceUnitPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnInvoiceUnitPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

