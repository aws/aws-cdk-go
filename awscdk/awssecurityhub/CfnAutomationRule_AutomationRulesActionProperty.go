package awssecurityhub


// One or more actions to update finding fields if a finding matches the defined criteria of the rule.
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
type CfnAutomationRule_AutomationRulesActionProperty struct {
	// Specifies that the automation rule action is an update to a finding field.
	FindingFieldsUpdate interface{} `field:"required" json:"findingFieldsUpdate" yaml:"findingFieldsUpdate"`
	// Specifies that the rule action should update the `Types` finding field.
	//
	// The `Types` finding field provides one or more finding types in the format of namespace/category/classifier that classify a finding. For more information, see [Types taxonomy for ASFF](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-findings-format-type-taxonomy.html) in the *AWS Security Hub User Guide* .
	Type *string `field:"required" json:"type" yaml:"type"`
}

