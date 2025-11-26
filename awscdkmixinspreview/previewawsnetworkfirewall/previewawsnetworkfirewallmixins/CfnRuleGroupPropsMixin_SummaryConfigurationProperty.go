package previewawsnetworkfirewallmixins


// A complex type that specifies which Suricata rule metadata fields to use when displaying threat information. Contains:.
//
// - `RuleOptions` - The Suricata rule options fields to extract and display
//
// These settings affect how threat information appears in both the console and API responses. Summaries are available for rule groups you manage and for active threat defense AWS managed rule groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   summaryConfigurationProperty := &SummaryConfigurationProperty{
//   	RuleOptions: []*string{
//   		jsii.String("ruleOptions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-summaryconfiguration.html
//
type CfnRuleGroupPropsMixin_SummaryConfigurationProperty struct {
	// Specifies the selected rule options returned by `DescribeRuleGroupSummary` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-summaryconfiguration.html#cfn-networkfirewall-rulegroup-summaryconfiguration-ruleoptions
	//
	RuleOptions *[]*string `field:"optional" json:"ruleOptions" yaml:"ruleOptions"`
}

