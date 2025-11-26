package previewawsquicksightmixins


// The configuration of an insight visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-insightconfiguration.html
//
type CfnAnalysisPropsMixin_InsightConfigurationProperty struct {
	// The computations configurations of the insight visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-insightconfiguration.html#cfn-quicksight-analysis-insightconfiguration-computations
	//
	Computations interface{} `field:"optional" json:"computations" yaml:"computations"`
	// The custom narrative of the insight visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-insightconfiguration.html#cfn-quicksight-analysis-insightconfiguration-customnarrative
	//
	CustomNarrative interface{} `field:"optional" json:"customNarrative" yaml:"customNarrative"`
	// The general visual interactions setup for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-insightconfiguration.html#cfn-quicksight-analysis-insightconfiguration-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
}

