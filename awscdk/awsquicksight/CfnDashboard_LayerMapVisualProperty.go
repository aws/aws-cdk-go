package awsquicksight


// A layer map visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layermapvisual.html
//
type CfnDashboard_LayerMapVisualProperty struct {
	// The ID of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layermapvisual.html#cfn-quicksight-dashboard-layermapvisual-visualid
	//
	VisualId *string `field:"required" json:"visualId" yaml:"visualId"`
	// The configuration settings of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layermapvisual.html#cfn-quicksight-dashboard-layermapvisual-chartconfiguration
	//
	ChartConfiguration interface{} `field:"optional" json:"chartConfiguration" yaml:"chartConfiguration"`
	// The dataset that is used to create the layer map visual.
	//
	// You can't create a visual without a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layermapvisual.html#cfn-quicksight-dashboard-layermapvisual-datasetidentifier
	//
	// Default: - "".
	//
	DataSetIdentifier *string `field:"optional" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layermapvisual.html#cfn-quicksight-dashboard-layermapvisual-subtitle
	//
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layermapvisual.html#cfn-quicksight-dashboard-layermapvisual-title
	//
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layermapvisual.html#cfn-quicksight-dashboard-layermapvisual-topicidentifier
	//
	TopicIdentifier *string `field:"optional" json:"topicIdentifier" yaml:"topicIdentifier"`
	// The alt text for the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layermapvisual.html#cfn-quicksight-dashboard-layermapvisual-visualcontentalttext
	//
	VisualContentAltText *string `field:"optional" json:"visualContentAltText" yaml:"visualContentAltText"`
}

