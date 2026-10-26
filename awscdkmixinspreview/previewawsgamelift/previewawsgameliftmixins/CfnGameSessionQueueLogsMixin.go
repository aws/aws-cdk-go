package previewawsgameliftmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgamelift/previewawsgameliftmixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
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
//   var logsDelivery ILogsDelivery
//
//   cfnGameSessionQueueLogsMixin := awscdkmixinspreview.Mixins.NewCfnGameSessionQueueLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html
//
type CfnGameSessionQueueLogsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGameSessionQueueLogsMixin
type jsiiProxy_CfnGameSessionQueueLogsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnGameSessionQueueLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGameSessionQueueLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::GameLift::GameSessionQueue`.
func NewCfnGameSessionQueueLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnGameSessionQueueLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnGameSessionQueueLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGameSessionQueueLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::GameLift::GameSessionQueue`.
func NewCfnGameSessionQueueLogsMixin_Override(c CfnGameSessionQueueLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnGameSessionQueueLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGameSessionQueueLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGameSessionQueueLogsMixin_QUEUE_PLACEMENT_LOGS() CfnGameSessionQueueQueuePlacementLogs {
	_init_.Initialize()
	var returns CfnGameSessionQueueQueuePlacementLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueLogsMixin",
		"QUEUE_PLACEMENT_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGameSessionQueueLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnGameSessionQueueLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

