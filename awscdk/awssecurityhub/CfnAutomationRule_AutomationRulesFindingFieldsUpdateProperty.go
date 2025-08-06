package awssecurityhub


// Identifies the finding fields that the automation rule action updates when a finding matches the defined criteria.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   automationRulesFindingFieldsUpdateProperty := &AutomationRulesFindingFieldsUpdateProperty{
//   	Confidence: jsii.Number(123),
//   	Criticality: jsii.Number(123),
//   	Note: &NoteUpdateProperty{
//   		Text: jsii.String("text"),
//   		UpdatedBy: jsii.String("updatedBy"),
//   	},
//   	RelatedFindings: []interface{}{
//   		&RelatedFindingProperty{
//   			Id: jsii.String("id"),
//   			ProductArn: jsii.String("productArn"),
//   		},
//   	},
//   	Severity: &SeverityUpdateProperty{
//   		Label: jsii.String("label"),
//   		Normalized: jsii.Number(123),
//   		Product: jsii.Number(123),
//   	},
//   	Types: []*string{
//   		jsii.String("types"),
//   	},
//   	UserDefinedFields: map[string]*string{
//   		"userDefinedFieldsKey": jsii.String("userDefinedFields"),
//   	},
//   	VerificationState: jsii.String("verificationState"),
//   	Workflow: &WorkflowUpdateProperty{
//   		Status: jsii.String("status"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfieldsupdate.html
//
type CfnAutomationRule_AutomationRulesFindingFieldsUpdateProperty struct {
	// The rule action updates the `Confidence` field of a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfieldsupdate.html#cfn-securityhub-automationrule-automationrulesfindingfieldsupdate-confidence
	//
	Confidence *float64 `field:"optional" json:"confidence" yaml:"confidence"`
	// The rule action updates the `Criticality` field of a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfieldsupdate.html#cfn-securityhub-automationrule-automationrulesfindingfieldsupdate-criticality
	//
	Criticality *float64 `field:"optional" json:"criticality" yaml:"criticality"`
	// The rule action will update the `Note` field of a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfieldsupdate.html#cfn-securityhub-automationrule-automationrulesfindingfieldsupdate-note
	//
	Note interface{} `field:"optional" json:"note" yaml:"note"`
	// The rule action will update the `RelatedFindings` field of a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfieldsupdate.html#cfn-securityhub-automationrule-automationrulesfindingfieldsupdate-relatedfindings
	//
	RelatedFindings interface{} `field:"optional" json:"relatedFindings" yaml:"relatedFindings"`
	// The rule action will update the `Severity` field of a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfieldsupdate.html#cfn-securityhub-automationrule-automationrulesfindingfieldsupdate-severity
	//
	Severity interface{} `field:"optional" json:"severity" yaml:"severity"`
	// The rule action updates the `Types` field of a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfieldsupdate.html#cfn-securityhub-automationrule-automationrulesfindingfieldsupdate-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
	// The rule action updates the `UserDefinedFields` field of a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfieldsupdate.html#cfn-securityhub-automationrule-automationrulesfindingfieldsupdate-userdefinedfields
	//
	UserDefinedFields interface{} `field:"optional" json:"userDefinedFields" yaml:"userDefinedFields"`
	// The rule action updates the `VerificationState` field of a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfieldsupdate.html#cfn-securityhub-automationrule-automationrulesfindingfieldsupdate-verificationstate
	//
	VerificationState *string `field:"optional" json:"verificationState" yaml:"verificationState"`
	// The rule action will update the `Workflow` field of a finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-automationrulesfindingfieldsupdate.html#cfn-securityhub-automationrule-automationrulesfindingfieldsupdate-workflow
	//
	Workflow interface{} `field:"optional" json:"workflow" yaml:"workflow"`
}

