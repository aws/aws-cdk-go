package awsquicksight


// The aggregated field well configuration of a `RadarChartVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartaggregatedfieldwells.html
//
type CfnDashboard_RadarChartAggregatedFieldWellsProperty struct {
	// The aggregated field well categories of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartaggregatedfieldwells.html#cfn-quicksight-dashboard-radarchartaggregatedfieldwells-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The color that are assigned to the aggregated field wells of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartaggregatedfieldwells.html#cfn-quicksight-dashboard-radarchartaggregatedfieldwells-color
	//
	Color interface{} `field:"optional" json:"color" yaml:"color"`
	// The values that are assigned to the aggregated field wells of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartaggregatedfieldwells.html#cfn-quicksight-dashboard-radarchartaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

