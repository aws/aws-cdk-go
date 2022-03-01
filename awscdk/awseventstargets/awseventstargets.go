package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Use an API Destination rule target.
//
// TODO: EXAMPLE
//
// Experimental.
type ApiDestination interface {
	awsevents.IRuleTarget
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for ApiDestination
type jsiiProxy_ApiDestination struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewApiDestination(apiDestination awsevents.IApiDestination, props *ApiDestinationProps) ApiDestination {
	_init_.Initialize()

	j := jsiiProxy_ApiDestination{}

	_jsii_.Create(
		"monocdk.aws_events_targets.ApiDestination",
		[]interface{}{apiDestination, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApiDestination_Override(a ApiDestination, apiDestination awsevents.IApiDestination, props *ApiDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.ApiDestination",
		[]interface{}{apiDestination, props},
		a,
	)
}

// Returns a RuleTarget that can be used to trigger API destinations from an EventBridge event.
// Experimental.
func (a *jsiiProxy_ApiDestination) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

// Customize the EventBridge Api Destinations Target.
//
// TODO: EXAMPLE
//
// Experimental.
type ApiDestinationProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Experimental.
	MaxEventAge awscdk.Duration `json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
	// The event to send.
	// Experimental.
	Event awsevents.RuleTargetInput `json:"event" yaml:"event"`
	// The role to assume before invoking the target.
	// Experimental.
	EventRole awsiam.IRole `json:"eventRole" yaml:"eventRole"`
	// Additional headers sent to the API Destination.
	//
	// These are merged with headers specified on the Connection, with
	// the headers on the Connection taking precedence.
	//
	// You can only specify secret values on the Connection.
	// Experimental.
	HeaderParameters *map[string]*string `json:"headerParameters" yaml:"headerParameters"`
	// Path parameters to insert in place of path wildcards (`*`).
	//
	// If the API destination has a wilcard in the path, these path parts
	// will be inserted in that place.
	// Experimental.
	PathParameterValues *[]*string `json:"pathParameterValues" yaml:"pathParameterValues"`
	// Additional query string parameters sent to the API Destination.
	//
	// These are merged with headers specified on the Connection, with
	// the headers on the Connection taking precedence.
	//
	// You can only specify secret values on the Connection.
	// Experimental.
	QueryStringParameters *map[string]*string `json:"queryStringParameters" yaml:"queryStringParameters"`
}

// Use an API Gateway REST APIs as a target for Amazon EventBridge rules.
//
// TODO: EXAMPLE
//
// Experimental.
type ApiGateway interface {
	awsevents.IRuleTarget
	RestApi() awsapigateway.RestApi
	Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for ApiGateway
type jsiiProxy_ApiGateway struct {
	internal.Type__awseventsIRuleTarget
}

func (j *jsiiProxy_ApiGateway) RestApi() awsapigateway.RestApi {
	var returns awsapigateway.RestApi
	_jsii_.Get(
		j,
		"restApi",
		&returns,
	)
	return returns
}


// Experimental.
func NewApiGateway(restApi awsapigateway.RestApi, props *ApiGatewayProps) ApiGateway {
	_init_.Initialize()

	j := jsiiProxy_ApiGateway{}

	_jsii_.Create(
		"monocdk.aws_events_targets.ApiGateway",
		[]interface{}{restApi, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApiGateway_Override(a ApiGateway, restApi awsapigateway.RestApi, props *ApiGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.ApiGateway",
		[]interface{}{restApi, props},
		a,
	)
}

// Returns a RuleTarget that can be used to trigger this API Gateway REST APIs as a result from an EventBridge event.
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/resource-based-policies-eventbridge.html#sqs-permissions
//
// Experimental.
func (a *jsiiProxy_ApiGateway) Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{rule, _id},
		&returns,
	)

	return returns
}

// Customize the API Gateway Event Target.
//
// TODO: EXAMPLE
//
// Experimental.
type ApiGatewayProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Experimental.
	MaxEventAge awscdk.Duration `json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
	// The role to assume before invoking the target (i.e., the pipeline) when the given rule is triggered.
	// Experimental.
	EventRole awsiam.IRole `json:"eventRole" yaml:"eventRole"`
	// The headers to be set when requesting API.
	// Experimental.
	HeaderParameters *map[string]*string `json:"headerParameters" yaml:"headerParameters"`
	// The method for api resource invoked by the rule.
	// Experimental.
	Method *string `json:"method" yaml:"method"`
	// The api resource invoked by the rule.
	//
	// We can use wildcards('*') to specify the path. In that case,
	// an equal number of real values must be specified for pathParameterValues.
	// Experimental.
	Path *string `json:"path" yaml:"path"`
	// The path parameter values to be used to populate to wildcards("*") of requesting api path.
	// Experimental.
	PathParameterValues *[]*string `json:"pathParameterValues" yaml:"pathParameterValues"`
	// This will be the post request body send to the API.
	// Experimental.
	PostBody awsevents.RuleTargetInput `json:"postBody" yaml:"postBody"`
	// The query parameters to be set when requesting API.
	// Experimental.
	QueryStringParameters *map[string]*string `json:"queryStringParameters" yaml:"queryStringParameters"`
	// The deploy stage of api gateway invoked by the rule.
	// Experimental.
	Stage *string `json:"stage" yaml:"stage"`
}

// Use an AWS Lambda function that makes API calls as an event rule target.
//
// TODO: EXAMPLE
//
// Experimental.
type AwsApi interface {
	awsevents.IRuleTarget
	Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for AwsApi
type jsiiProxy_AwsApi struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewAwsApi(props *AwsApiProps) AwsApi {
	_init_.Initialize()

	j := jsiiProxy_AwsApi{}

	_jsii_.Create(
		"monocdk.aws_events_targets.AwsApi",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApi_Override(a AwsApi, props *AwsApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.AwsApi",
		[]interface{}{props},
		a,
	)
}

// Returns a RuleTarget that can be used to trigger this AwsApi as a result from an EventBridge event.
// Experimental.
func (a *jsiiProxy_AwsApi) Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{rule, id},
		&returns,
	)

	return returns
}

// Rule target input for an AwsApi target.
//
// TODO: EXAMPLE
//
// Experimental.
type AwsApiInput struct {
	// The service action to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Action *string `json:"action" yaml:"action"`
	// The service to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Service *string `json:"service" yaml:"service"`
	// API version to use for the service.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/locking-api-versions.html
	//
	// Experimental.
	ApiVersion *string `json:"apiVersion" yaml:"apiVersion"`
	// The regex pattern to use to catch API errors.
	//
	// The `code` property of the
	// `Error` object will be tested against this pattern. If there is a match an
	// error will not be thrown.
	// Experimental.
	CatchErrorPattern *string `json:"catchErrorPattern" yaml:"catchErrorPattern"`
	// The parameters for the service action.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
}

// Properties for an AwsApi target.
//
// TODO: EXAMPLE
//
// Experimental.
type AwsApiProps struct {
	// The service action to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Action *string `json:"action" yaml:"action"`
	// The service to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Service *string `json:"service" yaml:"service"`
	// API version to use for the service.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/locking-api-versions.html
	//
	// Experimental.
	ApiVersion *string `json:"apiVersion" yaml:"apiVersion"`
	// The regex pattern to use to catch API errors.
	//
	// The `code` property of the
	// `Error` object will be tested against this pattern. If there is a match an
	// error will not be thrown.
	// Experimental.
	CatchErrorPattern *string `json:"catchErrorPattern" yaml:"catchErrorPattern"`
	// The parameters for the service action.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// The IAM policy statement to allow the API call.
	//
	// Use only if
	// resource restriction is needed.
	// Experimental.
	PolicyStatement awsiam.PolicyStatement `json:"policyStatement" yaml:"policyStatement"`
}

// Use an AWS Batch Job / Queue as an event rule target.
//
// Most likely the code will look something like this:
// `new BatchJob(jobQueue.jobQueueArn, jobQueue, jobDefinition.jobDefinitionArn, jobDefinition)`
//
// In the future this API will be improved to be fully typed
//
// TODO: EXAMPLE
//
// Experimental.
type BatchJob interface {
	awsevents.IRuleTarget
	Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for BatchJob
type jsiiProxy_BatchJob struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewBatchJob(jobQueueArn *string, jobQueueScope awscdk.IConstruct, jobDefinitionArn *string, jobDefinitionScope awscdk.IConstruct, props *BatchJobProps) BatchJob {
	_init_.Initialize()

	j := jsiiProxy_BatchJob{}

	_jsii_.Create(
		"monocdk.aws_events_targets.BatchJob",
		[]interface{}{jobQueueArn, jobQueueScope, jobDefinitionArn, jobDefinitionScope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBatchJob_Override(b BatchJob, jobQueueArn *string, jobQueueScope awscdk.IConstruct, jobDefinitionArn *string, jobDefinitionScope awscdk.IConstruct, props *BatchJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.BatchJob",
		[]interface{}{jobQueueArn, jobQueueScope, jobDefinitionArn, jobDefinitionScope, props},
		b,
	)
}

// Returns a RuleTarget that can be used to trigger queue this batch job as a result from an EventBridge event.
// Experimental.
func (b *jsiiProxy_BatchJob) Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{rule, _id},
		&returns,
	)

	return returns
}

// Customize the Batch Job Event Target.
//
// TODO: EXAMPLE
//
// Experimental.
type BatchJobProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Experimental.
	MaxEventAge awscdk.Duration `json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
	// The number of times to attempt to retry, if the job fails.
	//
	// Valid values are 1–10.
	// Experimental.
	Attempts *float64 `json:"attempts" yaml:"attempts"`
	// The event to send to the Lambda.
	//
	// This will be the payload sent to the Lambda Function.
	// Experimental.
	Event awsevents.RuleTargetInput `json:"event" yaml:"event"`
	// The name of the submitted job.
	// Experimental.
	JobName *string `json:"jobName" yaml:"jobName"`
	// The size of the array, if this is an array batch job.
	//
	// Valid values are integers between 2 and 10,000.
	// Experimental.
	Size *float64 `json:"size" yaml:"size"`
}

// Use an AWS CloudWatch LogGroup as an event rule target.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudWatchLogGroup interface {
	awsevents.IRuleTarget
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for CloudWatchLogGroup
type jsiiProxy_CloudWatchLogGroup struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewCloudWatchLogGroup(logGroup awslogs.ILogGroup, props *LogGroupProps) CloudWatchLogGroup {
	_init_.Initialize()

	j := jsiiProxy_CloudWatchLogGroup{}

	_jsii_.Create(
		"monocdk.aws_events_targets.CloudWatchLogGroup",
		[]interface{}{logGroup, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchLogGroup_Override(c CloudWatchLogGroup, logGroup awslogs.ILogGroup, props *LogGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.CloudWatchLogGroup",
		[]interface{}{logGroup, props},
		c,
	)
}

// Returns a RuleTarget that can be used to log an event into a CloudWatch LogGroup.
// Experimental.
func (c *jsiiProxy_CloudWatchLogGroup) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

// Start a CodeBuild build when an Amazon EventBridge rule is triggered.
//
// TODO: EXAMPLE
//
// Experimental.
type CodeBuildProject interface {
	awsevents.IRuleTarget
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for CodeBuildProject
type jsiiProxy_CodeBuildProject struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewCodeBuildProject(project awscodebuild.IProject, props *CodeBuildProjectProps) CodeBuildProject {
	_init_.Initialize()

	j := jsiiProxy_CodeBuildProject{}

	_jsii_.Create(
		"monocdk.aws_events_targets.CodeBuildProject",
		[]interface{}{project, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCodeBuildProject_Override(c CodeBuildProject, project awscodebuild.IProject, props *CodeBuildProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.CodeBuildProject",
		[]interface{}{project, props},
		c,
	)
}

// Allows using build projects as event rule targets.
// Experimental.
func (c *jsiiProxy_CodeBuildProject) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

// Customize the CodeBuild Event Target.
//
// TODO: EXAMPLE
//
// Experimental.
type CodeBuildProjectProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Experimental.
	MaxEventAge awscdk.Duration `json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
	// The event to send to CodeBuild.
	//
	// This will be the payload for the StartBuild API.
	// Experimental.
	Event awsevents.RuleTargetInput `json:"event" yaml:"event"`
	// The role to assume before invoking the target (i.e., the codebuild) when the given rule is triggered.
	// Experimental.
	EventRole awsiam.IRole `json:"eventRole" yaml:"eventRole"`
}

// Allows the pipeline to be used as an EventBridge rule target.
//
// TODO: EXAMPLE
//
// Experimental.
type CodePipeline interface {
	awsevents.IRuleTarget
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for CodePipeline
type jsiiProxy_CodePipeline struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewCodePipeline(pipeline awscodepipeline.IPipeline, options *CodePipelineTargetOptions) CodePipeline {
	_init_.Initialize()

	j := jsiiProxy_CodePipeline{}

	_jsii_.Create(
		"monocdk.aws_events_targets.CodePipeline",
		[]interface{}{pipeline, options},
		&j,
	)

	return &j
}

// Experimental.
func NewCodePipeline_Override(c CodePipeline, pipeline awscodepipeline.IPipeline, options *CodePipelineTargetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.CodePipeline",
		[]interface{}{pipeline, options},
		c,
	)
}

// Returns the rule target specification.
//
// NOTE: Do not use the various `inputXxx` options. They can be set in a call to `addTarget`.
// Experimental.
func (c *jsiiProxy_CodePipeline) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

// Customization options when creating a {@link CodePipeline} event target.
//
// TODO: EXAMPLE
//
// Experimental.
type CodePipelineTargetOptions struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Experimental.
	MaxEventAge awscdk.Duration `json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
	// The role to assume before invoking the target (i.e., the pipeline) when the given rule is triggered.
	// Experimental.
	EventRole awsiam.IRole `json:"eventRole" yaml:"eventRole"`
}

// TODO: EXAMPLE
//
// Experimental.
type ContainerOverride struct {
	// Name of the container inside the task definition.
	// Experimental.
	ContainerName *string `json:"containerName" yaml:"containerName"`
	// Command to run inside the container.
	// Experimental.
	Command *[]*string `json:"command" yaml:"command"`
	// The number of cpu units reserved for the container.
	// Experimental.
	Cpu *float64 `json:"cpu" yaml:"cpu"`
	// Variables to set in the container's environment.
	// Experimental.
	Environment *[]*TaskEnvironmentVariable `json:"environment" yaml:"environment"`
	// Hard memory limit on the container.
	// Experimental.
	MemoryLimit *float64 `json:"memoryLimit" yaml:"memoryLimit"`
	// Soft memory limit on the container.
	// Experimental.
	MemoryReservation *float64 `json:"memoryReservation" yaml:"memoryReservation"`
}

// Start a task on an ECS cluster.
//
// TODO: EXAMPLE
//
// Experimental.
type EcsTask interface {
	awsevents.IRuleTarget
	SecurityGroup() awsec2.ISecurityGroup
	SecurityGroups() *[]awsec2.ISecurityGroup
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for EcsTask
type jsiiProxy_EcsTask struct {
	internal.Type__awseventsIRuleTarget
}

func (j *jsiiProxy_EcsTask) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTask) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}


// Experimental.
func NewEcsTask(props *EcsTaskProps) EcsTask {
	_init_.Initialize()

	j := jsiiProxy_EcsTask{}

	_jsii_.Create(
		"monocdk.aws_events_targets.EcsTask",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcsTask_Override(e EcsTask, props *EcsTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.EcsTask",
		[]interface{}{props},
		e,
	)
}

// Allows using tasks as target of EventBridge events.
// Experimental.
func (e *jsiiProxy_EcsTask) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

// Properties to define an ECS Event Task.
//
// TODO: EXAMPLE
//
// Experimental.
type EcsTaskProps struct {
	// Cluster where service will be deployed.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster" yaml:"cluster"`
	// Task Definition of the task that should be started.
	// Experimental.
	TaskDefinition awsecs.ITaskDefinition `json:"taskDefinition" yaml:"taskDefinition"`
	// Container setting overrides.
	//
	// Key is the name of the container to override, value is the
	// values you want to override.
	// Experimental.
	ContainerOverrides *[]*ContainerOverride `json:"containerOverrides" yaml:"containerOverrides"`
	// The platform version on which to run your task.
	//
	// Unless you have specific compatibility requirements, you don't need to specify this.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html
	//
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `json:"platformVersion" yaml:"platformVersion"`
	// Existing IAM role to run the ECS task.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// Existing security group to use for the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking)
	// Deprecated: use securityGroups instead
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
	// Existing security groups to use for the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking)
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// In what subnets to place the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking)
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// How many tasks should be started when this event is triggered.
	// Experimental.
	TaskCount *float64 `json:"taskCount" yaml:"taskCount"`
}

// Notify an existing Event Bus of an event.
//
// TODO: EXAMPLE
//
// Experimental.
type EventBus interface {
	awsevents.IRuleTarget
	Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for EventBus
type jsiiProxy_EventBus struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewEventBus(eventBus awsevents.IEventBus, props *EventBusProps) EventBus {
	_init_.Initialize()

	j := jsiiProxy_EventBus{}

	_jsii_.Create(
		"monocdk.aws_events_targets.EventBus",
		[]interface{}{eventBus, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEventBus_Override(e EventBus, eventBus awsevents.IEventBus, props *EventBusProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.EventBus",
		[]interface{}{eventBus, props},
		e,
	)
}

// Returns the rule target specification.
//
// NOTE: Do not use the various `inputXxx` options. They can be set in a call to `addTarget`.
// Experimental.
func (e *jsiiProxy_EventBus) Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{rule, _id},
		&returns,
	)

	return returns
}

// Configuration properties of an Event Bus event.
//
// Cannot extend TargetBaseProps. Retry policy is not supported for Event bus targets.
//
// TODO: EXAMPLE
//
// Experimental.
type EventBusProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Role to be used to publish the event.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// Customize the Firehose Stream Event Target.
//
// TODO: EXAMPLE
//
// Experimental.
type KinesisFirehoseStream interface {
	awsevents.IRuleTarget
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for KinesisFirehoseStream
type jsiiProxy_KinesisFirehoseStream struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewKinesisFirehoseStream(stream awskinesisfirehose.CfnDeliveryStream, props *KinesisFirehoseStreamProps) KinesisFirehoseStream {
	_init_.Initialize()

	j := jsiiProxy_KinesisFirehoseStream{}

	_jsii_.Create(
		"monocdk.aws_events_targets.KinesisFirehoseStream",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisFirehoseStream_Override(k KinesisFirehoseStream, stream awskinesisfirehose.CfnDeliveryStream, props *KinesisFirehoseStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.KinesisFirehoseStream",
		[]interface{}{stream, props},
		k,
	)
}

// Returns a RuleTarget that can be used to trigger this Firehose Stream as a result from a Event Bridge event.
// Experimental.
func (k *jsiiProxy_KinesisFirehoseStream) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

// Customize the Firehose Stream Event Target.
//
// TODO: EXAMPLE
//
// Experimental.
type KinesisFirehoseStreamProps struct {
	// The message to send to the stream.
	//
	// Must be a valid JSON text passed to the target stream.
	// Experimental.
	Message awsevents.RuleTargetInput `json:"message" yaml:"message"`
}

// Use a Kinesis Stream as a target for AWS CloudWatch event rules.
//
// TODO: EXAMPLE
//
// Experimental.
type KinesisStream interface {
	awsevents.IRuleTarget
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for KinesisStream
type jsiiProxy_KinesisStream struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewKinesisStream(stream awskinesis.IStream, props *KinesisStreamProps) KinesisStream {
	_init_.Initialize()

	j := jsiiProxy_KinesisStream{}

	_jsii_.Create(
		"monocdk.aws_events_targets.KinesisStream",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisStream_Override(k KinesisStream, stream awskinesis.IStream, props *KinesisStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.KinesisStream",
		[]interface{}{stream, props},
		k,
	)
}

// Returns a RuleTarget that can be used to trigger this Kinesis Stream as a result from a CloudWatch event.
// Experimental.
func (k *jsiiProxy_KinesisStream) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

// Customize the Kinesis Stream Event Target.
//
// TODO: EXAMPLE
//
// Experimental.
type KinesisStreamProps struct {
	// The message to send to the stream.
	//
	// Must be a valid JSON text passed to the target stream.
	// Experimental.
	Message awsevents.RuleTargetInput `json:"message" yaml:"message"`
	// Partition Key Path for records sent to this stream.
	// Experimental.
	PartitionKeyPath *string `json:"partitionKeyPath" yaml:"partitionKeyPath"`
}

// Use an AWS Lambda function as an event rule target.
//
// TODO: EXAMPLE
//
// Experimental.
type LambdaFunction interface {
	awsevents.IRuleTarget
	Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for LambdaFunction
type jsiiProxy_LambdaFunction struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewLambdaFunction(handler awslambda.IFunction, props *LambdaFunctionProps) LambdaFunction {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunction{}

	_jsii_.Create(
		"monocdk.aws_events_targets.LambdaFunction",
		[]interface{}{handler, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaFunction_Override(l LambdaFunction, handler awslambda.IFunction, props *LambdaFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.LambdaFunction",
		[]interface{}{handler, props},
		l,
	)
}

// Returns a RuleTarget that can be used to trigger this Lambda as a result from an EventBridge event.
// Experimental.
func (l *jsiiProxy_LambdaFunction) Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{rule, _id},
		&returns,
	)

	return returns
}

// Customize the Lambda Event Target.
//
// TODO: EXAMPLE
//
// Experimental.
type LambdaFunctionProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Experimental.
	MaxEventAge awscdk.Duration `json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
	// The event to send to the Lambda.
	//
	// This will be the payload sent to the Lambda Function.
	// Experimental.
	Event awsevents.RuleTargetInput `json:"event" yaml:"event"`
}

// Customize the CloudWatch LogGroup Event Target.
//
// TODO: EXAMPLE
//
// Experimental.
type LogGroupProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Experimental.
	MaxEventAge awscdk.Duration `json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
	// The event to send to the CloudWatch LogGroup.
	//
	// This will be the event logged into the CloudWatch LogGroup
	// Experimental.
	Event awsevents.RuleTargetInput `json:"event" yaml:"event"`
}

// Use a StepFunctions state machine as a target for Amazon EventBridge rules.
//
// TODO: EXAMPLE
//
// Experimental.
type SfnStateMachine interface {
	awsevents.IRuleTarget
	Machine() awsstepfunctions.IStateMachine
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for SfnStateMachine
type jsiiProxy_SfnStateMachine struct {
	internal.Type__awseventsIRuleTarget
}

func (j *jsiiProxy_SfnStateMachine) Machine() awsstepfunctions.IStateMachine {
	var returns awsstepfunctions.IStateMachine
	_jsii_.Get(
		j,
		"machine",
		&returns,
	)
	return returns
}


// Experimental.
func NewSfnStateMachine(machine awsstepfunctions.IStateMachine, props *SfnStateMachineProps) SfnStateMachine {
	_init_.Initialize()

	j := jsiiProxy_SfnStateMachine{}

	_jsii_.Create(
		"monocdk.aws_events_targets.SfnStateMachine",
		[]interface{}{machine, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSfnStateMachine_Override(s SfnStateMachine, machine awsstepfunctions.IStateMachine, props *SfnStateMachineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.SfnStateMachine",
		[]interface{}{machine, props},
		s,
	)
}

// Returns a properties that are used in an Rule to trigger this State Machine.
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/resource-based-policies-eventbridge.html#sns-permissions
//
// Experimental.
func (s *jsiiProxy_SfnStateMachine) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

// Customize the Step Functions State Machine target.
//
// TODO: EXAMPLE
//
// Experimental.
type SfnStateMachineProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Experimental.
	MaxEventAge awscdk.Duration `json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
	// The input to the state machine execution.
	// Experimental.
	Input awsevents.RuleTargetInput `json:"input" yaml:"input"`
	// The IAM role to be assumed to execute the State Machine.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// Use an SNS topic as a target for Amazon EventBridge rules.
//
// TODO: EXAMPLE
//
// Experimental.
type SnsTopic interface {
	awsevents.IRuleTarget
	Topic() awssns.ITopic
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for SnsTopic
type jsiiProxy_SnsTopic struct {
	internal.Type__awseventsIRuleTarget
}

func (j *jsiiProxy_SnsTopic) Topic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}


// Experimental.
func NewSnsTopic(topic awssns.ITopic, props *SnsTopicProps) SnsTopic {
	_init_.Initialize()

	j := jsiiProxy_SnsTopic{}

	_jsii_.Create(
		"monocdk.aws_events_targets.SnsTopic",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsTopic_Override(s SnsTopic, topic awssns.ITopic, props *SnsTopicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.SnsTopic",
		[]interface{}{topic, props},
		s,
	)
}

// Returns a RuleTarget that can be used to trigger this SNS topic as a result from an EventBridge event.
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/resource-based-policies-eventbridge.html#sns-permissions
//
// Experimental.
func (s *jsiiProxy_SnsTopic) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}

// Customize the SNS Topic Event Target.
//
// TODO: EXAMPLE
//
// Experimental.
type SnsTopicProps struct {
	// The message to send to the topic.
	// Experimental.
	Message awsevents.RuleTargetInput `json:"message" yaml:"message"`
}

// Use an SQS Queue as a target for Amazon EventBridge rules.
//
// TODO: EXAMPLE
//
// Experimental.
type SqsQueue interface {
	awsevents.IRuleTarget
	Queue() awssqs.IQueue
	Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for SqsQueue
type jsiiProxy_SqsQueue struct {
	internal.Type__awseventsIRuleTarget
}

func (j *jsiiProxy_SqsQueue) Queue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"queue",
		&returns,
	)
	return returns
}


// Experimental.
func NewSqsQueue(queue awssqs.IQueue, props *SqsQueueProps) SqsQueue {
	_init_.Initialize()

	j := jsiiProxy_SqsQueue{}

	_jsii_.Create(
		"monocdk.aws_events_targets.SqsQueue",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsQueue_Override(s SqsQueue, queue awssqs.IQueue, props *SqsQueueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.SqsQueue",
		[]interface{}{queue, props},
		s,
	)
}

// Returns a RuleTarget that can be used to trigger this SQS queue as a result from an EventBridge event.
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/resource-based-policies-eventbridge.html#sqs-permissions
//
// Experimental.
func (s *jsiiProxy_SqsQueue) Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{rule, _id},
		&returns,
	)

	return returns
}

// Customize the SQS Queue Event Target.
//
// TODO: EXAMPLE
//
// Experimental.
type SqsQueueProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Experimental.
	MaxEventAge awscdk.Duration `json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
	// The message to send to the queue.
	//
	// Must be a valid JSON text passed to the target queue.
	// Experimental.
	Message awsevents.RuleTargetInput `json:"message" yaml:"message"`
	// Message Group ID for messages sent to this queue.
	//
	// Required for FIFO queues, leave empty for regular queues.
	// Experimental.
	MessageGroupId *string `json:"messageGroupId" yaml:"messageGroupId"`
}

// The generic properties for an RuleTarget.
//
// TODO: EXAMPLE
//
// Experimental.
type TargetBaseProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Experimental.
	MaxEventAge awscdk.Duration `json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
}

// An environment variable to be set in the container run as a task.
//
// TODO: EXAMPLE
//
// Experimental.
type TaskEnvironmentVariable struct {
	// Name for the environment variable.
	//
	// Exactly one of `name` and `namePath` must be specified.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Value of the environment variable.
	//
	// Exactly one of `value` and `valuePath` must be specified.
	// Experimental.
	Value *string `json:"value" yaml:"value"`
}

