package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a new API Gateway WebSocket API endpoint.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var messageHandler function
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
//   apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &WebSocketStageProps{
//   	WebSocketApi: WebSocketApi,
//   	StageName: jsii.String("dev"),
//   	AutoDeploy: jsii.Boolean(true),
//   })
//   webSocketApi.AddRoute(jsii.String("sendMessage"), &WebSocketRouteOptions{
//   	Integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("SendMessageIntegration"), messageHandler),
//   })
//
type WebSocketApi interface {
	awscdk.Resource
	IApi
	IWebSocketApi
	// The default endpoint for an API.
	ApiEndpoint() *string
	// The identifier of this API Gateway API.
	ApiId() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// A human friendly name for this WebSocket API.
	//
	// Note that this is different from `webSocketApiId`.
	WebSocketApiName() *string
	// Add a new route.
	AddRoute(routeKey *string, options *WebSocketRouteOptions) WebSocketRoute
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Get the "execute-api" ARN.
	// Deprecated: Use `arnForExecuteApiV2()` instead.
	ArnForExecuteApi(method *string, path *string, stage *string) *string
	// Get the "execute-api" ARN.
	// Default: - The default behavior applies when no specific route, or stage is provided.
	// In this case, the ARN will cover all routes, and all stages of this API.
	// Specifically, if 'route' is not specified, it defaults to '*', representing all routes.
	// If 'stage' is not specified, it also defaults to '*', representing all stages.
	//
	ArnForExecuteApiV2(route *string, stage *string) *string
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant access to the API Gateway management API for this WebSocket API to an IAM principal (Role/Group/User).
	GrantManageConnections(identity awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Api Gateway.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for WebSocketApi
type jsiiProxy_WebSocketApi struct {
	internal.Type__awscdkResource
	jsiiProxy_IApi
	jsiiProxy_IWebSocketApi
}

func (j *jsiiProxy_WebSocketApi) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketApi) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketApi) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketApi) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketApi) WebSocketApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webSocketApiName",
		&returns,
	)
	return returns
}


func NewWebSocketApi(scope constructs.Construct, id *string, props *WebSocketApiProps) WebSocketApi {
	_init_.Initialize()

	if err := validateNewWebSocketApiParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebSocketApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewWebSocketApi_Override(w WebSocketApi, scope constructs.Construct, id *string, props *WebSocketApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApi",
		[]interface{}{scope, id, props},
		w,
	)
}

// Import an existing WebSocket API into this CDK app.
func WebSocketApi_FromWebSocketApiAttributes(scope constructs.Construct, id *string, attrs *WebSocketApiAttributes) IWebSocketApi {
	_init_.Initialize()

	if err := validateWebSocketApi_FromWebSocketApiAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IWebSocketApi

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApi",
		"fromWebSocketApiAttributes",
		[]interface{}{scope, id, attrs},
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
func WebSocketApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWebSocketApi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func WebSocketApi_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateWebSocketApi_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApi",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func WebSocketApi_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateWebSocketApi_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApi",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func WebSocketApi_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApi",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WebSocketApi) AddRoute(routeKey *string, options *WebSocketRouteOptions) WebSocketRoute {
	if err := w.validateAddRouteParameters(routeKey, options); err != nil {
		panic(err)
	}
	var returns WebSocketRoute

	_jsii_.Invoke(
		w,
		"addRoute",
		[]interface{}{routeKey, options},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := w.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (w *jsiiProxy_WebSocketApi) ArnForExecuteApi(method *string, path *string, stage *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"arnForExecuteApi",
		[]interface{}{method, path, stage},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketApi) ArnForExecuteApiV2(route *string, stage *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"arnForExecuteApiV2",
		[]interface{}{route, stage},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketApi) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketApi) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := w.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketApi) GetResourceNameAttribute(nameAttr *string) *string {
	if err := w.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketApi) GrantManageConnections(identity awsiam.IGrantable) awsiam.Grant {
	if err := w.validateGrantManageConnectionsParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		w,
		"grantManageConnections",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketApi) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := w.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		w,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

