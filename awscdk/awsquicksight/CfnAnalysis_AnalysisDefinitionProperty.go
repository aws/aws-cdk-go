package awsquicksight


// The definition of an analysis.
//
// Example:
//
//
type CfnAnalysis_AnalysisDefinitionProperty struct {
	// An array of dataset identifier declarations.
	//
	// This mapping allows the usage of dataset identifiers instead of dataset ARNs throughout analysis sub-structures.
	DataSetIdentifierDeclarations interface{} `field:"required" json:"dataSetIdentifierDeclarations" yaml:"dataSetIdentifierDeclarations"`
	// `CfnAnalysis.AnalysisDefinitionProperty.AnalysisDefaults`.
	AnalysisDefaults interface{} `field:"optional" json:"analysisDefaults" yaml:"analysisDefaults"`
	// An array of calculated field definitions for the analysis.
	CalculatedFields interface{} `field:"optional" json:"calculatedFields" yaml:"calculatedFields"`
	// An array of analysis-level column configurations.
	//
	// Column configurations can be used to set default formatting for a column to be used throughout an analysis.
	ColumnConfigurations interface{} `field:"optional" json:"columnConfigurations" yaml:"columnConfigurations"`
	// Filter definitions for an analysis.
	//
	// For more information, see [Filtering Data in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/adding-a-filter.html) in the *Amazon QuickSight User Guide* .
	FilterGroups interface{} `field:"optional" json:"filterGroups" yaml:"filterGroups"`
	// An array of parameter declarations for an analysis.
	//
	// Parameters are named variables that can transfer a value for use by an action or an object.
	//
	// For more information, see [Parameters in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-in-quicksight.html) in the *Amazon QuickSight User Guide* .
	ParameterDeclarations interface{} `field:"optional" json:"parameterDeclarations" yaml:"parameterDeclarations"`
	// An array of sheet definitions for an analysis.
	//
	// Each `SheetDefinition` provides detailed information about a sheet within this analysis.
	Sheets interface{} `field:"optional" json:"sheets" yaml:"sheets"`
}

