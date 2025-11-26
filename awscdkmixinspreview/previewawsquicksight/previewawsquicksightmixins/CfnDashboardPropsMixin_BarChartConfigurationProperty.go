package previewawsquicksightmixins


// The configuration of a `BarChartVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html
//
type CfnDashboardPropsMixin_BarChartConfigurationProperty struct {
	// Determines the arrangement of the bars.
	//
	// The orientation and arrangement of bars determine the type of bar that is used in the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-barsarrangement
	//
	BarsArrangement *string `field:"optional" json:"barsArrangement" yaml:"barsArrangement"`
	// The label display options (grid line, range, scale, axis step) for bar chart category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-categoryaxis
	//
	CategoryAxis interface{} `field:"optional" json:"categoryAxis" yaml:"categoryAxis"`
	// The label options (label text, label visibility and sort icon visibility) for a bar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-categorylabeloptions
	//
	CategoryLabelOptions interface{} `field:"optional" json:"categoryLabelOptions" yaml:"categoryLabelOptions"`
	// The label options (label text, label visibility and sort icon visibility) for a color that is used in a bar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-colorlabeloptions
	//
	ColorLabelOptions interface{} `field:"optional" json:"colorLabelOptions" yaml:"colorLabelOptions"`
	// The contribution analysis (anomaly configuration) setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-contributionanalysisdefaults
	//
	ContributionAnalysisDefaults interface{} `field:"optional" json:"contributionAnalysisDefaults" yaml:"contributionAnalysisDefaults"`
	// The options that determine if visual data labels are displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-datalabels
	//
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field wells of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The general visual interactions setup for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
	// The legend display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The orientation of the bars in a bar chart visual. There are two valid values in this structure:.
	//
	// - `HORIZONTAL` : Used for charts that have horizontal bars. Visuals that use this value are horizontal bar charts, horizontal stacked bar charts, and horizontal stacked 100% bar charts.
	// - `VERTICAL` : Used for charts that have vertical bars. Visuals that use this value are vertical bar charts, vertical stacked bar charts, and vertical stacked 100% bar charts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-orientation
	//
	Orientation *string `field:"optional" json:"orientation" yaml:"orientation"`
	// The reference line setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-referencelines
	//
	ReferenceLines interface{} `field:"optional" json:"referenceLines" yaml:"referenceLines"`
	// The small multiples setup for the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-smallmultiplesoptions
	//
	SmallMultiplesOptions interface{} `field:"optional" json:"smallMultiplesOptions" yaml:"smallMultiplesOptions"`
	// The sort configuration of a `BarChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// The label display options (grid line, range, scale, axis step) for a bar chart value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-valueaxis
	//
	ValueAxis interface{} `field:"optional" json:"valueAxis" yaml:"valueAxis"`
	// The label options (label text, label visibility and sort icon visibility) for a bar chart value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-valuelabeloptions
	//
	ValueLabelOptions interface{} `field:"optional" json:"valueLabelOptions" yaml:"valueLabelOptions"`
	// The palette (chart color) display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-barchartconfiguration.html#cfn-quicksight-dashboard-barchartconfiguration-visualpalette
	//
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
}

