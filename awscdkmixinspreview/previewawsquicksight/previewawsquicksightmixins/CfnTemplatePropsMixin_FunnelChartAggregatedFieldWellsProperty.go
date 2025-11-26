package previewawsquicksightmixins


// The field well configuration of a `FunnelChartVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartaggregatedfieldwells.html
//
type CfnTemplatePropsMixin_FunnelChartAggregatedFieldWellsProperty struct {
	// The category field wells of a funnel chart.
	//
	// Values are grouped by category fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartaggregatedfieldwells.html#cfn-quicksight-template-funnelchartaggregatedfieldwells-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The value field wells of a funnel chart.
	//
	// Values are aggregated based on categories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartaggregatedfieldwells.html#cfn-quicksight-template-funnelchartaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

