package interfacesawsses


// A reference to a ReceiptFilter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   receiptFilterReference := &ReceiptFilterReference{
//   	ReceiptFilterId: jsii.String("receiptFilterId"),
//   }
//
type ReceiptFilterReference struct {
	// The Id of the ReceiptFilter resource.
	ReceiptFilterId *string `field:"required" json:"receiptFilterId" yaml:"receiptFilterId"`
}

