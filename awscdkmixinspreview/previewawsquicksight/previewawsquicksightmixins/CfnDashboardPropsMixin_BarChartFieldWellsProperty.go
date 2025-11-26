package previewawsquicksightmixins


// The field wells of a `BarChartVisual` .
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartfieldwells.html
//
type CfnDashboardPropsMixin_BarChartFieldWellsProperty struct {
	// The aggregated field wells of a bar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartfieldwells.html#cfn-quicksight-dashboard-barchartfieldwells-barchartaggregatedfieldwells
	//
	BarChartAggregatedFieldWells interface{} `field:"optional" json:"barChartAggregatedFieldWells" yaml:"barChartAggregatedFieldWells"`
}

