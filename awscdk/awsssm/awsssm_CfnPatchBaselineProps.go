package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPatchBaseline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPatchBaselineProps := &cfnPatchBaselineProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	approvalRules: &ruleGroupProperty{
//   		patchRules: []interface{}{
//   			&ruleProperty{
//   				approveAfterDays: jsii.Number(123),
//   				approveUntilDate: jsii.String("approveUntilDate"),
//   				complianceLevel: jsii.String("complianceLevel"),
//   				enableNonSecurity: jsii.Boolean(false),
//   				patchFilterGroup: &patchFilterGroupProperty{
//   					patchFilters: []interface{}{
//   						&patchFilterProperty{
//   							key: jsii.String("key"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	approvedPatches: []*string{
//   		jsii.String("approvedPatches"),
//   	},
//   	approvedPatchesComplianceLevel: jsii.String("approvedPatchesComplianceLevel"),
//   	approvedPatchesEnableNonSecurity: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	globalFilters: &patchFilterGroupProperty{
//   		patchFilters: []interface{}{
//   			&patchFilterProperty{
//   				key: jsii.String("key"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	operatingSystem: jsii.String("operatingSystem"),
//   	patchGroups: []*string{
//   		jsii.String("patchGroups"),
//   	},
//   	rejectedPatches: []*string{
//   		jsii.String("rejectedPatches"),
//   	},
//   	rejectedPatchesAction: jsii.String("rejectedPatchesAction"),
//   	sources: []interface{}{
//   		&patchSourceProperty{
//   			configuration: jsii.String("configuration"),
//   			name: jsii.String("name"),
//   			products: []*string{
//   				jsii.String("products"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPatchBaselineProps struct {
	// The name of the patch baseline.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A set of rules used to include patches in the baseline.
	ApprovalRules interface{} `field:"optional" json:"approvalRules" yaml:"approvalRules"`
	// A list of explicitly approved patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and rejected patches, see [About package name formats for approved and rejected patch lists](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html) in the *AWS Systems Manager User Guide* .
	ApprovedPatches *[]*string `field:"optional" json:"approvedPatches" yaml:"approvedPatches"`
	// Defines the compliance level for approved patches.
	//
	// When an approved patch is reported as missing, this value describes the severity of the compliance violation. The default value is `UNSPECIFIED` .
	ApprovedPatchesComplianceLevel *string `field:"optional" json:"approvedPatchesComplianceLevel" yaml:"approvedPatchesComplianceLevel"`
	// Indicates whether the list of approved patches includes non-security updates that should be applied to the managed nodes.
	//
	// The default value is `false` . Applies to Linux managed nodes only.
	ApprovedPatchesEnableNonSecurity interface{} `field:"optional" json:"approvedPatchesEnableNonSecurity" yaml:"approvedPatchesEnableNonSecurity"`
	// A description of the patch baseline.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A set of global filters used to include patches in the baseline.
	GlobalFilters interface{} `field:"optional" json:"globalFilters" yaml:"globalFilters"`
	// Defines the operating system the patch baseline applies to.
	//
	// The default value is `WINDOWS` .
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// The name of the patch group to be registered with the patch baseline.
	PatchGroups *[]*string `field:"optional" json:"patchGroups" yaml:"patchGroups"`
	// A list of explicitly rejected patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and rejected patches, see [About package name formats for approved and rejected patch lists](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html) in the *AWS Systems Manager User Guide* .
	RejectedPatches *[]*string `field:"optional" json:"rejectedPatches" yaml:"rejectedPatches"`
	// The action for Patch Manager to take on patches included in the `RejectedPackages` list.
	//
	// - *`ALLOW_AS_DEPENDENCY`* : A package in the `Rejected` patches list is installed only if it is a dependency of another package. It is considered compliant with the patch baseline, and its status is reported as `InstalledOther` . This is the default action if no option is specified.
	// - *`BLOCK`* : Packages in the `RejectedPatches` list, and packages that include them as dependencies, aren't installed under any circumstances. If a package was installed before it was added to the Rejected patches list, it is considered non-compliant with the patch baseline, and its status is reported as `InstalledRejected` .
	RejectedPatchesAction *string `field:"optional" json:"rejectedPatchesAction" yaml:"rejectedPatchesAction"`
	// Information about the patches to use to update the managed nodes, including target operating systems and source repositories.
	//
	// Applies to Linux managed nodes only.
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// Optional metadata that you assign to a resource.
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a patch baseline to identify the severity level of patches it specifies and the operating system family it applies to.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

