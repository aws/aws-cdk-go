package mixinsawsquicksight


// The aggregated field wells of a heat map.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-heatmapaggregatedfieldwells.html
//
type CfnTemplatePropsMixin_HeatMapAggregatedFieldWellsProperty struct {
	// The columns field well of a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-heatmapaggregatedfieldwells.html#cfn-quicksight-template-heatmapaggregatedfieldwells-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// The rows field well of a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-heatmapaggregatedfieldwells.html#cfn-quicksight-template-heatmapaggregatedfieldwells-rows
	//
	Rows interface{} `field:"optional" json:"rows" yaml:"rows"`
	// The values field well of a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-heatmapaggregatedfieldwells.html#cfn-quicksight-template-heatmapaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

