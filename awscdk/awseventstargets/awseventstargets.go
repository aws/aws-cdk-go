package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use an API Destination rule target.
//
// Example:
//   connection := events.NewConnection(this, jsii.String("Connection"), &connectionProps{
//   	authorization: events.authorization.apiKey(jsii.String("x-api-key"), awscdk.SecretValue.secretsManager(jsii.String("ApiSecretName"))),
//   	description: jsii.String("Connection with API Key x-api-key"),
//   })
//
//   destination := events.NewApiDestination(this, jsii.String("Destination"), &apiDestinationProps{
//   	connection: connection,
//   	endpoint: jsii.String("https://example.com"),
//   	description: jsii.String("Calling example.com with API key x-api-key"),
//   })
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.minutes(jsii.Number(1))),
//   	targets: []iRuleTarget{
//   		targets.NewApiDestination(destination),
//   	},
//   })
//
type ApiDestination interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger API destinations from an EventBridge event.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for ApiDestination
type jsiiProxy_ApiDestination struct {
	internal.Type__awseventsIRuleTarget
}

func NewApiDestination(apiDestination awsevents.IApiDestination, props *ApiDestinationProps) ApiDestination {
	_init_.Initialize()

	j := jsiiProxy_ApiDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.ApiDestination",
		[]interface{}{apiDestination, props},
		&j,
	)

	return &j
}

