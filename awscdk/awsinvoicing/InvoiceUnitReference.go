package awsinvoicing


// A reference to a InvoiceUnit resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   invoiceUnitReference := &InvoiceUnitReference{
//   	InvoiceUnitArn: jsii.String("invoiceUnitArn"),
//   }
//
type InvoiceUnitReference struct {
	// The InvoiceUnitArn of the InvoiceUnit resource.
	InvoiceUnitArn *string `field:"required" json:"invoiceUnitArn" yaml:"invoiceUnitArn"`
}

