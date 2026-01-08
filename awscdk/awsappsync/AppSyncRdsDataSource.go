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
//   import secretsmanager "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc Vpc
//
//   databaseName := "mydb"
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Cluster"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraPostgres(&AuroraPostgresClusterEngineProps{
//   		Version: rds.AuroraPostgresEngineVersion_VER_16_6(),
//   	}),
//   	Writer: rds.ClusterInstance_ServerlessV2(jsii.String("writer")),
//   	Vpc: vpc,
//   	Credentials: map[string]*string{
//   		"username": jsii.String("clusteradmin"),
//   	},
//   	DefaultDatabaseName: databaseName,
//   	EnableDataApi: jsii.Boolean(true),
//   })
//
//   secret := secretsmanager.Secret_FromSecretNameV2(this, jsii.String("Secret"), jsii.String("db-secretName"))
//
//   api := appsync.NewEventApi(this, jsii.String("EventApiRds"), &EventApiProps{
//   	ApiName: jsii.String("RdsEventApi"),
//   })
//
//   dataSource := api.AddRdsDataSource(jsii.String("rdsds"), cluster, secret, databaseName)
//
type AppSyncRdsDataSource interface {
	AppSyncBackedDataSource
	// The API this data source is attached to Set the API this data source is attached to.
	Api() IApi
	SetApi(val IApi)
	// The principal of the data source to be IGrantable.
	GrantPrincipal() awsiam.IPrincipal
	// The name of the data source.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// The underlying CFN data source resource.
	Resource() CfnDataSource
	ServiceRole() awsiam.IRole
	SetServiceRole(val awsiam.IRole)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AppSyncRdsDataSource
type jsiiProxy_AppSyncRdsDataSource struct {
	jsiiProxy_AppSyncBackedDataSource
}

func (j *jsiiProxy_AppSyncRdsDataSource) Api() IApi {
	var returns IApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncRdsDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncRdsDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncRdsDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncRdsDataSource) Resource() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncRdsDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


func NewAppSyncRdsDataSource(scope constructs.Construct, id *string, props *AppSyncRdsDataSourceProps) AppSyncRdsDataSource {
	_init_.Initialize()

	if err := validateNewAppSyncRdsDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSyncRdsDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.AppSyncRdsDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAppSyncRdsDataSource_Override(a AppSyncRdsDataSource, scope constructs.Construct, id *string, props *AppSyncRdsDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.AppSyncRdsDataSource",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AppSyncRdsDataSource)SetApi(val IApi) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_AppSyncRdsDataSource)SetServiceRole(val awsiam.IRole) {
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
func AppSyncRdsDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSyncRdsDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.AppSyncRdsDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncRdsDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

