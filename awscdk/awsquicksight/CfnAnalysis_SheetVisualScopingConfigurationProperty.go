package awsquicksight


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
type CfnAnalysis_SheetVisualScopingConfigurationProperty struct {
	// `CfnAnalysis.SheetVisualScopingConfigurationProperty.Scope`.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// `CfnAnalysis.SheetVisualScopingConfigurationProperty.SheetId`.
	SheetId *string `field:"required" json:"sheetId" yaml:"sheetId"`
	// `CfnAnalysis.SheetVisualScopingConfigurationProperty.VisualIds`.
	VisualIds *[]*string `field:"optional" json:"visualIds" yaml:"visualIds"`
}

