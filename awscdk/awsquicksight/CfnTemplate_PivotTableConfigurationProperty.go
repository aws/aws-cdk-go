package awsquicksight


// The configuration for a `PivotTableVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableconfiguration.html
//
type CfnTemplate_PivotTableConfigurationProperty struct {
	// The field options for a pivot table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableconfiguration.html#cfn-quicksight-template-pivottableconfiguration-fieldoptions
	//
	FieldOptions interface{} `field:"optional" json:"fieldOptions" yaml:"fieldOptions"`
	// The field wells of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableconfiguration.html#cfn-quicksight-template-pivottableconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The paginated report options for a pivot table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableconfiguration.html#cfn-quicksight-template-pivottableconfiguration-paginatedreportoptions
	//
	PaginatedReportOptions interface{} `field:"optional" json:"paginatedReportOptions" yaml:"paginatedReportOptions"`
	// The sort configuration for a `PivotTableVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableconfiguration.html#cfn-quicksight-template-pivottableconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The table options for a pivot table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableconfiguration.html#cfn-quicksight-template-pivottableconfiguration-tableoptions
	//
	TableOptions interface{} `field:"optional" json:"tableOptions" yaml:"tableOptions"`
	// The total options for a pivot table visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableconfiguration.html#cfn-quicksight-template-pivottableconfiguration-totaloptions
	//
	TotalOptions interface{} `field:"optional" json:"totalOptions" yaml:"totalOptions"`
}

