package awsinvoicing

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsinvoicing/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates and manages a procurement portal preference configuration for e-invoice delivery and purchase order retrieval.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnProcurementPortalPreferencePropsMixin := awscdkcfnpropertymixins.Aws_invoicing.NewCfnProcurementPortalPreferencePropsMixin(&CfnProcurementPortalPreferenceMixinProps{
//   	BuyerDomain: jsii.String("buyerDomain"),
//   	BuyerIdentifier: jsii.String("buyerIdentifier"),
//   	Contacts: []interface{}{
//   		&ContactProperty{
//   			Email: jsii.String("email"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	EinvoiceDeliveryEnabled: jsii.Boolean(false),
//   	EinvoiceDeliveryPreference: &EinvoiceDeliveryPreferenceProperty{
//   		ConnectionTestingMethod: jsii.String("connectionTestingMethod"),
//   		EinvoiceDeliveryActivationDate: jsii.String("einvoiceDeliveryActivationDate"),
//   		EinvoiceDeliveryAttachmentTypes: []*string{
//   			jsii.String("einvoiceDeliveryAttachmentTypes"),
//   		},
//   		EinvoiceDeliveryDocumentTypes: []*string{
//   			jsii.String("einvoiceDeliveryDocumentTypes"),
//   		},
//   		Protocol: jsii.String("protocol"),
//   		PurchaseOrderDataSources: []interface{}{
//   			&PurchaseOrderDataSourceProperty{
//   				EinvoiceDeliveryDocumentType: jsii.String("einvoiceDeliveryDocumentType"),
//   				PurchaseOrderDataSourceType: jsii.String("purchaseOrderDataSourceType"),
//   			},
//   		},
//   	},
//   	ProcurementPortalInstanceEndpoint: jsii.String("procurementPortalInstanceEndpoint"),
//   	ProcurementPortalName: jsii.String("procurementPortalName"),
//   	ProcurementPortalSharedSecret: jsii.String("procurementPortalSharedSecret"),
//   	PurchaseOrderRetrievalEnabled: jsii.Boolean(false),
//   	Selector: &ProcurementPortalPreferenceSelectorProperty{
//   		InvoiceUnitArns: []*string{
//   			jsii.String("invoiceUnitArns"),
//   		},
//   	},
//   	SupplierDomain: jsii.String("supplierDomain"),
//   	SupplierIdentifier: jsii.String("supplierIdentifier"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TestEnvPreference: &TestEnvPreferenceProperty{
//   		BuyerDomain: jsii.String("buyerDomain"),
//   		BuyerIdentifier: jsii.String("buyerIdentifier"),
//   		ProcurementPortalInstanceEndpoint: jsii.String("procurementPortalInstanceEndpoint"),
//   		ProcurementPortalSharedSecret: jsii.String("procurementPortalSharedSecret"),
//   		SupplierDomain: jsii.String("supplierDomain"),
//   		SupplierIdentifier: jsii.String("supplierIdentifier"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html
//
type CfnProcurementPortalPreferencePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnProcurementPortalPreferenceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnProcurementPortalPreferencePropsMixin
type jsiiProxy_CfnProcurementPortalPreferencePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnProcurementPortalPreferencePropsMixin) Props() *CfnProcurementPortalPreferenceMixinProps {
	var returns *CfnProcurementPortalPreferenceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcurementPortalPreferencePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Invoicing::ProcurementPortalPreference`.
func NewCfnProcurementPortalPreferencePropsMixin(props *CfnProcurementPortalPreferenceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnProcurementPortalPreferencePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnProcurementPortalPreferencePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProcurementPortalPreferencePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnProcurementPortalPreferencePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Invoicing::ProcurementPortalPreference`.
func NewCfnProcurementPortalPreferencePropsMixin_Override(c CfnProcurementPortalPreferencePropsMixin, props *CfnProcurementPortalPreferenceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnProcurementPortalPreferencePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnProcurementPortalPreferencePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProcurementPortalPreferencePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnProcurementPortalPreferencePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProcurementPortalPreferencePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_invoicing.CfnProcurementPortalPreferencePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProcurementPortalPreferencePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnProcurementPortalPreferencePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

