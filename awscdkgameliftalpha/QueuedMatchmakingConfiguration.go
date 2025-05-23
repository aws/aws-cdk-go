package awscdkgameliftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

// A FlexMatch matchmaker process does the work of building a game match.
//
// It manages the pool of matchmaking requests received, forms teams for a match, processes and selects players to find the best possible player groups, and initiates the process of placing and starting a game session for the match.
// This topic describes the key aspects of a matchmaker and how to configure one customized for your game.
//
// Example:
//   var queue gameSessionQueue
//   var ruleSet matchmakingRuleSet
//
//
//   gamelift.NewQueuedMatchmakingConfiguration(this, jsii.String("QueuedMatchmakingConfiguration"), &QueuedMatchmakingConfigurationProps{
//   	MatchmakingConfigurationName: jsii.String("test-queued-config-name"),
//   	GameSessionQueues: []iGameSessionQueue{
//   		queue,
//   	},
//   	RuleSet: ruleSet,
//   })
//
// See: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-configuration.html
//
// Experimental.
type QueuedMatchmakingConfiguration interface {
	MatchmakingConfigurationBase
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The ARN of the matchmaking configuration.
	// Experimental.
	MatchmakingConfigurationArn() *string
	// The name of the matchmaking configuration.
	// Experimental.
	MatchmakingConfigurationName() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The notification target for matchmaking events.
	// Experimental.
	NotificationTarget() awssns.ITopic
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Adds a game session queue destination to the matchmaking configuration.
	// Experimental.
	AddGameSessionQueue(gameSessionQueue IGameSessionQueue)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Return the given named metric for this matchmaking configuration.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Matchmaking requests currently being processed or waiting to be processed.
	// Experimental.
	MetricCurrentTickets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// For matchmaking configurations that require acceptance, the potential matches that were accepted since the last report.
	// Experimental.
	MetricMatchesAccepted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Potential matches that were created since the last report.
	// Experimental.
	MetricMatchesCreated(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Matches that were successfully placed into a game session since the last report.
	// Experimental.
	MetricMatchesPlaced(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// For matchmaking configurations that require acceptance, the potential matches that were rejected by at least one player since the last report.
	// Experimental.
	MetricMatchesRejected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Players in matchmaking tickets that were added since the last report.
	// Experimental.
	MetricPlayersStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// For matchmaking requests that were put into a potential match before the last report, the amount of time between ticket creation and potential match creation.
	//
	// Units: seconds.
	// Experimental.
	MetricTimeToMatch(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QueuedMatchmakingConfiguration
type jsiiProxy_QueuedMatchmakingConfiguration struct {
	jsiiProxy_MatchmakingConfigurationBase
}

func (j *jsiiProxy_QueuedMatchmakingConfiguration) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueuedMatchmakingConfiguration) MatchmakingConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchmakingConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueuedMatchmakingConfiguration) MatchmakingConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchmakingConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueuedMatchmakingConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueuedMatchmakingConfiguration) NotificationTarget() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"notificationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueuedMatchmakingConfiguration) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueuedMatchmakingConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewQueuedMatchmakingConfiguration(scope constructs.Construct, id *string, props *QueuedMatchmakingConfigurationProps) QueuedMatchmakingConfiguration {
	_init_.Initialize()

	if err := validateNewQueuedMatchmakingConfigurationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_QueuedMatchmakingConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.QueuedMatchmakingConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewQueuedMatchmakingConfiguration_Override(q QueuedMatchmakingConfiguration, scope constructs.Construct, id *string, props *QueuedMatchmakingConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.QueuedMatchmakingConfiguration",
		[]interface{}{scope, id, props},
		q,
	)
}

// Import an existing matchmaking configuration from its attributes.
// Experimental.
func QueuedMatchmakingConfiguration_FromMatchmakingConfigurationAttributes(scope constructs.Construct, id *string, attrs *MatchmakingConfigurationAttributes) IMatchmakingConfiguration {
	_init_.Initialize()

	if err := validateQueuedMatchmakingConfiguration_FromMatchmakingConfigurationAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IMatchmakingConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.QueuedMatchmakingConfiguration",
		"fromMatchmakingConfigurationAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing matchmaking configuration from its ARN.
// Experimental.
func QueuedMatchmakingConfiguration_FromQueuedMatchmakingConfigurationArn(scope constructs.Construct, id *string, matchmakingConfigurationArn *string) IMatchmakingConfiguration {
	_init_.Initialize()

	if err := validateQueuedMatchmakingConfiguration_FromQueuedMatchmakingConfigurationArnParameters(scope, id, matchmakingConfigurationArn); err != nil {
		panic(err)
	}
	var returns IMatchmakingConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.QueuedMatchmakingConfiguration",
		"fromQueuedMatchmakingConfigurationArn",
		[]interface{}{scope, id, matchmakingConfigurationArn},
		&returns,
	)

	return returns
}

// Import an existing matchmaking configuration from its name.
// Experimental.
func QueuedMatchmakingConfiguration_FromQueuedMatchmakingConfigurationName(scope constructs.Construct, id *string, matchmakingConfigurationName *string) IMatchmakingConfiguration {
	_init_.Initialize()

	if err := validateQueuedMatchmakingConfiguration_FromQueuedMatchmakingConfigurationNameParameters(scope, id, matchmakingConfigurationName); err != nil {
		panic(err)
	}
	var returns IMatchmakingConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.QueuedMatchmakingConfiguration",
		"fromQueuedMatchmakingConfigurationName",
		[]interface{}{scope, id, matchmakingConfigurationName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func QueuedMatchmakingConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQueuedMatchmakingConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.QueuedMatchmakingConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func QueuedMatchmakingConfiguration_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateQueuedMatchmakingConfiguration_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.QueuedMatchmakingConfiguration",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func QueuedMatchmakingConfiguration_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateQueuedMatchmakingConfiguration_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.QueuedMatchmakingConfiguration",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func QueuedMatchmakingConfiguration_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-gamelift-alpha.QueuedMatchmakingConfiguration",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) AddGameSessionQueue(gameSessionQueue IGameSessionQueue) {
	if err := q.validateAddGameSessionQueueParameters(gameSessionQueue); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addGameSessionQueue",
		[]interface{}{gameSessionQueue},
	)
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := q.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := q.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) GetResourceNameAttribute(nameAttr *string) *string {
	if err := q.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) MetricCurrentTickets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricCurrentTicketsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricCurrentTickets",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) MetricMatchesAccepted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricMatchesAcceptedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricMatchesAccepted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) MetricMatchesCreated(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricMatchesCreatedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricMatchesCreated",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) MetricMatchesPlaced(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricMatchesPlacedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricMatchesPlaced",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) MetricMatchesRejected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricMatchesRejectedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricMatchesRejected",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) MetricPlayersStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricPlayersStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricPlayersStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) MetricTimeToMatch(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricTimeToMatchParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricTimeToMatch",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueuedMatchmakingConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

