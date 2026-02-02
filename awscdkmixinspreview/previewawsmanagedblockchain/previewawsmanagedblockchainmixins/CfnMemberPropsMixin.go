package previewawsmanagedblockchainmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmanagedblockchain/previewawsmanagedblockchainmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a member within a Managed Blockchain network.
//
// Applies only to Hyperledger Fabric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMemberPropsMixin := awscdkmixinspreview.Mixins.NewCfnMemberPropsMixin(&CfnMemberMixinProps{
//   	InvitationId: jsii.String("invitationId"),
//   	MemberConfiguration: &MemberConfigurationProperty{
//   		Description: jsii.String("description"),
//   		MemberFrameworkConfiguration: &MemberFrameworkConfigurationProperty{
//   			MemberFabricConfiguration: &MemberFabricConfigurationProperty{
//   				AdminPassword: jsii.String("adminPassword"),
//   				AdminUsername: jsii.String("adminUsername"),
//   			},
//   		},
//   		Name: jsii.String("name"),
//   	},
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		Description: jsii.String("description"),
//   		Framework: jsii.String("framework"),
//   		FrameworkVersion: jsii.String("frameworkVersion"),
//   		Name: jsii.String("name"),
//   		NetworkFrameworkConfiguration: &NetworkFrameworkConfigurationProperty{
//   			NetworkFabricConfiguration: &NetworkFabricConfigurationProperty{
//   				Edition: jsii.String("edition"),
//   			},
//   		},
//   		VotingPolicy: &VotingPolicyProperty{
//   			ApprovalThresholdPolicy: &ApprovalThresholdPolicyProperty{
//   				ProposalDurationInHours: jsii.Number(123),
//   				ThresholdComparator: jsii.String("thresholdComparator"),
//   				ThresholdPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   	NetworkId: jsii.String("networkId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-member.html
//
type CfnMemberPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMemberMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMemberPropsMixin
type jsiiProxy_CfnMemberPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMemberPropsMixin) Props() *CfnMemberMixinProps {
	var returns *CfnMemberMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMemberPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ManagedBlockchain::Member`.
func NewCfnMemberPropsMixin(props *CfnMemberMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMemberPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMemberPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMemberPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ManagedBlockchain::Member`.
func NewCfnMemberPropsMixin_Override(c CfnMemberPropsMixin, props *CfnMemberMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMemberPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMemberPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMemberPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_managedblockchain.mixins.CfnMemberPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMemberPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMemberPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

