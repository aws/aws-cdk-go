package awsquicksight


// The field well configuration of a scatter plot.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
type CfnTemplate_ScatterPlotFieldWellsProperty struct {
	// The aggregated field wells of a scatter plot.
	//
	// The x and y-axes of scatter plots with aggregated field wells are aggregated by category, label, or both.
	ScatterPlotCategoricallyAggregatedFieldWells interface{} `field:"optional" json:"scatterPlotCategoricallyAggregatedFieldWells" yaml:"scatterPlotCategoricallyAggregatedFieldWells"`
	// The unaggregated field wells of a scatter plot.
	//
	// The x and y-axes of these scatter plots are unaggregated.
	ScatterPlotUnaggregatedFieldWells interface{} `field:"optional" json:"scatterPlotUnaggregatedFieldWells" yaml:"scatterPlotUnaggregatedFieldWells"`
}

