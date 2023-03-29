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
type CfnTemplate_SheetVisualScopingConfigurationProperty struct {
	// `CfnTemplate.SheetVisualScopingConfigurationProperty.Scope`.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// `CfnTemplate.SheetVisualScopingConfigurationProperty.SheetId`.
	SheetId *string `field:"required" json:"sheetId" yaml:"sheetId"`
	// `CfnTemplate.SheetVisualScopingConfigurationProperty.VisualIds`.
	VisualIds *[]*string `field:"optional" json:"visualIds" yaml:"visualIds"`
}

