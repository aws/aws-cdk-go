package previewawsquicksightmixins


// The configuration of a line chart.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html
//
type CfnTemplatePropsMixin_LineChartConfigurationProperty struct {
	// The default configuration of a line chart's contribution analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-contributionanalysisdefaults
	//
	ContributionAnalysisDefaults interface{} `field:"optional" json:"contributionAnalysisDefaults" yaml:"contributionAnalysisDefaults"`
	// The data label configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-datalabels
	//
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The options that determine the default presentation of all line series in `LineChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-defaultseriessettings
	//
	DefaultSeriesSettings interface{} `field:"optional" json:"defaultSeriesSettings" yaml:"defaultSeriesSettings"`
	// The field well configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The forecast configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-forecastconfigurations
	//
	ForecastConfigurations interface{} `field:"optional" json:"forecastConfigurations" yaml:"forecastConfigurations"`
	// The general visual interactions setup for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
	// The legend configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The series axis configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-primaryyaxisdisplayoptions
	//
	PrimaryYAxisDisplayOptions interface{} `field:"optional" json:"primaryYAxisDisplayOptions" yaml:"primaryYAxisDisplayOptions"`
	// The options that determine the presentation of the y-axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-primaryyaxislabeloptions
	//
	PrimaryYAxisLabelOptions interface{} `field:"optional" json:"primaryYAxisLabelOptions" yaml:"primaryYAxisLabelOptions"`
	// The reference lines configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-referencelines
	//
	ReferenceLines interface{} `field:"optional" json:"referenceLines" yaml:"referenceLines"`
	// The series axis configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-secondaryyaxisdisplayoptions
	//
	SecondaryYAxisDisplayOptions interface{} `field:"optional" json:"secondaryYAxisDisplayOptions" yaml:"secondaryYAxisDisplayOptions"`
	// The options that determine the presentation of the secondary y-axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-secondaryyaxislabeloptions
	//
	SecondaryYAxisLabelOptions interface{} `field:"optional" json:"secondaryYAxisLabelOptions" yaml:"secondaryYAxisLabelOptions"`
	// The series item configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-series
	//
	Series interface{} `field:"optional" json:"series" yaml:"series"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-singleaxisoptions
	//
	SingleAxisOptions interface{} `field:"optional" json:"singleAxisOptions" yaml:"singleAxisOptions"`
	// The small multiples setup for the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-smallmultiplesoptions
	//
	SmallMultiplesOptions interface{} `field:"optional" json:"smallMultiplesOptions" yaml:"smallMultiplesOptions"`
	// The sort configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// Determines the type of the line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The visual palette configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-visualpalette
	//
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
	// The options that determine the presentation of the x-axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-xaxisdisplayoptions
	//
	XAxisDisplayOptions interface{} `field:"optional" json:"xAxisDisplayOptions" yaml:"xAxisDisplayOptions"`
	// The options that determine the presentation of the x-axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartconfiguration.html#cfn-quicksight-template-linechartconfiguration-xaxislabeloptions
	//
	XAxisLabelOptions interface{} `field:"optional" json:"xAxisLabelOptions" yaml:"xAxisLabelOptions"`
}

