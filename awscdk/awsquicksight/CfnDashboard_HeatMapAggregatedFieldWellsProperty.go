package awsquicksight


// The aggregated field wells of a heat map.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapaggregatedfieldwells.html
//
type CfnDashboard_HeatMapAggregatedFieldWellsProperty struct {
	// The columns field well of a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapaggregatedfieldwells.html#cfn-quicksight-dashboard-heatmapaggregatedfieldwells-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// The rows field well of a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapaggregatedfieldwells.html#cfn-quicksight-dashboard-heatmapaggregatedfieldwells-rows
	//
	Rows interface{} `field:"optional" json:"rows" yaml:"rows"`
	// The values field well of a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapaggregatedfieldwells.html#cfn-quicksight-dashboard-heatmapaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

