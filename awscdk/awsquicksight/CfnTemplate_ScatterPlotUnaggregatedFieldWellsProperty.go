package awsquicksight


// The unaggregated field wells of a scatter plot.
//
// Example:
//
//
type CfnTemplate_ScatterPlotUnaggregatedFieldWellsProperty struct {
	// The category field well of a scatter plot.
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The label field well of a scatter plot.
	Label interface{} `field:"optional" json:"label" yaml:"label"`
	// The size field well of a scatter plot.
	Size interface{} `field:"optional" json:"size" yaml:"size"`
	// The x-axis field well of a scatter plot.
	//
	// The x-axis is a dimension field and cannot be aggregated.
	XAxis interface{} `field:"optional" json:"xAxis" yaml:"xAxis"`
	// The y-axis field well of a scatter plot.
	//
	// The y-axis is a dimension field and cannot be aggregated.
	YAxis interface{} `field:"optional" json:"yAxis" yaml:"yAxis"`
}

