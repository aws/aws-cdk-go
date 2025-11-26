package previewawsquicksightmixins


// The aggregated field well configuration of a `RadarChartVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartaggregatedfieldwells.html
//
type CfnTemplatePropsMixin_RadarChartAggregatedFieldWellsProperty struct {
	// The aggregated field well categories of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartaggregatedfieldwells.html#cfn-quicksight-template-radarchartaggregatedfieldwells-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The color that are assigned to the aggregated field wells of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartaggregatedfieldwells.html#cfn-quicksight-template-radarchartaggregatedfieldwells-color
	//
	Color interface{} `field:"optional" json:"color" yaml:"color"`
	// The values that are assigned to the aggregated field wells of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-radarchartaggregatedfieldwells.html#cfn-quicksight-template-radarchartaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

