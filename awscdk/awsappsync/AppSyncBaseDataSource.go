package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Abstract AppSync datasource implementation.
//
// Do not use directly but use subclasses for concrete datasources.
type AppSyncBaseDataSource interface {
	constructs.Construct
	Api() IApi
	SetApi(val IApi)
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

// The jsii proxy struct for AppSyncBaseDataSource
type jsiiProxy_AppSyncBaseDataSource struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AppSyncBaseDataSource) Api() IApi {
	var returns IApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncBaseDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncBaseDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncBaseDataSource) Resource() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncBaseDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


func NewAppSyncBaseDataSource_Override(a AppSyncBaseDataSource, scope constructs.Construct, id *string, props *AppSyncBackedDataSourceProps, extended *AppSyncExtendedDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.AppSyncBaseDataSource",
		[]interface{}{scope, id, props, extended},
		a,
	)
}

func (j *jsiiProxy_AppSyncBaseDataSource)SetApi(val IApi) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_AppSyncBaseDataSource)SetServiceRole(val awsiam.IRole) {
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
func AppSyncBaseDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSyncBaseDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.AppSyncBaseDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncBaseDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

