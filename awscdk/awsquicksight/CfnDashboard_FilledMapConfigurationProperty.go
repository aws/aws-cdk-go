package awsquicksight


// The configuration for a `FilledMapVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapconfiguration.html
//
type CfnDashboard_FilledMapConfigurationProperty struct {
	// The field wells of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapconfiguration.html#cfn-quicksight-dashboard-filledmapconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The legend display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapconfiguration.html#cfn-quicksight-dashboard-filledmapconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The map style options of the filled map visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapconfiguration.html#cfn-quicksight-dashboard-filledmapconfiguration-mapstyleoptions
	//
	MapStyleOptions interface{} `field:"optional" json:"mapStyleOptions" yaml:"mapStyleOptions"`
	// The sort configuration of a `FilledMapVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapconfiguration.html#cfn-quicksight-dashboard-filledmapconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapconfiguration.html#cfn-quicksight-dashboard-filledmapconfiguration-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// The window options of the filled map visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapconfiguration.html#cfn-quicksight-dashboard-filledmapconfiguration-windowoptions
	//
	WindowOptions interface{} `field:"optional" json:"windowOptions" yaml:"windowOptions"`
}

