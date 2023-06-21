package awsquicksight


// The configuration for a waterfall visual.
//
// Example:
//
//
type CfnAnalysis_WaterfallChartConfigurationProperty struct {
	// The options that determine the presentation of the category axis.
	CategoryAxisDisplayOptions interface{} `field:"optional" json:"categoryAxisDisplayOptions" yaml:"categoryAxisDisplayOptions"`
	// The options that determine the presentation of the category axis label.
	CategoryAxisLabelOptions interface{} `field:"optional" json:"categoryAxisLabelOptions" yaml:"categoryAxisLabelOptions"`
	// The data label configuration of a waterfall visual.
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field well configuration of a waterfall visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The legend configuration of a waterfall visual.
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The options that determine the presentation of the y-axis.
	PrimaryYAxisDisplayOptions interface{} `field:"optional" json:"primaryYAxisDisplayOptions" yaml:"primaryYAxisDisplayOptions"`
	// The options that determine the presentation of the y-axis label.
	PrimaryYAxisLabelOptions interface{} `field:"optional" json:"primaryYAxisLabelOptions" yaml:"primaryYAxisLabelOptions"`
	// The sort configuration of a waterfall visual.
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The visual palette configuration of a waterfall visual.
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
	// The options that determine the presentation of a waterfall visual.
	WaterfallChartOptions interface{} `field:"optional" json:"waterfallChartOptions" yaml:"waterfallChartOptions"`
}

