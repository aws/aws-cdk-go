package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A parameter group.
//
// Represents both a cluster parameter group,
// and an instance parameter group.
//
// Example:
//   var plan backupPlan
//   var vpc vpc
//
//   myTable := dynamodb.Table_FromTableName(this, jsii.String("Table"), jsii.String("myTableName"))
//   myDatabaseInstance := rds.NewDatabaseInstance(this, jsii.String("DatabaseInstance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
//   		Version: rds.MysqlEngineVersion_VER_8_0_26(),
//   	}),
//   	Vpc: Vpc,
//   })
//   myDatabaseCluster := rds.NewDatabaseCluster(this, jsii.String("DatabaseCluster"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_2_08_1(),
//   	}),
//   	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("clusteradmin")),
//   	InstanceProps: &InstanceProps{
//   		Vpc: *Vpc,
//   	},
//   })
//   myServerlessCluster := rds.NewServerlessCluster(this, jsii.String("ServerlessCluster"), &ServerlessClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA_POSTGRESQL(),
//   	ParameterGroup: rds.ParameterGroup_FromParameterGroupName(this, jsii.String("ParameterGroup"), jsii.String("default.aurora-postgresql11")),
//   	Vpc: Vpc,
//   })
//   myCoolConstruct := constructs.NewConstruct(this, jsii.String("MyCoolConstruct"))
//
//   plan.AddSelection(jsii.String("Selection"), &BackupSelectionOptions{
//   	Resources: []backupResource{
//   		backup.*backupResource_FromDynamoDbTable(myTable),
//   		backup.*backupResource_FromRdsDatabaseInstance(myDatabaseInstance),
//   		backup.*backupResource_FromRdsDatabaseCluster(myDatabaseCluster),
//   		backup.*backupResource_FromRdsServerlessCluster(myServerlessCluster),
//   		backup.*backupResource_FromTag(jsii.String("stage"), jsii.String("prod")),
//   		backup.*backupResource_FromConstruct(myCoolConstruct),
//   	},
//   })
//
type ParameterGroup interface {
	awscdk.Resource
	IParameterGroup
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Add a parameter to this parameter group.
	AddParameter(key *string, value *string) *bool
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Method called when this Parameter Group is used when defining a database cluster.
	BindToCluster(_options *ParameterGroupClusterBindOptions) *ParameterGroupClusterConfig
	// Method called when this Parameter Group is used when defining a database instance.
	BindToInstance(_options *ParameterGroupInstanceBindOptions) *ParameterGroupInstanceConfig
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ParameterGroup
type jsiiProxy_ParameterGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IParameterGroup
}

func (j *jsiiProxy_ParameterGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ParameterGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ParameterGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ParameterGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewParameterGroup(scope constructs.Construct, id *string, props *ParameterGroupProps) ParameterGroup {
	_init_.Initialize()

	if err := validateNewParameterGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ParameterGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewParameterGroup_Override(p ParameterGroup, scope constructs.Construct, id *string, props *ParameterGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		[]interface{}{scope, id, props},
		p,
	)
}

// Imports a parameter group.
func ParameterGroup_FromParameterGroupName(scope constructs.Construct, id *string, parameterGroupName *string) IParameterGroup {
	_init_.Initialize()

	if err := validateParameterGroup_FromParameterGroupNameParameters(scope, id, parameterGroupName); err != nil {
		panic(err)
	}
	var returns IParameterGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		"fromParameterGroupName",
		[]interface{}{scope, id, parameterGroupName},
		&returns,
	)

	return returns
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
func ParameterGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateParameterGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ParameterGroup_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateParameterGroup_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ParameterGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateParameterGroup_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func ParameterGroup_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.ParameterGroup",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_ParameterGroup) AddParameter(key *string, value *string) *bool {
	if err := p.validateAddParameterParameters(key, value); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		p,
		"addParameter",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := p.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_ParameterGroup) BindToCluster(_options *ParameterGroupClusterBindOptions) *ParameterGroupClusterConfig {
	if err := p.validateBindToClusterParameters(_options); err != nil {
		panic(err)
	}
	var returns *ParameterGroupClusterConfig

	_jsii_.Invoke(
		p,
		"bindToCluster",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterGroup) BindToInstance(_options *ParameterGroupInstanceBindOptions) *ParameterGroupInstanceConfig {
	if err := p.validateBindToInstanceParameters(_options); err != nil {
		panic(err)
	}
	var returns *ParameterGroupInstanceConfig

	_jsii_.Invoke(
		p,
		"bindToInstance",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := p.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterGroup) GetResourceNameAttribute(nameAttr *string) *string {
	if err := p.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParameterGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

