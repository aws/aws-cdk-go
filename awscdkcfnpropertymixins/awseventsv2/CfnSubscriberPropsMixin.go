package awseventsv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awseventsv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource type definition for AWS::EventsV2::Subscriber, an Amazon EventBridge subscription that delivers events from an event bus to a target.
//
// Canonical identity is the combination of bus, name, and target. Replacement uses delete_then_create.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnSubscriberPropsMixin := awscdkcfnpropertymixins.Aws_eventsv2.NewCfnSubscriberPropsMixin(&CfnSubscriberMixinProps{
//   	BatchConfiguration: &BatchConfigurationProperty{
//   		MaxBatchSize: jsii.Number(123),
//   		MaxBatchWindowInSeconds: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	EventBusArn: jsii.String("eventBusArn"),
//   	FilterConfiguration: &FilterConfigurationProperty{
//   		Filters: []interface{}{
//   			&FilterProperty{
//   				Pattern: jsii.String("pattern"),
//   				Scope: jsii.String("scope"),
//   			},
//   		},
//   		Language: jsii.String("language"),
//   	},
//   	LogConfiguration: &LogConfigurationProperty{
//   		IncludePayload: jsii.String("includePayload"),
//   		Level: jsii.String("level"),
//   	},
//   	Name: jsii.String("name"),
//   	OnFailureConfiguration: &OnFailureConfigurationProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	PointInTimeConfiguration: &PointInTimeConfigurationProperty{
//   		EndPoint: jsii.Number(123),
//   		PointType: jsii.String("pointType"),
//   		StartingPoint: jsii.Number(123),
//   	},
//   	ResumePosition: jsii.String("resumePosition"),
//   	RetryPolicy: &RetryPolicyProperty{
//   		MaxEventAgeInSeconds: jsii.Number(123),
//   		MaxRetryAttempts: jsii.Number(123),
//   		RetryStrategy: jsii.String("retryStrategy"),
//   	},
//   	StartingPosition: jsii.String("startingPosition"),
//   	State: jsii.String("state"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Transformer: &TransformerProperty{
//   		JsonataConfiguration: &JsonataConfigurationProperty{
//   			Expression: jsii.String("expression"),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-subscriber.html
//
type CfnSubscriberPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnSubscriberMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSubscriberPropsMixin
type jsiiProxy_CfnSubscriberPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnSubscriberPropsMixin) Props() *CfnSubscriberMixinProps {
	var returns *CfnSubscriberMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriberPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EventsV2::Subscriber`.
func NewCfnSubscriberPropsMixin(props *CfnSubscriberMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnSubscriberPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSubscriberPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSubscriberPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EventsV2::Subscriber`.
func NewCfnSubscriberPropsMixin_Override(c CfnSubscriberPropsMixin, props *CfnSubscriberMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnSubscriberPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSubscriberPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSubscriberPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_eventsv2.CfnSubscriberPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSubscriberPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSubscriberPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

