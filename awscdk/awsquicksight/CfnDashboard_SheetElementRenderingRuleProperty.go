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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetelementrenderingrule.html
//
type CfnDashboard_SheetElementRenderingRuleProperty struct {
	// The override configuration of the rendering rules of a sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetelementrenderingrule.html#cfn-quicksight-dashboard-sheetelementrenderingrule-configurationoverrides
	//
	ConfigurationOverrides interface{} `field:"required" json:"configurationOverrides" yaml:"configurationOverrides"`
	// The expression of the rendering rules of a sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sheetelementrenderingrule.html#cfn-quicksight-dashboard-sheetelementrenderingrule-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

