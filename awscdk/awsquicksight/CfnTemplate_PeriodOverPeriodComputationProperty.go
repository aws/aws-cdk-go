package awsquicksight


// The period over period computation configuration.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-periodoverperiodcomputation.html
//
type CfnTemplate_PeriodOverPeriodComputationProperty struct {
	// The ID for a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-periodoverperiodcomputation.html#cfn-quicksight-template-periodoverperiodcomputation-computationid
	//
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The name of a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-periodoverperiodcomputation.html#cfn-quicksight-template-periodoverperiodcomputation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The time field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-periodoverperiodcomputation.html#cfn-quicksight-template-periodoverperiodcomputation-time
	//
	Time interface{} `field:"optional" json:"time" yaml:"time"`
	// The value field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-periodoverperiodcomputation.html#cfn-quicksight-template-periodoverperiodcomputation-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

