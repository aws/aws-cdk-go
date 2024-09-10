package awsssm


// The `Rule` property type specifies an approval rule for a Systems Manager patch baseline.
//
// The `PatchRules` property of the [RuleGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rulegroup.html) property type contains a list of `Rule` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &RuleProperty{
//   	ApproveAfterDays: jsii.Number(123),
//   	ApproveUntilDate: jsii.String("approveUntilDate"),
//   	ComplianceLevel: jsii.String("complianceLevel"),
//   	EnableNonSecurity: jsii.Boolean(false),
//   	PatchFilterGroup: &PatchFilterGroupProperty{
//   		PatchFilters: []interface{}{
//   			&PatchFilterProperty{
//   				Key: jsii.String("key"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html
//
type CfnPatchBaseline_RuleProperty struct {
	// The number of days after the release date of each patch matched by the rule that the patch is marked as approved in the patch baseline.
	//
	// For example, a value of `7` means that patches are approved seven days after they are released.
	//
	// This parameter is marked as `Required: No` , but your request must include a value for either `ApproveAfterDays` or `ApproveUntilDate` .
	//
	// Not supported for Debian Server or Ubuntu Server.
	//
	// > Use caution when setting this value for Windows Server patch baselines. Because patch updates that are replaced by later updates are removed, setting too broad a value for this parameter can result in crucial patches not being installed. For more information, see the *Windows Server* tab in the topic [How security patches are selected](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-selecting-patches.html) in the *AWS Systems Manager User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html#cfn-ssm-patchbaseline-rule-approveafterdays
	//
	ApproveAfterDays *float64 `field:"optional" json:"approveAfterDays" yaml:"approveAfterDays"`
	// The cutoff date for auto approval of released patches.
	//
	// Any patches released on or before this date are installed automatically.
	//
	// Enter dates in the format `YYYY-MM-DD` . For example, `2024-12-31` .
	//
	// This parameter is marked as `Required: No` , but your request must include a value for either `ApproveUntilDate` or `ApproveAfterDays` .
	//
	// Not supported for Debian Server or Ubuntu Server.
	//
	// > Use caution when setting this value for Windows Server patch baselines. Because patch updates that are replaced by later updates are removed, setting too broad a value for this parameter can result in crucial patches not being installed. For more information, see the *Windows Server* tab in the topic [How security patches are selected](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-selecting-patches.html) in the *AWS Systems Manager User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html#cfn-ssm-patchbaseline-rule-approveuntildate
	//
	ApproveUntilDate *string `field:"optional" json:"approveUntilDate" yaml:"approveUntilDate"`
	// A compliance severity level for all approved patches in a patch baseline.
	//
	// Valid compliance severity levels include the following: `UNSPECIFIED` , `CRITICAL` , `HIGH` , `MEDIUM` , `LOW` , and `INFORMATIONAL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html#cfn-ssm-patchbaseline-rule-compliancelevel
	//
	ComplianceLevel *string `field:"optional" json:"complianceLevel" yaml:"complianceLevel"`
	// For managed nodes identified by the approval rule filters, enables a patch baseline to apply non-security updates available in the specified repository.
	//
	// The default value is `false` . Applies to Linux managed nodes only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html#cfn-ssm-patchbaseline-rule-enablenonsecurity
	//
	// Default: - false.
	//
	EnableNonSecurity interface{} `field:"optional" json:"enableNonSecurity" yaml:"enableNonSecurity"`
	// The patch filter group that defines the criteria for the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-rule.html#cfn-ssm-patchbaseline-rule-patchfiltergroup
	//
	PatchFilterGroup interface{} `field:"optional" json:"patchFilterGroup" yaml:"patchFilterGroup"`
}

