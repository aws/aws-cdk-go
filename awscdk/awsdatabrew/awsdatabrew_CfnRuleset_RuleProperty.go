package awsdatabrew


// Represents a single data quality requirement that should be validated in the scope of this dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &ruleProperty{
//   	checkExpression: jsii.String("checkExpression"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	columnSelectors: []interface{}{
//   		&columnSelectorProperty{
//   			name: jsii.String("name"),
//   			regex: jsii.String("regex"),
//   		},
//   	},
//   	disabled: jsii.Boolean(false),
//   	substitutionMap: []interface{}{
//   		&substitutionValueProperty{
//   			value: jsii.String("value"),
//   			valueReference: jsii.String("valueReference"),
//   		},
//   	},
//   	threshold: &thresholdProperty{
//   		value: jsii.Number(123),
//
//   		// the properties below are optional
//   		type: jsii.String("type"),
//   		unit: jsii.String("unit"),
//   	},
//   }
//
type CfnRuleset_RuleProperty struct {
	// The expression which includes column references, condition names followed by variable references, possibly grouped and combined with other conditions.
	//
	// For example, `(:col1 starts_with :prefix1 or :col1 starts_with :prefix2) and (:col1 ends_with :suffix1 or :col1 ends_with :suffix2)` . Column and value references are substitution variables that should start with the ':' symbol. Depending on the context, substitution variables' values can be either an actual value or a column name. These values are defined in the SubstitutionMap. If a CheckExpression starts with a column reference, then ColumnSelectors in the rule should be null. If ColumnSelectors has been defined, then there should be no columnn reference in the left side of a condition, for example, `is_between :val1 and :val2` .
	CheckExpression *string `field:"required" json:"checkExpression" yaml:"checkExpression"`
	// The name of the rule.
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of column selectors.
	//
	// Selectors can be used to select columns using a name or regular expression from the dataset. Rule will be applied to selected columns.
	ColumnSelectors interface{} `field:"optional" json:"columnSelectors" yaml:"columnSelectors"`
	// A value that specifies whether the rule is disabled.
	//
	// Once a rule is disabled, a profile job will not validate it during a job run. Default value is false.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// The map of substitution variable names to their values used in a check expression.
	//
	// Variable names should start with a ':' (colon). Variable values can either be actual values or column names. To differentiate between the two, column names should be enclosed in backticks, for example, `":col1": "`Column A`".`
	SubstitutionMap interface{} `field:"optional" json:"substitutionMap" yaml:"substitutionMap"`
	// The threshold used with a non-aggregate check expression.
	//
	// Non-aggregate check expressions will be applied to each row in a specific column, and the threshold will be used to determine whether the validation succeeds.
	Threshold interface{} `field:"optional" json:"threshold" yaml:"threshold"`
}

