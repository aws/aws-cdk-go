package awsquicksight


// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualconfiguration.html
//
type CfnAnalysis_PluginVisualConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualconfiguration.html#cfn-quicksight-analysis-pluginvisualconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualconfiguration.html#cfn-quicksight-analysis-pluginvisualconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualconfiguration.html#cfn-quicksight-analysis-pluginvisualconfiguration-visualoptions
	//
	VisualOptions interface{} `field:"optional" json:"visualOptions" yaml:"visualOptions"`
}

