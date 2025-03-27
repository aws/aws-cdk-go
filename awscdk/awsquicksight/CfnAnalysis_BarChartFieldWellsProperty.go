package awsquicksight


// The field wells of a `BarChartVisual` .
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-barchartfieldwells.html
//
type CfnAnalysis_BarChartFieldWellsProperty struct {
	// The aggregated field wells of a bar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-barchartfieldwells.html#cfn-quicksight-analysis-barchartfieldwells-barchartaggregatedfieldwells
	//
	BarChartAggregatedFieldWells interface{} `field:"optional" json:"barChartAggregatedFieldWells" yaml:"barChartAggregatedFieldWells"`
}

