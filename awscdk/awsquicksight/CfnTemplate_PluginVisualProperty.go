package awsquicksight


// A flexible visualization type that allows engineers to create new custom charts in Amazon QuickSight.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisual.html
//
type CfnTemplate_PluginVisualProperty struct {
	// The Amazon Resource Name (ARN) that reflects the plugin and version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisual.html#cfn-quicksight-template-pluginvisual-pluginarn
	//
	PluginArn *string `field:"required" json:"pluginArn" yaml:"pluginArn"`
	// The ID of the visual that you want to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisual.html#cfn-quicksight-template-pluginvisual-visualid
	//
	VisualId *string `field:"required" json:"visualId" yaml:"visualId"`
	// A description of the plugin field wells and their persisted properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisual.html#cfn-quicksight-template-pluginvisual-chartconfiguration
	//
	ChartConfiguration interface{} `field:"optional" json:"chartConfiguration" yaml:"chartConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisual.html#cfn-quicksight-template-pluginvisual-subtitle
	//
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisual.html#cfn-quicksight-template-pluginvisual-title
	//
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// The alt text for the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisual.html#cfn-quicksight-template-pluginvisual-visualcontentalttext
	//
	VisualContentAltText *string `field:"optional" json:"visualContentAltText" yaml:"visualContentAltText"`
}

