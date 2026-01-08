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
//   api := appsync.NewEventApi(this, jsii.String("EventApiEventBridge"), &EventApiProps{
//   	ApiName: jsii.String("EventBridgeEventApi"),
//   })
//
//   eventBus := events.NewEventBus(this, jsii.String("test-bus"))
//
//   dataSource := api.AddEventBridgeDataSource(jsii.String("eventbridgeds"), eventBus)
//
type AppSyncEventBridgeDataSource interface {
	AppSyncBackedDataSource
	// The API this data source is attached to Set the API this data source is attached to.
	Api() IApi
	SetApi(val IApi)
	// The principal of the data source to be IGrantable.
	GrantPrincipal() awsiam.IPrincipal
	// The name of the data source.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// The underlying CFN data source resource.
	Resource() CfnDataSource
	ServiceRole() awsiam.IRole
	SetServiceRole(val awsiam.IRole)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AppSyncEventBridgeDataSource
type jsiiProxy_AppSyncEventBridgeDataSource struct {
	jsiiProxy_AppSyncBackedDataSource
}

func (j *jsiiProxy_AppSyncEventBridgeDataSource) Api() IApi {
	var returns IApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncEventBridgeDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncEventBridgeDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncEventBridgeDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncEventBridgeDataSource) Resource() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncEventBridgeDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


func NewAppSyncEventBridgeDataSource(scope constructs.Construct, id *string, props *AppSyncEventBridgeDataSourceProps) AppSyncEventBridgeDataSource {
	_init_.Initialize()

	if err := validateNewAppSyncEventBridgeDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSyncEventBridgeDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.AppSyncEventBridgeDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAppSyncEventBridgeDataSource_Override(a AppSyncEventBridgeDataSource, scope constructs.Construct, id *string, props *AppSyncEventBridgeDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.AppSyncEventBridgeDataSource",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AppSyncEventBridgeDataSource)SetApi(val IApi) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_AppSyncEventBridgeDataSource)SetServiceRole(val awsiam.IRole) {
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
func AppSyncEventBridgeDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSyncEventBridgeDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.AppSyncEventBridgeDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncEventBridgeDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

