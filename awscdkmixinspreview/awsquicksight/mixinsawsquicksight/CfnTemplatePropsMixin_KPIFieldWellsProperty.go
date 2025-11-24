package mixinsawsquicksight


// The field well configuration of a KPI visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpifieldwells.html
//
type CfnTemplatePropsMixin_KPIFieldWellsProperty struct {
	// The target value field wells of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpifieldwells.html#cfn-quicksight-template-kpifieldwells-targetvalues
	//
	TargetValues interface{} `field:"optional" json:"targetValues" yaml:"targetValues"`
	// The trend group field wells of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpifieldwells.html#cfn-quicksight-template-kpifieldwells-trendgroups
	//
	TrendGroups interface{} `field:"optional" json:"trendGroups" yaml:"trendGroups"`
	// The value field wells of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpifieldwells.html#cfn-quicksight-template-kpifieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

