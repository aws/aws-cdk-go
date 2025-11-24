package mixinsawsquicksight


// The field well configuration of a `GaugeChartVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartfieldwells.html
//
type CfnTemplatePropsMixin_GaugeChartFieldWellsProperty struct {
	// The target value field wells of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartfieldwells.html#cfn-quicksight-template-gaugechartfieldwells-targetvalues
	//
	TargetValues interface{} `field:"optional" json:"targetValues" yaml:"targetValues"`
	// The value field wells of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartfieldwells.html#cfn-quicksight-template-gaugechartfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

