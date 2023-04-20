package awsquicksight


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
type CfnTemplate_SheetElementRenderingRuleProperty struct {
	// `CfnTemplate.SheetElementRenderingRuleProperty.ConfigurationOverrides`.
	ConfigurationOverrides interface{} `field:"required" json:"configurationOverrides" yaml:"configurationOverrides"`
	// `CfnTemplate.SheetElementRenderingRuleProperty.Expression`.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

