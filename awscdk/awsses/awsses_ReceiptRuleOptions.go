package awsses


// Options to add a receipt rule to a receipt rule set.
//
// Example:
//   ruleSet := ses.NewReceiptRuleSet(this, jsii.String("RuleSet"))
//
//   awsRule := ruleSet.addRule(jsii.String("Aws"), &receiptRuleOptions{
//   	recipients: []*string{
//   		jsii.String("aws.com"),
//   	},
//   })
//
type ReceiptRuleOptions struct {
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
}

