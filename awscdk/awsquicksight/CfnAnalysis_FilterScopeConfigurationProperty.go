package awsquicksight


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
type CfnAnalysis_FilterScopeConfigurationProperty struct {
	// `CfnAnalysis.FilterScopeConfigurationProperty.SelectedSheets`.
	SelectedSheets interface{} `field:"optional" json:"selectedSheets" yaml:"selectedSheets"`
}

