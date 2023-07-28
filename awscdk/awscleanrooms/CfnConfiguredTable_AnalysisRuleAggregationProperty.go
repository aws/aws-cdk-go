package awscleanrooms


// Enables query structure and specified queries that produce aggregate statistics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisRuleAggregationProperty := &AnalysisRuleAggregationProperty{
//   	AggregateColumns: []interface{}{
//   		&AggregateColumnProperty{
//   			ColumnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			Function: jsii.String("function"),
//   		},
//   	},
//   	DimensionColumns: []*string{
//   		jsii.String("dimensionColumns"),
//   	},
//   	JoinColumns: []*string{
//   		jsii.String("joinColumns"),
//   	},
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
//
//   	// the properties below are optional
//   	AllowedJoinOperators: []*string{
//   		jsii.String("allowedJoinOperators"),
//   	},
//   	JoinRequired: jsii.String("joinRequired"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html
//
type CfnConfiguredTable_AnalysisRuleAggregationProperty struct {
	// The columns that query runners are allowed to use in aggregation queries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-aggregatecolumns
	//
	AggregateColumns interface{} `field:"required" json:"aggregateColumns" yaml:"aggregateColumns"`
	// The columns that query runners are allowed to select, group by, or filter by.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-dimensioncolumns
	//
	DimensionColumns *[]*string `field:"required" json:"dimensionColumns" yaml:"dimensionColumns"`
	// Columns in configured table that can be used in join statements and/or as aggregate columns.
	//
	// They can never be outputted directly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-joincolumns
	//
	JoinColumns *[]*string `field:"required" json:"joinColumns" yaml:"joinColumns"`
	// Columns that must meet a specific threshold value (after an aggregation function is applied to it) for each output row to be returned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-outputconstraints
	//
	OutputConstraints interface{} `field:"required" json:"outputConstraints" yaml:"outputConstraints"`
	// Set of scalar functions that are allowed to be used on dimension columns and the output of aggregation of metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-scalarfunctions
	//
	ScalarFunctions *[]*string `field:"required" json:"scalarFunctions" yaml:"scalarFunctions"`
	// Which logical operators (if any) are to be used in an INNER JOIN match condition.
	//
	// Default is `AND` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-allowedjoinoperators
	//
	AllowedJoinOperators *[]*string `field:"optional" json:"allowedJoinOperators" yaml:"allowedJoinOperators"`
	// Control that requires member who runs query to do a join with their configured table and/or other configured table in query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-analysisruleaggregation.html#cfn-cleanrooms-configuredtable-analysisruleaggregation-joinrequired
	//
	JoinRequired *string `field:"optional" json:"joinRequired" yaml:"joinRequired"`
}

