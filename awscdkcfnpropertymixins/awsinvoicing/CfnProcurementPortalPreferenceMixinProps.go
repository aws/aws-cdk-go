package awsinvoicing

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnProcurementPortalPreferencePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnProcurementPortalPreferenceMixinProps := &CfnProcurementPortalPreferenceMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html
//
type CfnProcurementPortalPreferenceMixinProps struct {
	// The domain identifier for the buyer in the procurement portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-buyerdomain
	//
	BuyerDomain *string `field:"optional" json:"buyerDomain" yaml:"buyerDomain"`
	// The unique identifier for the buyer in the procurement portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-buyeridentifier
	//
	BuyerIdentifier *string `field:"optional" json:"buyerIdentifier" yaml:"buyerIdentifier"`
	// List of contact information for portal administrators and technical contacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-contacts
	//
	Contacts interface{} `field:"optional" json:"contacts" yaml:"contacts"`
	// Indicates whether e-invoice delivery is enabled for this procurement portal preference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-einvoicedeliveryenabled
	//
	EinvoiceDeliveryEnabled interface{} `field:"optional" json:"einvoiceDeliveryEnabled" yaml:"einvoiceDeliveryEnabled"`
	// Specifies the preferences for e-invoice delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-einvoicedeliverypreference
	//
	EinvoiceDeliveryPreference interface{} `field:"optional" json:"einvoiceDeliveryPreference" yaml:"einvoiceDeliveryPreference"`
	// The endpoint URL where e-invoices are delivered to the procurement portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-procurementportalinstanceendpoint
	//
	ProcurementPortalInstanceEndpoint *string `field:"optional" json:"procurementPortalInstanceEndpoint" yaml:"procurementPortalInstanceEndpoint"`
	// The name of the procurement portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-procurementportalname
	//
	ProcurementPortalName *string `field:"optional" json:"procurementPortalName" yaml:"procurementPortalName"`
	// The shared secret or authentication credential used for secure communication with the procurement portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-procurementportalsharedsecret
	//
	ProcurementPortalSharedSecret *string `field:"optional" json:"procurementPortalSharedSecret" yaml:"procurementPortalSharedSecret"`
	// Indicates whether purchase order retrieval is enabled for this procurement portal preference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-purchaseorderretrievalenabled
	//
	PurchaseOrderRetrievalEnabled interface{} `field:"optional" json:"purchaseOrderRetrievalEnabled" yaml:"purchaseOrderRetrievalEnabled"`
	// Specifies criteria for selecting which invoices should be processed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-selector
	//
	Selector interface{} `field:"optional" json:"selector" yaml:"selector"`
	// The domain identifier for the supplier in the procurement portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-supplierdomain
	//
	SupplierDomain *string `field:"optional" json:"supplierDomain" yaml:"supplierDomain"`
	// The unique identifier for the supplier in the procurement portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-supplieridentifier
	//
	SupplierIdentifier *string `field:"optional" json:"supplierIdentifier" yaml:"supplierIdentifier"`
	// The tags associated with this procurement portal preference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Configuration settings for the test environment of the procurement portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-procurementportalpreference.html#cfn-invoicing-procurementportalpreference-testenvpreference
	//
	TestEnvPreference interface{} `field:"optional" json:"testEnvPreference" yaml:"testEnvPreference"`
}

