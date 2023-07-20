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
//   filterScopeConfigurationProperty := &FilterScopeConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filterscopeconfiguration.html
//
type CfnAnalysis_FilterScopeConfigurationProperty struct {
	// The configuration for applying a filter to specific sheets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filterscopeconfiguration.html#cfn-quicksight-analysis-filterscopeconfiguration-selectedsheets
	//
	SelectedSheets interface{} `field:"optional" json:"selectedSheets" yaml:"selectedSheets"`
}

