package interfacesawsses


// A reference to a ReceiptRuleSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   receiptRuleSetReference := &ReceiptRuleSetReference{
//   	ReceiptRuleSetId: jsii.String("receiptRuleSetId"),
//   }
//
type ReceiptRuleSetReference struct {
	// The Id of the ReceiptRuleSet resource.
	ReceiptRuleSetId *string `field:"required" json:"receiptRuleSetId" yaml:"receiptRuleSetId"`
}

