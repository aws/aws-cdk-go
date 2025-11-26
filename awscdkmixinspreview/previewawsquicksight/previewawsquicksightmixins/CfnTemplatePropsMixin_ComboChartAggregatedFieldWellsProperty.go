package previewawsquicksightmixins


// The aggregated field wells of a combo chart.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartaggregatedfieldwells.html
//
type CfnTemplatePropsMixin_ComboChartAggregatedFieldWellsProperty struct {
	// The aggregated `BarValues` field well of a combo chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartaggregatedfieldwells.html#cfn-quicksight-template-combochartaggregatedfieldwells-barvalues
	//
	BarValues interface{} `field:"optional" json:"barValues" yaml:"barValues"`
	// The aggregated category field wells of a combo chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartaggregatedfieldwells.html#cfn-quicksight-template-combochartaggregatedfieldwells-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The aggregated colors field well of a combo chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartaggregatedfieldwells.html#cfn-quicksight-template-combochartaggregatedfieldwells-colors
	//
	Colors interface{} `field:"optional" json:"colors" yaml:"colors"`
	// The aggregated `LineValues` field well of a combo chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartaggregatedfieldwells.html#cfn-quicksight-template-combochartaggregatedfieldwells-linevalues
	//
	LineValues interface{} `field:"optional" json:"lineValues" yaml:"lineValues"`
}

