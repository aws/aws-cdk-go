package awsquicksight


// The aggregated field well for a box plot.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-boxplotaggregatedfieldwells.html
//
type CfnAnalysis_BoxPlotAggregatedFieldWellsProperty struct {
	// The group by field well of a box plot chart.
	//
	// Values are grouped based on group by fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-boxplotaggregatedfieldwells.html#cfn-quicksight-analysis-boxplotaggregatedfieldwells-groupby
	//
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// The value field well of a box plot chart.
	//
	// Values are aggregated based on group by fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-boxplotaggregatedfieldwells.html#cfn-quicksight-analysis-boxplotaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

