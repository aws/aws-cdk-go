package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses"
)

// Properties of a reference to an existing receipt rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var receiptRuleSetRef IReceiptRuleSetRef
//
//   receiptRuleAttributes := &ReceiptRuleAttributes{
//   	ReceiptRuleName: jsii.String("receiptRuleName"),
//   	RuleSet: receiptRuleSetRef,
//   }
//
type ReceiptRuleAttributes struct {
	// The name of the receipt rule.
	ReceiptRuleName *string `field:"required" json:"receiptRuleName" yaml:"receiptRuleName"`
	// The rule set that the receipt rule belongs to.
	RuleSet interfacesawsses.IReceiptRuleSetRef `field:"required" json:"ruleSet" yaml:"ruleSet"`
}

