package awsquicksight


// The configuration of a `ComboChartVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html
//
type CfnTemplate_ComboChartConfigurationProperty struct {
	// The options that determine if visual data labels are displayed.
	//
	// The data label options for a bar in a combo chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-bardatalabels
	//
	BarDataLabels interface{} `field:"optional" json:"barDataLabels" yaml:"barDataLabels"`
	// Determines the bar arrangement in a combo chart. The following are valid values in this structure:.
	//
	// - `CLUSTERED` : For clustered bar combo charts.
	// - `STACKED` : For stacked bar combo charts.
	// - `STACKED_PERCENT` : Do not use. If you use this value, the operation returns a validation error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-barsarrangement
	//
	BarsArrangement *string `field:"optional" json:"barsArrangement" yaml:"barsArrangement"`
	// The category axis of a combo chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-categoryaxis
	//
	CategoryAxis interface{} `field:"optional" json:"categoryAxis" yaml:"categoryAxis"`
	// The label options (label text, label visibility, and sort icon visibility) of a combo chart category (group/color) field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-categorylabeloptions
	//
	CategoryLabelOptions interface{} `field:"optional" json:"categoryLabelOptions" yaml:"categoryLabelOptions"`
	// The label options (label text, label visibility, and sort icon visibility) of a combo chart's color field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-colorlabeloptions
	//
	ColorLabelOptions interface{} `field:"optional" json:"colorLabelOptions" yaml:"colorLabelOptions"`
	// The field wells of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The legend display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The options that determine if visual data labels are displayed.
	//
	// The data label options for a line in a combo chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-linedatalabels
	//
	LineDataLabels interface{} `field:"optional" json:"lineDataLabels" yaml:"lineDataLabels"`
	// The label display options (grid line, range, scale, and axis step) of a combo chart's primary y-axis (bar) field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-primaryyaxisdisplayoptions
	//
	PrimaryYAxisDisplayOptions interface{} `field:"optional" json:"primaryYAxisDisplayOptions" yaml:"primaryYAxisDisplayOptions"`
	// The label options (label text, label visibility, and sort icon visibility) of a combo chart's primary y-axis (bar) field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-primaryyaxislabeloptions
	//
	PrimaryYAxisLabelOptions interface{} `field:"optional" json:"primaryYAxisLabelOptions" yaml:"primaryYAxisLabelOptions"`
	// The reference line setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-referencelines
	//
	ReferenceLines interface{} `field:"optional" json:"referenceLines" yaml:"referenceLines"`
	// The label display options (grid line, range, scale, axis step) of a combo chart's secondary y-axis (line) field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-secondaryyaxisdisplayoptions
	//
	SecondaryYAxisDisplayOptions interface{} `field:"optional" json:"secondaryYAxisDisplayOptions" yaml:"secondaryYAxisDisplayOptions"`
	// The label options (label text, label visibility, and sort icon visibility) of a combo chart's secondary y-axis(line) field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-secondaryyaxislabeloptions
	//
	SecondaryYAxisLabelOptions interface{} `field:"optional" json:"secondaryYAxisLabelOptions" yaml:"secondaryYAxisLabelOptions"`
	// The sort configuration of a `ComboChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The legend display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// The palette (chart color) display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-combochartconfiguration.html#cfn-quicksight-template-combochartconfiguration-visualpalette
	//
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
}

