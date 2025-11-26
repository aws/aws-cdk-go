package previewawsquicksightmixins


// The plugin visual configuration.
//
// This includes the field wells, sorting options, and persisted options of the plugin visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualconfiguration.html
//
type CfnDashboardPropsMixin_PluginVisualConfigurationProperty struct {
	// The field wells configuration of the plugin visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualconfiguration.html#cfn-quicksight-dashboard-pluginvisualconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The sort configuration of the plugin visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualconfiguration.html#cfn-quicksight-dashboard-pluginvisualconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The persisted properties of the plugin visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualconfiguration.html#cfn-quicksight-dashboard-pluginvisualconfiguration-visualoptions
	//
	VisualOptions interface{} `field:"optional" json:"visualOptions" yaml:"visualOptions"`
}

