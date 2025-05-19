package awsquicksight


// The field wells of a `BoxPlotVisual` .
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-boxplotfieldwells.html
//
type CfnAnalysis_BoxPlotFieldWellsProperty struct {
	// The aggregated field wells of a box plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-boxplotfieldwells.html#cfn-quicksight-analysis-boxplotfieldwells-boxplotaggregatedfieldwells
	//
	BoxPlotAggregatedFieldWells interface{} `field:"optional" json:"boxPlotAggregatedFieldWells" yaml:"boxPlotAggregatedFieldWells"`
}

