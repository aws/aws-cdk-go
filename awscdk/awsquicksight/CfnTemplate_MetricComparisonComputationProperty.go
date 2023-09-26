package awsquicksight


// The metric comparison computation configuration.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-metriccomparisoncomputation.html
//
type CfnTemplate_MetricComparisonComputationProperty struct {
	// The ID for a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-metriccomparisoncomputation.html#cfn-quicksight-template-metriccomparisoncomputation-computationid
	//
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The field that is used in a metric comparison from value setup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-metriccomparisoncomputation.html#cfn-quicksight-template-metriccomparisoncomputation-fromvalue
	//
	FromValue interface{} `field:"optional" json:"fromValue" yaml:"fromValue"`
	// The name of a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-metriccomparisoncomputation.html#cfn-quicksight-template-metriccomparisoncomputation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The field that is used in a metric comparison to value setup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-metriccomparisoncomputation.html#cfn-quicksight-template-metriccomparisoncomputation-targetvalue
	//
	TargetValue interface{} `field:"optional" json:"targetValue" yaml:"targetValue"`
	// The time field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-metriccomparisoncomputation.html#cfn-quicksight-template-metriccomparisoncomputation-time
	//
	Time interface{} `field:"optional" json:"time" yaml:"time"`
}

