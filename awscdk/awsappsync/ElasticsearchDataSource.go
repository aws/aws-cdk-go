package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Appsync datasource backed by Elasticsearch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var domain domain
//   var graphqlApi graphqlApi
//   var role role
//
//   elasticsearchDataSource := awscdk.Aws_appsync.NewElasticsearchDataSource(this, jsii.String("MyElasticsearchDataSource"), &ElasticsearchDataSourceProps{
//   	Api: graphqlApi,
//   	Domain: domain,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ServiceRole: role,
//   })
//
// Deprecated: - use `OpenSearchDataSource`.
type ElasticsearchDataSource interface {
	BackedDataSource
	// Deprecated: - use `OpenSearchDataSource`.
	Api() IGraphqlApi
	// Deprecated: - use `OpenSearchDataSource`.
	SetApi(val IGraphqlApi)
	// the underlying CFN data source resource.
	// Deprecated: - use `OpenSearchDataSource`.
	Ds() CfnDataSource
	// the principal of the data source to be IGrantable.
	// Deprecated: - use `OpenSearchDataSource`.
	GrantPrincipal() awsiam.IPrincipal
	// the name of the data source.
	// Deprecated: - use `OpenSearchDataSource`.
	Name() *string
	// The tree node.
	// Deprecated: - use `OpenSearchDataSource`.
	Node() constructs.Node
	// Deprecated: - use `OpenSearchDataSource`.
	ServiceRole() awsiam.IRole
	// Deprecated: - use `OpenSearchDataSource`.
	SetServiceRole(val awsiam.IRole)
	// creates a new appsync function for this datasource and API using the given properties.
	// Deprecated: - use `OpenSearchDataSource`.
	CreateFunction(id *string, props *BaseAppsyncFunctionProps) AppsyncFunction
	// creates a new resolver for this datasource and API using the given properties.
	// Deprecated: - use `OpenSearchDataSource`.
	CreateResolver(id *string, props *BaseResolverProps) Resolver
	// Returns a string representation of this construct.
	// Deprecated: - use `OpenSearchDataSource`.
	ToString() *string
}

// The jsii proxy struct for ElasticsearchDataSource
type jsiiProxy_ElasticsearchDataSource struct {
	jsiiProxy_BackedDataSource
}

func (j *jsiiProxy_ElasticsearchDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


// Deprecated: - use `OpenSearchDataSource`.
func NewElasticsearchDataSource(scope constructs.Construct, id *string, props *ElasticsearchDataSourceProps) ElasticsearchDataSource {
	_init_.Initialize()

	if err := validateNewElasticsearchDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticsearchDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.ElasticsearchDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: - use `OpenSearchDataSource`.
func NewElasticsearchDataSource_Override(e ElasticsearchDataSource, scope constructs.Construct, id *string, props *ElasticsearchDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.ElasticsearchDataSource",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDataSource)SetApi(val IGraphqlApi) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDataSource)SetServiceRole(val awsiam.IRole) {
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
// Deprecated: - use `OpenSearchDataSource`.
func ElasticsearchDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticsearchDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.ElasticsearchDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDataSource) CreateFunction(id *string, props *BaseAppsyncFunctionProps) AppsyncFunction {
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

func (e *jsiiProxy_ElasticsearchDataSource) CreateResolver(id *string, props *BaseResolverProps) Resolver {
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

func (e *jsiiProxy_ElasticsearchDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