func NewApiDestination_Override(a ApiDestination, apiDestination awsevents.IApiDestination, props *ApiDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.ApiDestination",
		[]interface{}{apiDestination, props},
		a,
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//   var role role
//   var ruleTargetInput ruleTargetInput
//
//   apiDestinationProps := &apiDestinationProps{
//   	deadLetterQueue: queue,
//   	event: ruleTargetInput,
//   	eventRole: role,
//   	headerParameters: map[string]*string{
//   		"headerParametersKey": jsii.String("headerParameters"),
//   	},
//   	maxEventAge: cdk.duration.minutes(jsii.Number(30)),
//   	pathParameterValues: []*string{
//   		jsii.String("pathParameterValues"),
//   	},
//   	queryStringParameters: map[string]*string{
//   		"queryStringParametersKey": jsii.String("queryStringParameters"),
//   	},
//   	retryAttempts: jsii.Number(123),
//   }
//
type ApiDestinationProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The event to send.
	Event awsevents.RuleTargetInput `field:"optional" json:"event" yaml:"event"`
	// The role to assume before invoking the target.
	EventRole awsiam.IRole `field:"optional" json:"eventRole" yaml:"eventRole"`
	// Additional headers sent to the API Destination.
	//
	// These are merged with headers specified on the Connection, with
	// the headers on the Connection taking precedence.
	//
	// You can only specify secret values on the Connection.
	HeaderParameters *map[string]*string `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// Path parameters to insert in place of path wildcards (`*`).
	//
	// If the API destination has a wilcard in the path, these path parts
	// will be inserted in that place.
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// Additional query string parameters sent to the API Destination.
	//
	// These are merged with headers specified on the Connection, with
	// the headers on the Connection taking precedence.
	//
	// You can only specify secret values on the Connection.
	QueryStringParameters *map[string]*string `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

// Use an API Gateway REST APIs as a target for Amazon EventBridge rules.
//
// Example:
//   import api "github.com/aws/aws-cdk-go/awscdk"
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.minutes(jsii.Number(1))),
//   })
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &functionProps{
//   	handler: jsii.String("index.handler"),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	code: lambda.code.fromInline(jsii.String("exports.handler = e => {}")),
//   })
//
//   restApi := api.NewLambdaRestApi(this, jsii.String("MyRestAPI"), &lambdaRestApiProps{
//   	handler: fn,
//   })
//
//   dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   rule.addTarget(
//   targets.NewApiGateway(restApi, &apiGatewayProps{
//   	path: jsii.String("/*/test"),
//   	method: jsii.String("GET"),
//   	stage: jsii.String("prod"),
//   	pathParameterValues: []*string{
//   		jsii.String("path-value"),
//   	},
//   	headerParameters: map[string]*string{
//   		"Header1": jsii.String("header1"),
//   	},
//   	queryStringParameters: map[string]*string{
//   		"QueryParam1": jsii.String("query-param-1"),
//   	},
//   	deadLetterQueue: dlq,
//   }))
//
type ApiGateway interface {
	awsevents.IRuleTarget
	RestApi() awsapigateway.RestApi
	// Returns a RuleTarget that can be used to trigger this API Gateway REST APIs as a result from an EventBridge event.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/resource-based-policies-eventbridge.html#sqs-permissions
	//
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


func NewApiGateway(restApi awsapigateway.RestApi, props *ApiGatewayProps) ApiGateway {
	_init_.Initialize()

	j := jsiiProxy_ApiGateway{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.ApiGateway",
		[]interface{}{restApi, props},
		&j,
	)

	return &j
}

func NewApiGateway_Override(a ApiGateway, restApi awsapigateway.RestApi, props *ApiGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.ApiGateway",
		[]interface{}{restApi, props},
		a,
	)
}

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
// Example:
//   import api "github.com/aws/aws-cdk-go/awscdk"
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.minutes(jsii.Number(1))),
//   })
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &functionProps{
//   	handler: jsii.String("index.handler"),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	code: lambda.code.fromInline(jsii.String("exports.handler = e => {}")),
//   })
//
//   restApi := api.NewLambdaRestApi(this, jsii.String("MyRestAPI"), &lambdaRestApiProps{
//   	handler: fn,
//   })
//
//   dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   rule.addTarget(
//   targets.NewApiGateway(restApi, &apiGatewayProps{
//   	path: jsii.String("/*/test"),
//   	method: jsii.String("GET"),
//   	stage: jsii.String("prod"),
//   	pathParameterValues: []*string{
//   		jsii.String("path-value"),
//   	},
//   	headerParameters: map[string]*string{
//   		"Header1": jsii.String("header1"),
//   	},
//   	queryStringParameters: map[string]*string{
//   		"QueryParam1": jsii.String("query-param-1"),
//   	},
//   	deadLetterQueue: dlq,
//   }))
//
type ApiGatewayProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The role to assume before invoking the target (i.e., the pipeline) when the given rule is triggered.
	EventRole awsiam.IRole `field:"optional" json:"eventRole" yaml:"eventRole"`
	// The headers to be set when requesting API.
	HeaderParameters *map[string]*string `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// The method for api resource invoked by the rule.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The api resource invoked by the rule.
	//
	// We can use wildcards('*') to specify the path. In that case,
	// an equal number of real values must be specified for pathParameterValues.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The path parameter values to be used to populate to wildcards("*") of requesting api path.
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// This will be the post request body send to the API.
	PostBody awsevents.RuleTargetInput `field:"optional" json:"postBody" yaml:"postBody"`
	// The query parameters to be set when requesting API.
	QueryStringParameters *map[string]*string `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
	// The deploy stage of api gateway invoked by the rule.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

// Use an AWS Lambda function that makes API calls as an event rule target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var policyStatement policyStatement
//
//   awsApi := awscdk.Aws_events_targets.NewAwsApi(&awsApiProps{
//   	action: jsii.String("action"),
//   	service: jsii.String("service"),
//
//   	// the properties below are optional
//   	apiVersion: jsii.String("apiVersion"),
//   	catchErrorPattern: jsii.String("catchErrorPattern"),
//   	parameters: parameters,
//   	policyStatement: policyStatement,
//   })
//
type AwsApi interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this AwsApi as a result from an EventBridge event.
	Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for AwsApi
type jsiiProxy_AwsApi struct {
	internal.Type__awseventsIRuleTarget
}

func NewAwsApi(props *AwsApiProps) AwsApi {
	_init_.Initialize()

	j := jsiiProxy_AwsApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.AwsApi",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAwsApi_Override(a AwsApi, props *AwsApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.AwsApi",
		[]interface{}{props},
		a,
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   awsApiInput := &awsApiInput{
//   	action: jsii.String("action"),
//   	service: jsii.String("service"),
//
//   	// the properties below are optional
//   	apiVersion: jsii.String("apiVersion"),
//   	catchErrorPattern: jsii.String("catchErrorPattern"),
//   	parameters: parameters,
//   }
//
type AwsApiInput struct {
	// The service action to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The service to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	Service *string `field:"required" json:"service" yaml:"service"`
	// API version to use for the service.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/locking-api-versions.html
	//
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The regex pattern to use to catch API errors.
	//
	// The `code` property of the
	// `Error` object will be tested against this pattern. If there is a match an
	// error will not be thrown.
	CatchErrorPattern *string `field:"optional" json:"catchErrorPattern" yaml:"catchErrorPattern"`
	// The parameters for the service action.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

// Properties for an AwsApi target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var policyStatement policyStatement
//
//   awsApiProps := &awsApiProps{
//   	action: jsii.String("action"),
//   	service: jsii.String("service"),
//
//   	// the properties below are optional
//   	apiVersion: jsii.String("apiVersion"),
//   	catchErrorPattern: jsii.String("catchErrorPattern"),
//   	parameters: parameters,
//   	policyStatement: policyStatement,
//   }
//
type AwsApiProps struct {
	// The service action to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The service to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	Service *string `field:"required" json:"service" yaml:"service"`
	// API version to use for the service.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/locking-api-versions.html
	//
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The regex pattern to use to catch API errors.
	//
	// The `code` property of the
	// `Error` object will be tested against this pattern. If there is a match an
	// error will not be thrown.
	CatchErrorPattern *string `field:"optional" json:"catchErrorPattern" yaml:"catchErrorPattern"`
	// The parameters for the service action.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The IAM policy statement to allow the API call.
	//
	// Use only if
	// resource restriction is needed.
	PolicyStatement awsiam.PolicyStatement `field:"optional" json:"policyStatement" yaml:"policyStatement"`
}

// Use an AWS Batch Job / Queue as an event rule target.
//
// Most likely the code will look something like this:
// `new BatchJob(jobQueue.jobQueueArn, jobQueue, jobDefinition.jobDefinitionArn, jobDefinition)`
//
// In the future this API will be improved to be fully typed.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import batch "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   jobQueue := batch.NewJobQueue(this, jsii.String("MyQueue"), map[string][]map[string]interface{}{
//   	"computeEnvironments": []map[string]interface{}{
//   		map[string]interface{}{
//   			"computeEnvironment": batch.NewComputeEnvironment(this, jsii.String("ComputeEnvironment"), map[string]*bool{
//   				"managed": jsii.Boolean(false),
//   			}),
//   			"order": jsii.Number(1),
//   		},
//   	},
//   })
//
//   jobDefinition := batch.NewJobDefinition(this, jsii.String("MyJob"), map[string]map[string]repositoryImage{
//   	"container": map[string]repositoryImage{
//   		"image": awscdk.ContainerImage.fromRegistry(jsii.String("test-repo")),
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.hours(jsii.Number(1))),
//   })
//
//   rule.addTarget(targets.NewBatchJob(jobQueue.jobQueueArn, jobQueue, jobDefinition.jobDefinitionArn, jobDefinition, &batchJobProps{
//   	deadLetterQueue: queue,
//   	event: events.ruleTargetInput.fromObject(map[string]*string{
//   		"SomeParam": jsii.String("SomeValue"),
//   	}),
//   	retryAttempts: jsii.Number(2),
//   	maxEventAge: cdk.*duration.hours(jsii.Number(2)),
//   }))
//
type BatchJob interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger queue this batch job as a result from an EventBridge event.
	Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for BatchJob
type jsiiProxy_BatchJob struct {
	internal.Type__awseventsIRuleTarget
}

func NewBatchJob(jobQueueArn *string, jobQueueScope constructs.IConstruct, jobDefinitionArn *string, jobDefinitionScope constructs.IConstruct, props *BatchJobProps) BatchJob {
	_init_.Initialize()

	j := jsiiProxy_BatchJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.BatchJob",
		[]interface{}{jobQueueArn, jobQueueScope, jobDefinitionArn, jobDefinitionScope, props},
		&j,
	)

	return &j
}

func NewBatchJob_Override(b BatchJob, jobQueueArn *string, jobQueueScope constructs.IConstruct, jobDefinitionArn *string, jobDefinitionScope constructs.IConstruct, props *BatchJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.BatchJob",
		[]interface{}{jobQueueArn, jobQueueScope, jobDefinitionArn, jobDefinitionScope, props},
		b,
	)
}

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
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import batch "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   jobQueue := batch.NewJobQueue(this, jsii.String("MyQueue"), map[string][]map[string]interface{}{
//   	"computeEnvironments": []map[string]interface{}{
//   		map[string]interface{}{
//   			"computeEnvironment": batch.NewComputeEnvironment(this, jsii.String("ComputeEnvironment"), map[string]*bool{
//   				"managed": jsii.Boolean(false),
//   			}),
//   			"order": jsii.Number(1),
//   		},
//   	},
//   })
//
//   jobDefinition := batch.NewJobDefinition(this, jsii.String("MyJob"), map[string]map[string]repositoryImage{
//   	"container": map[string]repositoryImage{
//   		"image": awscdk.ContainerImage.fromRegistry(jsii.String("test-repo")),
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.hours(jsii.Number(1))),
//   })
//
//   rule.addTarget(targets.NewBatchJob(jobQueue.jobQueueArn, jobQueue, jobDefinition.jobDefinitionArn, jobDefinition, &batchJobProps{
//   	deadLetterQueue: queue,
//   	event: events.ruleTargetInput.fromObject(map[string]*string{
//   		"SomeParam": jsii.String("SomeValue"),
//   	}),
//   	retryAttempts: jsii.Number(2),
//   	maxEventAge: cdk.*duration.hours(jsii.Number(2)),
//   }))
//
type BatchJobProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The number of times to attempt to retry, if the job fails.
	//
	// Valid values are 1–10.
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
	// The event to send to the Lambda.
	//
	// This will be the payload sent to the Lambda Function.
	Event awsevents.RuleTargetInput `field:"optional" json:"event" yaml:"event"`
	// The name of the submitted job.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// The size of the array, if this is an array batch job.
	//
	// Valid values are integers between 2 and 10,000.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

