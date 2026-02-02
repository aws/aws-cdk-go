package previewawsgameliftmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgamelift/previewawsgameliftmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::GameLift::GameSessionQueue` resource creates a placement queue that processes requests for new game sessions.
//
// A queue uses FleetIQ algorithms to determine the best placement locations and find an available game server, then prompts the game server to start a new game session. Queues can have destinations (GameLift fleets or aliases), which determine where the queue can place new game sessions. A queue can have destinations with varied fleet type (Spot and On-Demand), instance type, and AWS Region .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGameSessionQueuePropsMixin := awscdkmixinspreview.Mixins.NewCfnGameSessionQueuePropsMixin(&CfnGameSessionQueueMixinProps{
//   	CustomEventData: jsii.String("customEventData"),
//   	Destinations: []interface{}{
//   		&DestinationProperty{
//   			DestinationArn: jsii.String("destinationArn"),
//   		},
//   	},
//   	FilterConfiguration: &FilterConfigurationProperty{
//   		AllowedLocations: []*string{
//   			jsii.String("allowedLocations"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	NotificationTarget: jsii.String("notificationTarget"),
//   	PlayerLatencyPolicies: []interface{}{
//   		&PlayerLatencyPolicyProperty{
//   			MaximumIndividualPlayerLatencyMilliseconds: jsii.Number(123),
//   			PolicyDurationSeconds: jsii.Number(123),
//   		},
//   	},
//   	PriorityConfiguration: &PriorityConfigurationProperty{
//   		LocationOrder: []*string{
//   			jsii.String("locationOrder"),
//   		},
//   		PriorityOrder: []*string{
//   			jsii.String("priorityOrder"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeoutInSeconds: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html
//
type CfnGameSessionQueuePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGameSessionQueueMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGameSessionQueuePropsMixin
type jsiiProxy_CfnGameSessionQueuePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGameSessionQueuePropsMixin) Props() *CfnGameSessionQueueMixinProps {
	var returns *CfnGameSessionQueueMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueuePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GameLift::GameSessionQueue`.
func NewCfnGameSessionQueuePropsMixin(props *CfnGameSessionQueueMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGameSessionQueuePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGameSessionQueuePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGameSessionQueuePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueuePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GameLift::GameSessionQueue`.
func NewCfnGameSessionQueuePropsMixin_Override(c CfnGameSessionQueuePropsMixin, props *CfnGameSessionQueueMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueuePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGameSessionQueuePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGameSessionQueuePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueuePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGameSessionQueuePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueuePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGameSessionQueuePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnGameSessionQueuePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

