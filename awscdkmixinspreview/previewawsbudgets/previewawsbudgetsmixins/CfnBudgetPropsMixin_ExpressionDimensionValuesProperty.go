package previewawsbudgetsmixins


// Contains the specifications for the filters to use for your request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   expressionDimensionValuesProperty := &ExpressionDimensionValuesProperty{
//   	Key: jsii.String("key"),
//   	MatchOptions: []*string{
//   		jsii.String("matchOptions"),
//   	},
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-expressiondimensionvalues.html
//
type CfnBudgetPropsMixin_ExpressionDimensionValuesProperty struct {
	// The name of the dimension that you want to filter on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-expressiondimensionvalues.html#cfn-budgets-budget-expressiondimensionvalues-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The match options that you can use to filter your results.
	//
	// You can specify only one of these values in the array.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-expressiondimensionvalues.html#cfn-budgets-budget-expressiondimensionvalues-matchoptions
	//
	MatchOptions *[]*string `field:"optional" json:"matchOptions" yaml:"matchOptions"`
	// The metadata values you can specify to filter upon, so that the results all match at least one of the specified values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-expressiondimensionvalues.html#cfn-budgets-budget-expressiondimensionvalues-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