// Use an AWS CloudWatch LogGroup as an event rule target.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   var logGroup logGroup
//   var rule rule
//
//
//   rule.addTarget(targets.NewCloudWatchLogGroup(logGroup, &logGroupProps{
//   	logEvent: targets.logGroupTargetInput(map[string]*string{
//   		"message": JSON.stringify(map[string]*string{
//   			"CustomField": jsii.String("CustomValue"),
//   		}),
//   	}),
//   }))
//
type CloudWatchLogGroup interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to log an event into a CloudWatch LogGroup.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for CloudWatchLogGroup
type jsiiProxy_CloudWatchLogGroup struct {
	internal.Type__awseventsIRuleTarget
}

func NewCloudWatchLogGroup(logGroup awslogs.ILogGroup, props *LogGroupProps) CloudWatchLogGroup {
	_init_.Initialize()

	j := jsiiProxy_CloudWatchLogGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.CloudWatchLogGroup",
		[]interface{}{logGroup, props},
		&j,
	)

	return &j
}

func NewCloudWatchLogGroup_Override(c CloudWatchLogGroup, logGroup awslogs.ILogGroup, props *LogGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.CloudWatchLogGroup",
		[]interface{}{logGroup, props},
		c,
	)
}

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
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//
//   var repo repository
//   var project pipelineProject
//   var myTopic topic
//
//
//   // starts a CodeBuild project when a commit is pushed to the "main" branch of the repo
//   repo.onCommit(jsii.String("CommitToMain"), &onCommitOptions{
//   	target: targets.NewCodeBuildProject(project),
//   	branches: []*string{
//   		jsii.String("main"),
//   	},
//   })
//
//   // publishes a message to an Amazon SNS topic when a comment is made on a pull request
//   rule := repo.onCommentOnPullRequest(jsii.String("CommentOnPullRequest"), &onEventOptions{
//   	target: targets.NewSnsTopic(myTopic),
//   })
//
type CodeBuildProject interface {
	awsevents.IRuleTarget
	// Allows using build projects as event rule targets.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for CodeBuildProject
type jsiiProxy_CodeBuildProject struct {
	internal.Type__awseventsIRuleTarget
}

func NewCodeBuildProject(project awscodebuild.IProject, props *CodeBuildProjectProps) CodeBuildProject {
	_init_.Initialize()

	j := jsiiProxy_CodeBuildProject{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.CodeBuildProject",
		[]interface{}{project, props},
		&j,
	)

	return &j
}

func NewCodeBuildProject_Override(c CodeBuildProject, project awscodebuild.IProject, props *CodeBuildProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.CodeBuildProject",
		[]interface{}{project, props},
		c,
	)
}

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
// Example:
//   import codebuild "github.com/aws/aws-cdk-go/awscdk"
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//
//
//   repo := codecommit.NewRepository(this, jsii.String("MyRepo"), &repositoryProps{
//   	repositoryName: jsii.String("aws-cdk-codebuild-events"),
//   })
//
//   project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
//   	source: codebuild.source.codeCommit(&codeCommitSourceProps{
//   		repository: repo,
//   	}),
//   })
//
//   deadLetterQueue := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   // trigger a build when a commit is pushed to the repo
//   onCommitRule := repo.onCommit(jsii.String("OnCommit"), &onCommitOptions{
//   	target: targets.NewCodeBuildProject(project, &codeBuildProjectProps{
//   		deadLetterQueue: deadLetterQueue,
//   	}),
//   	branches: []*string{
//   		jsii.String("master"),
//   	},
//   })
//
type CodeBuildProjectProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The event to send to CodeBuild.
	//
	// This will be the payload for the StartBuild API.
	Event awsevents.RuleTargetInput `field:"optional" json:"event" yaml:"event"`
	// The role to assume before invoking the target (i.e., the codebuild) when the given rule is triggered.
	EventRole awsiam.IRole `field:"optional" json:"eventRole" yaml:"eventRole"`
}

