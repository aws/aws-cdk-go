package awsquicksight


// The detailed definition of a template.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversiondefinition.html
//
type CfnTemplate_TemplateVersionDefinitionProperty struct {
	// An array of dataset configurations.
	//
	// These configurations define the required columns for each dataset used within a template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversiondefinition.html#cfn-quicksight-template-templateversiondefinition-datasetconfigurations
	//
	DataSetConfigurations interface{} `field:"required" json:"dataSetConfigurations" yaml:"dataSetConfigurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversiondefinition.html#cfn-quicksight-template-templateversiondefinition-analysisdefaults
	//
	AnalysisDefaults interface{} `field:"optional" json:"analysisDefaults" yaml:"analysisDefaults"`
	// An array of calculated field definitions for the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversiondefinition.html#cfn-quicksight-template-templateversiondefinition-calculatedfields
	//
	CalculatedFields interface{} `field:"optional" json:"calculatedFields" yaml:"calculatedFields"`
	// An array of template-level column configurations.
	//
	// Column configurations are used to set default formatting for a column that's used throughout a template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversiondefinition.html#cfn-quicksight-template-templateversiondefinition-columnconfigurations
	//
	ColumnConfigurations interface{} `field:"optional" json:"columnConfigurations" yaml:"columnConfigurations"`
	// Filter definitions for a template.
	//
	// For more information, see [Filtering Data](https://docs.aws.amazon.com/quicksight/latest/user/filtering-visual-data.html) in the *Amazon Quick Suite User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversiondefinition.html#cfn-quicksight-template-templateversiondefinition-filtergroups
	//
	FilterGroups interface{} `field:"optional" json:"filterGroups" yaml:"filterGroups"`
	// An array of option definitions for a template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversiondefinition.html#cfn-quicksight-template-templateversiondefinition-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// An array of parameter declarations for a template.
	//
	// *Parameters* are named variables that can transfer a value for use by an action or an object.
	//
	// For more information, see [Parameters in Amazon Quick Sight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-in-quicksight.html) in the *Amazon Quick Suite User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversiondefinition.html#cfn-quicksight-template-templateversiondefinition-parameterdeclarations
	//
	ParameterDeclarations interface{} `field:"optional" json:"parameterDeclarations" yaml:"parameterDeclarations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversiondefinition.html#cfn-quicksight-template-templateversiondefinition-queryexecutionoptions
	//
	QueryExecutionOptions interface{} `field:"optional" json:"queryExecutionOptions" yaml:"queryExecutionOptions"`
	// An array of sheet definitions for a template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversiondefinition.html#cfn-quicksight-template-templateversiondefinition-sheets
	//
	Sheets interface{} `field:"optional" json:"sheets" yaml:"sheets"`
}

