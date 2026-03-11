package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::BedrockAgentCore::PolicyEngine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPolicyEnginePropsMixin := awscdkcfnpropertymixins.Aws_bedrockagentcore.NewCfnPolicyEnginePropsMixin(&CfnPolicyEngineMixinProps{
//   	Description: jsii.String("description"),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policyengine.html
//
type CfnPolicyEnginePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPolicyEngineMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPolicyEnginePropsMixin
type jsiiProxy_CfnPolicyEnginePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPolicyEnginePropsMixin) Props() *CfnPolicyEngineMixinProps {
	var returns *CfnPolicyEngineMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyEnginePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BedrockAgentCore::PolicyEngine`.
func NewCfnPolicyEnginePropsMixin(props *CfnPolicyEngineMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPolicyEnginePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPolicyEnginePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPolicyEnginePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPolicyEnginePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BedrockAgentCore::PolicyEngine`.
func NewCfnPolicyEnginePropsMixin_Override(c CfnPolicyEnginePropsMixin, props *CfnPolicyEngineMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPolicyEnginePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPolicyEnginePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPolicyEnginePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPolicyEnginePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPolicyEnginePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnPolicyEnginePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPolicyEnginePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPolicyEnginePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

