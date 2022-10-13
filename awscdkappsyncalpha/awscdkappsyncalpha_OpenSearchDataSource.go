// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Appsync datasource backed by OpenSearch.
//
// Example:
//   import opensearch "github.com/aws/aws-cdk-go/awscdk"
//
//   var api graphqlApi
//
//
//   user := iam.NewUser(this, jsii.String("User"))
//   domain := opensearch.NewDomain(this, jsii.String("Domain"), &domainProps{
//   	version: opensearch.engineVersion_OPENSEARCH_1_3(),
//   	removalPolicy: awscdk.RemovalPolicy_DESTROY,
//   	fineGrainedAccessControl: &advancedSecurityOptions{
//   		masterUserArn: user.userArn,
//   	},
//   	encryptionAtRest: &encryptionAtRestOptions{
//   		enabled: jsii.Boolean(true),
//   	},
//   	nodeToNodeEncryption: jsii.Boolean(true),
//   	enforceHttps: jsii.Boolean(true),
//   })
//   ds := api.addOpenSearchDataSource(jsii.String("ds"), domain)
//
//   ds.createResolver(&baseResolverProps{
//   	typeName: jsii.String("Query"),
//   	fieldName: jsii.String("getTests"),
//   	requestMappingTemplate: appsync.mappingTemplate.fromString(jSON.stringify(map[string]interface{}{
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
//   	responseMappingTemplate: appsync.*mappingTemplate.fromString(jsii.String("[\n    #foreach($entry in $context.result.hits.hits)\n    #if( $velocityCount > 1 ) , #end\n    $utils.toJson($entry.get(\"_source\"))\n    #end\n  ]")),
//   })
//
// Experimental.
type OpenSearchDataSource interface {
	BackedDataSource
	// Experimental.
	Api() IGraphqlApi
	// Experimental.
	SetApi(val IGraphqlApi)
	// the underlying CFN data source resource.
	// Experimental.
	Ds() awsappsync.CfnDataSource
	// the principal of the data source to be IGrantable.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
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

func (j *jsiiProxy_OpenSearchDataSource) Ds() awsappsync.CfnDataSource {
	var returns awsappsync.CfnDataSource
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


// Experimental.
func NewOpenSearchDataSource(scope constructs.Construct, id *string, props *OpenSearchDataSourceProps) OpenSearchDataSource {
	_init_.Initialize()

	if err := validateNewOpenSearchDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenSearchDataSource{}

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.OpenSearchDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewOpenSearchDataSource_Override(o OpenSearchDataSource, scope constructs.Construct, id *string, props *OpenSearchDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appsync-alpha.OpenSearchDataSource",
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
// Experimental.
func OpenSearchDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpenSearchDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.OpenSearchDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchDataSource) CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction {
	if err := o.validateCreateFunctionParameters(props); err != nil {
		panic(err)
	}
	var returns AppsyncFunction

	_jsii_.Invoke(
		o,
		"createFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchDataSource) CreateResolver(props *BaseResolverProps) Resolver {
	if err := o.validateCreateResolverParameters(props); err != nil {
		panic(err)
	}
	var returns Resolver

	_jsii_.Invoke(
		o,
		"createResolver",
		[]interface{}{props},
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

