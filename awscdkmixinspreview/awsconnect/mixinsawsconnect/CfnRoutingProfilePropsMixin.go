package mixinsawsconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsconnect/mixinsawsconnect/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new routing profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRoutingProfilePropsMixin := awscdkmixinspreview.Mixins.NewCfnRoutingProfilePropsMixin(&CfnRoutingProfileMixinProps{
//   	AgentAvailabilityTimer: jsii.String("agentAvailabilityTimer"),
//   	DefaultOutboundQueueArn: jsii.String("defaultOutboundQueueArn"),
//   	Description: jsii.String("description"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	ManualAssignmentQueueConfigs: []interface{}{
//   		&RoutingProfileManualAssignmentQueueConfigProperty{
//   			QueueReference: &RoutingProfileQueueReferenceProperty{
//   				Channel: jsii.String("channel"),
//   				QueueArn: jsii.String("queueArn"),
//   			},
//   		},
//   	},
//   	MediaConcurrencies: []interface{}{
//   		&MediaConcurrencyProperty{
//   			Channel: jsii.String("channel"),
//   			Concurrency: jsii.Number(123),
//   			CrossChannelBehavior: &CrossChannelBehaviorProperty{
//   				BehaviorType: jsii.String("behaviorType"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	QueueConfigs: []interface{}{
//   		&RoutingProfileQueueConfigProperty{
//   			Delay: jsii.Number(123),
//   			Priority: jsii.Number(123),
//   			QueueReference: &RoutingProfileQueueReferenceProperty{
//   				Channel: jsii.String("channel"),
//   				QueueArn: jsii.String("queueArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-routingprofile.html
//
type CfnRoutingProfilePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRoutingProfileMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRoutingProfilePropsMixin
type jsiiProxy_CfnRoutingProfilePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRoutingProfilePropsMixin) Props() *CfnRoutingProfileMixinProps {
	var returns *CfnRoutingProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoutingProfilePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::RoutingProfile`.
func NewCfnRoutingProfilePropsMixin(props *CfnRoutingProfileMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRoutingProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRoutingProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRoutingProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnRoutingProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::RoutingProfile`.
func NewCfnRoutingProfilePropsMixin_Override(c CfnRoutingProfilePropsMixin, props *CfnRoutingProfileMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnRoutingProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRoutingProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRoutingProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnRoutingProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRoutingProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnRoutingProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRoutingProfilePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRoutingProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

