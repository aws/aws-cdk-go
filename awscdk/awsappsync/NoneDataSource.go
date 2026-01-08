package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AppSync dummy datasource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var graphQLApiRef IGraphQLApiRef
//
//   noneDataSource := awscdk.Aws_appsync.NewNoneDataSource(this, jsii.String("MyNoneDataSource"), &NoneDataSourceProps{
//   	Api: graphQLApiRef,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   })
//
type NoneDataSource interface {
	BaseDataSource
	// The API this data source is attached to Set the API this data source is attached to.
	Api() IGraphqlApi
	SetApi(val IGraphqlApi)
	// the underlying CFN data source resource.
	Ds() CfnDataSource
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

// The jsii proxy struct for NoneDataSource
type jsiiProxy_NoneDataSource struct {
	jsiiProxy_BaseDataSource
}

func (j *jsiiProxy_NoneDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NoneDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NoneDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NoneDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NoneDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


func NewNoneDataSource(scope constructs.Construct, id *string, props *NoneDataSourceProps) NoneDataSource {
	_init_.Initialize()

	if err := validateNewNoneDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NoneDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.NoneDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNoneDataSource_Override(n NoneDataSource, scope constructs.Construct, id *string, props *NoneDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.NoneDataSource",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_NoneDataSource)SetApi(val IGraphqlApi) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_NoneDataSource)SetServiceRole(val awsiam.IRole) {
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
func NoneDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNoneDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.NoneDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NoneDataSource) CreateFunction(id *string, props *BaseAppsyncFunctionProps) AppsyncFunction {
	if err := n.validateCreateFunctionParameters(id, props); err != nil {
		panic(err)
	}
	var returns AppsyncFunction

	_jsii_.Invoke(
		n,
		"createFunction",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NoneDataSource) CreateResolver(id *string, props *BaseResolverProps) Resolver {
	if err := n.validateCreateResolverParameters(id, props); err != nil {
		panic(err)
	}
	var returns Resolver

	_jsii_.Invoke(
		n,
		"createResolver",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NoneDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

