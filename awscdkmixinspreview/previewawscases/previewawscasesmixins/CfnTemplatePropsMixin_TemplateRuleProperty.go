package previewawscasesmixins


// An association representing a case rule acting upon a field.
//
// In the Amazon Connect admin website, case rules are known as *case field conditions* . For more information about case field conditions, see [Add case field conditions to a case template](https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   templateRuleProperty := &TemplateRuleProperty{
//   	CaseRuleId: jsii.String("caseRuleId"),
//   	FieldId: jsii.String("fieldId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-template-templaterule.html
//
type CfnTemplatePropsMixin_TemplateRuleProperty struct {
	// Unique identifier of a case rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-template-templaterule.html#cfn-cases-template-templaterule-caseruleid
	//
	CaseRuleId *string `field:"optional" json:"caseRuleId" yaml:"caseRuleId"`
	// Unique identifier of a field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-template-templaterule.html#cfn-cases-template-templaterule-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
}

