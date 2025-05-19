package awsquicksight


// The configuration for a `FilledMapVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html
//
type CfnAnalysis_FilledMapConfigurationProperty struct {
	// The field wells of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html#cfn-quicksight-analysis-filledmapconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The general visual interactions setup for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html#cfn-quicksight-analysis-filledmapconfiguration-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
	// The legend display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html#cfn-quicksight-analysis-filledmapconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The map style options of the filled map visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html#cfn-quicksight-analysis-filledmapconfiguration-mapstyleoptions
	//
	MapStyleOptions interface{} `field:"optional" json:"mapStyleOptions" yaml:"mapStyleOptions"`
	// The sort configuration of a `FilledMapVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html#cfn-quicksight-analysis-filledmapconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html#cfn-quicksight-analysis-filledmapconfiguration-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// The window options of the filled map visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconfiguration.html#cfn-quicksight-analysis-filledmapconfiguration-windowoptions
	//
	WindowOptions interface{} `field:"optional" json:"windowOptions" yaml:"windowOptions"`
}

