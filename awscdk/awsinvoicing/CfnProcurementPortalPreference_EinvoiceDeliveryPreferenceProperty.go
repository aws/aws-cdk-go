package awsinvoicing


// Specifies the preferences for e-invoice delivery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   einvoiceDeliveryPreferenceProperty := &EinvoiceDeliveryPreferenceProperty{
//   	ConnectionTestingMethod: jsii.String("connectionTestingMethod"),
//   	EinvoiceDeliveryActivationDate: jsii.String("einvoiceDeliveryActivationDate"),
//   	EinvoiceDeliveryAttachmentTypes: []*string{
//   		jsii.String("einvoiceDeliveryAttachmentTypes"),
//   	},
//   	EinvoiceDeliveryDocumentTypes: []*string{
//   		jsii.String("einvoiceDeliveryDocumentTypes"),
//   	},
//   	Protocol: jsii.String("protocol"),
//   	PurchaseOrderDataSources: []interface{}{
//   		&PurchaseOrderDataSourceProperty{
//   			EinvoiceDeliveryDocumentType: jsii.String("einvoiceDeliveryDocumentType"),
//   			PurchaseOrderDataSourceType: jsii.String("purchaseOrderDataSourceType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-einvoicedeliverypreference.html
//
type CfnProcurementPortalPreference_EinvoiceDeliveryPreferenceProperty struct {
	// The method to use for testing the connection to the procurement portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-einvoicedeliverypreference.html#cfn-invoicing-procurementportalpreference-einvoicedeliverypreference-connectiontestingmethod
	//
	ConnectionTestingMethod *string `field:"optional" json:"connectionTestingMethod" yaml:"connectionTestingMethod"`
	// The ISO 8601 date-time when e-invoice delivery should be activated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-einvoicedeliverypreference.html#cfn-invoicing-procurementportalpreference-einvoicedeliverypreference-einvoicedeliveryactivationdate
	//
	EinvoiceDeliveryActivationDate *string `field:"optional" json:"einvoiceDeliveryActivationDate" yaml:"einvoiceDeliveryActivationDate"`
	// The types of attachments to include with the e-invoice delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-einvoicedeliverypreference.html#cfn-invoicing-procurementportalpreference-einvoicedeliverypreference-einvoicedeliveryattachmenttypes
	//
	EinvoiceDeliveryAttachmentTypes *[]*string `field:"optional" json:"einvoiceDeliveryAttachmentTypes" yaml:"einvoiceDeliveryAttachmentTypes"`
	// The types of e-invoice documents to be delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-einvoicedeliverypreference.html#cfn-invoicing-procurementportalpreference-einvoicedeliverypreference-einvoicedeliverydocumenttypes
	//
	EinvoiceDeliveryDocumentTypes *[]*string `field:"optional" json:"einvoiceDeliveryDocumentTypes" yaml:"einvoiceDeliveryDocumentTypes"`
	// The communication protocol to use for e-invoice delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-einvoicedeliverypreference.html#cfn-invoicing-procurementportalpreference-einvoicedeliverypreference-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The sources of purchase order data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-einvoicedeliverypreference.html#cfn-invoicing-procurementportalpreference-einvoicedeliverypreference-purchaseorderdatasources
	//
	PurchaseOrderDataSources interface{} `field:"optional" json:"purchaseOrderDataSources" yaml:"purchaseOrderDataSources"`
}

