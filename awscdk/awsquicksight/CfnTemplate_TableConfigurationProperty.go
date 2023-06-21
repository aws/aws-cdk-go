package awsquicksight


// The configuration for a `TableVisual` .
//
// Example:
//
//
type CfnTemplate_TableConfigurationProperty struct {
	// The field options for a table visual.
	FieldOptions interface{} `field:"optional" json:"fieldOptions" yaml:"fieldOptions"`
	// The field wells of the visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The paginated report options for a table visual.
	PaginatedReportOptions interface{} `field:"optional" json:"paginatedReportOptions" yaml:"paginatedReportOptions"`
	// The sort configuration for a `TableVisual` .
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// A collection of inline visualizations to display within a chart.
	TableInlineVisualizations interface{} `field:"optional" json:"tableInlineVisualizations" yaml:"tableInlineVisualizations"`
	// The table options for a table visual.
	TableOptions interface{} `field:"optional" json:"tableOptions" yaml:"tableOptions"`
	// The total options for a table visual.
	TotalOptions interface{} `field:"optional" json:"totalOptions" yaml:"totalOptions"`
}

