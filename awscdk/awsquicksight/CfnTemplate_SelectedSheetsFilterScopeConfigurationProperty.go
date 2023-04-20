package awsquicksight


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
type CfnTemplate_SelectedSheetsFilterScopeConfigurationProperty struct {
	// `CfnTemplate.SelectedSheetsFilterScopeConfigurationProperty.SheetVisualScopingConfigurations`.
	SheetVisualScopingConfigurations interface{} `field:"optional" json:"sheetVisualScopingConfigurations" yaml:"sheetVisualScopingConfigurations"`
}

