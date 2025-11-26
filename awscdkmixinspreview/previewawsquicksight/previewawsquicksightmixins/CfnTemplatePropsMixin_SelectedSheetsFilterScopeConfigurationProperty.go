package previewawsquicksightmixins


// The configuration for applying a filter to specific sheets or visuals.
//
// You can apply this filter to multiple visuals that are on one sheet or to all visuals on a sheet.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   selectedSheetsFilterScopeConfigurationProperty := &SelectedSheetsFilterScopeConfigurationProperty{
//   	SheetVisualScopingConfigurations: []interface{}{
//   		&SheetVisualScopingConfigurationProperty{
//   			Scope: jsii.String("scope"),
//   			SheetId: jsii.String("sheetId"),
//   			VisualIds: []*string{
//   				jsii.String("visualIds"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-selectedsheetsfilterscopeconfiguration.html
//
type CfnTemplatePropsMixin_SelectedSheetsFilterScopeConfigurationProperty struct {
	// The sheet ID and visual IDs of the sheet and visuals that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-selectedsheetsfilterscopeconfiguration.html#cfn-quicksight-template-selectedsheetsfilterscopeconfiguration-sheetvisualscopingconfigurations
	//
	SheetVisualScopingConfigurations interface{} `field:"optional" json:"sheetVisualScopingConfigurations" yaml:"sheetVisualScopingConfigurations"`
}

