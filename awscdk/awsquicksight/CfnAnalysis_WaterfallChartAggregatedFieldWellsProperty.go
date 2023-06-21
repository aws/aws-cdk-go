package awsquicksight


// The field well configuration of a waterfall visual.
//
// Example:
//
//
type CfnAnalysis_WaterfallChartAggregatedFieldWellsProperty struct {
	// The breakdown field wells of a waterfall visual.
	Breakdowns interface{} `field:"optional" json:"breakdowns" yaml:"breakdowns"`
	// The category field wells of a waterfall visual.
	Categories interface{} `field:"optional" json:"categories" yaml:"categories"`
	// The value field wells of a waterfall visual.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

