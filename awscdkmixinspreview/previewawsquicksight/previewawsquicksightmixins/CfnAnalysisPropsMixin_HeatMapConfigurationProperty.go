package previewawsquicksightmixins


// The configuration of a heat map.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapconfiguration.html
//
type CfnAnalysisPropsMixin_HeatMapConfigurationProperty struct {
	// The color options (gradient color, point of divergence) in a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapconfiguration.html#cfn-quicksight-analysis-heatmapconfiguration-colorscale
	//
	ColorScale interface{} `field:"optional" json:"colorScale" yaml:"colorScale"`
	// The label options of the column that is displayed in a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapconfiguration.html#cfn-quicksight-analysis-heatmapconfiguration-columnlabeloptions
	//
	ColumnLabelOptions interface{} `field:"optional" json:"columnLabelOptions" yaml:"columnLabelOptions"`
	// The options that determine if visual data labels are displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapconfiguration.html#cfn-quicksight-analysis-heatmapconfiguration-datalabels
	//
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field wells of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapconfiguration.html#cfn-quicksight-analysis-heatmapconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The general visual interactions setup for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapconfiguration.html#cfn-quicksight-analysis-heatmapconfiguration-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
	// The legend display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapconfiguration.html#cfn-quicksight-analysis-heatmapconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The label options of the row that is displayed in a `heat map` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapconfiguration.html#cfn-quicksight-analysis-heatmapconfiguration-rowlabeloptions
	//
	RowLabelOptions interface{} `field:"optional" json:"rowLabelOptions" yaml:"rowLabelOptions"`
	// The sort configuration of a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapconfiguration.html#cfn-quicksight-analysis-heatmapconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-heatmapconfiguration.html#cfn-quicksight-analysis-heatmapconfiguration-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
}

