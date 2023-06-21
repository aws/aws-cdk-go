package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AppSync datasource backed by a Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//   var graphqlApi graphqlApi
//   var role role
//
//   lambdaDataSource := awscdk.Aws_appsync.NewLambdaDataSource(this, jsii.String("MyLambdaDataSource"), &LambdaDataSourceProps{
//   	Api: graphqlApi,
//   	LambdaFunction: function_,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ServiceRole: role,
//   })
//
type LambdaDataSource interface {
	BackedDataSource
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

// The jsii proxy struct for LambdaDataSource
type jsiiProxy_LambdaDataSource struct {
	jsiiProxy_BackedDataSource
}

func (j *jsiiProxy_LambdaDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


func NewLambdaDataSource(scope constructs.Construct, id *string, props *LambdaDataSourceProps) LambdaDataSource {
	_init_.Initialize()

	if err := validateNewLambdaDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.LambdaDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaDataSource_Override(l LambdaDataSource, scope constructs.Construct, id *string, props *LambdaDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.LambdaDataSource",
		[]interface{}{scope, id, props},
		l,
	)
}

func (j *jsiiProxy_LambdaDataSource)SetApi(val IGraphqlApi) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_LambdaDataSource)SetServiceRole(val awsiam.IRole) {
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
func LambdaDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.LambdaDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaDataSource) CreateFunction(id *string, props *BaseAppsyncFunctionProps) AppsyncFunction {
	if err := l.validateCreateFunctionParameters(id, props); err != nil {
		panic(err)
	}
	var returns AppsyncFunction

	_jsii_.Invoke(
		l,
		"createFunction",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaDataSource) CreateResolver(id *string, props *BaseResolverProps) Resolver {
	if err := l.validateCreateResolverParameters(id, props); err != nil {
		panic(err)
	}
	var returns Resolver

	_jsii_.Invoke(
		l,
		"createResolver",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

