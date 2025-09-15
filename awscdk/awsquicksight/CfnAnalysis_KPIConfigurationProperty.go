package awsquicksight


// The configuration of a KPI visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpiconfiguration.html
//
type CfnAnalysis_KPIConfigurationProperty struct {
	// The field well configuration of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpiconfiguration.html#cfn-quicksight-analysis-kpiconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The general visual interactions setup for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpiconfiguration.html#cfn-quicksight-analysis-kpiconfiguration-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
	// The options that determine the presentation of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpiconfiguration.html#cfn-quicksight-analysis-kpiconfiguration-kpioptions
	//
	KpiOptions interface{} `field:"optional" json:"kpiOptions" yaml:"kpiOptions"`
	// The sort configuration of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpiconfiguration.html#cfn-quicksight-analysis-kpiconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
}

