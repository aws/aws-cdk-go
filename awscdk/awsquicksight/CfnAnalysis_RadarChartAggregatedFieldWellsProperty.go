package awsquicksight


// The aggregated field well configuration of a `RadarChartVisual` .
//
// Example:
//
//
type CfnAnalysis_RadarChartAggregatedFieldWellsProperty struct {
	// The aggregated field well categories of a radar chart.
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The color that are assigned to the aggregated field wells of a radar chart.
	Color interface{} `field:"optional" json:"color" yaml:"color"`
	// The values that are assigned to the aggregated field wells of a radar chart.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

