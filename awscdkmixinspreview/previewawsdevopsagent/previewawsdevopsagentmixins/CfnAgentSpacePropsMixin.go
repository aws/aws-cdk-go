package previewawsdevopsagentmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdevopsagent/previewawsdevopsagentmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DevOpsAgent::AgentSpace` resource specifies an Agent Space for the AWS DevOps Agent Service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentSpacePropsMixin := awscdkmixinspreview.Mixins.NewCfnAgentSpacePropsMixin(&CfnAgentSpaceMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	OperatorApp: &OperatorAppProperty{
//   		Iam: &IamAuthConfigurationProperty{
//   			CreatedAt: jsii.String("createdAt"),
//   			OperatorAppRoleArn: jsii.String("operatorAppRoleArn"),
//   			UpdatedAt: jsii.String("updatedAt"),
//   		},
//   		Idc: &IdcAuthConfigurationProperty{
//   			CreatedAt: jsii.String("createdAt"),
//   			IdcApplicationArn: jsii.String("idcApplicationArn"),
//   			IdcInstanceArn: jsii.String("idcInstanceArn"),
//   			OperatorAppRoleArn: jsii.String("operatorAppRoleArn"),
//   			UpdatedAt: jsii.String("updatedAt"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html
//
type CfnAgentSpacePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAgentSpaceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAgentSpacePropsMixin
type jsiiProxy_CfnAgentSpacePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAgentSpacePropsMixin) Props() *CfnAgentSpaceMixinProps {
	var returns *CfnAgentSpaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgentSpacePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DevOpsAgent::AgentSpace`.
func NewCfnAgentSpacePropsMixin(props *CfnAgentSpaceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAgentSpacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAgentSpacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAgentSpacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DevOpsAgent::AgentSpace`.
func NewCfnAgentSpacePropsMixin_Override(c CfnAgentSpacePropsMixin, props *CfnAgentSpaceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAgentSpacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAgentSpacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAgentSpacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAgentSpacePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAgentSpacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

