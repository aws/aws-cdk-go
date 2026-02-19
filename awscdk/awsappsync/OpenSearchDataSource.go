package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Appsync datasource backed by OpenSearch.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var api GraphqlApi
//
//
//   user := iam.NewUser(this, jsii.String("User"))
//   domain := opensearch.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: opensearch.EngineVersion_OPENSEARCH_2_3(),
//   	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
//   	FineGrainedAccessControl: &AdvancedSecurityOptions{
//   		MasterUserArn: user.userArn,
//   	},
//   	EncryptionAtRest: &EncryptionAtRestOptions{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	NodeToNodeEncryption: jsii.Boolean(true),
//   	EnforceHttps: jsii.Boolean(true),
//   })
//   ds := api.AddOpenSearchDataSource(jsii.String("ds"), domain)
//
//   ds.CreateResolver(jsii.String("QueryGetTestsResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Query"),
//   	FieldName: jsii.String("getTests"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromString(jSON.stringify(map[string]interface{}{
//   		"version": jsii.String("2017-02-28"),
//   		"operation": jsii.String("GET"),
//   		"path": jsii.String("/id/post/_search"),
//   		"params": map[string]map[string]interface{}{
//   			"headers": map[string]interface{}{
//   			},
//   			"queryString": map[string]interface{}{
//   			},
//   			"body": map[string]*f64{
//   				"from": jsii.Number(0),
//   				"size": jsii.Number(50),
//   			},
//   		},
//   	})),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`[
//   	    #foreach($entry in $context.result.hits.hits)
//   	    #if( $velocityCount > 1 ) , #end
//   	    $utils.toJson($entry.get("_source"))
//   	    #end
//   	  ]`)),
//   })
//
type OpenSearchDataSource interface {
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for OpenSearchDataSource
type jsiiProxy_OpenSearchDataSource struct {
	jsiiProxy_BackedDataSource
}

func (j *jsiiProxy_OpenSearchDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


func NewOpenSearchDataSource(scope constructs.Construct, id *string, props *OpenSearchDataSourceProps) OpenSearchDataSource {
	_init_.Initialize()

	if err := validateNewOpenSearchDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenSearchDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.OpenSearchDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewOpenSearchDataSource_Override(o OpenSearchDataSource, scope constructs.Construct, id *string, props *OpenSearchDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.OpenSearchDataSource",
		[]interface{}{scope, id, props},
		o,
	)
}

func (j *jsiiProxy_OpenSearchDataSource)SetApi(val IGraphqlApi) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_OpenSearchDataSource)SetServiceRole(val awsiam.IRole) {
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
func OpenSearchDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpenSearchDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.OpenSearchDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpenSearchDataSource_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.OpenSearchDataSource",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OpenSearchDataSource) CreateFunction(id *string, props *BaseAppsyncFunctionProps) AppsyncFunction {
	if err := o.validateCreateFunctionParameters(id, props); err != nil {
		panic(err)
	}
	var returns AppsyncFunction

	_jsii_.Invoke(
		o,
		"createFunction",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchDataSource) CreateResolver(id *string, props *BaseResolverProps) Resolver {
	if err := o.validateCreateResolverParameters(id, props); err != nil {
		panic(err)
	}
	var returns Resolver

	_jsii_.Invoke(
		o,
		"createResolver",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchDataSource) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		o,
		"with",
		args,
		&returns,
	)

	return returns
}

