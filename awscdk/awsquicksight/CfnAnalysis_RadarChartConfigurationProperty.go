package awsquicksight


// The configuration of a `RadarChartVisual` .
//
// Example:
//
//
type CfnAnalysis_RadarChartConfigurationProperty struct {
	// Determines the visibility of the colors of alternatign bands in a radar chart.
	AlternateBandColorsVisibility *string `field:"optional" json:"alternateBandColorsVisibility" yaml:"alternateBandColorsVisibility"`
	// The color of the even-numbered alternate bands of a radar chart.
	AlternateBandEvenColor *string `field:"optional" json:"alternateBandEvenColor" yaml:"alternateBandEvenColor"`
	// The color of the odd-numbered alternate bands of a radar chart.
	AlternateBandOddColor *string `field:"optional" json:"alternateBandOddColor" yaml:"alternateBandOddColor"`
	// The axis behavior options of a radar chart.
	AxesRangeScale *string `field:"optional" json:"axesRangeScale" yaml:"axesRangeScale"`
	// The base sreies settings of a radar chart.
	BaseSeriesSettings interface{} `field:"optional" json:"baseSeriesSettings" yaml:"baseSeriesSettings"`
	// The category axis of a radar chart.
	CategoryAxis interface{} `field:"optional" json:"categoryAxis" yaml:"categoryAxis"`
	// The category label options of a radar chart.
	CategoryLabelOptions interface{} `field:"optional" json:"categoryLabelOptions" yaml:"categoryLabelOptions"`
	// The color axis of a radar chart.
	ColorAxis interface{} `field:"optional" json:"colorAxis" yaml:"colorAxis"`
	// The color label options of a radar chart.
	ColorLabelOptions interface{} `field:"optional" json:"colorLabelOptions" yaml:"colorLabelOptions"`
	// The field well configuration of a `RadarChartVisual` .
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The legend display setup of the visual.
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The shape of the radar chart.
	Shape *string `field:"optional" json:"shape" yaml:"shape"`
	// The sort configuration of a `RadarChartVisual` .
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The start angle of a radar chart's axis.
	StartAngle *float64 `field:"optional" json:"startAngle" yaml:"startAngle"`
	// The palette (chart color) display setup of the visual.
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
}

