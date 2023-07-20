package awsquicksight


// The field well configuration of a pie chart.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartaggregatedfieldwells.html
//
type CfnAnalysis_PieChartAggregatedFieldWellsProperty struct {
	// The category (group/color) field wells of a pie chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartaggregatedfieldwells.html#cfn-quicksight-analysis-piechartaggregatedfieldwells-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The small multiples field well of a pie chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartaggregatedfieldwells.html#cfn-quicksight-analysis-piechartaggregatedfieldwells-smallmultiples
	//
	SmallMultiples interface{} `field:"optional" json:"smallMultiples" yaml:"smallMultiples"`
	// The value field wells of a pie chart.
	//
	// Values are aggregated based on categories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartaggregatedfieldwells.html#cfn-quicksight-analysis-piechartaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

