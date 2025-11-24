package mixinsawsquicksight


// The color field that defines a gradient or categorical style.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayercolorfield.html
//
type CfnDashboardPropsMixin_GeospatialLayerColorFieldProperty struct {
	// A list of color dimension fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayercolorfield.html#cfn-quicksight-dashboard-geospatiallayercolorfield-colordimensionsfields
	//
	ColorDimensionsFields interface{} `field:"optional" json:"colorDimensionsFields" yaml:"colorDimensionsFields"`
	// A list of color measure fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayercolorfield.html#cfn-quicksight-dashboard-geospatiallayercolorfield-colorvaluesfields
	//
	ColorValuesFields interface{} `field:"optional" json:"colorValuesFields" yaml:"colorValuesFields"`
}

