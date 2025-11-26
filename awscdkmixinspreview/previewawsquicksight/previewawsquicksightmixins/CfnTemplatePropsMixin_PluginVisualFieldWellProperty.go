package previewawsquicksightmixins


// A collection of field wells for a plugin visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualfieldwell.html
//
type CfnTemplatePropsMixin_PluginVisualFieldWellProperty struct {
	// The semantic axis name for the field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualfieldwell.html#cfn-quicksight-template-pluginvisualfieldwell-axisname
	//
	AxisName *string `field:"optional" json:"axisName" yaml:"axisName"`
	// A list of dimensions for the field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualfieldwell.html#cfn-quicksight-template-pluginvisualfieldwell-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// A list of measures that exist in the field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualfieldwell.html#cfn-quicksight-template-pluginvisualfieldwell-measures
	//
	Measures interface{} `field:"optional" json:"measures" yaml:"measures"`
	// A list of unaggregated fields that exist in the field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualfieldwell.html#cfn-quicksight-template-pluginvisualfieldwell-unaggregated
	//
	Unaggregated interface{} `field:"optional" json:"unaggregated" yaml:"unaggregated"`
}

