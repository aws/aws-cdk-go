package awsquicksight


// The configuration of a scatter plot.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotconfiguration.html
//
type CfnTemplate_ScatterPlotConfigurationProperty struct {
	// The options that determine if visual data labels are displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotconfiguration.html#cfn-quicksight-template-scatterplotconfiguration-datalabels
	//
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field wells of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotconfiguration.html#cfn-quicksight-template-scatterplotconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The legend display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotconfiguration.html#cfn-quicksight-template-scatterplotconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The legend display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotconfiguration.html#cfn-quicksight-template-scatterplotconfiguration-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// The palette (chart color) display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotconfiguration.html#cfn-quicksight-template-scatterplotconfiguration-visualpalette
	//
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
	// The label display options (grid line, range, scale, and axis step) of the scatter plot's x-axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotconfiguration.html#cfn-quicksight-template-scatterplotconfiguration-xaxisdisplayoptions
	//
	XAxisDisplayOptions interface{} `field:"optional" json:"xAxisDisplayOptions" yaml:"xAxisDisplayOptions"`
	// The label options (label text, label visibility, and sort icon visibility) of the scatter plot's x-axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotconfiguration.html#cfn-quicksight-template-scatterplotconfiguration-xaxislabeloptions
	//
	XAxisLabelOptions interface{} `field:"optional" json:"xAxisLabelOptions" yaml:"xAxisLabelOptions"`
	// The label display options (grid line, range, scale, and axis step) of the scatter plot's y-axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotconfiguration.html#cfn-quicksight-template-scatterplotconfiguration-yaxisdisplayoptions
	//
	YAxisDisplayOptions interface{} `field:"optional" json:"yAxisDisplayOptions" yaml:"yAxisDisplayOptions"`
	// The label options (label text, label visibility, and sort icon visibility) of the scatter plot's y-axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-scatterplotconfiguration.html#cfn-quicksight-template-scatterplotconfiguration-yaxislabeloptions
	//
	YAxisLabelOptions interface{} `field:"optional" json:"yAxisLabelOptions" yaml:"yAxisLabelOptions"`
}

