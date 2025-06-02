package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AppSync datasource backed by RDS.
//
// Example:
//   // Build a data source for AppSync to access the database.
//   var api graphqlApi
//   // Create username and password secret for DB Cluster
//   secret := rds.NewDatabaseSecret(this, jsii.String("AuroraSecret"), &DatabaseSecretProps{
//   	Username: jsii.String("clusteradmin"),
//   })
//
//   // The VPC to place the cluster in
//   vpc := ec2.NewVpc(this, jsii.String("AuroraVpc"))
//
//   // Create the serverless cluster, provide all values needed to customise the database.
//   cluster := rds.NewServerlessCluster(this, jsii.String("AuroraCluster"), &ServerlessClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA_MYSQL(),
//   	Vpc: Vpc,
//   	Credentials: map[string]*string{
//   		"username": jsii.String("clusteradmin"),
//   	},
//   	ClusterIdentifier: jsii.String("db-endpoint-test"),
//   	DefaultDatabaseName: jsii.String("demos"),
//   })
//   rdsDS := api.AddRdsDataSource(jsii.String("rds"), cluster, secret, jsii.String("demos"))
//
//   // Set up a resolver for an RDS query.
//   rdsDS.CreateResolver(jsii.String("QueryGetDemosRdsResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Query"),
//   	FieldName: jsii.String("getDemosRds"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromString(jsii.String(`
//   	  {
//   	    "version": "2018-05-29",
//   	    "statements": [
//   	      "SELECT * FROM demos"
//   	    ]
//   	  }
//   	  `)),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
//   	    $utils.toJson($utils.rds.toJsonObject($ctx.result)[0])
//   	  `)),
//   })
//
//   // Set up a resolver for an RDS mutation.
//   rdsDS.CreateResolver(jsii.String("MutationAddDemoRdsResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("addDemoRds"),
//   	RequestMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
//   	  {
//   	    "version": "2018-05-29",
//   	    "statements": [
//   	      "INSERT INTO demos VALUES (:id, :version)",
//   	      "SELECT * WHERE id = :id"
//   	    ],
//   	    "variableMap": {
//   	      ":id": $util.toJson($util.autoId()),
//   	      ":version": $util.toJson($ctx.args.version)
//   	    }
//   	  }
//   	  `)),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromString(jsii.String(`
//   	    $utils.toJson($utils.rds.toJsonObject($ctx.result)[1][0])
//   	  `)),
//   })
//
type RdsDataSource interface {
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

// The jsii proxy struct for RdsDataSource
type jsiiProxy_RdsDataSource struct {
	jsiiProxy_BackedDataSource
}

func (j *jsiiProxy_RdsDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


func NewRdsDataSource(scope constructs.Construct, id *string, props *RdsDataSourceProps) RdsDataSource {
	_init_.Initialize()

	if err := validateNewRdsDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.RdsDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRdsDataSource_Override(r RdsDataSource, scope constructs.Construct, id *string, props *RdsDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.RdsDataSource",
		[]interface{}{scope, id, props},
		r,
	)
}

func (j *jsiiProxy_RdsDataSource)SetApi(val IGraphqlApi) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_RdsDataSource)SetServiceRole(val awsiam.IRole) {
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
func RdsDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.RdsDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RdsDataSource_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appsync.RdsDataSource",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RdsDataSource) CreateFunction(id *string, props *BaseAppsyncFunctionProps) AppsyncFunction {
	if err := r.validateCreateFunctionParameters(id, props); err != nil {
		panic(err)
	}
	var returns AppsyncFunction

	_jsii_.Invoke(
		r,
		"createFunction",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDataSource) CreateResolver(id *string, props *BaseResolverProps) Resolver {
	if err := r.validateCreateResolverParameters(id, props); err != nil {
		panic(err)
	}
	var returns Resolver

	_jsii_.Invoke(
		r,
		"createResolver",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

