package mixinsawsbedrock

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsbedrock/mixinsawsbedrock/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Automated Reasoning policy for Amazon Bedrock Guardrails.
//
// Automated Reasoning policies use mathematical techniques to detect hallucinations, suggest corrections, and highlight unstated assumptions in the responses of your GenAI application.
//
// To create a policy, you upload a source document that describes the rules that you're encoding. Automated Reasoning extracts important concepts from the source document that will become variables in the policy and infers policy rules.
//
// To learn more about creating Automated Reasoning policies, see [Minimize AI hallucinations and deliver up to 99% verification accuracy with Automated Reasoning checks: Now available](https://docs.aws.amazon.com/aws/minimize-ai-hallucinations-and-deliver-up-to-99-verification-accuracy-with-automated-reasoning-checks-now-available/) in the *AWS News Blog* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAutomatedReasoningPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnAutomatedReasoningPolicyPropsMixin(&CfnAutomatedReasoningPolicyMixinProps{
//   	Description: jsii.String("description"),
//   	ForceDelete: jsii.Boolean(false),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Name: jsii.String("name"),
//   	PolicyDefinition: &PolicyDefinitionProperty{
//   		Rules: []interface{}{
//   			&PolicyDefinitionRuleProperty{
//   				AlternateExpression: jsii.String("alternateExpression"),
//   				Expression: jsii.String("expression"),
//   				Id: jsii.String("id"),
//   			},
//   		},
//   		Types: []interface{}{
//   			&PolicyDefinitionTypeProperty{
//   				Description: jsii.String("description"),
//   				Name: jsii.String("name"),
//   				Values: []interface{}{
//   					&PolicyDefinitionTypeValueProperty{
//   						Description: jsii.String("description"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		Variables: []interface{}{
//   			&PolicyDefinitionVariableProperty{
//   				Description: jsii.String("description"),
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Version: jsii.String("version"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-automatedreasoningpolicy.html
//
type CfnAutomatedReasoningPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAutomatedReasoningPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAutomatedReasoningPolicyPropsMixin
type jsiiProxy_CfnAutomatedReasoningPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAutomatedReasoningPolicyPropsMixin) Props() *CfnAutomatedReasoningPolicyMixinProps {
	var returns *CfnAutomatedReasoningPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutomatedReasoningPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Bedrock::AutomatedReasoningPolicy`.
func NewCfnAutomatedReasoningPolicyPropsMixin(props *CfnAutomatedReasoningPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAutomatedReasoningPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAutomatedReasoningPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAutomatedReasoningPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAutomatedReasoningPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Bedrock::AutomatedReasoningPolicy`.
func NewCfnAutomatedReasoningPolicyPropsMixin_Override(c CfnAutomatedReasoningPolicyPropsMixin, props *CfnAutomatedReasoningPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAutomatedReasoningPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAutomatedReasoningPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAutomatedReasoningPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAutomatedReasoningPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAutomatedReasoningPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAutomatedReasoningPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAutomatedReasoningPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAutomatedReasoningPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

