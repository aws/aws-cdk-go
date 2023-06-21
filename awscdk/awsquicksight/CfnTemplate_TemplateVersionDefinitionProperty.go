package awsquicksight


// The detailed definition of a template.
//
// Example:
//
//
type CfnTemplate_TemplateVersionDefinitionProperty struct {
	// An array of dataset configurations.
	//
	// These configurations define the required columns for each dataset used within a template.
	DataSetConfigurations interface{} `field:"required" json:"dataSetConfigurations" yaml:"dataSetConfigurations"`
	// `CfnTemplate.TemplateVersionDefinitionProperty.AnalysisDefaults`.
	AnalysisDefaults interface{} `field:"optional" json:"analysisDefaults" yaml:"analysisDefaults"`
	// An array of calculated field definitions for the template.
	CalculatedFields interface{} `field:"optional" json:"calculatedFields" yaml:"calculatedFields"`
	// An array of template-level column configurations.
	//
	// Column configurations are used to set default formatting for a column that's used throughout a template.
	ColumnConfigurations interface{} `field:"optional" json:"columnConfigurations" yaml:"columnConfigurations"`
	// Filter definitions for a template.
	//
	// For more information, see [Filtering Data](https://docs.aws.amazon.com/quicksight/latest/user/filtering-visual-data.html) in the *Amazon QuickSight User Guide* .
	FilterGroups interface{} `field:"optional" json:"filterGroups" yaml:"filterGroups"`
	// An array of parameter declarations for a template.
	//
	// *Parameters* are named variables that can transfer a value for use by an action or an object.
	//
	// For more information, see [Parameters in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-in-quicksight.html) in the *Amazon QuickSight User Guide* .
	ParameterDeclarations interface{} `field:"optional" json:"parameterDeclarations" yaml:"parameterDeclarations"`
	// An array of sheet definitions for a template.
	Sheets interface{} `field:"optional" json:"sheets" yaml:"sheets"`
}

