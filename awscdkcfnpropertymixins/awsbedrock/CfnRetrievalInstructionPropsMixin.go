package awsbedrock

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Bedrock::RetrievalInstruction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnRetrievalInstructionPropsMixin := awscdkcfnpropertymixins.Aws_bedrock.NewCfnRetrievalInstructionPropsMixin(&CfnRetrievalInstructionMixinProps{
//   	CustomerEncryptionKeyArn: jsii.String("customerEncryptionKeyArn"),
//   	Description: jsii.String("description"),
//   	Instruction: &InstructionProperty{
//   		AnswerFormat: jsii.String("answerFormat"),
//   		AnswerPolicy: jsii.String("answerPolicy"),
//   		Persona: jsii.String("persona"),
//   		RetrievalGuidance: jsii.String("retrievalGuidance"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-retrievalinstruction.html
//
type CfnRetrievalInstructionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRetrievalInstructionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRetrievalInstructionPropsMixin
type jsiiProxy_CfnRetrievalInstructionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRetrievalInstructionPropsMixin) Props() *CfnRetrievalInstructionMixinProps {
	var returns *CfnRetrievalInstructionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRetrievalInstructionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Bedrock::RetrievalInstruction`.
func NewCfnRetrievalInstructionPropsMixin(props *CfnRetrievalInstructionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRetrievalInstructionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRetrievalInstructionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRetrievalInstructionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrock.CfnRetrievalInstructionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Bedrock::RetrievalInstruction`.
func NewCfnRetrievalInstructionPropsMixin_Override(c CfnRetrievalInstructionPropsMixin, props *CfnRetrievalInstructionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrock.CfnRetrievalInstructionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRetrievalInstructionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRetrievalInstructionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bedrock.CfnRetrievalInstructionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRetrievalInstructionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bedrock.CfnRetrievalInstructionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRetrievalInstructionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRetrievalInstructionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

