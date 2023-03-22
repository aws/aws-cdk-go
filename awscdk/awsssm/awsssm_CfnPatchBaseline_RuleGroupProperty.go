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
//   ruleGroupProperty := &ruleGroupProperty{
//   	patchRules: []interface{}{
//   		&ruleProperty{
//   			approveAfterDays: jsii.Number(123),
//   			approveUntilDate: jsii.String("approveUntilDate"),
//   			complianceLevel: jsii.String("complianceLevel"),
//   			enableNonSecurity: jsii.Boolean(false),
//   			patchFilterGroup: &patchFilterGroupProperty{
//   				patchFilters: []interface{}{
//   					&patchFilterProperty{
//   						key: jsii.String("key"),
//   						values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnPatchBaseline_RuleGroupProperty struct {
	// The rules that make up the rule group.
	PatchRules interface{} `field:"optional" json:"patchRules" yaml:"patchRules"`
}

