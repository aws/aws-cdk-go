package awsquicksight


// The field well configuration of a scatter plot.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotfieldwells.html
//
type CfnDashboard_ScatterPlotFieldWellsProperty struct {
	// The aggregated field wells of a scatter plot.
	//
	// The x and y-axes of scatter plots with aggregated field wells are aggregated by category, label, or both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotfieldwells.html#cfn-quicksight-dashboard-scatterplotfieldwells-scatterplotcategoricallyaggregatedfieldwells
	//
	ScatterPlotCategoricallyAggregatedFieldWells interface{} `field:"optional" json:"scatterPlotCategoricallyAggregatedFieldWells" yaml:"scatterPlotCategoricallyAggregatedFieldWells"`
	// The unaggregated field wells of a scatter plot.
	//
	// The x and y-axes of these scatter plots are unaggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotfieldwells.html#cfn-quicksight-dashboard-scatterplotfieldwells-scatterplotunaggregatedfieldwells
	//
	ScatterPlotUnaggregatedFieldWells interface{} `field:"optional" json:"scatterPlotUnaggregatedFieldWells" yaml:"scatterPlotUnaggregatedFieldWells"`
}

