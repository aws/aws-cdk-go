package awsquicksight


// The growth rate computation configuration.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-growthratecomputation.html
//
type CfnTemplate_GrowthRateComputationProperty struct {
	// The ID for a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-growthratecomputation.html#cfn-quicksight-template-growthratecomputation-computationid
	//
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The name of a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-growthratecomputation.html#cfn-quicksight-template-growthratecomputation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The period size setup of a growth rate computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-growthratecomputation.html#cfn-quicksight-template-growthratecomputation-periodsize
	//
	// Default: - 0.
	//
	PeriodSize *float64 `field:"optional" json:"periodSize" yaml:"periodSize"`
	// The time field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-growthratecomputation.html#cfn-quicksight-template-growthratecomputation-time
	//
	Time interface{} `field:"optional" json:"time" yaml:"time"`
	// The value field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-growthratecomputation.html#cfn-quicksight-template-growthratecomputation-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

