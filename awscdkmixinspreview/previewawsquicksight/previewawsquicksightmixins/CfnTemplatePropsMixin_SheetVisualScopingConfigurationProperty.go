package previewawsquicksightmixins


// The filter that is applied to the options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sheetVisualScopingConfigurationProperty := &SheetVisualScopingConfigurationProperty{
//   	Scope: jsii.String("scope"),
//   	SheetId: jsii.String("sheetId"),
//   	VisualIds: []*string{
//   		jsii.String("visualIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetvisualscopingconfiguration.html
//
type CfnTemplatePropsMixin_SheetVisualScopingConfigurationProperty struct {
	// The scope of the applied entities. Choose one of the following options:.
	//
	// - `ALL_VISUALS`
	// - `SELECTED_VISUALS`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetvisualscopingconfiguration.html#cfn-quicksight-template-sheetvisualscopingconfiguration-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The selected sheet that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetvisualscopingconfiguration.html#cfn-quicksight-template-sheetvisualscopingconfiguration-sheetid
	//
	SheetId *string `field:"optional" json:"sheetId" yaml:"sheetId"`
	// The selected visuals that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetvisualscopingconfiguration.html#cfn-quicksight-template-sheetvisualscopingconfiguration-visualids
	//
	VisualIds *[]*string `field:"optional" json:"visualIds" yaml:"visualIds"`
}

