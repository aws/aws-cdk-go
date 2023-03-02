package awsnetworkfirewall


// Additional settings for a stateful rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleOptionProperty := &ruleOptionProperty{
//   	keyword: jsii.String("keyword"),
//
//   	// the properties below are optional
//   	settings: []*string{
//   		jsii.String("settings"),
//   	},
//   }
//
type CfnRuleGroup_RuleOptionProperty struct {
	// `CfnRuleGroup.RuleOptionProperty.Keyword`.
	Keyword *string `field:"required" json:"keyword" yaml:"keyword"`
	// `CfnRuleGroup.RuleOptionProperty.Settings`.
	Settings *[]*string `field:"optional" json:"settings" yaml:"settings"`
}

