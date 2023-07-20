package awsquicksight


// The configuration of an insight visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-insightconfiguration.html
//
type CfnAnalysis_InsightConfigurationProperty struct {
	// The computations configurations of the insight visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-insightconfiguration.html#cfn-quicksight-analysis-insightconfiguration-computations
	//
	Computations interface{} `field:"optional" json:"computations" yaml:"computations"`
	// The custom narrative of the insight visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-insightconfiguration.html#cfn-quicksight-analysis-insightconfiguration-customnarrative
	//
	CustomNarrative interface{} `field:"optional" json:"customNarrative" yaml:"customNarrative"`
}

