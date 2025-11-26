package previewawsquicksightmixins


// The maximum and minimum computation configuration.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-maximumminimumcomputation.html
//
type CfnAnalysisPropsMixin_MaximumMinimumComputationProperty struct {
	// The ID for a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-maximumminimumcomputation.html#cfn-quicksight-analysis-maximumminimumcomputation-computationid
	//
	ComputationId *string `field:"optional" json:"computationId" yaml:"computationId"`
	// The name of a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-maximumminimumcomputation.html#cfn-quicksight-analysis-maximumminimumcomputation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The time field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-maximumminimumcomputation.html#cfn-quicksight-analysis-maximumminimumcomputation-time
	//
	Time interface{} `field:"optional" json:"time" yaml:"time"`
	// The type of computation. Choose one of the following options:.
	//
	// - MAXIMUM: A maximum computation.
	// - MINIMUM: A minimum computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-maximumminimumcomputation.html#cfn-quicksight-analysis-maximumminimumcomputation-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The value field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-maximumminimumcomputation.html#cfn-quicksight-analysis-maximumminimumcomputation-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

