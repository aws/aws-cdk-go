package interfacesawsses


// A reference to a ReceiptRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   receiptRuleReference := &ReceiptRuleReference{
//   	RuleName: jsii.String("ruleName"),
//   	RuleSetName: jsii.String("ruleSetName"),
//   }
//
type ReceiptRuleReference struct {
	// The RuleName of the ReceiptRule resource.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// The RuleSetName of the ReceiptRule resource.
	RuleSetName *string `field:"required" json:"ruleSetName" yaml:"ruleSetName"`
}

