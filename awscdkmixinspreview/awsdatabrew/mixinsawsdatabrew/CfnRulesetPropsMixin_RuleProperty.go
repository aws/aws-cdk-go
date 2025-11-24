package mixinsawsdatabrew


// Represents a single data quality requirement that should be validated in the scope of this dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleProperty := &RuleProperty{
//   	CheckExpression: jsii.String("checkExpression"),
//   	ColumnSelectors: []interface{}{
//   		&ColumnSelectorProperty{
//   			Name: jsii.String("name"),
//   			Regex: jsii.String("regex"),
//   		},
//   	},
//   	Disabled: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	SubstitutionMap: []interface{}{
//   		&SubstitutionValueProperty{
//   			Value: jsii.String("value"),
//   			ValueReference: jsii.String("valueReference"),
//   		},
//   	},
//   	Threshold: &ThresholdProperty{
//   		Type: jsii.String("type"),
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-rule.html
//
type CfnRulesetPropsMixin_RuleProperty struct {
	// The expression which includes column references, condition names followed by variable references, possibly grouped and combined with other conditions.
	//
	// For example, `(:col1 starts_with :prefix1 or :col1 starts_with :prefix2) and (:col1 ends_with :suffix1 or :col1 ends_with :suffix2)` . Column and value references are substitution variables that should start with the ':' symbol. Depending on the context, substitution variables' values can be either an actual value or a column name. These values are defined in the SubstitutionMap. If a CheckExpression starts with a column reference, then ColumnSelectors in the rule should be null. If ColumnSelectors has been defined, then there should be no columnn reference in the left side of a condition, for example, `is_between :val1 and :val2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-rule.html#cfn-databrew-ruleset-rule-checkexpression
	//
	CheckExpression *string `field:"optional" json:"checkExpression" yaml:"checkExpression"`
	// List of column selectors.
	//
	// Selectors can be used to select columns using a name or regular expression from the dataset. Rule will be applied to selected columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-rule.html#cfn-databrew-ruleset-rule-columnselectors
	//
	ColumnSelectors interface{} `field:"optional" json:"columnSelectors" yaml:"columnSelectors"`
	// A value that specifies whether the rule is disabled.
	//
	// Once a rule is disabled, a profile job will not validate it during a job run. Default value is false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-rule.html#cfn-databrew-ruleset-rule-disabled
	//
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// The name of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-rule.html#cfn-databrew-ruleset-rule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The map of substitution variable names to their values used in a check expression.
	//
	// Variable names should start with a ':' (colon). Variable values can either be actual values or column names. To differentiate between the two, column names should be enclosed in backticks, for example, `":col1": "`Column A`".`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-rule.html#cfn-databrew-ruleset-rule-substitutionmap
	//
	SubstitutionMap interface{} `field:"optional" json:"substitutionMap" yaml:"substitutionMap"`
	// The threshold used with a non-aggregate check expression.
	//
	// Non-aggregate check expressions will be applied to each row in a specific column, and the threshold will be used to determine whether the validation succeeds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-rule.html#cfn-databrew-ruleset-rule-threshold
	//
	Threshold interface{} `field:"optional" json:"threshold" yaml:"threshold"`
}

