package awsquicksight


// The maximum and minimum computation configuration.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-maximumminimumcomputation.html
//
type CfnTemplate_MaximumMinimumComputationProperty struct {
	// The ID for a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-maximumminimumcomputation.html#cfn-quicksight-template-maximumminimumcomputation-computationid
	//
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The type of computation. Choose one of the following options:.
	//
	// - MAXIMUM: A maximum computation.
	// - MINIMUM: A minimum computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-maximumminimumcomputation.html#cfn-quicksight-template-maximumminimumcomputation-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-maximumminimumcomputation.html#cfn-quicksight-template-maximumminimumcomputation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The time field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-maximumminimumcomputation.html#cfn-quicksight-template-maximumminimumcomputation-time
	//
	Time interface{} `field:"optional" json:"time" yaml:"time"`
	// The value field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-maximumminimumcomputation.html#cfn-quicksight-template-maximumminimumcomputation-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

