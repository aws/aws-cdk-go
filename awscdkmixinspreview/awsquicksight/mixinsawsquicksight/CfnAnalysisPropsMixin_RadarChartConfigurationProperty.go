package mixinsawsquicksight


// The configuration of a `RadarChartVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html
//
type CfnAnalysisPropsMixin_RadarChartConfigurationProperty struct {
	// Determines the visibility of the colors of alternatign bands in a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-alternatebandcolorsvisibility
	//
	AlternateBandColorsVisibility *string `field:"optional" json:"alternateBandColorsVisibility" yaml:"alternateBandColorsVisibility"`
	// The color of the even-numbered alternate bands of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-alternatebandevencolor
	//
	AlternateBandEvenColor *string `field:"optional" json:"alternateBandEvenColor" yaml:"alternateBandEvenColor"`
	// The color of the odd-numbered alternate bands of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-alternatebandoddcolor
	//
	AlternateBandOddColor *string `field:"optional" json:"alternateBandOddColor" yaml:"alternateBandOddColor"`
	// The axis behavior options of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-axesrangescale
	//
	AxesRangeScale *string `field:"optional" json:"axesRangeScale" yaml:"axesRangeScale"`
	// The base sreies settings of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-baseseriessettings
	//
	BaseSeriesSettings interface{} `field:"optional" json:"baseSeriesSettings" yaml:"baseSeriesSettings"`
	// The category axis of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-categoryaxis
	//
	CategoryAxis interface{} `field:"optional" json:"categoryAxis" yaml:"categoryAxis"`
	// The category label options of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-categorylabeloptions
	//
	CategoryLabelOptions interface{} `field:"optional" json:"categoryLabelOptions" yaml:"categoryLabelOptions"`
	// The color axis of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-coloraxis
	//
	ColorAxis interface{} `field:"optional" json:"colorAxis" yaml:"colorAxis"`
	// The color label options of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-colorlabeloptions
	//
	ColorLabelOptions interface{} `field:"optional" json:"colorLabelOptions" yaml:"colorLabelOptions"`
	// The field well configuration of a `RadarChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The general visual interactions setup for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
	// The legend display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The shape of the radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-shape
	//
	Shape *string `field:"optional" json:"shape" yaml:"shape"`
	// The sort configuration of a `RadarChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The start angle of a radar chart's axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-startangle
	//
	StartAngle *float64 `field:"optional" json:"startAngle" yaml:"startAngle"`
	// The palette (chart color) display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-radarchartconfiguration.html#cfn-quicksight-analysis-radarchartconfiguration-visualpalette
	//
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
}

