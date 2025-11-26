package previewawsquicksightmixins


// The field well configuration of a line chart.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartaggregatedfieldwells.html
//
type CfnDashboardPropsMixin_LineChartAggregatedFieldWellsProperty struct {
	// The category field wells of a line chart.
	//
	// Values are grouped by category fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartaggregatedfieldwells.html#cfn-quicksight-dashboard-linechartaggregatedfieldwells-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The color field wells of a line chart.
	//
	// Values are grouped by category fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartaggregatedfieldwells.html#cfn-quicksight-dashboard-linechartaggregatedfieldwells-colors
	//
	Colors interface{} `field:"optional" json:"colors" yaml:"colors"`
	// The small multiples field well of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartaggregatedfieldwells.html#cfn-quicksight-dashboard-linechartaggregatedfieldwells-smallmultiples
	//
	SmallMultiples interface{} `field:"optional" json:"smallMultiples" yaml:"smallMultiples"`
	// The value field wells of a line chart.
	//
	// Values are aggregated based on categories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartaggregatedfieldwells.html#cfn-quicksight-dashboard-linechartaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

