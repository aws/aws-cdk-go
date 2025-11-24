package mixinsawscleanrooms


// A type of analysis rule that enables query structure and specified queries that produce aggregate statistics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analysisRuleAggregationProperty := &AnalysisRuleAggregationProperty{
//   	AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   	AggregateColumns: []interface{}{
//   		&AggregateColumnProperty{
//   			ColumnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			Function: jsii.String("function"),
//   		},
//   	},
//   	AllowedJoinOperators: []*string{
//   		jsii.String("allowedJoinOperators"),
//   	},
//   	DimensionColumns: []*string{
//   		jsii.String("dimensionColumns"),
//   	},
//   	JoinColumns: []*string{
//   		jsii.String("joinColumns"),
//   	},
//   	JoinRequired: jsii.String("joinRequired"),
//   	OutputConstraints: []interface{}{
//   		&AggregationConstraintProperty{
//   			ColumnName: jsii.String("columnName"),
//   			Minimum: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ScalarFunctions: []*string{
//   		jsii.String("scalarFunctions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html
//
type CfnConfiguredTablePropsMixin_AnalysisRuleAggregationProperty struct {
	// An indicator as to whether additional analyses (such as AWS Clean Rooms ML) can be applied to the output of the direct query.
	//
	// The `additionalAnalyses` parameter is currently supported for the list analysis rule ( `AnalysisRuleList` ) and the custom analysis rule ( `AnalysisRuleCustom` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-additionalanalyses
	//
	AdditionalAnalyses *string `field:"optional" json:"additionalAnalyses" yaml:"additionalAnalyses"`
	// The columns that query runners are allowed to use in aggregation queries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-aggregatecolumns
	//
	AggregateColumns interface{} `field:"optional" json:"aggregateColumns" yaml:"aggregateColumns"`
	// Which logical operators (if any) are to be used in an INNER JOIN match condition.
	//
	// Default is `AND` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-allowedjoinoperators
	//
	AllowedJoinOperators *[]*string `field:"optional" json:"allowedJoinOperators" yaml:"allowedJoinOperators"`
	// The columns that query runners are allowed to select, group by, or filter by.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-dimensioncolumns
	//
	DimensionColumns *[]*string `field:"optional" json:"dimensionColumns" yaml:"dimensionColumns"`
	// Columns in configured table that can be used in join statements and/or as aggregate columns.
	//
	// They can never be outputted directly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-joincolumns
	//
	JoinColumns *[]*string `field:"optional" json:"joinColumns" yaml:"joinColumns"`
	// Control that requires member who runs query to do a join with their configured table and/or other configured table in query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-joinrequired
	//
	JoinRequired *string `field:"optional" json:"joinRequired" yaml:"joinRequired"`
	// Columns that must meet a specific threshold value (after an aggregation function is applied to it) for each output row to be returned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-outputconstraints
	//
	OutputConstraints interface{} `field:"optional" json:"outputConstraints" yaml:"outputConstraints"`
	// Set of scalar functions that are allowed to be used on dimension columns and the output of aggregation of metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-scalarfunctions
	//
	ScalarFunctions *[]*string `field:"optional" json:"scalarFunctions" yaml:"scalarFunctions"`
}

