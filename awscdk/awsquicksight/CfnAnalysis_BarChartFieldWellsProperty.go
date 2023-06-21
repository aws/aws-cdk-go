package awsquicksight


// The field wells of a `BarChartVisual` .
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
type CfnAnalysis_BarChartFieldWellsProperty struct {
	// The aggregated field wells of a bar chart.
	BarChartAggregatedFieldWells interface{} `field:"optional" json:"barChartAggregatedFieldWells" yaml:"barChartAggregatedFieldWells"`
}

