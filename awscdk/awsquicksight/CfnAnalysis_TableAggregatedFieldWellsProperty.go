package awsquicksight


// The aggregated field well for the table.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tableaggregatedfieldwells.html
//
type CfnAnalysis_TableAggregatedFieldWellsProperty struct {
	// The group by field well for a pivot table.
	//
	// Values are grouped by group by fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tableaggregatedfieldwells.html#cfn-quicksight-analysis-tableaggregatedfieldwells-groupby
	//
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// The values field well for a pivot table.
	//
	// Values are aggregated based on group by fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tableaggregatedfieldwells.html#cfn-quicksight-analysis-tableaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

