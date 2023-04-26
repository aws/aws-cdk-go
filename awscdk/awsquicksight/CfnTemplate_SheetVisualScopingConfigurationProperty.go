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
type CfnTemplate_SheetVisualScopingConfigurationProperty struct {
	// The scope of the applied entities. Choose one of the following options:.
	//
	// - `ALL_VISUALS`
	// - `SELECTED_VISUALS`.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// The selected sheet that the filter is applied to.
	SheetId *string `field:"required" json:"sheetId" yaml:"sheetId"`
	// The selected visuals that the filter is applied to.
	VisualIds *[]*string `field:"optional" json:"visualIds" yaml:"visualIds"`
}

