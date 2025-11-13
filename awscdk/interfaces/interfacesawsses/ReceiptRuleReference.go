package interfacesawsses


// A reference to a ReceiptRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   receiptRuleReference := &ReceiptRuleReference{
//   	ReceiptRuleId: jsii.String("receiptRuleId"),
//   }
//
type ReceiptRuleReference struct {
	// The Id of the ReceiptRule resource.
	ReceiptRuleId *string `field:"required" json:"receiptRuleId" yaml:"receiptRuleId"`
}

