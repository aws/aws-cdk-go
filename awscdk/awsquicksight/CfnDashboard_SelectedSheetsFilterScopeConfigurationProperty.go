package awsquicksight


// The configuration for applying a filter to specific sheets or visuals.
//
// You can apply this filter to multiple visuals that are on one sheet or to all visuals on a sheet.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selectedSheetsFilterScopeConfigurationProperty := &SelectedSheetsFilterScopeConfigurationProperty{
//   	SheetVisualScopingConfigurations: []interface{}{
//   		&SheetVisualScopingConfigurationProperty{
//   			Scope: jsii.String("scope"),
//   			SheetId: jsii.String("sheetId"),
//
//   			// the properties below are optional
//   			VisualIds: []*string{
//   				jsii.String("visualIds"),
//   			},
//   		},
//   	},
//   }
//
type CfnDashboard_SelectedSheetsFilterScopeConfigurationProperty struct {
	// The sheet ID and visual IDs of the sheet and visuals that the filter is applied to.
	SheetVisualScopingConfigurations interface{} `field:"optional" json:"sheetVisualScopingConfigurations" yaml:"sheetVisualScopingConfigurations"`
}

