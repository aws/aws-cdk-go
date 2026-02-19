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
//
//   domain := opensearch.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: opensearch.EngineVersion_OPENSEARCH_2_17(),
//   	EncryptionAtRest: &EncryptionAtRestOptions{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	NodeToNodeEncryption: jsii.Boolean(true),
//   	EnforceHttps: jsii.Boolean(true),
//   	Capacity: &CapacityConfig{
//   		MultiAzWithStandbyEnabled: jsii.Boolean(false),
//   	},
//   	Ebs: &EbsOptions{
//   		Enabled: jsii.Boolean(true),
//   		VolumeSize: jsii.Number(10),
//   	},
//   })
//   api := appsync.NewEventApi(this, jsii.String("EventApiOpenSearch"), &EventApiProps{
//   	ApiName: jsii.String("OpenSearchEventApi"),
//   })
//
//   dataSource := api.AddOpenSearchDataSource(jsii.String("opensearchds"), domain)
//
type AppSyncOpenSearchDataSource interface {
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

// The jsii proxy struct for AppSyncOpenSearchDataSource
type jsiiProxy_AppSyncOpenSearchDataSource struct {
	jsiiProxy_AppSyncBackedDataSource
}

func (j *jsiiProxy_AppSyncOpenSearchDataSource) Api() IApi {
	var returns IApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncOpenSearchDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncOpenSearchDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncOpenSearchDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncOpenSearchDataSource) Resource() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncOpenSearchDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


func NewAppSyncOpenSearchDataSource(scope constructs.Construct, id *string, props *AppSyncOpenSearchDataSourceProps) AppSyncOpenSearchDataSource {
	_init_.Initialize()

	if err := validateNewAppSyncOpenSearchDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSyncOpenSearchDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.AppSyncOpenSearchDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAppSyncOpenSearchDataSource_Override(a AppSyncOpenSearchDataSource, scope constructs.Construct, id *string, props *AppSyncOpenSearchDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.AppSyncOpenSearchDataSource",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AppSyncOpenSearchDataSource)SetApi(val IApi) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_AppSyncOpenSearchDataSource)SetServiceRole(val awsiam.IRole) {
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
func AppSyncOpenSearchDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSyncOpenSearchDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.AppSyncOpenSearchDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncOpenSearchDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncOpenSearchDataSource) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

