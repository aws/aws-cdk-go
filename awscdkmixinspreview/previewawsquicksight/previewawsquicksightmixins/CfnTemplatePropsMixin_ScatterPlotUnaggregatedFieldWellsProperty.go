package previewawsquicksightmixins


// The unaggregated field wells of a scatter plot.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotunaggregatedfieldwells.html
//
type CfnTemplatePropsMixin_ScatterPlotUnaggregatedFieldWellsProperty struct {
	// The category field well of a scatter plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotunaggregatedfieldwells.html#cfn-quicksight-template-scatterplotunaggregatedfieldwells-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The label field well of a scatter plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotunaggregatedfieldwells.html#cfn-quicksight-template-scatterplotunaggregatedfieldwells-label
	//
	Label interface{} `field:"optional" json:"label" yaml:"label"`
	// The size field well of a scatter plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotunaggregatedfieldwells.html#cfn-quicksight-template-scatterplotunaggregatedfieldwells-size
	//
	Size interface{} `field:"optional" json:"size" yaml:"size"`
	// The x-axis field well of a scatter plot.
	//
	// The x-axis is a dimension field and cannot be aggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotunaggregatedfieldwells.html#cfn-quicksight-template-scatterplotunaggregatedfieldwells-xaxis
	//
	XAxis interface{} `field:"optional" json:"xAxis" yaml:"xAxis"`
	// The y-axis field well of a scatter plot.
	//
	// The y-axis is a dimension field and cannot be aggregated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotunaggregatedfieldwells.html#cfn-quicksight-template-scatterplotunaggregatedfieldwells-yaxis
	//
	YAxis interface{} `field:"optional" json:"yAxis" yaml:"yAxis"`
}

