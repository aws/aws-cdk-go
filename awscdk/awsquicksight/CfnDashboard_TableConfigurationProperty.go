package awsquicksight


// The configuration for a `TableVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableconfiguration.html
//
type CfnDashboard_TableConfigurationProperty struct {
	// The field options for a table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableconfiguration.html#cfn-quicksight-dashboard-tableconfiguration-fieldoptions
	//
	FieldOptions interface{} `field:"optional" json:"fieldOptions" yaml:"fieldOptions"`
	// The field wells of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableconfiguration.html#cfn-quicksight-dashboard-tableconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The paginated report options for a table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableconfiguration.html#cfn-quicksight-dashboard-tableconfiguration-paginatedreportoptions
	//
	PaginatedReportOptions interface{} `field:"optional" json:"paginatedReportOptions" yaml:"paginatedReportOptions"`
	// The sort configuration for a `TableVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableconfiguration.html#cfn-quicksight-dashboard-tableconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// A collection of inline visualizations to display within a chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableconfiguration.html#cfn-quicksight-dashboard-tableconfiguration-tableinlinevisualizations
	//
	TableInlineVisualizations interface{} `field:"optional" json:"tableInlineVisualizations" yaml:"tableInlineVisualizations"`
	// The table options for a table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableconfiguration.html#cfn-quicksight-dashboard-tableconfiguration-tableoptions
	//
	TableOptions interface{} `field:"optional" json:"tableOptions" yaml:"tableOptions"`
	// The total options for a table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableconfiguration.html#cfn-quicksight-dashboard-tableconfiguration-totaloptions
	//
	TotalOptions interface{} `field:"optional" json:"totalOptions" yaml:"totalOptions"`
}

