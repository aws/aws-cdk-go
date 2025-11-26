package previewawsquicksightmixins


// The period over period computation configuration.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodoverperiodcomputation.html
//
type CfnAnalysisPropsMixin_PeriodOverPeriodComputationProperty struct {
	// The ID for a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodoverperiodcomputation.html#cfn-quicksight-analysis-periodoverperiodcomputation-computationid
	//
	ComputationId *string `field:"optional" json:"computationId" yaml:"computationId"`
	// The name of a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodoverperiodcomputation.html#cfn-quicksight-analysis-periodoverperiodcomputation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The time field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodoverperiodcomputation.html#cfn-quicksight-analysis-periodoverperiodcomputation-time
	//
	Time interface{} `field:"optional" json:"time" yaml:"time"`
	// The value field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-periodoverperiodcomputation.html#cfn-quicksight-analysis-periodoverperiodcomputation-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

