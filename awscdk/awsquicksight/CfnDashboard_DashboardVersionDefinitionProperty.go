package awsquicksight


// The contents of a dashboard.
//
// Example:
//
//
type CfnDashboard_DashboardVersionDefinitionProperty struct {
	// An array of dataset identifier declarations.
	//
	// With this mapping,you can use dataset identifiers instead of dataset Amazon Resource Names (ARNs) throughout the dashboard's sub-structures.
	DataSetIdentifierDeclarations interface{} `field:"required" json:"dataSetIdentifierDeclarations" yaml:"dataSetIdentifierDeclarations"`
	// `CfnDashboard.DashboardVersionDefinitionProperty.AnalysisDefaults`.
	AnalysisDefaults interface{} `field:"optional" json:"analysisDefaults" yaml:"analysisDefaults"`
	// An array of calculated field definitions for the dashboard.
	CalculatedFields interface{} `field:"optional" json:"calculatedFields" yaml:"calculatedFields"`
	// An array of dashboard-level column configurations.
	//
	// Column configurations are used to set the default formatting for a column that is used throughout a dashboard.
	ColumnConfigurations interface{} `field:"optional" json:"columnConfigurations" yaml:"columnConfigurations"`
	// The filter definitions for a dashboard.
	//
	// For more information, see [Filtering Data in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/adding-a-filter.html) in the *Amazon QuickSight User Guide* .
	FilterGroups interface{} `field:"optional" json:"filterGroups" yaml:"filterGroups"`
	// The parameter declarations for a dashboard.
	//
	// Parameters are named variables that can transfer a value for use by an action or an object.
	//
	// For more information, see [Parameters in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-in-quicksight.html) in the *Amazon QuickSight User Guide* .
	ParameterDeclarations interface{} `field:"optional" json:"parameterDeclarations" yaml:"parameterDeclarations"`
	// An array of sheet definitions for a dashboard.
	Sheets interface{} `field:"optional" json:"sheets" yaml:"sheets"`
}

