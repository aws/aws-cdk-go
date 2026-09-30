package awsquicksight


// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layermapvisual.html
//
type CfnTemplate_LayerMapVisualProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layermapvisual.html#cfn-quicksight-template-layermapvisual-visualid
	//
	VisualId *string `field:"required" json:"visualId" yaml:"visualId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layermapvisual.html#cfn-quicksight-template-layermapvisual-chartconfiguration
	//
	ChartConfiguration interface{} `field:"optional" json:"chartConfiguration" yaml:"chartConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layermapvisual.html#cfn-quicksight-template-layermapvisual-datasetidentifier
	//
	// Default: - "".
	//
	DataSetIdentifier *string `field:"optional" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layermapvisual.html#cfn-quicksight-template-layermapvisual-subtitle
	//
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layermapvisual.html#cfn-quicksight-template-layermapvisual-title
	//
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layermapvisual.html#cfn-quicksight-template-layermapvisual-topicidentifier
	//
	TopicIdentifier *string `field:"optional" json:"topicIdentifier" yaml:"topicIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layermapvisual.html#cfn-quicksight-template-layermapvisual-visualcontentalttext
	//
	VisualContentAltText *string `field:"optional" json:"visualContentAltText" yaml:"visualContentAltText"`
}

