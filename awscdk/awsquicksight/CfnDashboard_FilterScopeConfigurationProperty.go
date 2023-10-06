package awsquicksight


// The scope configuration for a `FilterGroup` .
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allSheets interface{}
//
//   filterScopeConfigurationProperty := &FilterScopeConfigurationProperty{
//   	AllSheets: allSheets,
//   	SelectedSheets: &SelectedSheetsFilterScopeConfigurationProperty{
//   		SheetVisualScopingConfigurations: []interface{}{
//   			&SheetVisualScopingConfigurationProperty{
//   				Scope: jsii.String("scope"),
//   				SheetId: jsii.String("sheetId"),
//
//   				// the properties below are optional
//   				VisualIds: []*string{
//   					jsii.String("visualIds"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterscopeconfiguration.html
//
type CfnDashboard_FilterScopeConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterscopeconfiguration.html#cfn-quicksight-dashboard-filterscopeconfiguration-allsheets
	//
	AllSheets interface{} `field:"optional" json:"allSheets" yaml:"allSheets"`
	// The configuration for applying a filter to specific sheets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filterscopeconfiguration.html#cfn-quicksight-dashboard-filterscopeconfiguration-selectedsheets
	//
	SelectedSheets interface{} `field:"optional" json:"selectedSheets" yaml:"selectedSheets"`
}

