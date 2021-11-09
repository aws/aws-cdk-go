// The CDK Construct Library for AWS::IoT
package awscdkiotalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Properties for an topic rule action.
// Experimental.
type ActionConfig struct {
	// The configuration for this action.
	// Experimental.
	Configuration *awsiot.CfnTopicRule_ActionProperty `json:"configuration"`
}

// An abstract action for TopicRule.
// Experimental.
type IAction interface {
	// Returns the topic rule action specification.
	// Experimental.
	Bind(topicRule ITopicRule) *ActionConfig
}

// The jsii proxy for IAction
type jsiiProxy_IAction struct {
	_ byte // padding
}

func (i *jsiiProxy_IAction) Bind(topicRule ITopicRule) *ActionConfig {
	var returns *ActionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{topicRule},
		&returns,
	)

	return returns
}

// Represents an AWS IoT Rule.
// Experimental.
type ITopicRule interface {
	awscdk.IResource
	// The value of the topic rule Amazon Resource Name (ARN), such as arn:aws:iot:us-east-2:123456789012:rule/rule_name.
	// Experimental.
	TopicRuleArn() *string
	// The name topic rule.
	// Experimental.
	TopicRuleName() *string
}

// The jsii proxy for ITopicRule
type jsiiProxy_ITopicRule struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ITopicRule) TopicRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopicRule) TopicRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicRuleName",
		&returns,
	)
	return returns
}

// Defines AWS IoT SQL.
// Experimental.
type IotSql interface {
	Bind(scope constructs.Construct) *IotSqlConfig
}

// The jsii proxy struct for IotSql
type jsiiProxy_IotSql struct {
	_ byte // padding
}

// Experimental.
func NewIotSql_Override(i IotSql) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-alpha.IotSql",
		nil, // no parameters
		i,
	)
}

// Uses the original SQL version built on 2015-10-08.
//
// Returns: Instance of IotSql
// Experimental.
func IotSql_FromStringAsVer20151008(sql *string) IotSql {
	_init_.Initialize()

	var returns IotSql

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.IotSql",
		"fromStringAsVer20151008",
		[]interface{}{sql},
		&returns,
	)

	return returns
}

// Uses the SQL version built on 2016-03-23.
//
// Returns: Instance of IotSql
// Experimental.
func IotSql_FromStringAsVer20160323(sql *string) IotSql {
	_init_.Initialize()

	var returns IotSql

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.IotSql",
		"fromStringAsVer20160323",
		[]interface{}{sql},
		&returns,
	)

	return returns
}

// Uses the most recent beta SQL version.
//
// If you use this version, it might
// introduce breaking changes to your rules.
//
// Returns: Instance of IotSql
// Experimental.
func IotSql_FromStringAsVerNewestUnstable(sql *string) IotSql {
	_init_.Initialize()

	var returns IotSql

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.IotSql",
		"fromStringAsVerNewestUnstable",
		[]interface{}{sql},
		&returns,
	)

	return returns
}

// Returns the IoT SQL configuration.
// Experimental.
func (i *jsiiProxy_IotSql) Bind(scope constructs.Construct) *IotSqlConfig {
	var returns *IotSqlConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// The type returned from the `bind()` method in {@link IotSql}.
// Experimental.
type IotSqlConfig struct {
	// The version of the SQL rules engine to use when evaluating the rule.
	// Experimental.
	AwsIotSqlVersion *string `json:"awsIotSqlVersion"`
	// The SQL statement used to query the topic.
	// Experimental.
	Sql *string `json:"sql"`
}

// Defines an AWS IoT Rule in this stack.
// Experimental.
type TopicRule interface {
	awscdk.Resource
	ITopicRule
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	TopicRuleArn() *string
	TopicRuleName() *string
	AddAction(action IAction)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for TopicRule
type jsiiProxy_TopicRule struct {
	internal.Type__awscdkResource
	jsiiProxy_ITopicRule
}

func (j *jsiiProxy_TopicRule) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicRule) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicRule) TopicRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicRule) TopicRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicRuleName",
		&returns,
	)
	return returns
}


// Experimental.
func NewTopicRule(scope constructs.Construct, id *string, props *TopicRuleProps) TopicRule {
	_init_.Initialize()

	j := jsiiProxy_TopicRule{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-alpha.TopicRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTopicRule_Override(t TopicRule, scope constructs.Construct, id *string, props *TopicRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-alpha.TopicRule",
		[]interface{}{scope, id, props},
		t,
	)
}

// Import an existing AWS IoT Rule provided an ARN.
// Experimental.
func TopicRule_FromTopicRuleArn(scope constructs.Construct, id *string, topicRuleArn *string) ITopicRule {
	_init_.Initialize()

	var returns ITopicRule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.TopicRule",
		"fromTopicRuleArn",
		[]interface{}{scope, id, topicRuleArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TopicRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.TopicRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func TopicRule_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iot-alpha.TopicRule",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a action to the topic rule.
// Experimental.
func (t *jsiiProxy_TopicRule) AddAction(action IAction) {
	_jsii_.InvokeVoid(
		t,
		"addAction",
		[]interface{}{action},
	)
}

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
func (t *jsiiProxy_TopicRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (t *jsiiProxy_TopicRule) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (t *jsiiProxy_TopicRule) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (t *jsiiProxy_TopicRule) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TopicRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for defining an AWS IoT Rule.
// Experimental.
type TopicRuleProps struct {
	// A simplified SQL syntax to filter messages received on an MQTT topic and push the data elsewhere.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-sql-reference.html
	//
	// Experimental.
	Sql IotSql `json:"sql"`
	// The actions associated with the topic rule.
	// Experimental.
	Actions *[]IAction `json:"actions"`
	// A textual description of the topic rule.
	// Experimental.
	Description *string `json:"description"`
	// Specifies whether the rule is enabled.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// The action AWS IoT performs when it is unable to perform a rule's action.
	// Experimental.
	ErrorAction IAction `json:"errorAction"`
	// The name of the topic rule.
	// Experimental.
	TopicRuleName *string `json:"topicRuleName"`
}

