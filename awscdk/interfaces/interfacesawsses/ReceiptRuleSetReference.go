package interfacesawsses


// A reference to a ReceiptRuleSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   receiptRuleSetReference := &ReceiptRuleSetReference{
//   	RuleSetName: jsii.String("ruleSetName"),
//   }
//
type ReceiptRuleSetReference struct {
	// The RuleSetName of the ReceiptRuleSet resource.
	RuleSetName *string `field:"required" json:"ruleSetName" yaml:"ruleSetName"`
}

