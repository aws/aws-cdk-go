package awsquicksight


// The configuration for a `HistogramVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html
//
type CfnAnalysis_HistogramConfigurationProperty struct {
	// The options that determine the presentation of histogram bins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-binoptions
	//
	BinOptions interface{} `field:"optional" json:"binOptions" yaml:"binOptions"`
	// The data label configuration of a histogram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-datalabels
	//
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field well configuration of a histogram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The tooltip configuration of a histogram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// The visual palette configuration of a histogram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-visualpalette
	//
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
	// The options that determine the presentation of the x-axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-xaxisdisplayoptions
	//
	XAxisDisplayOptions interface{} `field:"optional" json:"xAxisDisplayOptions" yaml:"xAxisDisplayOptions"`
	// The options that determine the presentation of the x-axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-xaxislabeloptions
	//
	XAxisLabelOptions interface{} `field:"optional" json:"xAxisLabelOptions" yaml:"xAxisLabelOptions"`
	// The options that determine the presentation of the y-axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-histogramconfiguration.html#cfn-quicksight-analysis-histogramconfiguration-yaxisdisplayoptions
	//
	YAxisDisplayOptions interface{} `field:"optional" json:"yAxisDisplayOptions" yaml:"yAxisDisplayOptions"`
}

