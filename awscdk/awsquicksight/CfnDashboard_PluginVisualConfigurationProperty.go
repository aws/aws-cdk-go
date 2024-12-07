package awsquicksight


// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualconfiguration.html
//
type CfnDashboard_PluginVisualConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualconfiguration.html#cfn-quicksight-dashboard-pluginvisualconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualconfiguration.html#cfn-quicksight-dashboard-pluginvisualconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualconfiguration.html#cfn-quicksight-dashboard-pluginvisualconfiguration-visualoptions
	//
	VisualOptions interface{} `field:"optional" json:"visualOptions" yaml:"visualOptions"`
}

