package awsinvoicing


// Specifies the source configuration for retrieving purchase order data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   purchaseOrderDataSourceProperty := &PurchaseOrderDataSourceProperty{
//   	EinvoiceDeliveryDocumentType: jsii.String("einvoiceDeliveryDocumentType"),
//   	PurchaseOrderDataSourceType: jsii.String("purchaseOrderDataSourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-purchaseorderdatasource.html
//
type CfnProcurementPortalPreference_PurchaseOrderDataSourceProperty struct {
	// The type of e-invoice document that requires purchase order data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-purchaseorderdatasource.html#cfn-invoicing-procurementportalpreference-purchaseorderdatasource-einvoicedeliverydocumenttype
	//
	EinvoiceDeliveryDocumentType *string `field:"optional" json:"einvoiceDeliveryDocumentType" yaml:"einvoiceDeliveryDocumentType"`
	// The type of source for purchase order data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-purchaseorderdatasource.html#cfn-invoicing-procurementportalpreference-purchaseorderdatasource-purchaseorderdatasourcetype
	//
	PurchaseOrderDataSourceType *string `field:"optional" json:"purchaseOrderDataSourceType" yaml:"purchaseOrderDataSourceType"`
}

