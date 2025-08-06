package awsssm


// The `RuleGroup` property type specifies a set of rules that define the approval rules for an AWS Systems Manager patch baseline.
//
// `RuleGroup` is the property type for the `ApprovalRules` property of the [AWS::SSM::PatchBaseline](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleGroupProperty := &RuleGroupProperty{
//   	PatchRules: []interface{}{
//   		&RuleProperty{
//   			ApproveAfterDays: jsii.Number(123),
//   			ApproveUntilDate: jsii.String("approveUntilDate"),
//   			ComplianceLevel: jsii.String("complianceLevel"),
//   			EnableNonSecurity: jsii.Boolean(false),
//   			PatchFilterGroup: &PatchFilterGroupProperty{
//   				PatchFilters: []interface{}{
//   					&PatchFilterProperty{
//   						Key: jsii.String("key"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rulegroup.html
//
type CfnPatchBaseline_RuleGroupProperty struct {
	// The rules that make up the rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rulegroup.html#cfn-ssm-patchbaseline-rulegroup-patchrules
	//
	PatchRules interface{} `field:"optional" json:"patchRules" yaml:"patchRules"`
}

