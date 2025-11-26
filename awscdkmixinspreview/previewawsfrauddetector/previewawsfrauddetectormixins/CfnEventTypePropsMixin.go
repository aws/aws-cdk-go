package previewawsfrauddetectormixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsfrauddetector/previewawsfrauddetectormixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Manages an event type.
//
// An event is a business activity that is evaluated for fraud risk. With Amazon Fraud Detector, you generate fraud predictions for events. An event type defines the structure for an event sent to Amazon Fraud Detector. This includes the variables sent as part of the event, the entity performing the event (such as a customer), and the labels that classify the event. Example event types include online payment transactions, account registrations, and authentications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventTypePropsMixin := awscdkmixinspreview.Mixins.NewCfnEventTypePropsMixin(&CfnEventTypeMixinProps{
//   	Description: jsii.String("description"),
//   	EntityTypes: []interface{}{
//   		&EntityTypeProperty{
//   			Arn: jsii.String("arn"),
//   			CreatedTime: jsii.String("createdTime"),
//   			Description: jsii.String("description"),
//   			Inline: jsii.Boolean(false),
//   			LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			Name: jsii.String("name"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	EventVariables: []interface{}{
//   		&EventVariableProperty{
//   			Arn: jsii.String("arn"),
//   			CreatedTime: jsii.String("createdTime"),
//   			DataSource: jsii.String("dataSource"),
//   			DataType: jsii.String("dataType"),
//   			DefaultValue: jsii.String("defaultValue"),
//   			Description: jsii.String("description"),
//   			Inline: jsii.Boolean(false),
//   			LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			Name: jsii.String("name"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			VariableType: jsii.String("variableType"),
//   		},
//   	},
//   	Labels: []interface{}{
//   		&LabelProperty{
//   			Arn: jsii.String("arn"),
//   			CreatedTime: jsii.String("createdTime"),
//   			Description: jsii.String("description"),
//   			Inline: jsii.Boolean(false),
//   			LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			Name: jsii.String("name"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-eventtype.html
//
type CfnEventTypePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEventTypeMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEventTypePropsMixin
type jsiiProxy_CfnEventTypePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEventTypePropsMixin) Props() *CfnEventTypeMixinProps {
	var returns *CfnEventTypeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventTypePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::FraudDetector::EventType`.
func NewCfnEventTypePropsMixin(props *CfnEventTypeMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEventTypePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEventTypePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEventTypePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_frauddetector.mixins.CfnEventTypePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::FraudDetector::EventType`.
func NewCfnEventTypePropsMixin_Override(c CfnEventTypePropsMixin, props *CfnEventTypeMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_frauddetector.mixins.CfnEventTypePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEventTypePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventTypePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_frauddetector.mixins.CfnEventTypePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventTypePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_frauddetector.mixins.CfnEventTypePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventTypePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEventTypePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

