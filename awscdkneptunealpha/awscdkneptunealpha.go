// The CDK Construct Library for AWS::Neptune
package awscdkneptunealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkneptunealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdkneptunealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A cluster parameter group.
//
// TODO: EXAMPLE
//
// Experimental.
type ClusterParameterGroup interface {
	awscdk.Resource
	IClusterParameterGroup
	ClusterParameterGroupName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for ClusterParameterGroup
type jsiiProxy_ClusterParameterGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IClusterParameterGroup
}

func (j *jsiiProxy_ClusterParameterGroup) ClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterParameterGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterParameterGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterParameterGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterParameterGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewClusterParameterGroup(scope constructs.Construct, id *string, props *ClusterParameterGroupProps) ClusterParameterGroup {
	_init_.Initialize()

	j := jsiiProxy_ClusterParameterGroup{}

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.ClusterParameterGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewClusterParameterGroup_Override(c ClusterParameterGroup, scope constructs.Construct, id *string, props *ClusterParameterGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.ClusterParameterGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

// Imports a parameter group.
// Experimental.
func ClusterParameterGroup_FromClusterParameterGroupName(scope constructs.Construct, id *string, clusterParameterGroupName *string) IClusterParameterGroup {
	_init_.Initialize()

	var returns IClusterParameterGroup

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.ClusterParameterGroup",
		"fromClusterParameterGroupName",
		[]interface{}{scope, id, clusterParameterGroupName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ClusterParameterGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.ClusterParameterGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ClusterParameterGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.ClusterParameterGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_ClusterParameterGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (c *jsiiProxy_ClusterParameterGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (c *jsiiProxy_ClusterParameterGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (c *jsiiProxy_ClusterParameterGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_ClusterParameterGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Marker class for cluster parameter group.
//
// TODO: EXAMPLE
//
// Experimental.
type ClusterParameterGroupProps struct {
	// The parameters in this parameter group.
	// Experimental.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// The name of the parameter group.
	// Experimental.
	ClusterParameterGroupName *string `json:"clusterParameterGroupName" yaml:"clusterParameterGroupName"`
	// Description for this parameter group.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
}

// Create a clustered database with a given number of instances.
//
// TODO: EXAMPLE
//
// Experimental.
type DatabaseCluster interface {
	DatabaseClusterBase
	IDatabaseCluster
	ClusterEndpoint() Endpoint
	ClusterIdentifier() *string
	ClusterReadEndpoint() Endpoint
	ClusterResourceIdentifier() *string
	Connections() awsec2.Connections
	EnableIamAuthentication() *bool
	SetEnableIamAuthentication(val *bool)
	Env() *awscdk.ResourceEnvironment
	InstanceEndpoints() *[]Endpoint
	InstanceIdentifiers() *[]*string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	SubnetGroup() ISubnetGroup
	Vpc() awsec2.IVpc
	VpcSubnets() *awsec2.SubnetSelection
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantConnect(grantee awsiam.IGrantable) awsiam.Grant
	ToString() *string
}

// The jsii proxy struct for DatabaseCluster
type jsiiProxy_DatabaseCluster struct {
	jsiiProxy_DatabaseClusterBase
	jsiiProxy_IDatabaseCluster
}

func (j *jsiiProxy_DatabaseCluster) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) ClusterResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) EnableIamAuthentication() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableIamAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) InstanceEndpoints() *[]Endpoint {
	var returns *[]Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) InstanceIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) SubnetGroup() ISubnetGroup {
	var returns ISubnetGroup
	_jsii_.Get(
		j,
		"subnetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) VpcSubnets() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"vpcSubnets",
		&returns,
	)
	return returns
}


// Experimental.
func NewDatabaseCluster(scope constructs.Construct, id *string, props *DatabaseClusterProps) DatabaseCluster {
	_init_.Initialize()

	j := jsiiProxy_DatabaseCluster{}

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.DatabaseCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDatabaseCluster_Override(d DatabaseCluster, scope constructs.Construct, id *string, props *DatabaseClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.DatabaseCluster",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DatabaseCluster) SetEnableIamAuthentication(val *bool) {
	_jsii_.Set(
		j,
		"enableIamAuthentication",
		val,
	)
}

// Import an existing DatabaseCluster from properties.
// Experimental.
func DatabaseCluster_FromDatabaseClusterAttributes(scope constructs.Construct, id *string, attrs *DatabaseClusterAttributes) IDatabaseCluster {
	_init_.Initialize()

	var returns IDatabaseCluster

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.DatabaseCluster",
		"fromDatabaseClusterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DatabaseCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.DatabaseCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DatabaseCluster_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.DatabaseCluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func DatabaseCluster_DEFAULT_NUM_INSTANCES() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.DatabaseCluster",
		"DEFAULT_NUM_INSTANCES",
		&returns,
	)
	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_DatabaseCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (d *jsiiProxy_DatabaseCluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (d *jsiiProxy_DatabaseCluster) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (d *jsiiProxy_DatabaseCluster) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity connection access to the database.
// Experimental.
func (d *jsiiProxy_DatabaseCluster) GrantConnect(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantConnect",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DatabaseCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties that describe an existing cluster instance.
//
// TODO: EXAMPLE
//
// Experimental.
type DatabaseClusterAttributes struct {
	// Cluster endpoint address.
	// Experimental.
	ClusterEndpointAddress *string `json:"clusterEndpointAddress" yaml:"clusterEndpointAddress"`
	// Identifier for the cluster.
	// Experimental.
	ClusterIdentifier *string `json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Resource Identifier for the cluster.
	// Experimental.
	ClusterResourceIdentifier *string `json:"clusterResourceIdentifier" yaml:"clusterResourceIdentifier"`
	// The database port.
	// Experimental.
	Port *float64 `json:"port" yaml:"port"`
	// Reader endpoint address.
	// Experimental.
	ReaderEndpointAddress *string `json:"readerEndpointAddress" yaml:"readerEndpointAddress"`
	// The security group of the database cluster.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
}

// A new or imported database cluster.
//
// TODO: EXAMPLE
//
// Experimental.
type DatabaseClusterBase interface {
	awscdk.Resource
	IDatabaseCluster
	ClusterEndpoint() Endpoint
	ClusterIdentifier() *string
	ClusterReadEndpoint() Endpoint
	ClusterResourceIdentifier() *string
	Connections() awsec2.Connections
	EnableIamAuthentication() *bool
	SetEnableIamAuthentication(val *bool)
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantConnect(grantee awsiam.IGrantable) awsiam.Grant
	ToString() *string
}

// The jsii proxy struct for DatabaseClusterBase
type jsiiProxy_DatabaseClusterBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IDatabaseCluster
}

func (j *jsiiProxy_DatabaseClusterBase) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) ClusterResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) EnableIamAuthentication() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableIamAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDatabaseClusterBase_Override(d DatabaseClusterBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.DatabaseClusterBase",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DatabaseClusterBase) SetEnableIamAuthentication(val *bool) {
	_jsii_.Set(
		j,
		"enableIamAuthentication",
		val,
	)
}

// Import an existing DatabaseCluster from properties.
// Experimental.
func DatabaseClusterBase_FromDatabaseClusterAttributes(scope constructs.Construct, id *string, attrs *DatabaseClusterAttributes) IDatabaseCluster {
	_init_.Initialize()

	var returns IDatabaseCluster

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.DatabaseClusterBase",
		"fromDatabaseClusterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DatabaseClusterBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.DatabaseClusterBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DatabaseClusterBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.DatabaseClusterBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_DatabaseClusterBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (d *jsiiProxy_DatabaseClusterBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (d *jsiiProxy_DatabaseClusterBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (d *jsiiProxy_DatabaseClusterBase) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity connection access to the database.
// Experimental.
func (d *jsiiProxy_DatabaseClusterBase) GrantConnect(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantConnect",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DatabaseClusterBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a new database cluster.
//
// TODO: EXAMPLE
//
// Experimental.
type DatabaseClusterProps struct {
	// What type of instance to start for the replicas.
	// Experimental.
	InstanceType InstanceType `json:"instanceType" yaml:"instanceType"`
	// What subnets to run the Neptune instances in.
	//
	// Must be at least 2 subnets in two different AZs.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// A list of AWS Identity and Access Management (IAM) role that can be used by the cluster to access other AWS services.
	// Experimental.
	AssociatedRoles *[]awsiam.IRole `json:"associatedRoles" yaml:"associatedRoles"`
	// If set to true, Neptune will automatically update the engine of the entire cluster to the latest minor version after a stabilization window of 2 to 3 weeks.
	// Experimental.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// How many days to retain the backup.
	// Experimental.
	BackupRetention awscdk.Duration `json:"backupRetention" yaml:"backupRetention"`
	// Additional parameters to pass to the database engine.
	// Experimental.
	ClusterParameterGroup IClusterParameterGroup `json:"clusterParameterGroup" yaml:"clusterParameterGroup"`
	// An optional identifier for the cluster.
	// Experimental.
	DbClusterName *string `json:"dbClusterName" yaml:"dbClusterName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	// Experimental.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// What version of the database to start.
	// Experimental.
	EngineVersion EngineVersion `json:"engineVersion" yaml:"engineVersion"`
	// Map AWS Identity and Access Management (IAM) accounts to database accounts.
	// Experimental.
	IamAuthentication *bool `json:"iamAuthentication" yaml:"iamAuthentication"`
	// Base identifier for instances.
	//
	// Every replica is named by appending the replica number to this string, 1-based.
	// Experimental.
	InstanceIdentifierBase *string `json:"instanceIdentifierBase" yaml:"instanceIdentifierBase"`
	// Number of Neptune compute instances.
	// Experimental.
	Instances *float64 `json:"instances" yaml:"instances"`
	// The KMS key for storage encryption.
	// Experimental.
	KmsKey awskms.IKey `json:"kmsKey" yaml:"kmsKey"`
	// The DB parameter group to associate with the instance.
	// Experimental.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The port the Neptune cluster will listen on.
	// Experimental.
	Port *float64 `json:"port" yaml:"port"`
	// A daily time range in 24-hours UTC format in which backups preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: '01:00-02:00'
	// Experimental.
	PreferredBackupWindow *string `json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// A weekly time range in which maintenance should preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: 'tue:04:17-tue:04:47'
	// Experimental.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The removal policy to apply when the cluster and its instances are removed or replaced during a stack update, or when the stack is deleted.
	//
	// This
	// removal policy also applies to the implicit security group created for the
	// cluster if one is not supplied as a parameter.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// Security group.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// Whether to enable storage encryption.
	// Experimental.
	StorageEncrypted *bool `json:"storageEncrypted" yaml:"storageEncrypted"`
	// Existing subnet group for the cluster.
	// Experimental.
	SubnetGroup ISubnetGroup `json:"subnetGroup" yaml:"subnetGroup"`
	// Where to place the instances within the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// A database instance.
//
// TODO: EXAMPLE
//
// Experimental.
type DatabaseInstance interface {
	awscdk.Resource
	IDatabaseInstance
	Cluster() IDatabaseCluster
	DbInstanceEndpointAddress() *string
	DbInstanceEndpointPort() *string
	Env() *awscdk.ResourceEnvironment
	InstanceEndpoint() Endpoint
	InstanceIdentifier() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for DatabaseInstance
type jsiiProxy_DatabaseInstance struct {
	internal.Type__awscdkResource
	jsiiProxy_IDatabaseInstance
}

func (j *jsiiProxy_DatabaseInstance) Cluster() IDatabaseCluster {
	var returns IDatabaseCluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) DbInstanceEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) DbInstanceEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDatabaseInstance(scope constructs.Construct, id *string, props *DatabaseInstanceProps) DatabaseInstance {
	_init_.Initialize()

	j := jsiiProxy_DatabaseInstance{}

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.DatabaseInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDatabaseInstance_Override(d DatabaseInstance, scope constructs.Construct, id *string, props *DatabaseInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.DatabaseInstance",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing database instance.
// Experimental.
func DatabaseInstance_FromDatabaseInstanceAttributes(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) IDatabaseInstance {
	_init_.Initialize()

	var returns IDatabaseInstance

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.DatabaseInstance",
		"fromDatabaseInstanceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DatabaseInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.DatabaseInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DatabaseInstance_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.DatabaseInstance",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_DatabaseInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (d *jsiiProxy_DatabaseInstance) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (d *jsiiProxy_DatabaseInstance) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (d *jsiiProxy_DatabaseInstance) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DatabaseInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties that describe an existing instance.
//
// TODO: EXAMPLE
//
// Experimental.
type DatabaseInstanceAttributes struct {
	// The endpoint address.
	// Experimental.
	InstanceEndpointAddress *string `json:"instanceEndpointAddress" yaml:"instanceEndpointAddress"`
	// The instance identifier.
	// Experimental.
	InstanceIdentifier *string `json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// The database port.
	// Experimental.
	Port *float64 `json:"port" yaml:"port"`
}

// Construction properties for a DatabaseInstanceNew.
//
// TODO: EXAMPLE
//
// Experimental.
type DatabaseInstanceProps struct {
	// The Neptune database cluster the instance should launch into.
	// Experimental.
	Cluster IDatabaseCluster `json:"cluster" yaml:"cluster"`
	// What type of instance to start for the replicas.
	// Experimental.
	InstanceType InstanceType `json:"instanceType" yaml:"instanceType"`
	// The name of the Availability Zone where the DB instance will be located.
	// Experimental.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	// Experimental.
	DbInstanceName *string `json:"dbInstanceName" yaml:"dbInstanceName"`
	// The DB parameter group to associate with the instance.
	// Experimental.
	ParameterGroup IParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
}

// Connection endpoint of a neptune cluster or instance.
//
// Consists of a combination of hostname and port.
//
// TODO: EXAMPLE
//
// Experimental.
type Endpoint interface {
	Hostname() *string
	Port() *float64
	SocketAddress() *string
}

// The jsii proxy struct for Endpoint
type jsiiProxy_Endpoint struct {
	_ byte // padding
}

func (j *jsiiProxy_Endpoint) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) SocketAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketAddress",
		&returns,
	)
	return returns
}


// Experimental.
func NewEndpoint(address *string, port *float64) Endpoint {
	_init_.Initialize()

	j := jsiiProxy_Endpoint{}

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.Endpoint",
		[]interface{}{address, port},
		&j,
	)

	return &j
}

// Experimental.
func NewEndpoint_Override(e Endpoint, address *string, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.Endpoint",
		[]interface{}{address, port},
		e,
	)
}

// Possible Instances Types to use in Neptune cluster used for defining {@link DatabaseClusterProps.engineVersion}.
//
// TODO: EXAMPLE
//
// Experimental.
type EngineVersion interface {
	Version() *string
}

// The jsii proxy struct for EngineVersion
type jsiiProxy_EngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_EngineVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Constructor for specifying a custom engine version.
// Experimental.
func NewEngineVersion(version *string) EngineVersion {
	_init_.Initialize()

	j := jsiiProxy_EngineVersion{}

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		[]interface{}{version},
		&j,
	)

	return &j
}

