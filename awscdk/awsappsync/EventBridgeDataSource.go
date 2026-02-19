package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AppSync datasource backed by EventBridge.
//
// Example:
//   import events "github.com/aws/aws-cdk-go/awscdk"
//
//
//   api := appsync.NewGraphqlApi(this, jsii.String("EventBridgeApi"), &GraphqlApiProps{
//   	Name: jsii.String("EventBridgeApi"),
//   	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.eventbridge.graphql"))),
//   })
//
//   bus := events.NewEventBus(this, jsii.String("DestinationEventBus"), &EventBusProps{
//   })
//
//   dataSource := api.AddEventBridgeDataSource(jsii.String("NoneDS"), bus)
//
//   dataSource.CreateResolver(jsii.String("EventResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("emitEvent"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("request.vtl")),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("response.vtl")),
//   })
//
type EventBridgeDataSource interface {
	BackedDataSource
	// The API this data source is attached to Set the API this data source is attached to.
	Api() IGraphqlApi
	SetApi(val IGraphqlApi)
	// the underlying CFN data source resource.
	Ds() CfnDataSource
	// the principal of the data source to be IGrantable.
	GrantPrincipal() awsiam.IPrincipal
	// the name of the data source.
	Name() *string
	// The tree node.
	Node() constructs.Node
	ServiceRole() awsiam.IRole
	SetServiceRole(val awsiam.IRole)
	// creates a new appsync function for this datasource and API using the given properties.
	CreateFunction(id *string, props *BaseAppsyncFunctionProps) AppsyncFunction
	// creates a new resolver for this datasource and API using the given properties.
	CreateResolver(id *string, props *BaseResolverProps) Resolver
	// Returns a string representation of this construct.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for EventBridgeDataSource
type jsiiProxy_EventBridgeDataSource struct {
	jsiiProxy_BackedDataSource
}

func (j *jsiiProxy_EventBridgeDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgeDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgeDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgeDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgeDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgeDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


func NewEventBridgeDataSource(scope constructs.Construct, id *string, props *EventBridgeDataSourceProps) EventBridgeDataSource {
	_init_.Initialize()

	if err := validateNewEventBridgeDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventBridgeDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.EventBridgeDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEventBridgeDataSource_Override(e EventBridgeDataSource, scope constructs.Construct, id *string, props *EventBridgeDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.EventBridgeDataSource",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EventBridgeDataSource)SetApi(val IGraphqlApi) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_EventBridgeDataSource)SetServiceRole(val awsiam.IRole) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
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
func EventBridgeDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventBridgeDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.EventBridgeDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EventBridgeDataSource_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.EventBridgeDataSource",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EventBridgeDataSource) CreateFunction(id *string, props *BaseAppsyncFunctionProps) AppsyncFunction {
	if err := e.validateCreateFunctionParameters(id, props); err != nil {
		panic(err)
	}
	var returns AppsyncFunction

	_jsii_.Invoke(
		e,
		"createFunction",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventBridgeDataSource) CreateResolver(id *string, props *BaseResolverProps) Resolver {
	if err := e.validateCreateResolverParameters(id, props); err != nil {
		panic(err)
	}
	var returns Resolver

	_jsii_.Invoke(
		e,
		"createResolver",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventBridgeDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventBridgeDataSource) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

