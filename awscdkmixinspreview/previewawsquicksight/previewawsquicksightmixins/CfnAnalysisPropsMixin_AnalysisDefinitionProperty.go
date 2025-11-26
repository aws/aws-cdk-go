package previewawsquicksightmixins


// The definition of an analysis.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysisdefinition.html
//
type CfnAnalysisPropsMixin_AnalysisDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysisdefinition.html#cfn-quicksight-analysis-analysisdefinition-analysisdefaults
	//
	AnalysisDefaults interface{} `field:"optional" json:"analysisDefaults" yaml:"analysisDefaults"`
	// An array of calculated field definitions for the analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysisdefinition.html#cfn-quicksight-analysis-analysisdefinition-calculatedfields
	//
	CalculatedFields interface{} `field:"optional" json:"calculatedFields" yaml:"calculatedFields"`
	// An array of analysis-level column configurations.
	//
	// Column configurations can be used to set default formatting for a column to be used throughout an analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysisdefinition.html#cfn-quicksight-analysis-analysisdefinition-columnconfigurations
	//
	ColumnConfigurations interface{} `field:"optional" json:"columnConfigurations" yaml:"columnConfigurations"`
	// An array of dataset identifier declarations.
	//
	// This mapping allows the usage of dataset identifiers instead of dataset ARNs throughout analysis sub-structures.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysisdefinition.html#cfn-quicksight-analysis-analysisdefinition-datasetidentifierdeclarations
	//
	DataSetIdentifierDeclarations interface{} `field:"optional" json:"dataSetIdentifierDeclarations" yaml:"dataSetIdentifierDeclarations"`
	// Filter definitions for an analysis.
	//
	// For more information, see [Filtering Data in Amazon Quick Sight](https://docs.aws.amazon.com/quicksight/latest/user/adding-a-filter.html) in the *Amazon Quick Suite User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysisdefinition.html#cfn-quicksight-analysis-analysisdefinition-filtergroups
	//
	FilterGroups interface{} `field:"optional" json:"filterGroups" yaml:"filterGroups"`
	// An array of option definitions for an analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysisdefinition.html#cfn-quicksight-analysis-analysisdefinition-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// An array of parameter declarations for an analysis.
	//
	// Parameters are named variables that can transfer a value for use by an action or an object.
	//
	// For more information, see [Parameters in Amazon Quick Sight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-in-quicksight.html) in the *Amazon Quick Suite User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysisdefinition.html#cfn-quicksight-analysis-analysisdefinition-parameterdeclarations
	//
	ParameterDeclarations interface{} `field:"optional" json:"parameterDeclarations" yaml:"parameterDeclarations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysisdefinition.html#cfn-quicksight-analysis-analysisdefinition-queryexecutionoptions
	//
	QueryExecutionOptions interface{} `field:"optional" json:"queryExecutionOptions" yaml:"queryExecutionOptions"`
	// An array of sheet definitions for an analysis.
	//
	// Each `SheetDefinition` provides detailed information about a sheet within this analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysisdefinition.html#cfn-quicksight-analysis-analysisdefinition-sheets
	//
	Sheets interface{} `field:"optional" json:"sheets" yaml:"sheets"`
	// The static files for the definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-analysisdefinition.html#cfn-quicksight-analysis-analysisdefinition-staticfiles
	//
	StaticFiles interface{} `field:"optional" json:"staticFiles" yaml:"staticFiles"`
}

