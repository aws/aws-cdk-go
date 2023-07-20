package awsquicksight


// The unaggregated field wells of a scatter plot.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotunaggregatedfieldwells.html
//
type CfnDashboard_ScatterPlotUnaggregatedFieldWellsProperty struct {
	// The category field well of a scatter plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotunaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotunaggregatedfieldwells-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The label field well of a scatter plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotunaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotunaggregatedfieldwells-label
	//
	Label interface{} `field:"optional" json:"label" yaml:"label"`
	// The size field well of a scatter plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotunaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotunaggregatedfieldwells-size
	//
	Size interface{} `field:"optional" json:"size" yaml:"size"`
	// The x-axis field well of a scatter plot.
	//
	// The x-axis is a dimension field and cannot be aggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotunaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotunaggregatedfieldwells-xaxis
	//
	XAxis interface{} `field:"optional" json:"xAxis" yaml:"xAxis"`
	// The y-axis field well of a scatter plot.
	//
	// The y-axis is a dimension field and cannot be aggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotunaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotunaggregatedfieldwells-yaxis
	//
	YAxis interface{} `field:"optional" json:"yAxis" yaml:"yAxis"`
}

