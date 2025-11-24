package mixinsawsmpa

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsmpa/mixinsawsmpa/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new approval team.
//
// For more information, see [Approval team](https://docs.aws.amazon.com/mpa/latest/userguide/mpa-concepts.html) in the *Multi-party approval User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApprovalTeamPropsMixin := awscdkmixinspreview.Mixins.NewCfnApprovalTeamPropsMixin(&CfnApprovalTeamMixinProps{
//   	ApprovalStrategy: &ApprovalStrategyProperty{
//   		MofN: &MofNApprovalStrategyProperty{
//   			MinApprovalsRequired: jsii.Number(123),
//   		},
//   	},
//   	Approvers: []interface{}{
//   		&ApproverProperty{
//   			ApproverId: jsii.String("approverId"),
//   			PrimaryIdentityId: jsii.String("primaryIdentityId"),
//   			PrimaryIdentitySourceArn: jsii.String("primaryIdentitySourceArn"),
//   			PrimaryIdentityStatus: jsii.String("primaryIdentityStatus"),
//   			ResponseTime: jsii.String("responseTime"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Policies: []interface{}{
//   		&PolicyProperty{
//   			PolicyArn: jsii.String("policyArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-approvalteam.html
//
type CfnApprovalTeamPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApprovalTeamMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApprovalTeamPropsMixin
type jsiiProxy_CfnApprovalTeamPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApprovalTeamPropsMixin) Props() *CfnApprovalTeamMixinProps {
	var returns *CfnApprovalTeamMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApprovalTeamPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MPA::ApprovalTeam`.
func NewCfnApprovalTeamPropsMixin(props *CfnApprovalTeamMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApprovalTeamPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApprovalTeamPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApprovalTeamPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnApprovalTeamPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MPA::ApprovalTeam`.
func NewCfnApprovalTeamPropsMixin_Override(c CfnApprovalTeamPropsMixin, props *CfnApprovalTeamMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnApprovalTeamPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApprovalTeamPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApprovalTeamPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnApprovalTeamPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApprovalTeamPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mpa.mixins.CfnApprovalTeamPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApprovalTeamPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnApprovalTeamPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

