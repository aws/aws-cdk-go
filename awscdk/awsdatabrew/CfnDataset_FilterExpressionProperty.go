package awsdatabrew


// Represents a structure for defining parameter conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterExpressionProperty := &FilterExpressionProperty{
//   	Expression: jsii.String("expression"),
//   	ValuesMap: []interface{}{
//   		&FilterValueProperty{
//   			Value: jsii.String("value"),
//   			ValueReference: jsii.String("valueReference"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-filterexpression.html
//
type CfnDataset_FilterExpressionProperty struct {
	// The expression which includes condition names followed by substitution variables, possibly grouped and combined with other conditions.
	//
	// For example, "(starts_with :prefix1 or starts_with :prefix2) and (ends_with :suffix1 or ends_with :suffix2)". Substitution variables should start with ':' symbol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-filterexpression.html#cfn-databrew-dataset-filterexpression-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The map of substitution variable names to their values used in this filter expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-filterexpression.html#cfn-databrew-dataset-filterexpression-valuesmap
	//
	ValuesMap interface{} `field:"required" json:"valuesMap" yaml:"valuesMap"`
}

