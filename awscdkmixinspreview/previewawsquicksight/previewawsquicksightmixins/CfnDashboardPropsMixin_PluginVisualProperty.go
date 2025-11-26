package previewawsquicksightmixins


// A flexible visualization type that allows engineers to create new custom charts in Quick Sight.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisual.html
//
type CfnDashboardPropsMixin_PluginVisualProperty struct {
	// A description of the plugin field wells and their persisted properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisual.html#cfn-quicksight-dashboard-pluginvisual-chartconfiguration
	//
	ChartConfiguration interface{} `field:"optional" json:"chartConfiguration" yaml:"chartConfiguration"`
	// The Amazon Resource Name (ARN) that reflects the plugin and version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisual.html#cfn-quicksight-dashboard-pluginvisual-pluginarn
	//
	PluginArn *string `field:"optional" json:"pluginArn" yaml:"pluginArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisual.html#cfn-quicksight-dashboard-pluginvisual-subtitle
	//
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisual.html#cfn-quicksight-dashboard-pluginvisual-title
	//
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// The alt text for the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisual.html#cfn-quicksight-dashboard-pluginvisual-visualcontentalttext
	//
	VisualContentAltText *string `field:"optional" json:"visualContentAltText" yaml:"visualContentAltText"`
	// The ID of the visual that you want to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisual.html#cfn-quicksight-dashboard-pluginvisual-visualid
	//
	VisualId *string `field:"optional" json:"visualId" yaml:"visualId"`
}

