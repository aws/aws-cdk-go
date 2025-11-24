package mixinsawsquicksight


// The field well configuration of a `GaugeChartVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartfieldwells.html
//
type CfnAnalysisPropsMixin_GaugeChartFieldWellsProperty struct {
	// The target value field wells of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartfieldwells.html#cfn-quicksight-analysis-gaugechartfieldwells-targetvalues
	//
	TargetValues interface{} `field:"optional" json:"targetValues" yaml:"targetValues"`
	// The value field wells of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartfieldwells.html#cfn-quicksight-analysis-gaugechartfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

