package awsquicksight


// The configuration of an insight visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-insightconfiguration.html
//
type CfnDashboard_InsightConfigurationProperty struct {
	// The computations configurations of the insight visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-insightconfiguration.html#cfn-quicksight-dashboard-insightconfiguration-computations
	//
	Computations interface{} `field:"optional" json:"computations" yaml:"computations"`
	// The custom narrative of the insight visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-insightconfiguration.html#cfn-quicksight-dashboard-insightconfiguration-customnarrative
	//
	CustomNarrative interface{} `field:"optional" json:"customNarrative" yaml:"customNarrative"`
}

