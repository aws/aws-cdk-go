package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
)

// An AppSync datasource backed by a http endpoint.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
//   	name: jsii.String("api"),
//   	schema: appsync.schema.fromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
//   })
//
//   httpDs := api.addHttpDataSource(jsii.String("ds"), jsii.String("https://states.amazonaws.com"), &httpDataSourceOptions{
//   	name: jsii.String("httpDsWithStepF"),
//   	description: jsii.String("from appsync to StepFunctions Workflow"),
//   	authorizationConfig: &awsIamConfig{
//   		signingRegion: jsii.String("us-east-1"),
//   		signingServiceName: jsii.String("states"),
//   	},
//   })
//
//   httpDs.createResolver(&baseResolverProps{
//   	typeName: jsii.String("Mutation"),
//   	fieldName: jsii.String("callStepFunction"),
//   	requestMappingTemplate: appsync.mappingTemplate.fromFile(jsii.String("request.vtl")),
//   	responseMappingTemplate: appsync.*mappingTemplate.fromFile(jsii.String("response.vtl")),
//   })
//
// Experimental.
type HttpDataSource interface {
	BackedDataSource
	// Experimental.
	Api() IGraphqlApi
	// Experimental.
	SetApi(val IGraphqlApi)
	// the underlying CFN data source resource.
	// Experimental.
	Ds() CfnDataSource
	// the principal of the data source to be IGrantable.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// the name of the data source.
	// Experimental.
	Name() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
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

func (j *jsiiProxy_HttpDataSource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Experimental.
func NewHttpDataSource(scope constructs.Construct, id *string, props *HttpDataSourceProps) HttpDataSource {
	_init_.Initialize()

	if err := validateNewHttpDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpDataSource{}

	_jsii_.Create(
		"monocdk.aws_appsync.HttpDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpDataSource_Override(h HttpDataSource, scope constructs.Construct, id *string, props *HttpDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.HttpDataSource",
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

// Return whether the given object is a Construct.
// Experimental.
func HttpDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHttpDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.HttpDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpDataSource) CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction {
	if err := h.validateCreateFunctionParameters(props); err != nil {
		panic(err)
	}
	var returns AppsyncFunction

	_jsii_.Invoke(
		h,
		"createFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpDataSource) CreateResolver(props *BaseResolverProps) Resolver {
	if err := h.validateCreateResolverParameters(props); err != nil {
		panic(err)
	}
	var returns Resolver

	_jsii_.Invoke(
		h,
		"createResolver",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpDataSource) OnPrepare() {
	_jsii_.InvokeVoid(
		h,
		"onPrepare",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HttpDataSource) OnSynthesize(session constructs.ISynthesisSession) {
	if err := h.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (h *jsiiProxy_HttpDataSource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpDataSource) Prepare() {
	_jsii_.InvokeVoid(
		h,
		"prepare",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HttpDataSource) Synthesize(session awscdk.ISynthesisSession) {
	if err := h.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"synthesize",
		[]interface{}{session},
	)
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

func (h *jsiiProxy_HttpDataSource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

