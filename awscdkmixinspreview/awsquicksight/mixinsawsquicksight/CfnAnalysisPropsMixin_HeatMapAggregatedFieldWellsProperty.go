package mixinsawsquicksight


// The aggregated field wells of a heat map.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapaggregatedfieldwells.html
//
type CfnAnalysisPropsMixin_HeatMapAggregatedFieldWellsProperty struct {
	// The columns field well of a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapaggregatedfieldwells.html#cfn-quicksight-analysis-heatmapaggregatedfieldwells-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// The rows field well of a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapaggregatedfieldwells.html#cfn-quicksight-analysis-heatmapaggregatedfieldwells-rows
	//
	Rows interface{} `field:"optional" json:"rows" yaml:"rows"`
	// The values field well of a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapaggregatedfieldwells.html#cfn-quicksight-analysis-heatmapaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

