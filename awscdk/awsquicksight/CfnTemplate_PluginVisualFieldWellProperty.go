package awsquicksight


// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualfieldwell.html
//
type CfnTemplate_PluginVisualFieldWellProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualfieldwell.html#cfn-quicksight-template-pluginvisualfieldwell-axisname
	//
	AxisName *string `field:"optional" json:"axisName" yaml:"axisName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualfieldwell.html#cfn-quicksight-template-pluginvisualfieldwell-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualfieldwell.html#cfn-quicksight-template-pluginvisualfieldwell-measures
	//
	Measures interface{} `field:"optional" json:"measures" yaml:"measures"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualfieldwell.html#cfn-quicksight-template-pluginvisualfieldwell-unaggregated
	//
	Unaggregated interface{} `field:"optional" json:"unaggregated" yaml:"unaggregated"`
}

