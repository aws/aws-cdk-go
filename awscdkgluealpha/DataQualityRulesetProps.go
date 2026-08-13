package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for `DataQualityRuleset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataQualityTargetTable DataQualityTargetTable
//
//   dataQualityRulesetProps := &DataQualityRulesetProps{
//   	RulesetDqdl: jsii.String("rulesetDqdl"),
//   	TargetTable: dataQualityTargetTable,
//
//   	// the properties below are optional
//   	ClientToken: jsii.String("clientToken"),
//   	Description: jsii.String("description"),
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	RulesetName: jsii.String("rulesetName"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
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

