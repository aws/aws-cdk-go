package awsinvoicing


// Specifies criteria for selecting which invoices should be processed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   procurementPortalPreferenceSelectorProperty := &ProcurementPortalPreferenceSelectorProperty{
//   	InvoiceUnitArns: []*string{
//   		jsii.String("invoiceUnitArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-procurementportalpreferenceselector.html
//
type CfnProcurementPortalPreference_ProcurementPortalPreferenceSelectorProperty struct {
	// The Amazon Resource Name (ARN) of invoice unit identifiers to which this preference applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-procurementportalpreference-procurementportalpreferenceselector.html#cfn-invoicing-procurementportalpreference-procurementportalpreferenceselector-invoiceunitarns
	//
	InvoiceUnitArns *[]*string `field:"optional" json:"invoiceUnitArns" yaml:"invoiceUnitArns"`
}

