package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses"
)

// Construction properties for a ReceiptRule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var receiptRuleAction IReceiptRuleAction
//   var receiptRuleRef IReceiptRuleRef
//   var receiptRuleSetRef IReceiptRuleSetRef
//
//   receiptRuleProps := &ReceiptRuleProps{
//   	RuleSet: receiptRuleSetRef,
//
//   	// the properties below are optional
//   	Actions: []IReceiptRuleAction{
//   		receiptRuleAction,
//   	},
//   	After: receiptRuleRef,
//   	Enabled: jsii.Boolean(false),
//   	ReceiptRuleName: jsii.String("receiptRuleName"),
//   	Recipients: []*string{
//   		jsii.String("recipients"),
//   	},
//   	ScanEnabled: jsii.Boolean(false),
//   	TlsPolicy: awscdk.Aws_ses.TlsPolicy_OPTIONAL,
//   }
//
type ReceiptRuleProps struct {
	// An ordered list of actions to perform on messages that match at least one of the recipient email addresses or domains specified in the receipt rule.
	// Default: - No actions.
	//
	Actions *[]IReceiptRuleAction `field:"optional" json:"actions" yaml:"actions"`
	// An existing rule after which the new rule will be placed.
	// Default: - The new rule is inserted at the beginning of the rule list.
	//
	After interfacesawsses.IReceiptRuleRef `field:"optional" json:"after" yaml:"after"`
	// Whether the rule is active.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The name for the rule.
	// Default: - A CloudFormation generated name.
	//
	ReceiptRuleName *string `field:"optional" json:"receiptRuleName" yaml:"receiptRuleName"`
	// The recipient domains and email addresses that the receipt rule applies to.
	// Default: - Match all recipients under all verified domains.
	//
	Recipients *[]*string `field:"optional" json:"recipients" yaml:"recipients"`
	// Whether to scan for spam and viruses.
	// Default: false.
	//
	ScanEnabled *bool `field:"optional" json:"scanEnabled" yaml:"scanEnabled"`
	// Whether Amazon SES should require that incoming email is delivered over a connection encrypted with Transport Layer Security (TLS).
	// Default: - Optional which will not check for TLS.
	//
	TlsPolicy TlsPolicy `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
	// The name of the rule set that the receipt rule will be added to.
	RuleSet interfacesawsses.IReceiptRuleSetRef `field:"required" json:"ruleSet" yaml:"ruleSet"`
}

