package mixinsawsquicksight


// A sheet is an object that contains a set of visuals that are viewed together on one page in a paginated report.
//
// Every analysis and dashboard must contain at least one sheet.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html
//
type CfnAnalysisPropsMixin_SheetDefinitionProperty struct {
	// The layout content type of the sheet. Choose one of the following options:.
	//
	// - `PAGINATED` : Creates a sheet for a paginated report.
	// - `INTERACTIVE` : Creates a sheet for an interactive dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// A description of the sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The list of filter controls that are on a sheet.
	//
	// For more information, see [Adding filter controls to analysis sheets](https://docs.aws.amazon.com/quicksight/latest/user/filter-controls.html) in the *Amazon Quick Suite User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-filtercontrols
	//
	FilterControls interface{} `field:"optional" json:"filterControls" yaml:"filterControls"`
	// A list of images on a sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-images
	//
	Images interface{} `field:"optional" json:"images" yaml:"images"`
	// Layouts define how the components of a sheet are arranged.
	//
	// For more information, see [Types of layout](https://docs.aws.amazon.com/quicksight/latest/user/types-of-layout.html) in the *Amazon Quick Suite User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-layouts
	//
	Layouts interface{} `field:"optional" json:"layouts" yaml:"layouts"`
	// The name of the sheet.
	//
	// This name is displayed on the sheet's tab in the Quick Suite console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The list of parameter controls that are on a sheet.
	//
	// For more information, see [Using a Control with a Parameter in Amazon Quick Sight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-controls.html) in the *Amazon Quick Suite User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-parametercontrols
	//
	ParameterControls interface{} `field:"optional" json:"parameterControls" yaml:"parameterControls"`
	// The control layouts of the sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-sheetcontrollayouts
	//
	SheetControlLayouts interface{} `field:"optional" json:"sheetControlLayouts" yaml:"sheetControlLayouts"`
	// The unique identifier of a sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-sheetid
	//
	SheetId *string `field:"optional" json:"sheetId" yaml:"sheetId"`
	// The text boxes that are on a sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-textboxes
	//
	TextBoxes interface{} `field:"optional" json:"textBoxes" yaml:"textBoxes"`
	// The title of the sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// A list of the visuals that are on a sheet.
	//
	// Visual placement is determined by the layout of the sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-visuals
	//
	Visuals interface{} `field:"optional" json:"visuals" yaml:"visuals"`
}

