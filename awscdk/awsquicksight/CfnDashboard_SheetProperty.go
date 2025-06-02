package awsquicksight


// A *sheet* , which is an object that contains a set of visuals that are viewed together on one page in Amazon QuickSight.
//
// Every analysis and dashboard contains at least one sheet. Each sheet contains at least one visualization widget, for example a chart, pivot table, or narrative insight. Sheets can be associated with other components, such as controls, filters, and so on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetProperty := &SheetProperty{
//   	Name: jsii.String("name"),
//   	SheetId: jsii.String("sheetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheet.html
//
type CfnDashboard_SheetProperty struct {
	// The name of a sheet.
	//
	// This name is displayed on the sheet's tab in the Amazon QuickSight console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheet.html#cfn-quicksight-dashboard-sheet-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The unique identifier associated with a sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheet.html#cfn-quicksight-dashboard-sheet-sheetid
	//
	SheetId *string `field:"optional" json:"sheetId" yaml:"sheetId"`
}

