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
//   ruleProperty := &ruleProperty{
//   	approveAfterDays: jsii.Number(123),
//   	approveUntilDate: jsii.String("approveUntilDate"),
//   	complianceLevel: jsii.String("complianceLevel"),
//   	enableNonSecurity: jsii.Boolean(false),
//   	patchFilterGroup: &patchFilterGroupProperty{
//   		patchFilters: []interface{}{
//   			&patchFilterProperty{
//   				key: jsii.String("key"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnPatchBaseline_RuleProperty struct {
	// The number of days after the release date of each patch matched by the rule that the patch is marked as approved in the patch baseline.
	//
	// For example, a value of `7` means that patches are approved seven days after they are released.
	//
	// You must specify a value for `ApproveAfterDays` .
	//
	// Exception: Not supported on Debian Server or Ubuntu Server.
	ApproveAfterDays *float64 `field:"optional" json:"approveAfterDays" yaml:"approveAfterDays"`
	// The cutoff date for auto approval of released patches.
	//
	// Any patches released on or before this date are installed automatically. Not supported on Debian Server or Ubuntu Server.
	//
	// Enter dates in the format `YYYY-MM-DD` . For example, `2021-12-31` .
	ApproveUntilDate *string `field:"optional" json:"approveUntilDate" yaml:"approveUntilDate"`
	// A compliance severity level for all approved patches in a patch baseline.
	//
	// Valid compliance severity levels include the following: `UNSPECIFIED` , `CRITICAL` , `HIGH` , `MEDIUM` , `LOW` , and `INFORMATIONAL` .
	ComplianceLevel *string `field:"optional" json:"complianceLevel" yaml:"complianceLevel"`
	// For managed nodes identified by the approval rule filters, enables a patch baseline to apply non-security updates available in the specified repository.
	//
	// The default value is `false` . Applies to Linux managed nodes only.
	EnableNonSecurity interface{} `field:"optional" json:"enableNonSecurity" yaml:"enableNonSecurity"`
	// The patch filter group that defines the criteria for the rule.
	PatchFilterGroup interface{} `field:"optional" json:"patchFilterGroup" yaml:"patchFilterGroup"`
}

