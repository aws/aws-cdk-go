package mixinsawsssm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsssm/mixinsawsssm/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SSM::PatchBaseline` resource defines the basic information for an AWS Systems Manager patch baseline.
//
// A patch baseline defines which patches are approved for installation on your instances.
//
// For more information, see [CreatePatchBaseline](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_CreatePatchBaseline.html) in the *AWS Systems Manager API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPatchBaselinePropsMixin := awscdkmixinspreview.Mixins.NewCfnPatchBaselinePropsMixin(&CfnPatchBaselineMixinProps{
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
//   	Name: jsii.String("name"),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html
//
type CfnPatchBaselinePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPatchBaselineMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPatchBaselinePropsMixin
type jsiiProxy_CfnPatchBaselinePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPatchBaselinePropsMixin) Props() *CfnPatchBaselineMixinProps {
	var returns *CfnPatchBaselineMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPatchBaselinePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSM::PatchBaseline`.
func NewCfnPatchBaselinePropsMixin(props *CfnPatchBaselineMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPatchBaselinePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPatchBaselinePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPatchBaselinePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnPatchBaselinePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSM::PatchBaseline`.
func NewCfnPatchBaselinePropsMixin_Override(c CfnPatchBaselinePropsMixin, props *CfnPatchBaselineMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnPatchBaselinePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPatchBaselinePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPatchBaselinePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnPatchBaselinePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPatchBaselinePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnPatchBaselinePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPatchBaselinePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPatchBaselinePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

