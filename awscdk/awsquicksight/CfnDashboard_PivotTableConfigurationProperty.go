package awsquicksight


// The configuration for a `PivotTableVisual` .
//
// Example:
//
//
type CfnDashboard_PivotTableConfigurationProperty struct {
	// The field options for a pivot table visual.
	FieldOptions interface{} `field:"optional" json:"fieldOptions" yaml:"fieldOptions"`
	// The field wells of the visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The paginated report options for a pivot table visual.
	PaginatedReportOptions interface{} `field:"optional" json:"paginatedReportOptions" yaml:"paginatedReportOptions"`
	// The sort configuration for a `PivotTableVisual` .
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The table options for a pivot table visual.
	TableOptions interface{} `field:"optional" json:"tableOptions" yaml:"tableOptions"`
	// The total options for a pivot table visual.
	TotalOptions interface{} `field:"optional" json:"totalOptions" yaml:"totalOptions"`
}

