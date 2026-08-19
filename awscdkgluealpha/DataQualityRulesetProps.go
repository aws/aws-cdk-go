package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for `DataQualityRuleset`.
//
// Example:
//   glue.NewDataQualityRuleset(this, jsii.String("MyRuleset"), &DataQualityRulesetProps{
//   	RulesetName: jsii.String("my_ruleset"),
//   	Dqdl: glue.Dqdl_FromString(jsii.String("Rules = [ RowCount > 100, IsComplete \"order_id\" ]")),
//   	TargetTable: glue.NewDataQualityTargetTable(jsii.String("my_database"), jsii.String("my_table")),
//   })
//
// Experimental.
type DataQualityRulesetProps struct {
	// The DQDL document defining the ruleset's data quality rules.
	//
	// Build it with `Dqdl.fromString(...)`.
	// Experimental.
	Dqdl Dqdl `field:"required" json:"dqdl" yaml:"dqdl"`
	// The target table of the ruleset.
	// Experimental.
	TargetTable DataQualityTargetTable `field:"required" json:"targetTable" yaml:"targetTable"`
	// The client token of the ruleset.
	// Experimental.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The description of the ruleset.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Policy to apply when the ruleset is removed from the stack.
	// Default: - resource will be destroyed.
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
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

