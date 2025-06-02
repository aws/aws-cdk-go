package awsquicksight


// The configuration for a `PivotTableVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html
//
type CfnDashboard_PivotTableConfigurationProperty struct {
	// The field options for a pivot table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html#cfn-quicksight-dashboard-pivottableconfiguration-fieldoptions
	//
	FieldOptions interface{} `field:"optional" json:"fieldOptions" yaml:"fieldOptions"`
	// The field wells of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html#cfn-quicksight-dashboard-pivottableconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The general visual interactions setup for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html#cfn-quicksight-dashboard-pivottableconfiguration-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
	// The paginated report options for a pivot table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html#cfn-quicksight-dashboard-pivottableconfiguration-paginatedreportoptions
	//
	PaginatedReportOptions interface{} `field:"optional" json:"paginatedReportOptions" yaml:"paginatedReportOptions"`
	// The sort configuration for a `PivotTableVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html#cfn-quicksight-dashboard-pivottableconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The table options for a pivot table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html#cfn-quicksight-dashboard-pivottableconfiguration-tableoptions
	//
	TableOptions interface{} `field:"optional" json:"tableOptions" yaml:"tableOptions"`
	// The total options for a pivot table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconfiguration.html#cfn-quicksight-dashboard-pivottableconfiguration-totaloptions
	//
	TotalOptions interface{} `field:"optional" json:"totalOptions" yaml:"totalOptions"`
}

