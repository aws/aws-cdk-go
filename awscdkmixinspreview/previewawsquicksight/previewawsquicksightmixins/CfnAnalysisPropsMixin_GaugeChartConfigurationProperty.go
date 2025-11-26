package previewawsquicksightmixins


// The configuration of a `GaugeChartVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartconfiguration.html
//
type CfnAnalysisPropsMixin_GaugeChartConfigurationProperty struct {
	// The color configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartconfiguration.html#cfn-quicksight-analysis-gaugechartconfiguration-colorconfiguration
	//
	ColorConfiguration interface{} `field:"optional" json:"colorConfiguration" yaml:"colorConfiguration"`
	// The data label configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartconfiguration.html#cfn-quicksight-analysis-gaugechartconfiguration-datalabels
	//
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field well configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartconfiguration.html#cfn-quicksight-analysis-gaugechartconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The options that determine the presentation of the `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartconfiguration.html#cfn-quicksight-analysis-gaugechartconfiguration-gaugechartoptions
	//
	GaugeChartOptions interface{} `field:"optional" json:"gaugeChartOptions" yaml:"gaugeChartOptions"`
	// The general visual interactions setup for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartconfiguration.html#cfn-quicksight-analysis-gaugechartconfiguration-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
	// The tooltip configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartconfiguration.html#cfn-quicksight-analysis-gaugechartconfiguration-tooltipoptions
	//
	TooltipOptions interface{} `field:"optional" json:"tooltipOptions" yaml:"tooltipOptions"`
	// The visual palette configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartconfiguration.html#cfn-quicksight-analysis-gaugechartconfiguration-visualpalette
	//
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
}

