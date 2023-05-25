package awsnetworkfirewall


// Additional settings for a stateful rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleOptionProperty := &RuleOptionProperty{
//   	Keyword: jsii.String("keyword"),
//
//   	// the properties below are optional
//   	Settings: []*string{
//   		jsii.String("settings"),
//   	},
//   }
//
type CfnRuleGroup_RuleOptionProperty struct {
	// The Suricata rule option keywords.
	//
	// For Network Firewall , the keyword signature ID (sid) is required in the format `sid: 112233` . The sid must be unique within the rule group. For information about Suricata rule option keywords, see [Rule options](https://docs.aws.amazon.com/https://suricata.readthedocs.io/en/suricata-6.0.9/rules/intro.html#rule-options) .
	Keyword *string `field:"required" json:"keyword" yaml:"keyword"`
	// The Suricata rule option settings.
	//
	// Settings have zero or more values, and the number of possible settings and required settings depends on the keyword. The format for Settings is `number` . For information about Suricata rule option settings, see [Rule options](https://docs.aws.amazon.com/https://suricata.readthedocs.io/en/suricata-6.0.9/rules/intro.html#rule-options) .
	Settings *[]*string `field:"optional" json:"settings" yaml:"settings"`
}

