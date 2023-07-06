package awsquicksight


// The aggregated field well of a scatter plot.
//
// Example:
//
//
type CfnTemplate_ScatterPlotCategoricallyAggregatedFieldWellsProperty struct {
	// The category field well of a scatter plot.
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The label field well of a scatter plot.
	Label interface{} `field:"optional" json:"label" yaml:"label"`
	// The size field well of a scatter plot.
	Size interface{} `field:"optional" json:"size" yaml:"size"`
	// The x-axis field well of a scatter plot.
	//
	// The x-axis is aggregated by category.
	XAxis interface{} `field:"optional" json:"xAxis" yaml:"xAxis"`
	// The y-axis field well of a scatter plot.
	//
	// The y-axis is aggregated by category.
	YAxis interface{} `field:"optional" json:"yAxis" yaml:"yAxis"`
}

