package previewawsiotmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiot/previewawsiotmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the resource definition of AWS IoT Command.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCommandPropsMixin := awscdkmixinspreview.Mixins.NewCfnCommandPropsMixin(&CfnCommandMixinProps{
//   	CommandId: jsii.String("commandId"),
//   	CreatedAt: jsii.String("createdAt"),
//   	Deprecated: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	LastUpdatedAt: jsii.String("lastUpdatedAt"),
//   	MandatoryParameters: []interface{}{
//   		&CommandParameterProperty{
//   			DefaultValue: &CommandParameterValueProperty{
//   				B: jsii.Boolean(false),
//   				Bin: jsii.String("bin"),
//   				D: jsii.Number(123),
//   				I: jsii.Number(123),
//   				L: jsii.String("l"),
//   				S: jsii.String("s"),
//   				Ul: jsii.String("ul"),
//   			},
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   			Value: &CommandParameterValueProperty{
//   				B: jsii.Boolean(false),
//   				Bin: jsii.String("bin"),
//   				D: jsii.Number(123),
//   				I: jsii.Number(123),
//   				L: jsii.String("l"),
//   				S: jsii.String("s"),
//   				Ul: jsii.String("ul"),
//   			},
//   			ValueConditions: []interface{}{
//   				&CommandParameterValueConditionProperty{
//   					ComparisonOperator: jsii.String("comparisonOperator"),
//   					Operand: &CommandParameterValueComparisonOperandProperty{
//   						Number: jsii.String("number"),
//   						NumberRange: &CommandParameterValueNumberRangeProperty{
//   							Max: jsii.String("max"),
//   							Min: jsii.String("min"),
//   						},
//   						Numbers: []*string{
//   							jsii.String("numbers"),
//   						},
//   						String: jsii.String("string"),
//   						Strings: []*string{
//   							jsii.String("strings"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Namespace: jsii.String("namespace"),
//   	Payload: &CommandPayloadProperty{
//   		Content: jsii.String("content"),
//   		ContentType: jsii.String("contentType"),
//   	},
//   	PayloadTemplate: jsii.String("payloadTemplate"),
//   	PendingDeletion: jsii.Boolean(false),
//   	Preprocessor: &CommandPreprocessorProperty{
//   		AwsJsonSubstitution: &AwsJsonSubstitutionCommandPreprocessorConfigProperty{
//   			OutputFormat: jsii.String("outputFormat"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html
//
type CfnCommandPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCommandMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCommandPropsMixin
type jsiiProxy_CfnCommandPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCommandPropsMixin) Props() *CfnCommandMixinProps {
	var returns *CfnCommandMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCommandPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::Command`.
func NewCfnCommandPropsMixin(props *CfnCommandMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCommandPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCommandPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCommandPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnCommandPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::Command`.
func NewCfnCommandPropsMixin_Override(c CfnCommandPropsMixin, props *CfnCommandMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnCommandPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCommandPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCommandPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnCommandPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCommandPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnCommandPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCommandPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCommandPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

