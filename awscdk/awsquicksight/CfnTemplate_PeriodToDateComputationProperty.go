package awsquicksight


// The period to date computation configuration.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-periodtodatecomputation.html
//
type CfnTemplate_PeriodToDateComputationProperty struct {
	// The ID for a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-periodtodatecomputation.html#cfn-quicksight-template-periodtodatecomputation-computationid
	//
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The name of a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-periodtodatecomputation.html#cfn-quicksight-template-periodtodatecomputation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The time granularity setup of period to date computation. Choose from the following options:.
	//
	// - YEAR: Year to date.
	// - MONTH: Month to date.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-periodtodatecomputation.html#cfn-quicksight-template-periodtodatecomputation-periodtimegranularity
	//
	PeriodTimeGranularity *string `field:"optional" json:"periodTimeGranularity" yaml:"periodTimeGranularity"`
	// The time field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-periodtodatecomputation.html#cfn-quicksight-template-periodtodatecomputation-time
	//
	Time interface{} `field:"optional" json:"time" yaml:"time"`
	// The value field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-periodtodatecomputation.html#cfn-quicksight-template-periodtodatecomputation-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

