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
//   	AvailableSecurityUpdatesComplianceStatus: jsii.String("availableSecurityUpdatesComplianceStatus"),
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
	// For information about accepted formats for lists of approved patches and rejected patches, see [Package name formats for approved and rejected patch lists](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html) in the *AWS Systems Manager User Guide* .
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
	// Indicates the status you want to assign to security patches that are available but not approved because they don't meet the installation criteria specified in the patch baseline.
	//
	// Example scenario: Security patches that you might want installed can be skipped if you have specified a long period to wait after a patch is released before installation. If an update to the patch is released during your specified waiting period, the waiting period for installing the patch starts over. If the waiting period is too long, multiple versions of the patch could be released but never installed.
	//
	// Supported for Windows Server managed nodes only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-availablesecurityupdatescompliancestatus
	//
	AvailableSecurityUpdatesComplianceStatus *string `field:"optional" json:"availableSecurityUpdatesComplianceStatus" yaml:"availableSecurityUpdatesComplianceStatus"`
	// Indicates whether this is the default baseline.
	//
	// AWS Systems Manager supports creating multiple default patch baselines. For example, you can create a default patch baseline for each operating system.
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
	//
	// > The `GlobalFilters` parameter can be configured only by using the AWS CLI or an AWS SDK. It can't be configured from the Patch Manager console, and its value isn't displayed in the console.
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
	// For information about accepted formats for lists of approved patches and rejected patches, see [Package name formats for approved and rejected patch lists](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html) in the *AWS Systems Manager User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-rejectedpatches
	//
	RejectedPatches *[]*string `field:"optional" json:"rejectedPatches" yaml:"rejectedPatches"`
	// The action for Patch Manager to take on patches included in the `RejectedPackages` list.
	//
	// - **ALLOW_AS_DEPENDENCY** - *Linux and macOS* : A package in the rejected patches list is installed only if it is a dependency of another package. It is considered compliant with the patch baseline, and its status is reported as `INSTALLED_OTHER` . This is the default action if no option is specified.
	//
	// *Windows Server* : Windows Server doesn't support the concept of package dependencies. If a package in the rejected patches list and already installed on the node, its status is reported as `INSTALLED_OTHER` . Any package not already installed on the node is skipped. This is the default action if no option is specified.
	// - **BLOCK** - *All OSs* : Packages in the rejected patches list, and packages that include them as dependencies, aren't installed by Patch Manager under any circumstances. If a package was installed before it was added to the rejected patches list, or is installed outside of Patch Manager afterward, it's considered noncompliant with the patch baseline and its status is reported as `INSTALLED_REJECTED` .
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

