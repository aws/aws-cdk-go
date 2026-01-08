package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AppSync datasource backed by a http endpoint.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("api"),
//   	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("schema.graphql"))),
//   })
//
//   httpDs := api.AddHttpDataSource(jsii.String("ds"), jsii.String("https://states.amazonaws.com"), &HttpDataSourceOptions{
//   	Name: jsii.String("httpDsWithStepF"),
//   	Description: jsii.String("from appsync to StepFunctions Workflow"),
//   	AuthorizationConfig: &AwsIamConfig{
//   		SigningRegion: jsii.String("us-east-1"),
//   		SigningServiceName: jsii.String("states"),
//   	},
//   })
//
//   httpDs.CreateResolver(jsii.String("MutationCallStepFunctionResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("callStepFunction"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("request.vtl")),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("response.vtl")),
//   })
//
type HttpDataSource interface {
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
}

// The jsii proxy struct for HttpDataSource
type jsiiProxy_HttpDataSource struct {
	jsiiProxy_BackedDataSource
}

func (j *jsiiProxy_HttpDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


func NewHttpDataSource(scope constructs.Construct, id *string, props *HttpDataSourceProps) HttpDataSource {
	_init_.Initialize()

	if err := validateNewHttpDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.HttpDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewHttpDataSource_Override(h HttpDataSource, scope constructs.Construct, id *string, props *HttpDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.HttpDataSource",
		[]interface{}{scope, id, props},
		h,
	)
}

func (j *jsiiProxy_HttpDataSource)SetApi(val IGraphqlApi) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_HttpDataSource)SetServiceRole(val awsiam.IRole) {
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
func HttpDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHttpDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.HttpDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HttpDataSource_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.HttpDataSource",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HttpDataSource) CreateFunction(id *string, props *BaseAppsyncFunctionProps) AppsyncFunction {
	if err := h.validateCreateFunctionParameters(id, props); err != nil {
		panic(err)
	}
	var returns AppsyncFunction

	_jsii_.Invoke(
		h,
		"createFunction",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpDataSource) CreateResolver(id *string, props *BaseResolverProps) Resolver {
	if err := h.validateCreateResolverParameters(id, props); err != nil {
		panic(err)
	}
	var returns Resolver

	_jsii_.Invoke(
		h,
		"createResolver",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

