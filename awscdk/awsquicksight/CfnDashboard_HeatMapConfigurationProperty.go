package awsquicksight


// The configuration of a heat map.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapconfiguration.html
//
type CfnDashboard_HeatMapConfigurationProperty struct {
	// The color options (gradient color, point of divergence) in a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapconfiguration.html#cfn-quicksight-dashboard-heatmapconfiguration-colorscale
	//
	ColorScale interface{} `field:"optional" json:"colorScale" yaml:"colorScale"`
	// The label options of the column that is displayed in a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapconfiguration.html#cfn-quicksight-dashboard-heatmapconfiguration-columnlabeloptions
	//
	ColumnLabelOptions interface{} `field:"optional" json:"columnLabelOptions" yaml:"columnLabelOptions"`
	// The options that determine if visual data labels are displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapconfiguration.html#cfn-quicksight-dashboard-heatmapconfiguration-datalabels
	//
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field wells of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapconfiguration.html#cfn-quicksight-dashboard-heatmapconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The legend display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapconfiguration.html#cfn-quicksight-dashboard-heatmapconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The label options of the row that is displayed in a `heat map` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapconfiguration.html#cfn-quicksight-dashboard-heatmapconfiguration-rowlabeloptions
	//
	RowLabelOptions interface{} `field:"optional" json:"rowLabelOptions" yaml:"rowLabelOptions"`
	// The sort configuration of a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapconfiguration.html#cfn-quicksight-dashboard-heatmapconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapconfiguration.html#cfn-quicksight-dashboard-heatmapconfiguration-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
}