// Constructor for specifying a custom engine version.
// Experimental.
func NewEngineVersion_Override(e EngineVersion, version *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		[]interface{}{version},
		e,
	)
}

func EngineVersion_V1_0_1_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_1_0",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_1_1() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_1_1",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_1_2() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_1_2",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_2_1() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_2_1",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_2_2() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_2_2",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_3_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_3_0",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_4_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_4_0",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_4_1() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_4_1",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_5_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_5_0",
		&returns,
	)
	return returns
}

func EngineVersion_V1_1_0_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_1_0_0",
		&returns,
	)
	return returns
}

// A parameter group.
// Experimental.
type IClusterParameterGroup interface {
	awscdk.IResource
	// The name of this parameter group.
	// Experimental.
	ClusterParameterGroupName() *string
}

// The jsii proxy for IClusterParameterGroup
type jsiiProxy_IClusterParameterGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IClusterParameterGroup) ClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterParameterGroupName",
		&returns,
	)
	return returns
}

// Create a clustered database with a given number of instances.
// Experimental.
type IDatabaseCluster interface {
	awsec2.IConnectable
	awscdk.IResource
	// Grant the given identity connection access to the database.
	// Experimental.
	GrantConnect(grantee awsiam.IGrantable) awsiam.Grant
	// The endpoint to use for read/write operations.
	// Experimental.
	ClusterEndpoint() Endpoint
	// Identifier of the cluster.
	// Experimental.
	ClusterIdentifier() *string
	// Endpoint to use for load-balanced read-only operations.
	// Experimental.
	ClusterReadEndpoint() Endpoint
	// Resource identifier of the cluster.
	// Experimental.
	ClusterResourceIdentifier() *string
}

