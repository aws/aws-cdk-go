package awsquicksight


// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualfieldwell.html
//
type CfnDashboard_PluginVisualFieldWellProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualfieldwell.html#cfn-quicksight-dashboard-pluginvisualfieldwell-axisname
	//
	AxisName *string `field:"optional" json:"axisName" yaml:"axisName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualfieldwell.html#cfn-quicksight-dashboard-pluginvisualfieldwell-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualfieldwell.html#cfn-quicksight-dashboard-pluginvisualfieldwell-measures
	//
	Measures interface{} `field:"optional" json:"measures" yaml:"measures"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualfieldwell.html#cfn-quicksight-dashboard-pluginvisualfieldwell-unaggregated
	//
	Unaggregated interface{} `field:"optional" json:"unaggregated" yaml:"unaggregated"`
}

