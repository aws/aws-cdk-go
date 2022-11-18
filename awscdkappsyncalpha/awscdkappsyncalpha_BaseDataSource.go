// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Abstract AppSync datasource implementation.
//
// Do not use directly but use subclasses for concrete datasources.
//
// Example:
//   var api graphqlApi
//   var dummyRequest mappingTemplate
//   var dummyResponse mappingTemplate
//
//   info := appsync.NewObjectType(jsii.String("Info"), &objectTypeOptions{
//   	definition: map[string]iField{
//   		"node": appsync.NewResolvableField(&ResolvableFieldOptions{
//   			"returnType": appsync.GraphqlType.string(),
//   			"args": map[string]GraphqlType{
//   				"id": appsync.GraphqlType.string(),
//   			},
//   			"dataSource": api.addNoneDataSource(jsii.String("none")),
//   			"requestMappingTemplate": dummyRequest,
//   			"responseMappingTemplate": dummyResponse,
//   		}),
//   	},
//   })
//
// Experimental.
type BaseDataSource interface {
	constructs.Construct
	// Experimental.
	Api() IGraphqlApi
	// Experimental.
	SetApi(val IGraphqlApi)
	// the underlying CFN data source resource.
	// Experimental.
	Ds() awsappsync.CfnDataSource
	// the name of the data source.
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	ServiceRole() awsiam.IRole
	// Experimental.
	SetServiceRole(val awsiam.IRole)
	// creates a new appsync function for this datasource and API using the given properties.
	// Experimental.
	CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction
	// creates a new resolver for this datasource and API using the given properties.
	// Experimental.
	CreateResolver(props *BaseResolverProps) Resolver
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BaseDataSource
type jsiiProxy_BaseDataSource struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_BaseDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDataSource) Ds() awsappsync.CfnDataSource {
	var returns awsappsync.CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


// Experimental.
func NewBaseDataSource_Override(b BaseDataSource, scope constructs.Construct, id *string, props *BackedDataSourceProps, extended *ExtendedDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.BaseDataSource",
		[]interface{}{scope, id, props, extended},
		b,
	)
}

func (j *jsiiProxy_BaseDataSource)SetApi(val IGraphqlApi) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_BaseDataSource)SetServiceRole(val awsiam.IRole) {
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
// Experimental.
func BaseDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBaseDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.BaseDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseDataSource) CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction {
	if err := b.validateCreateFunctionParameters(props); err != nil {
		panic(err)
	}
	var returns AppsyncFunction

	_jsii_.Invoke(
		b,
		"createFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseDataSource) CreateResolver(props *BaseResolverProps) Resolver {
	if err := b.validateCreateResolverParameters(props); err != nil {
		panic(err)
	}
	var returns Resolver

	_jsii_.Invoke(
		b,
		"createResolver",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

