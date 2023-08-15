package awscdkgluealpha


// Construction properties for `DataQualityRuleset`.
//
// Example:
//   glue.NewDataQualityRuleset(this, jsii.String("MyDataQualityRuleset"), &DataQualityRulesetProps{
//   	ClientToken: jsii.String("client_token"),
//   	Description: jsii.String("description"),
//   	RulesetName: jsii.String("ruleset_name"),
//   	RulesetDqdl: jsii.String("ruleset_dqdl"),
//   	Tags: map[string]*string{
//   		"key1": jsii.String("value1"),
//   		"key2": jsii.String("value2"),
//   	},
//   	TargetTable: glue.NewDataQualityTargetTable(jsii.String("database_name"), jsii.String("table_name")),
//   })
//
// Experimental.
type DataQualityRulesetProps struct {
	// The dqdl of the ruleset.
	// Experimental.
	RulesetDqdl *string `field:"required" json:"rulesetDqdl" yaml:"rulesetDqdl"`
	// The target table of the ruleset.
	// Experimental.
	TargetTable DataQualityTargetTable `field:"required" json:"targetTable" yaml:"targetTable"`
	// The client token of the ruleset.
	// Experimental.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The description of the ruleset.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the ruleset.
	// Default: cloudformation generated name.
	//
	// Experimental.
	RulesetName *string `field:"optional" json:"rulesetName" yaml:"rulesetName"`
	// Key-Value pairs that define tags for the ruleset.
	// Default: empty tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

