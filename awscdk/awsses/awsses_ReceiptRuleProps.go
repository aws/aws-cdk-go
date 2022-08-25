package awsses


// Construction properties for a ReceiptRule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var receiptRule receiptRule
//   var receiptRuleAction iReceiptRuleAction
//   var receiptRuleSet receiptRuleSet
//
//   receiptRuleProps := &receiptRuleProps{
//   	ruleSet: receiptRuleSet,
//
//   	// the properties below are optional
//   	actions: []*iReceiptRuleAction{
//   		receiptRuleAction,
//   	},
//   	after: receiptRule,
//   	enabled: jsii.Boolean(false),
//   	receiptRuleName: jsii.String("receiptRuleName"),
//   	recipients: []*string{
//   		jsii.String("recipients"),
//   	},
//   	scanEnabled: jsii.Boolean(false),
//   	tlsPolicy: awscdk.Aws_ses.tlsPolicy_OPTIONAL,
//   }
//
type ReceiptRuleProps struct {
	// An ordered list of actions to perform on messages that match at least one of the recipient email addresses or domains specified in the receipt rule.
	Actions *[]IReceiptRuleAction `field:"optional" json:"actions" yaml:"actions"`
	// An existing rule after which the new rule will be placed.
	After IReceiptRule `field:"optional" json:"after" yaml:"after"`
	// Whether the rule is active.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The name for the rule.
	ReceiptRuleName *string `field:"optional" json:"receiptRuleName" yaml:"receiptRuleName"`
	// The recipient domains and email addresses that the receipt rule applies to.
	Recipients *[]*string `field:"optional" json:"recipients" yaml:"recipients"`
	// Whether to scan for spam and viruses.
	ScanEnabled *bool `field:"optional" json:"scanEnabled" yaml:"scanEnabled"`
	// Whether Amazon SES should require that incoming email is delivered over a connection encrypted with Transport Layer Security (TLS).
	TlsPolicy TlsPolicy `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
	// The name of the rule set that the receipt rule will be added to.
	RuleSet IReceiptRuleSet `field:"required" json:"ruleSet" yaml:"ruleSet"`
}

