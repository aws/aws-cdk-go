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
// Do not use directly but use subclasses for resource backed datasources.
type BackedDataSource interface {
	BaseDataSource
	awsiam.IGrantable
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
}

// The jsii proxy struct for BackedDataSource
type jsiiProxy_BackedDataSource struct {
	jsiiProxy_BaseDataSource
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_BackedDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackedDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackedDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackedDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackedDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackedDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


func NewBackedDataSource_Override(b BackedDataSource, scope constructs.Construct, id *string, props *BackedDataSourceProps, extended *ExtendedDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.BackedDataSource",
		[]interface{}{scope, id, props, extended},
		b,
	)
}

func (j *jsiiProxy_BackedDataSource)SetApi(val IGraphqlApi) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_BackedDataSource)SetServiceRole(val awsiam.IRole) {
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
func BackedDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackedDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.BackedDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackedDataSource) CreateFunction(id *string, props *BaseAppsyncFunctionProps) AppsyncFunction {
	if err := b.validateCreateFunctionParameters(id, props); err != nil {
		panic(err)
	}
	var returns AppsyncFunction

	_jsii_.Invoke(
		b,
		"createFunction",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackedDataSource) CreateResolver(id *string, props *BaseResolverProps) Resolver {
	if err := b.validateCreateResolverParameters(id, props); err != nil {
		panic(err)
	}
	var returns Resolver

	_jsii_.Invoke(
		b,
		"createResolver",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackedDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