// Allows the pipeline to be used as an EventBridge rule target.
//
// Example:
//   // A pipeline being used as a target for a CloudWatch event rule.
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//   import events "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipeline pipeline
//
//
//   // kick off the pipeline every day
//   rule := events.NewRule(this, jsii.String("Daily"), &ruleProps{
//   	schedule: events.schedule.rate(awscdk.Duration.days(jsii.Number(1))),
//   })
//   rule.addTarget(targets.NewCodePipeline(pipeline))
//
type CodePipeline interface {
	awsevents.IRuleTarget
	// Returns the rule target specification.
	//
	// NOTE: Do not use the various `inputXxx` options. They can be set in a call to `addTarget`.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for CodePipeline
type jsiiProxy_CodePipeline struct {
	internal.Type__awseventsIRuleTarget
}

func NewCodePipeline(pipeline awscodepipeline.IPipeline, options *CodePipelineTargetOptions) CodePipeline {
	_init_.Initialize()

	j := jsiiProxy_CodePipeline{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.CodePipeline",
		[]interface{}{pipeline, options},
		&j,
	)

	return &j
}

func NewCodePipeline_Override(c CodePipeline, pipeline awscodepipeline.IPipeline, options *CodePipelineTargetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.CodePipeline",
		[]interface{}{pipeline, options},
		c,
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//   var role role
//
//   codePipelineTargetOptions := &codePipelineTargetOptions{
//   	deadLetterQueue: queue,
//   	eventRole: role,
//   	maxEventAge: cdk.duration.minutes(jsii.Number(30)),
//   	retryAttempts: jsii.Number(123),
//   }
//
type CodePipelineTargetOptions struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The role to assume before invoking the target (i.e., the pipeline) when the given rule is triggered.
	EventRole awsiam.IRole `field:"optional" json:"eventRole" yaml:"eventRole"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerOverride := &containerOverride{
//   	containerName: jsii.String("containerName"),
//
//   	// the properties below are optional
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	cpu: jsii.Number(123),
//   	environment: []taskEnvironmentVariable{
//   		&taskEnvironmentVariable{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	memoryLimit: jsii.Number(123),
//   	memoryReservation: jsii.Number(123),
//   }
//
type ContainerOverride struct {
	// Name of the container inside the task definition.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Command to run inside the container.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The number of cpu units reserved for the container.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Variables to set in the container's environment.
	Environment *[]*TaskEnvironmentVariable `field:"optional" json:"environment" yaml:"environment"`
	// Hard memory limit on the container.
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// Soft memory limit on the container.
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
}

// Start a task on an ECS cluster.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var role role
//
//
//   ecsTaskTarget := awscdk.NewEcsTask(&ecsTaskProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	role: role,
//   })
//
//   awscdk.NewRule(this, jsii.String("ScheduleRule"), &ruleProps{
//   	schedule: awscdk.Schedule.cron(&cronOptions{
//   		minute: jsii.String("0"),
//   		hour: jsii.String("4"),
//   	}),
//   	targets: []iRuleTarget{
//   		ecsTaskTarget,
//   	},
//   })
//
type EcsTask interface {
	awsevents.IRuleTarget
	// The security groups associated with the task.
	//
	// Only applicable with awsvpc network mode.
	SecurityGroups() *[]awsec2.ISecurityGroup
	// Allows using tasks as target of EventBridge events.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for EcsTask
type jsiiProxy_EcsTask struct {
	internal.Type__awseventsIRuleTarget
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


func NewEcsTask(props *EcsTaskProps) EcsTask {
	_init_.Initialize()

	j := jsiiProxy_EcsTask{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.EcsTask",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewEcsTask_Override(e EcsTask, props *EcsTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.EcsTask",
		[]interface{}{props},
		e,
	)
}

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
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var role role
//
//
//   ecsTaskTarget := awscdk.NewEcsTask(&ecsTaskProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	role: role,
//   })
//
//   awscdk.NewRule(this, jsii.String("ScheduleRule"), &ruleProps{
//   	schedule: awscdk.Schedule.cron(&cronOptions{
//   		minute: jsii.String("0"),
//   		hour: jsii.String("4"),
//   	}),
//   	targets: []iRuleTarget{
//   		ecsTaskTarget,
//   	},
//   })
//
type EcsTaskProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Cluster where service will be deployed.
	Cluster awsecs.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// Task Definition of the task that should be started.
	TaskDefinition awsecs.ITaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// Container setting overrides.
	//
	// Key is the name of the container to override, value is the
	// values you want to override.
	ContainerOverrides *[]*ContainerOverride `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// The platform version on which to run your task.
	//
	// Unless you have specific compatibility requirements, you don't need to specify this.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html
	//
	PlatformVersion awsecs.FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Existing IAM role to run the ECS task.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Existing security groups to use for the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking).
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// In what subnets to place the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking).
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// How many tasks should be started when this event is triggered.
	TaskCount *float64 `field:"optional" json:"taskCount" yaml:"taskCount"`
}

