package awssecurityhub


// One or more actions that AWS Security Hub takes when a finding matches the defined criteria of a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var id interface{}
//   var updatedBy interface{}
//
//   automationRulesActionProperty := &AutomationRulesActionProperty{
//   	FindingFieldsUpdate: &AutomationRulesFindingFieldsUpdateProperty{
//   		Confidence: jsii.Number(123),
//   		Criticality: jsii.Number(123),
//   		Note: &NoteUpdateProperty{
//   			Text: jsii.String("text"),
//   			UpdatedBy: updatedBy,
//   		},
//   		RelatedFindings: []interface{}{
//   			&RelatedFindingProperty{
//   				Id: id,
//   				ProductArn: jsii.String("productArn"),
//   			},
//   		},
//   		Severity: &SeverityUpdateProperty{
//   			Label: jsii.String("label"),
//   			Normalized: jsii.Number(123),
//   			Product: jsii.Number(123),
//   		},
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   		UserDefinedFields: map[string]*string{
//   			"userDefinedFieldsKey": jsii.String("userDefinedFields"),
//   		},
//   		VerificationState: jsii.String("verificationState"),
//   		Workflow: &WorkflowUpdateProperty{
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesaction.html
//
type CfnAutomationRule_AutomationRulesActionProperty struct {
	// Specifies that the automation rule action is an update to a finding field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesaction.html#cfn-securityhub-automationrule-automationrulesaction-findingfieldsupdate
	//
	FindingFieldsUpdate interface{} `field:"required" json:"findingFieldsUpdate" yaml:"findingFieldsUpdate"`
	// Specifies the type of action that Security Hub takes when a finding matches the defined criteria of a rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesaction.html#cfn-securityhub-automationrule-automationrulesaction-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

