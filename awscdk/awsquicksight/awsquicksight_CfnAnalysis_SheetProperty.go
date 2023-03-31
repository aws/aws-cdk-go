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
//   sheetProperty := &sheetProperty{
//   	name: jsii.String("name"),
//   	sheetId: jsii.String("sheetId"),
//   }
//
type CfnAnalysis_SheetProperty struct {
	// The name of a sheet.
	//
	// This name is displayed on the sheet's tab in the Amazon QuickSight console.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The unique identifier associated with a sheet.
	SheetId *string `field:"optional" json:"sheetId" yaml:"sheetId"`
}

