package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
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
//   elasticsearchDataSource := awscdk.Aws_appsync.NewElasticsearchDataSource(this, jsii.String("MyElasticsearchDataSource"), &elasticsearchDataSourceProps{
//   	api: graphqlApi,
//   	domain: domain,
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	serviceRole: role,
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
	// The construct tree node associated with this construct.
	// Deprecated: - use `OpenSearchDataSource`.
	Node() awscdk.ConstructNode
	// Deprecated: - use `OpenSearchDataSource`.
	ServiceRole() awsiam.IRole
	// Deprecated: - use `OpenSearchDataSource`.
	SetServiceRole(val awsiam.IRole)
	// creates a new appsync function for this datasource and API using the given properties.
	// Deprecated: - use `OpenSearchDataSource`.
	CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction
	// creates a new resolver for this datasource and API using the given properties.
	// Deprecated: - use `OpenSearchDataSource`.
	CreateResolver(props *BaseResolverProps) Resolver
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: - use `OpenSearchDataSource`.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: - use `OpenSearchDataSource`.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: - use `OpenSearchDataSource`.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: - use `OpenSearchDataSource`.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: - use `OpenSearchDataSource`.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Deprecated: - use `OpenSearchDataSource`.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: - use `OpenSearchDataSource`.
	Validate() *[]*string
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

func (j *jsiiProxy_ElasticsearchDataSource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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

	j := jsiiProxy_ElasticsearchDataSource{}

	_jsii_.Create(
		"monocdk.aws_appsync.ElasticsearchDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: - use `OpenSearchDataSource`.
func NewElasticsearchDataSource_Override(e ElasticsearchDataSource, scope constructs.Construct, id *string, props *ElasticsearchDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.ElasticsearchDataSource",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDataSource) SetApi(val IGraphqlApi) {
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDataSource) SetServiceRole(val awsiam.IRole) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

// Return whether the given object is a Construct.
// Deprecated: - use `OpenSearchDataSource`.
func ElasticsearchDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ElasticsearchDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDataSource) CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction {
	var returns AppsyncFunction

	_jsii_.Invoke(
		e,
		"createFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDataSource) CreateResolver(props *BaseResolverProps) Resolver {
	var returns Resolver

	_jsii_.Invoke(
		e,
		"createResolver",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDataSource) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDataSource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (e *jsiiProxy_ElasticsearchDataSource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDataSource) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDataSource) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
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

func (e *jsiiProxy_ElasticsearchDataSource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

