package previewawsquicksightmixins


// A collection of field wells for a plugin visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualfieldwell.html
//
type CfnAnalysisPropsMixin_PluginVisualFieldWellProperty struct {
	// The semantic axis name for the field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualfieldwell.html#cfn-quicksight-analysis-pluginvisualfieldwell-axisname
	//
	AxisName *string `field:"optional" json:"axisName" yaml:"axisName"`
	// A list of dimensions for the field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualfieldwell.html#cfn-quicksight-analysis-pluginvisualfieldwell-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// A list of measures that exist in the field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualfieldwell.html#cfn-quicksight-analysis-pluginvisualfieldwell-measures
	//
	Measures interface{} `field:"optional" json:"measures" yaml:"measures"`
	// A list of unaggregated fields that exist in the field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pluginvisualfieldwell.html#cfn-quicksight-analysis-pluginvisualfieldwell-unaggregated
	//
	Unaggregated interface{} `field:"optional" json:"unaggregated" yaml:"unaggregated"`
}

