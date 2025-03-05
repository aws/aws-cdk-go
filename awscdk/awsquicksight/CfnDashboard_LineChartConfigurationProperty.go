package awsquicksight


// The configuration of a line chart.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html
//
type CfnDashboard_LineChartConfigurationProperty struct {
	// The default configuration of a line chart's contribution analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-contributionanalysisdefaults
	//
	ContributionAnalysisDefaults interface{} `field:"optional" json:"contributionAnalysisDefaults" yaml:"contributionAnalysisDefaults"`
	// The data label configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-datalabels
	//
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The options that determine the default presentation of all line series in `LineChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-defaultseriessettings
	//
	DefaultSeriesSettings interface{} `field:"optional" json:"defaultSeriesSettings" yaml:"defaultSeriesSettings"`
	// The field well configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The forecast configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-forecastconfigurations
	//
	ForecastConfigurations interface{} `field:"optional" json:"forecastConfigurations" yaml:"forecastConfigurations"`
	// The legend configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The series axis configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-primaryyaxisdisplayoptions
	//
	PrimaryYAxisDisplayOptions interface{} `field:"optional" json:"primaryYAxisDisplayOptions" yaml:"primaryYAxisDisplayOptions"`
	// The options that determine the presentation of the y-axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-primaryyaxislabeloptions
	//
	PrimaryYAxisLabelOptions interface{} `field:"optional" json:"primaryYAxisLabelOptions" yaml:"primaryYAxisLabelOptions"`
	// The reference lines configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-referencelines
	//
	ReferenceLines interface{} `field:"optional" json:"referenceLines" yaml:"referenceLines"`
	// The series axis configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-secondaryyaxisdisplayoptions
	//
	SecondaryYAxisDisplayOptions interface{} `field:"optional" json:"secondaryYAxisDisplayOptions" yaml:"secondaryYAxisDisplayOptions"`
	// The options that determine the presentation of the secondary y-axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-secondaryyaxislabeloptions
	//
	SecondaryYAxisLabelOptions interface{} `field:"optional" json:"secondaryYAxisLabelOptions" yaml:"secondaryYAxisLabelOptions"`
	// The series item configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-series
	//
	Series interface{} `field:"optional" json:"series" yaml:"series"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-singleaxisoptions
	//
	SingleAxisOptions interface{} `field:"optional" json:"singleAxisOptions" yaml:"singleAxisOptions"`
	// The small multiples setup for the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-smallmultiplesoptions
	//
	SmallMultiplesOptions interface{} `field:"optional" json:"smallMultiplesOptions" yaml:"smallMultiplesOptions"`
	// The sort configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// Determines the type of the line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The visual palette configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-visualpalette
	//
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
	// The options that determine the presentation of the x-axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-xaxisdisplayoptions
	//
	XAxisDisplayOptions interface{} `field:"optional" json:"xAxisDisplayOptions" yaml:"xAxisDisplayOptions"`
	// The options that determine the presentation of the x-axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartconfiguration.html#cfn-quicksight-dashboard-linechartconfiguration-xaxislabeloptions
	//
	XAxisLabelOptions interface{} `field:"optional" json:"xAxisLabelOptions" yaml:"xAxisLabelOptions"`
}

