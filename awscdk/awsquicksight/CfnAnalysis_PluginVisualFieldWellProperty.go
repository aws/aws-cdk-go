package awsquicksight


// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualfieldwell.html
//
type CfnAnalysis_PluginVisualFieldWellProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualfieldwell.html#cfn-quicksight-analysis-pluginvisualfieldwell-axisname
	//
	AxisName *string `field:"optional" json:"axisName" yaml:"axisName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualfieldwell.html#cfn-quicksight-analysis-pluginvisualfieldwell-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualfieldwell.html#cfn-quicksight-analysis-pluginvisualfieldwell-measures
	//
	Measures interface{} `field:"optional" json:"measures" yaml:"measures"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualfieldwell.html#cfn-quicksight-analysis-pluginvisualfieldwell-unaggregated
	//
	Unaggregated interface{} `field:"optional" json:"unaggregated" yaml:"unaggregated"`
}

