package previewawsquicksightmixins


// The field well configuration of a `FunnelChartVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartaggregatedfieldwells.html
//
type CfnAnalysisPropsMixin_FunnelChartAggregatedFieldWellsProperty struct {
	// The category field wells of a funnel chart.
	//
	// Values are grouped by category fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartaggregatedfieldwells.html#cfn-quicksight-analysis-funnelchartaggregatedfieldwells-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The value field wells of a funnel chart.
	//
	// Values are aggregated based on categories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-funnelchartaggregatedfieldwells.html#cfn-quicksight-analysis-funnelchartaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