// The jsii proxy for IDatabaseCluster
type jsiiProxy_IDatabaseCluster struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDatabaseCluster) GrantConnect(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantConnect",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IDatabaseCluster) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) ClusterResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// A database instance.
// Experimental.
type IDatabaseInstance interface {
	awscdk.IResource
	// The instance endpoint address.
	// Experimental.
	DbInstanceEndpointAddress() *string
	// The instance endpoint port.
	// Experimental.
	DbInstanceEndpointPort() *string
	// The instance endpoint.
	// Experimental.
	InstanceEndpoint() Endpoint
	// The instance identifier.
	// Experimental.
	InstanceIdentifier() *string
}

// The jsii proxy for IDatabaseInstance
type jsiiProxy_IDatabaseInstance struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDatabaseInstance) DbInstanceEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) DbInstanceEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdentifier",
		&returns,
	)
	return returns
}

// A parameter group.
// Experimental.
type IParameterGroup interface {
	awscdk.IResource
	// The name of this parameter group.
	// Experimental.
	ParameterGroupName() *string
}

// The jsii proxy for IParameterGroup
type jsiiProxy_IParameterGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IParameterGroup) ParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupName",
		&returns,
	)
	return returns
}

// Interface for a subnet group.
// Experimental.
type ISubnetGroup interface {
	awscdk.IResource
	// The name of the subnet group.
	// Experimental.
	SubnetGroupName() *string
}

