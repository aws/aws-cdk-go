package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPatchBaseline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPatchBaselineProps := &CfnPatchBaselineProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ApprovalRules: &RuleGroupProperty{
//   		PatchRules: []interface{}{
//   			&RuleProperty{
//   				ApproveAfterDays: jsii.Number(123),
//   				ApproveUntilDate: jsii.String("approveUntilDate"),
//   				ComplianceLevel: jsii.String("complianceLevel"),
//   				EnableNonSecurity: jsii.Boolean(false),
//   				PatchFilterGroup: &PatchFilterGroupProperty{
//   					PatchFilters: []interface{}{
//   						&PatchFilterProperty{
//   							Key: jsii.String("key"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ApprovedPatches: []*string{
//   		jsii.String("approvedPatches"),
//   	},
//   	ApprovedPatchesComplianceLevel: jsii.String("approvedPatchesComplianceLevel"),
//   	ApprovedPatchesEnableNonSecurity: jsii.Boolean(false),
//   	DefaultBaseline: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	GlobalFilters: &PatchFilterGroupProperty{
//   		PatchFilters: []interface{}{
//   			&PatchFilterProperty{
//   				Key: jsii.String("key"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	OperatingSystem: jsii.String("operatingSystem"),
//   	PatchGroups: []*string{
//   		jsii.String("patchGroups"),
//   	},
//   	RejectedPatches: []*string{
//   		jsii.String("rejectedPatches"),
//   	},
//   	RejectedPatchesAction: jsii.String("rejectedPatchesAction"),
//   	Sources: []interface{}{
//   		&PatchSourceProperty{
//   			Configuration: jsii.String("configuration"),
//   			Name: jsii.String("name"),
//   			Products: []*string{
//   				jsii.String("products"),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html
//
type CfnPatchBaselineProps struct {
	// The name of the patch baseline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A set of rules used to include patches in the baseline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-approvalrules
	//
	ApprovalRules interface{} `field:"optional" json:"approvalRules" yaml:"approvalRules"`
	// A list of explicitly approved patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and rejected patches, see [About package name formats for approved and rejected patch lists](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html) in the *AWS Systems Manager User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-approvedpatches
	//
	ApprovedPatches *[]*string `field:"optional" json:"approvedPatches" yaml:"approvedPatches"`
	// Defines the compliance level for approved patches.
	//
	// When an approved patch is reported as missing, this value describes the severity of the compliance violation. The default value is `UNSPECIFIED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-approvedpatchescompliancelevel
	//
	// Default: - "UNSPECIFIED".
	//
	ApprovedPatchesComplianceLevel *string `field:"optional" json:"approvedPatchesComplianceLevel" yaml:"approvedPatchesComplianceLevel"`
	// Indicates whether the list of approved patches includes non-security updates that should be applied to the managed nodes.
	//
	// The default value is `false` . Applies to Linux managed nodes only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-approvedpatchesenablenonsecurity
	//
	// Default: - false.
	//
	ApprovedPatchesEnableNonSecurity interface{} `field:"optional" json:"approvedPatchesEnableNonSecurity" yaml:"approvedPatchesEnableNonSecurity"`
	// Set the baseline as default baseline.
	//
	// Only registering to default patch baseline is allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-defaultbaseline
	//
	// Default: - false.
	//
	DefaultBaseline interface{} `field:"optional" json:"defaultBaseline" yaml:"defaultBaseline"`
	// A description of the patch baseline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A set of global filters used to include patches in the baseline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-globalfilters
	//
	GlobalFilters interface{} `field:"optional" json:"globalFilters" yaml:"globalFilters"`
	// Defines the operating system the patch baseline applies to.
	//
	// The default value is `WINDOWS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-operatingsystem
	//
	// Default: - "WINDOWS".
	//
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// The name of the patch group to be registered with the patch baseline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-patchgroups
	//
	PatchGroups *[]*string `field:"optional" json:"patchGroups" yaml:"patchGroups"`
	// A list of explicitly rejected patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and rejected patches, see [About package name formats for approved and rejected patch lists](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html) in the *AWS Systems Manager User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-rejectedpatches
	//
	RejectedPatches *[]*string `field:"optional" json:"rejectedPatches" yaml:"rejectedPatches"`
	// The action for Patch Manager to take on patches included in the `RejectedPackages` list.
	//
	// - *`ALLOW_AS_DEPENDENCY`* : A package in the `Rejected` patches list is installed only if it is a dependency of another package. It is considered compliant with the patch baseline, and its status is reported as `InstalledOther` . This is the default action if no option is specified.
	// - *BLOCK* : Packages in the *Rejected patches* list, and packages that include them as dependencies, aren't installed by Patch Manager under any circumstances. If a package was installed before it was added to the *Rejected patches* list, or is installed outside of Patch Manager afterward, it's considered noncompliant with the patch baseline and its status is reported as *InstalledRejected* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-rejectedpatchesaction
	//
	// Default: - "ALLOW_AS_DEPENDENCY".
	//
	RejectedPatchesAction *string `field:"optional" json:"rejectedPatchesAction" yaml:"rejectedPatchesAction"`
	// Information about the patches to use to update the managed nodes, including target operating systems and source repositories.
	//
	// Applies to Linux managed nodes only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-sources
	//
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// Optional metadata that you assign to a resource.
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a patch baseline to identify the severity level of patches it specifies and the operating system family it applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

