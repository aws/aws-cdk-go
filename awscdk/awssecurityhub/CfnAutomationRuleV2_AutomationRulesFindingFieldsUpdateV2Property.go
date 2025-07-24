package awssecurityhub


// Allows you to define the structure for modifying specific fields in security findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   automationRulesFindingFieldsUpdateV2Property := &AutomationRulesFindingFieldsUpdateV2Property{
//   	Comment: jsii.String("comment"),
//   	SeverityId: jsii.Number(123),
//   	StatusId: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-automationrulesfindingfieldsupdatev2.html
//
type CfnAutomationRuleV2_AutomationRulesFindingFieldsUpdateV2Property struct {
	// Notes or contextual information for findings that are modified by the automation rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-automationrulesfindingfieldsupdatev2.html#cfn-securityhub-automationrulev2-automationrulesfindingfieldsupdatev2-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The severity level to be assigned to findings that match the automation rule criteria.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-automationrulesfindingfieldsupdatev2.html#cfn-securityhub-automationrulev2-automationrulesfindingfieldsupdatev2-severityid
	//
	SeverityId *float64 `field:"optional" json:"severityId" yaml:"severityId"`
	// The status to be applied to findings that match automation rule criteria.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrulev2-automationrulesfindingfieldsupdatev2.html#cfn-securityhub-automationrulev2-automationrulesfindingfieldsupdatev2-statusid
	//
	StatusId *float64 `field:"optional" json:"statusId" yaml:"statusId"`
}

