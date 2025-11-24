package mixinsawsquicksight


// The aggregated field well of a scatter plot.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotcategoricallyaggregatedfieldwells.html
//
type CfnDashboardPropsMixin_ScatterPlotCategoricallyAggregatedFieldWellsProperty struct {
	// The category field well of a scatter plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotcategoricallyaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotcategoricallyaggregatedfieldwells-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The label field well of a scatter plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotcategoricallyaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotcategoricallyaggregatedfieldwells-label
	//
	Label interface{} `field:"optional" json:"label" yaml:"label"`
	// The size field well of a scatter plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotcategoricallyaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotcategoricallyaggregatedfieldwells-size
	//
	Size interface{} `field:"optional" json:"size" yaml:"size"`
	// The x-axis field well of a scatter plot.
	//
	// The x-axis is aggregated by category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotcategoricallyaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotcategoricallyaggregatedfieldwells-xaxis
	//
	XAxis interface{} `field:"optional" json:"xAxis" yaml:"xAxis"`
	// The y-axis field well of a scatter plot.
	//
	// The y-axis is aggregated by category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotcategoricallyaggregatedfieldwells.html#cfn-quicksight-dashboard-scatterplotcategoricallyaggregatedfieldwells-yaxis
	//
	YAxis interface{} `field:"optional" json:"yAxis" yaml:"yAxis"`
}