// The jsii proxy for ISubnetGroup
type jsiiProxy_ISubnetGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ISubnetGroup) SubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupName",
		&returns,
	)
	return returns
}

// Possible Instances Types to use in Neptune cluster used for defining {@link DatabaseInstanceProps.instanceType}.
//
// TODO: EXAMPLE
//
// Experimental.
type InstanceType interface {
}

// The jsii proxy struct for InstanceType
type jsiiProxy_InstanceType struct {
	_ byte // padding
}

// Build an InstanceType from given string or token, such as CfnParameter.
// Experimental.
func InstanceType_Of(instanceType *string) InstanceType {
	_init_.Initialize()

	var returns InstanceType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"of",
		[]interface{}{instanceType},
		&returns,
	)

	return returns
}

func InstanceType_R4_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R4_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R4_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R4_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R4_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R4_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R4_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R4_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_R4_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R4_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_24XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_24XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R5_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_16XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_16XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_R6G_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"R6G_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_T3_MEDIUM() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"T3_MEDIUM",
		&returns,
	)
	return returns
}

func InstanceType_T4G_MEDIUM() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.InstanceType",
		"T4G_MEDIUM",
		&returns,
	)
	return returns
}

// DB parameter group.
//
// TODO: EXAMPLE
//
// Experimental.
type ParameterGroup interface {
	awscdk.Resource
	IParameterGroup
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	ParameterGroupName() *string
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
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

func (j *jsiiProxy_ParameterGroup) ParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupName",
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


// Experimental.
func NewParameterGroup(scope constructs.Construct, id *string, props *ParameterGroupProps) ParameterGroup {
	_init_.Initialize()

	j := jsiiProxy_ParameterGroup{}

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.ParameterGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewParameterGroup_Override(p ParameterGroup, scope constructs.Construct, id *string, props *ParameterGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.ParameterGroup",
		[]interface{}{scope, id, props},
		p,
	)
}

// Imports a parameter group.
// Experimental.
func ParameterGroup_FromParameterGroupName(scope constructs.Construct, id *string, parameterGroupName *string) IParameterGroup {
	_init_.Initialize()

	var returns IParameterGroup

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.ParameterGroup",
		"fromParameterGroupName",
		[]interface{}{scope, id, parameterGroupName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ParameterGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.ParameterGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ParameterGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.ParameterGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (p *jsiiProxy_ParameterGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (p *jsiiProxy_ParameterGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (p *jsiiProxy_ParameterGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
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

// Marker class for cluster parameter group.
//
// TODO: EXAMPLE
//
// Experimental.
type ParameterGroupProps struct {
	// The parameters in this parameter group.
	// Experimental.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// Description for this parameter group.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// The name of the parameter group.
	// Experimental.
	ParameterGroupName *string `json:"parameterGroupName" yaml:"parameterGroupName"`
}

// Class for creating a RDS DB subnet group.
//
// TODO: EXAMPLE
//
// Experimental.
type SubnetGroup interface {
	awscdk.Resource
	ISubnetGroup
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	SubnetGroupName() *string
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for SubnetGroup
type jsiiProxy_SubnetGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_ISubnetGroup
}

func (j *jsiiProxy_SubnetGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubnetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubnetGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubnetGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubnetGroup) SubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupName",
		&returns,
	)
	return returns
}


// Experimental.
func NewSubnetGroup(scope constructs.Construct, id *string, props *SubnetGroupProps) SubnetGroup {
	_init_.Initialize()

	j := jsiiProxy_SubnetGroup{}

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.SubnetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSubnetGroup_Override(s SubnetGroup, scope constructs.Construct, id *string, props *SubnetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.SubnetGroup",
		[]interface{}{scope, id, props},
		s,
	)
}

// Imports an existing subnet group by name.
// Experimental.
func SubnetGroup_FromSubnetGroupName(scope constructs.Construct, id *string, subnetGroupName *string) ISubnetGroup {
	_init_.Initialize()

	var returns ISubnetGroup

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.SubnetGroup",
		"fromSubnetGroupName",
		[]interface{}{scope, id, subnetGroupName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SubnetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.SubnetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func SubnetGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-neptune-alpha.SubnetGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (s *jsiiProxy_SubnetGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (s *jsiiProxy_SubnetGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (s *jsiiProxy_SubnetGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (s *jsiiProxy_SubnetGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_SubnetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for creating a SubnetGroup.
//
// TODO: EXAMPLE
//
// Experimental.
type SubnetGroupProps struct {
	// The VPC to place the subnet group in.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Description of the subnet group.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// The removal policy to apply when the subnet group are removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// The name of the subnet group.
	// Experimental.
	SubnetGroupName *string `json:"subnetGroupName" yaml:"subnetGroupName"`
	// Which subnets within the VPC to associate with this group.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

