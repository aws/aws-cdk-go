package previewawsquicksightmixins


// The aggregated field wells of a bar chart.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartaggregatedfieldwells.html
//
type CfnDashboardPropsMixin_BarChartAggregatedFieldWellsProperty struct {
	// The category (y-axis) field well of a bar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartaggregatedfieldwells.html#cfn-quicksight-dashboard-barchartaggregatedfieldwells-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The color (group/color) field well of a bar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartaggregatedfieldwells.html#cfn-quicksight-dashboard-barchartaggregatedfieldwells-colors
	//
	Colors interface{} `field:"optional" json:"colors" yaml:"colors"`
	// The small multiples field well of a bar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartaggregatedfieldwells.html#cfn-quicksight-dashboard-barchartaggregatedfieldwells-smallmultiples
	//
	SmallMultiples interface{} `field:"optional" json:"smallMultiples" yaml:"smallMultiples"`
	// The value field wells of a bar chart.
	//
	// Values are aggregated by category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartaggregatedfieldwells.html#cfn-quicksight-dashboard-barchartaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

