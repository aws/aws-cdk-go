package awsquicksight


// The rendering rules of a sheet that uses a free-form layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetElementRenderingRuleProperty := &SheetElementRenderingRuleProperty{
//   	ConfigurationOverrides: &SheetElementConfigurationOverridesProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	Expression: jsii.String("expression"),
//   }
//
type CfnAnalysis_SheetElementRenderingRuleProperty struct {
	// The override configuration of the rendering rules of a sheet.
	ConfigurationOverrides interface{} `field:"required" json:"configurationOverrides" yaml:"configurationOverrides"`
	// The expression of the rendering rules of a sheet.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

