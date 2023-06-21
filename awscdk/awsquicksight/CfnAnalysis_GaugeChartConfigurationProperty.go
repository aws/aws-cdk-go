package awsquicksight


// The configuration of a `GaugeChartVisual` .
//
// Example:
//
//
type CfnAnalysis_GaugeChartConfigurationProperty struct {
	// The data label configuration of a `GaugeChartVisual` .
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field well configuration of a `GaugeChartVisual` .
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The options that determine the presentation of the `GaugeChartVisual` .
	GaugeChartOptions interface{} `field:"optional" json:"gaugeChartOptions" yaml:"gaugeChartOptions"`
	// The tooltip configuration of a `GaugeChartVisual` .
	TooltipOptions interface{} `field:"optional" json:"tooltipOptions" yaml:"tooltipOptions"`
	// The visual palette configuration of a `GaugeChartVisual` .
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
}

