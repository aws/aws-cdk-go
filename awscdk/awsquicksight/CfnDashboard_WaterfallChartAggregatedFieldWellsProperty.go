package awsquicksight


// The field well configuration of a waterfall visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartaggregatedfieldwells.html
//
type CfnDashboard_WaterfallChartAggregatedFieldWellsProperty struct {
	// The breakdown field wells of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartaggregatedfieldwells.html#cfn-quicksight-dashboard-waterfallchartaggregatedfieldwells-breakdowns
	//
	Breakdowns interface{} `field:"optional" json:"breakdowns" yaml:"breakdowns"`
	// The category field wells of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartaggregatedfieldwells.html#cfn-quicksight-dashboard-waterfallchartaggregatedfieldwells-categories
	//
	Categories interface{} `field:"optional" json:"categories" yaml:"categories"`
	// The value field wells of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-waterfallchartaggregatedfieldwells.html#cfn-quicksight-dashboard-waterfallchartaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

