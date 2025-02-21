package awsquicksight


// The configuration of a `GaugeChartVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconfiguration.html
//
type CfnTemplate_GaugeChartConfigurationProperty struct {
	// The data label configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconfiguration.html#cfn-quicksight-template-gaugechartconfiguration-datalabels
	//
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field well configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconfiguration.html#cfn-quicksight-template-gaugechartconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The options that determine the presentation of the `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconfiguration.html#cfn-quicksight-template-gaugechartconfiguration-gaugechartoptions
	//
	GaugeChartOptions interface{} `field:"optional" json:"gaugeChartOptions" yaml:"gaugeChartOptions"`
	// The tooltip configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconfiguration.html#cfn-quicksight-template-gaugechartconfiguration-tooltipoptions
	//
	TooltipOptions interface{} `field:"optional" json:"tooltipOptions" yaml:"tooltipOptions"`
	// The visual palette configuration of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconfiguration.html#cfn-quicksight-template-gaugechartconfiguration-visualpalette
	//
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
}

