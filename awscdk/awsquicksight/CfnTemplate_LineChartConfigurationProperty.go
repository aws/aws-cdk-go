package awsquicksight


// The configuration of a line chart.
//
// Example:
//
//
type CfnTemplate_LineChartConfigurationProperty struct {
	// The default configuration of a line chart's contribution analysis.
	ContributionAnalysisDefaults interface{} `field:"optional" json:"contributionAnalysisDefaults" yaml:"contributionAnalysisDefaults"`
	// The data label configuration of a line chart.
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The options that determine the default presentation of all line series in `LineChartVisual` .
	DefaultSeriesSettings interface{} `field:"optional" json:"defaultSeriesSettings" yaml:"defaultSeriesSettings"`
	// The field well configuration of a line chart.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The forecast configuration of a line chart.
	ForecastConfigurations interface{} `field:"optional" json:"forecastConfigurations" yaml:"forecastConfigurations"`
	// The legend configuration of a line chart.
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The series axis configuration of a line chart.
	PrimaryYAxisDisplayOptions interface{} `field:"optional" json:"primaryYAxisDisplayOptions" yaml:"primaryYAxisDisplayOptions"`
	// The options that determine the presentation of the y-axis label.
	PrimaryYAxisLabelOptions interface{} `field:"optional" json:"primaryYAxisLabelOptions" yaml:"primaryYAxisLabelOptions"`
	// The reference lines configuration of a line chart.
	ReferenceLines interface{} `field:"optional" json:"referenceLines" yaml:"referenceLines"`
	// The series axis configuration of a line chart.
	SecondaryYAxisDisplayOptions interface{} `field:"optional" json:"secondaryYAxisDisplayOptions" yaml:"secondaryYAxisDisplayOptions"`
	// The options that determine the presentation of the secondary y-axis label.
	SecondaryYAxisLabelOptions interface{} `field:"optional" json:"secondaryYAxisLabelOptions" yaml:"secondaryYAxisLabelOptions"`
	// The series item configuration of a line chart.
	Series interface{} `field:"optional" json:"series" yaml:"series"`
	// The small multiples setup for the visual.
	SmallMultiplesOptions interface{} `field:"optional" json:"smallMultiplesOptions" yaml:"smallMultiplesOptions"`
	// The sort configuration of a line chart.
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip configuration of a line chart.
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// Determines the type of the line chart.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The visual palette configuration of a line chart.
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
	// The options that determine the presentation of the x-axis.
	XAxisDisplayOptions interface{} `field:"optional" json:"xAxisDisplayOptions" yaml:"xAxisDisplayOptions"`
	// The options that determine the presentation of the x-axis label.
	XAxisLabelOptions interface{} `field:"optional" json:"xAxisLabelOptions" yaml:"xAxisLabelOptions"`
}

