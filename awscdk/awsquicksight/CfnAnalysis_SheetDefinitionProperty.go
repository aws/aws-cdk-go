package awsquicksight


// A sheet is an object that contains a set of visuals that are viewed together on one page in a paginated report.
//
// Every analysis and dashboard must contain at least one sheet.
//
// Example:
//
//
type CfnAnalysis_SheetDefinitionProperty struct {
	// The unique identifier of a sheet.
	SheetId *string `field:"required" json:"sheetId" yaml:"sheetId"`
	// The layout content type of the sheet. Choose one of the following options:.
	//
	// - `PAGINATED` : Creates a sheet for a paginated report.
	// - `INTERACTIVE` : Creates a sheet for an interactive dashboard.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// A description of the sheet.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The list of filter controls that are on a sheet.
	//
	// For more information, see [Adding filter controls to analysis sheets](https://docs.aws.amazon.com/quicksight/latest/user/filter-controls.html) in the *Amazon QuickSight User Guide* .
	FilterControls interface{} `field:"optional" json:"filterControls" yaml:"filterControls"`
	// Layouts define how the components of a sheet are arranged.
	//
	// For more information, see [Types of layout](https://docs.aws.amazon.com/quicksight/latest/user/types-of-layout.html) in the *Amazon QuickSight User Guide* .
	Layouts interface{} `field:"optional" json:"layouts" yaml:"layouts"`
	// The name of the sheet.
	//
	// This name is displayed on the sheet's tab in the Amazon QuickSight console.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The list of parameter controls that are on a sheet.
	//
	// For more information, see [Using a Control with a Parameter in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-controls.html) in the *Amazon QuickSight User Guide* .
	ParameterControls interface{} `field:"optional" json:"parameterControls" yaml:"parameterControls"`
	// The control layouts of the sheet.
	SheetControlLayouts interface{} `field:"optional" json:"sheetControlLayouts" yaml:"sheetControlLayouts"`
	// The text boxes that are on a sheet.
	TextBoxes interface{} `field:"optional" json:"textBoxes" yaml:"textBoxes"`
	// The title of the sheet.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// A list of the visuals that are on a sheet.
	//
	// Visual placement is determined by the layout of the sheet.
	Visuals interface{} `field:"optional" json:"visuals" yaml:"visuals"`
}

