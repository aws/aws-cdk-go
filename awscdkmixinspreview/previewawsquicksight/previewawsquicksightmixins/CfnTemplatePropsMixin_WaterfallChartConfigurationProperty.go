package previewawsquicksightmixins


// The configuration for a waterfall visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartconfiguration.html
//
type CfnTemplatePropsMixin_WaterfallChartConfigurationProperty struct {
	// The options that determine the presentation of the category axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartconfiguration.html#cfn-quicksight-template-waterfallchartconfiguration-categoryaxisdisplayoptions
	//
	CategoryAxisDisplayOptions interface{} `field:"optional" json:"categoryAxisDisplayOptions" yaml:"categoryAxisDisplayOptions"`
	// The options that determine the presentation of the category axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartconfiguration.html#cfn-quicksight-template-waterfallchartconfiguration-categoryaxislabeloptions
	//
	CategoryAxisLabelOptions interface{} `field:"optional" json:"categoryAxisLabelOptions" yaml:"categoryAxisLabelOptions"`
	// The color configuration of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartconfiguration.html#cfn-quicksight-template-waterfallchartconfiguration-colorconfiguration
	//
	ColorConfiguration interface{} `field:"optional" json:"colorConfiguration" yaml:"colorConfiguration"`
	// The data label configuration of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartconfiguration.html#cfn-quicksight-template-waterfallchartconfiguration-datalabels
	//
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field well configuration of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartconfiguration.html#cfn-quicksight-template-waterfallchartconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The general visual interactions setup for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartconfiguration.html#cfn-quicksight-template-waterfallchartconfiguration-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
	// The legend configuration of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartconfiguration.html#cfn-quicksight-template-waterfallchartconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The options that determine the presentation of the y-axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartconfiguration.html#cfn-quicksight-template-waterfallchartconfiguration-primaryyaxisdisplayoptions
	//
	PrimaryYAxisDisplayOptions interface{} `field:"optional" json:"primaryYAxisDisplayOptions" yaml:"primaryYAxisDisplayOptions"`
	// The options that determine the presentation of the y-axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartconfiguration.html#cfn-quicksight-template-waterfallchartconfiguration-primaryyaxislabeloptions
	//
	PrimaryYAxisLabelOptions interface{} `field:"optional" json:"primaryYAxisLabelOptions" yaml:"primaryYAxisLabelOptions"`
	// The sort configuration of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartconfiguration.html#cfn-quicksight-template-waterfallchartconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The visual palette configuration of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartconfiguration.html#cfn-quicksight-template-waterfallchartconfiguration-visualpalette
	//
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
	// The options that determine the presentation of a waterfall visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-waterfallchartconfiguration.html#cfn-quicksight-template-waterfallchartconfiguration-waterfallchartoptions
	//
	WaterfallChartOptions interface{} `field:"optional" json:"waterfallChartOptions" yaml:"waterfallChartOptions"`
}

