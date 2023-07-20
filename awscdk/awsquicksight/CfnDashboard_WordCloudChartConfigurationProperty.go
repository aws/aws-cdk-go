package awsquicksight


// The configuration of a word cloud visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-wordcloudchartconfiguration.html
//
type CfnDashboard_WordCloudChartConfigurationProperty struct {
	// The label options (label text, label visibility, and sort icon visibility) for the word cloud category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-wordcloudchartconfiguration.html#cfn-quicksight-dashboard-wordcloudchartconfiguration-categorylabeloptions
	//
	CategoryLabelOptions interface{} `field:"optional" json:"categoryLabelOptions" yaml:"categoryLabelOptions"`
	// The field wells of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-wordcloudchartconfiguration.html#cfn-quicksight-dashboard-wordcloudchartconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The sort configuration of a word cloud visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-wordcloudchartconfiguration.html#cfn-quicksight-dashboard-wordcloudchartconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The options for a word cloud visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-wordcloudchartconfiguration.html#cfn-quicksight-dashboard-wordcloudchartconfiguration-wordcloudoptions
	//
	WordCloudOptions interface{} `field:"optional" json:"wordCloudOptions" yaml:"wordCloudOptions"`
}

