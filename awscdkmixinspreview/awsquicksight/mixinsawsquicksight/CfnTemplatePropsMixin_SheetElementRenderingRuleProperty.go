package mixinsawsquicksight


// The rendering rules of a sheet that uses a free-form layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sheetElementRenderingRuleProperty := &SheetElementRenderingRuleProperty{
//   	ConfigurationOverrides: &SheetElementConfigurationOverridesProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	Expression: jsii.String("expression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetelementrenderingrule.html
//
type CfnTemplatePropsMixin_SheetElementRenderingRuleProperty struct {
	// The override configuration of the rendering rules of a sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetelementrenderingrule.html#cfn-quicksight-template-sheetelementrenderingrule-configurationoverrides
	//
	ConfigurationOverrides interface{} `field:"optional" json:"configurationOverrides" yaml:"configurationOverrides"`
	// The expression of the rendering rules of a sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetelementrenderingrule.html#cfn-quicksight-template-sheetelementrenderingrule-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

