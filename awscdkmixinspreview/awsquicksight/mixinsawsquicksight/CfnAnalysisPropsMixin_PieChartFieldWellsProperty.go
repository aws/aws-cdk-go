package mixinsawsquicksight


// The field well configuration of a pie chart.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartfieldwells.html
//
type CfnAnalysisPropsMixin_PieChartFieldWellsProperty struct {
	// The field well configuration of a pie chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartfieldwells.html#cfn-quicksight-analysis-piechartfieldwells-piechartaggregatedfieldwells
	//
	PieChartAggregatedFieldWells interface{} `field:"optional" json:"pieChartAggregatedFieldWells" yaml:"pieChartAggregatedFieldWells"`
}

