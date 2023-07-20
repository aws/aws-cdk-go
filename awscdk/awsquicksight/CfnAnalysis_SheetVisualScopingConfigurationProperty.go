package awsquicksight


// The filter that is applied to the options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetVisualScopingConfigurationProperty := &SheetVisualScopingConfigurationProperty{
//   	Scope: jsii.String("scope"),
//   	SheetId: jsii.String("sheetId"),
//
//   	// the properties below are optional
//   	VisualIds: []*string{
//   		jsii.String("visualIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetvisualscopingconfiguration.html
//
type CfnAnalysis_SheetVisualScopingConfigurationProperty struct {
	// The scope of the applied entities. Choose one of the following options:.
	//
	// - `ALL_VISUALS`
	// - `SELECTED_VISUALS`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetvisualscopingconfiguration.html#cfn-quicksight-analysis-sheetvisualscopingconfiguration-scope
	//
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// The selected sheet that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetvisualscopingconfiguration.html#cfn-quicksight-analysis-sheetvisualscopingconfiguration-sheetid
	//
	SheetId *string `field:"required" json:"sheetId" yaml:"sheetId"`
	// The selected visuals that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetvisualscopingconfiguration.html#cfn-quicksight-analysis-sheetvisualscopingconfiguration-visualids
	//
	VisualIds *[]*string `field:"optional" json:"visualIds" yaml:"visualIds"`
}

