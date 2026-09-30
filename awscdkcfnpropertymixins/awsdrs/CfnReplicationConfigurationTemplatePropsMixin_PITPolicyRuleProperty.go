package awsdrs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   pITPolicyRuleProperty := &PITPolicyRuleProperty{
//   	Enabled: jsii.Boolean(false),
//   	Interval: jsii.Number(123),
//   	RetentionDuration: jsii.Number(123),
//   	RuleId: jsii.Number(123),
//   	Units: jsii.String("units"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-drs-replicationconfigurationtemplate-pitpolicyrule.html
//
type CfnReplicationConfigurationTemplatePropsMixin_PITPolicyRuleProperty struct {
	// Whether this rule is enabled or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-drs-replicationconfigurationtemplate-pitpolicyrule.html#cfn-drs-replicationconfigurationtemplate-pitpolicyrule-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// How often, in the chosen units, a snapshot should be taken.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-drs-replicationconfigurationtemplate-pitpolicyrule.html#cfn-drs-replicationconfigurationtemplate-pitpolicyrule-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The duration to retain a snapshot for, in the chosen units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-drs-replicationconfigurationtemplate-pitpolicyrule.html#cfn-drs-replicationconfigurationtemplate-pitpolicyrule-retentionduration
	//
	RetentionDuration *float64 `field:"optional" json:"retentionDuration" yaml:"retentionDuration"`
	// The ID of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-drs-replicationconfigurationtemplate-pitpolicyrule.html#cfn-drs-replicationconfigurationtemplate-pitpolicyrule-ruleid
	//
	RuleId *float64 `field:"optional" json:"ruleId" yaml:"ruleId"`
	// The units used to measure the interval and retentionDuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-drs-replicationconfigurationtemplate-pitpolicyrule.html#cfn-drs-replicationconfigurationtemplate-pitpolicyrule-units
	//
	Units *string `field:"optional" json:"units" yaml:"units"`
}