// Notify an existing Event Bus of an event.
//
// Example:
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.expression(jsii.String("rate(1 minute)")),
//   })
//
//   rule.addTarget(targets.NewEventBus(events.eventBus.fromEventBusArn(this, jsii.String("External"), jsii.String("arn:aws:events:eu-west-1:999999999999:event-bus/test-bus"))))
//
type EventBus interface {
	awsevents.IRuleTarget
	// Returns the rule target specification.
	//
	// NOTE: Do not use the various `inputXxx` options. They can be set in a call to `addTarget`.
	Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for EventBus
type jsiiProxy_EventBus struct {
	internal.Type__awseventsIRuleTarget
}

func NewEventBus(eventBus awsevents.IEventBus, props *EventBusProps) EventBus {
	_init_.Initialize()

	j := jsiiProxy_EventBus{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.EventBus",
		[]interface{}{eventBus, props},
		&j,
	)

	return &j
}

func NewEventBus_Override(e EventBus, eventBus awsevents.IEventBus, props *EventBusProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.EventBus",
		[]interface{}{eventBus, props},
		e,
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//   var role role
//
//   eventBusProps := &eventBusProps{
//   	deadLetterQueue: queue,
//   	role: role,
//   }
//
type EventBusProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Role to be used to publish the event.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

// Customize the Firehose Stream Event Target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnDeliveryStream cfnDeliveryStream
//   var ruleTargetInput ruleTargetInput
//
//   kinesisFirehoseStream := awscdk.Aws_events_targets.NewKinesisFirehoseStream(cfnDeliveryStream, &kinesisFirehoseStreamProps{
//   	message: ruleTargetInput,
//   })
//
type KinesisFirehoseStream interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this Firehose Stream as a result from a Event Bridge event.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for KinesisFirehoseStream
type jsiiProxy_KinesisFirehoseStream struct {
	internal.Type__awseventsIRuleTarget
}

func NewKinesisFirehoseStream(stream awskinesisfirehose.CfnDeliveryStream, props *KinesisFirehoseStreamProps) KinesisFirehoseStream {
	_init_.Initialize()

	j := jsiiProxy_KinesisFirehoseStream{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.KinesisFirehoseStream",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

func NewKinesisFirehoseStream_Override(k KinesisFirehoseStream, stream awskinesisfirehose.CfnDeliveryStream, props *KinesisFirehoseStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.KinesisFirehoseStream",
		[]interface{}{stream, props},
		k,
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var ruleTargetInput ruleTargetInput
//
//   kinesisFirehoseStreamProps := &kinesisFirehoseStreamProps{
//   	message: ruleTargetInput,
//   }
//
type KinesisFirehoseStreamProps struct {
	// The message to send to the stream.
	//
	// Must be a valid JSON text passed to the target stream.
	Message awsevents.RuleTargetInput `field:"optional" json:"message" yaml:"message"`
}

// Use a Kinesis Stream as a target for AWS CloudWatch event rules.
//
// Example:
//   // put to a Kinesis stream every time code is committed
//   // to a CodeCommit repository
//   repository.onCommit(jsii.String("onCommit"), &onCommitOptions{
//   	target: targets.NewKinesisStream(stream),
//   })
//
type KinesisStream interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this Kinesis Stream as a result from a CloudWatch event.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for KinesisStream
type jsiiProxy_KinesisStream struct {
	internal.Type__awseventsIRuleTarget
}

func NewKinesisStream(stream awskinesis.IStream, props *KinesisStreamProps) KinesisStream {
	_init_.Initialize()

	j := jsiiProxy_KinesisStream{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.KinesisStream",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

func NewKinesisStream_Override(k KinesisStream, stream awskinesis.IStream, props *KinesisStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.KinesisStream",
		[]interface{}{stream, props},
		k,
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var ruleTargetInput ruleTargetInput
//
//   kinesisStreamProps := &kinesisStreamProps{
//   	message: ruleTargetInput,
//   	partitionKeyPath: jsii.String("partitionKeyPath"),
//   }
//
type KinesisStreamProps struct {
	// The message to send to the stream.
	//
	// Must be a valid JSON text passed to the target stream.
	Message awsevents.RuleTargetInput `field:"optional" json:"message" yaml:"message"`
	// Partition Key Path for records sent to this stream.
	PartitionKeyPath *string `field:"optional" json:"partitionKeyPath" yaml:"partitionKeyPath"`
}

// Use an AWS Lambda function as an event rule target.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("exports.handler = handler.toString()")),
//   })
//
//   rule := events.NewRule(this, jsii.String("rule"), &ruleProps{
//   	eventPattern: &eventPattern{
//   		source: []*string{
//   			jsii.String("aws.ec2"),
//   		},
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule.addTarget(targets.NewLambdaFunction(fn, &lambdaFunctionProps{
//   	deadLetterQueue: queue,
//   	 // Optional: add a dead letter queue
//   	maxEventAge: cdk.duration.hours(jsii.Number(2)),
//   	 // Optional: set the maxEventAge retry policy
//   	retryAttempts: jsii.Number(2),
//   }))
//
type LambdaFunction interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this Lambda as a result from an EventBridge event.
	Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for LambdaFunction
type jsiiProxy_LambdaFunction struct {
	internal.Type__awseventsIRuleTarget
}

func NewLambdaFunction(handler awslambda.IFunction, props *LambdaFunctionProps) LambdaFunction {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.LambdaFunction",
		[]interface{}{handler, props},
		&j,
	)

	return &j
}

func NewLambdaFunction_Override(l LambdaFunction, handler awslambda.IFunction, props *LambdaFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.LambdaFunction",
		[]interface{}{handler, props},
		l,
	)
}

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
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("exports.handler = handler.toString()")),
//   })
//
//   rule := events.NewRule(this, jsii.String("rule"), &ruleProps{
//   	eventPattern: &eventPattern{
//   		source: []*string{
//   			jsii.String("aws.ec2"),
//   		},
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule.addTarget(targets.NewLambdaFunction(fn, &lambdaFunctionProps{
//   	deadLetterQueue: queue,
//   	 // Optional: add a dead letter queue
//   	maxEventAge: cdk.duration.hours(jsii.Number(2)),
//   	 // Optional: set the maxEventAge retry policy
//   	retryAttempts: jsii.Number(2),
//   }))
//
type LambdaFunctionProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The event to send to the Lambda.
	//
	// This will be the payload sent to the Lambda Function.
	Event awsevents.RuleTargetInput `field:"optional" json:"event" yaml:"event"`
}

// Customize the CloudWatch LogGroup Event Target.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   var logGroup logGroup
//   var rule rule
//
//
//   rule.addTarget(targets.NewCloudWatchLogGroup(logGroup, &logGroupProps{
//   	logEvent: targets.logGroupTargetInput(map[string]interface{}{
//   		"timestamp": events.EventField.from(jsii.String("$.time")),
//   		"message": events.EventField.from(jsii.String("$.detail-type")),
//   	}),
//   }))
//
type LogGroupProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The event to send to the CloudWatch LogGroup.
	//
	// This will be the event logged into the CloudWatch LogGroup.
	// Deprecated: use logEvent instead.
	Event awsevents.RuleTargetInput `field:"optional" json:"event" yaml:"event"`
	// The event to send to the CloudWatch LogGroup.
	//
	// This will be the event logged into the CloudWatch LogGroup.
	LogEvent LogGroupTargetInput `field:"optional" json:"logEvent" yaml:"logEvent"`
}

// The input to send to the CloudWatch LogGroup target.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   var logGroup logGroup
//   var rule rule
//
//
//   rule.addTarget(targets.NewCloudWatchLogGroup(logGroup, &logGroupProps{
//   	logEvent: targets.logGroupTargetInput(map[string]interface{}{
//   		"timestamp": events.EventField.from(jsii.String("$.time")),
//   		"message": events.EventField.from(jsii.String("$.detail-type")),
//   	}),
//   }))
//
type LogGroupTargetInput interface {
	// Return the input properties for this input object.
	Bind(rule awsevents.IRule) *awsevents.RuleTargetInputProperties
}

// The jsii proxy struct for LogGroupTargetInput
type jsiiProxy_LogGroupTargetInput struct {
	_ byte // padding
}

func NewLogGroupTargetInput_Override(l LogGroupTargetInput) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.LogGroupTargetInput",
		nil, // no parameters
		l,
	)
}

// Pass a JSON object to the the log group event target.
//
// May contain strings returned by `EventField.from()` to substitute in parts of the
// matched event.
func LogGroupTargetInput_FromObject(options *LogGroupTargetInputOptions) awsevents.RuleTargetInput {
	_init_.Initialize()

	var returns awsevents.RuleTargetInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events_targets.LogGroupTargetInput",
		"fromObject",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroupTargetInput) Bind(rule awsevents.IRule) *awsevents.RuleTargetInputProperties {
	var returns *awsevents.RuleTargetInputProperties

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// Options used when creating a target input template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var message interface{}
//   var timestamp interface{}
//
//   logGroupTargetInputOptions := &logGroupTargetInputOptions{
//   	message: message,
//   	timestamp: timestamp,
//   }
//
type LogGroupTargetInputOptions struct {
	// The value provided here will be used in the Log "message" field.
	//
	// This field must be a string. If an object is passed (e.g. JSON data)
	// it will not throw an error, but the message that makes it to
	// CloudWatch logs will be incorrect. This is a likely scenario if
	// doing something like: EventField.fromPath('$.detail') since in most cases
	// the `detail` field contains JSON data.
	Message interface{} `field:"optional" json:"message" yaml:"message"`
	// The timestamp that will appear in the CloudWatch Logs record.
	Timestamp interface{} `field:"optional" json:"timestamp" yaml:"timestamp"`
}

// Use a StepFunctions state machine as a target for Amazon EventBridge rules.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   import sfn "github.com/aws/aws-cdk-go/awscdk"
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.minutes(jsii.Number(1))),
//   })
//
//   dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   role := iam.NewRole(this, jsii.String("Role"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("events.amazonaws.com")),
//   })
//   stateMachine := sfn.NewStateMachine(this, jsii.String("SM"), &stateMachineProps{
//   	definition: sfn.NewWait(this, jsii.String("Hello"), &waitProps{
//   		time: sfn.waitTime.duration(cdk.*duration.seconds(jsii.Number(10))),
//   	}),
//   })
//
//   rule.addTarget(targets.NewSfnStateMachine(stateMachine, &sfnStateMachineProps{
//   	input: events.ruleTargetInput.fromObject(map[string]*string{
//   		"SomeParam": jsii.String("SomeValue"),
//   	}),
//   	deadLetterQueue: dlq,
//   	role: role,
//   }))
//
type SfnStateMachine interface {
	awsevents.IRuleTarget
	Machine() awsstepfunctions.IStateMachine
	// Returns a properties that are used in an Rule to trigger this State Machine.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/resource-based-policies-eventbridge.html#sns-permissions
	//
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


func NewSfnStateMachine(machine awsstepfunctions.IStateMachine, props *SfnStateMachineProps) SfnStateMachine {
	_init_.Initialize()

	j := jsiiProxy_SfnStateMachine{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.SfnStateMachine",
		[]interface{}{machine, props},
		&j,
	)

	return &j
}

func NewSfnStateMachine_Override(s SfnStateMachine, machine awsstepfunctions.IStateMachine, props *SfnStateMachineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.SfnStateMachine",
		[]interface{}{machine, props},
		s,
	)
}

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
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   import sfn "github.com/aws/aws-cdk-go/awscdk"
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.minutes(jsii.Number(1))),
//   })
//
//   dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   role := iam.NewRole(this, jsii.String("Role"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("events.amazonaws.com")),
//   })
//   stateMachine := sfn.NewStateMachine(this, jsii.String("SM"), &stateMachineProps{
//   	definition: sfn.NewWait(this, jsii.String("Hello"), &waitProps{
//   		time: sfn.waitTime.duration(cdk.*duration.seconds(jsii.Number(10))),
//   	}),
//   })
//
//   rule.addTarget(targets.NewSfnStateMachine(stateMachine, &sfnStateMachineProps{
//   	input: events.ruleTargetInput.fromObject(map[string]*string{
//   		"SomeParam": jsii.String("SomeValue"),
//   	}),
//   	deadLetterQueue: dlq,
//   	role: role,
//   }))
//
type SfnStateMachineProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The input to the state machine execution.
	Input awsevents.RuleTargetInput `field:"optional" json:"input" yaml:"input"`
	// The IAM role to be assumed to execute the State Machine.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

// Use an SNS topic as a target for Amazon EventBridge rules.
//
// Example:
//   // publish to an SNS topic every time code is committed
//   // to a CodeCommit repository
//   repository.onCommit(jsii.String("onCommit"), &onCommitOptions{
//   	target: targets.NewSnsTopic(topic),
//   })
//
type SnsTopic interface {
	awsevents.IRuleTarget
	Topic() awssns.ITopic
	// Returns a RuleTarget that can be used to trigger this SNS topic as a result from an EventBridge event.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/resource-based-policies-eventbridge.html#sns-permissions
	//
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


func NewSnsTopic(topic awssns.ITopic, props *SnsTopicProps) SnsTopic {
	_init_.Initialize()

	j := jsiiProxy_SnsTopic{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.SnsTopic",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

func NewSnsTopic_Override(s SnsTopic, topic awssns.ITopic, props *SnsTopicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.SnsTopic",
		[]interface{}{topic, props},
		s,
	)
}

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
// Example:
//   var onCommitRule rule
//   var topic topic
//
//
//   onCommitRule.addTarget(targets.NewSnsTopic(topic, &snsTopicProps{
//   	message: events.ruleTargetInput.fromText(
//   	fmt.Sprintf("A commit was pushed to the repository %v on branch %v", codecommit.referenceEvent.repositoryName, codecommit.*referenceEvent.referenceName)),
//   }))
//
type SnsTopicProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The message to send to the topic.
	Message awsevents.RuleTargetInput `field:"optional" json:"message" yaml:"message"`
}

// Use an SQS Queue as a target for Amazon EventBridge rules.
//
// Example:
//   // publish to an SQS queue every time code is committed
//   // to a CodeCommit repository
//   repository.onCommit(jsii.String("onCommit"), &onCommitOptions{
//   	target: targets.NewSqsQueue(queue),
//   })
//
type SqsQueue interface {
	awsevents.IRuleTarget
	Queue() awssqs.IQueue
	// Returns a RuleTarget that can be used to trigger this SQS queue as a result from an EventBridge event.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/resource-based-policies-eventbridge.html#sqs-permissions
	//
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


func NewSqsQueue(queue awssqs.IQueue, props *SqsQueueProps) SqsQueue {
	_init_.Initialize()

	j := jsiiProxy_SqsQueue{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.SqsQueue",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

func NewSqsQueue_Override(s SqsQueue, queue awssqs.IQueue, props *SqsQueueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.SqsQueue",
		[]interface{}{queue, props},
		s,
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//   var ruleTargetInput ruleTargetInput
//
//   sqsQueueProps := &sqsQueueProps{
//   	deadLetterQueue: queue,
//   	maxEventAge: cdk.duration.minutes(jsii.Number(30)),
//   	message: ruleTargetInput,
//   	messageGroupId: jsii.String("messageGroupId"),
//   	retryAttempts: jsii.Number(123),
//   }
//
type SqsQueueProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The message to send to the queue.
	//
	// Must be a valid JSON text passed to the target queue.
	Message awsevents.RuleTargetInput `field:"optional" json:"message" yaml:"message"`
	// Message Group ID for messages sent to this queue.
	//
	// Required for FIFO queues, leave empty for regular queues.
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

// The generic properties for an RuleTarget.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//
//   targetBaseProps := &targetBaseProps{
//   	deadLetterQueue: queue,
//   	maxEventAge: cdk.duration.minutes(jsii.Number(30)),
//   	retryAttempts: jsii.Number(123),
//   }
//
type TargetBaseProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
}

// An environment variable to be set in the container run as a task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskEnvironmentVariable := &taskEnvironmentVariable{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type TaskEnvironmentVariable struct {
	// Name for the environment variable.
	//
	// Exactly one of `name` and `namePath` must be specified.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the environment variable.
	//
	// Exactly one of `value` and `valuePath` must be specified.
	Value *string `field:"required" json:"value" yaml:"value"`
}

