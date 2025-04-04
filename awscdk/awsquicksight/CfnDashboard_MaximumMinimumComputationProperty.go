package awsquicksight


// The maximum and minimum computation configuration.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-maximumminimumcomputation.html
//
type CfnDashboard_MaximumMinimumComputationProperty struct {
	// The ID for a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-maximumminimumcomputation.html#cfn-quicksight-dashboard-maximumminimumcomputation-computationid
	//
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The type of computation. Choose one of the following options:.
	//
	// - MAXIMUM: A maximum computation.
	// - MINIMUM: A minimum computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-maximumminimumcomputation.html#cfn-quicksight-dashboard-maximumminimumcomputation-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-maximumminimumcomputation.html#cfn-quicksight-dashboard-maximumminimumcomputation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The time field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-maximumminimumcomputation.html#cfn-quicksight-dashboard-maximumminimumcomputation-time
	//
	Time interface{} `field:"optional" json:"time" yaml:"time"`
	// The value field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-maximumminimumcomputation.html#cfn-quicksight-dashboard-maximumminimumcomputation-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

