package awscdkmskalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmskalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a MSK Cluster.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
//   	ClusterName: jsii.String("myCluster"),
//   	KafkaVersion: msk.KafkaVersion_V4_0_X_KRAFT(),
//   	Vpc: Vpc,
//   	EncryptionInTransit: &EncryptionInTransitConfig{
//   		ClientBroker: msk.ClientBrokerEncryption_TLS,
//   	},
//   	ClientAuthentication: msk.ClientAuthentication_Sasl(&SaslAuthProps{
//   		Scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type Cluster interface {
	ClusterBase
	// Get the list of brokers that a client application can use to bootstrap.
	//
	// Uses a Custom Resource to make an API call to `getBootstrapBrokers` using the Javascript SDK.
	//
	// Returns: - A string containing one or more hostname:port pairs.
	// Experimental.
	BootstrapBrokers() *string
	// Get the list of brokers that a SASL/IAM authenticated client application can use to bootstrap.
	//
	// Uses a Custom Resource to make an API call to `getBootstrapBrokers` using the Javascript SDK.
	//
	// Returns: - A string containing one or more DNS names (or IP) and TLS port pairs.
	// Experimental.
	BootstrapBrokersSaslIam() *string
	// Get the list of brokers that a SASL/SCRAM authenticated client application can use to bootstrap.
	//
	// Uses a Custom Resource to make an API call to `getBootstrapBrokers` using the Javascript SDK.
	//
	// Returns: - A string containing one or more dns name (or IP) and SASL SCRAM port pairs.
	// Experimental.
	BootstrapBrokersSaslScram() *string
	// Get the list of brokers that a TLS authenticated client application can use to bootstrap.
	//
	// Uses a Custom Resource to make an API call to `getBootstrapBrokers` using the Javascript SDK.
	//
	// Returns: - A string containing one or more DNS names (or IP) and TLS port pairs.
	// Experimental.
	BootstrapBrokersTls() *string
	// The ARN of cluster.
	// Experimental.
	ClusterArn() *string
	// The physical name of the cluster.
	// Experimental.
	ClusterName() *string
	// Manages connections for the cluster.
	// Experimental.
	Connections() awsec2.Connections
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// Key used to encrypt SASL/SCRAM users.
	// Experimental.
	SaslScramAuthenticationKey() awskms.IKey
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Get the ZooKeeper Connection string.
	//
	// Uses a Custom Resource to make an API call to `describeCluster` using the Javascript SDK.
	//
	// Returns: - The connection string to use to connect to the Apache ZooKeeper cluster.
	// Experimental.
	ZookeeperConnectionString() *string
	// Get the ZooKeeper Connection string for a TLS enabled cluster.
	//
	// Uses a Custom Resource to make an API call to `describeCluster` using the Javascript SDK.
	//
	// Returns: - The connection string to use to connect to zookeeper cluster on TLS port.
	// Experimental.
	ZookeeperConnectionStringTls() *string
	// A list of usersnames to register with the cluster.
	//
	// The password will automatically be generated using Secrets
	// Manager and the { username, password } JSON object stored in Secrets Manager as `AmazonMSK_username`.
	//
	// Must be using the SASL/SCRAM authentication mechanism.
	// Experimental.
	AddUser(usernames ...*string)
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Cluster
type jsiiProxy_Cluster struct {
	jsiiProxy_ClusterBase
}

func (j *jsiiProxy_Cluster) BootstrapBrokers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BootstrapBrokersSaslIam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersSaslIam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BootstrapBrokersSaslScram() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersSaslScram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BootstrapBrokersTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SaslScramAuthenticationKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"saslScramAuthenticationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ZookeeperConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zookeeperConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ZookeeperConnectionStringTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zookeeperConnectionStringTls",
		&returns,
	)
	return returns
}


// Experimental.
func NewCluster(scope constructs.Construct, id *string, props *ClusterProps) Cluster {
	_init_.Initialize()

	if err := validateNewClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cluster{}

	_jsii_.Create(
		"@aws-cdk/aws-msk-alpha.Cluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCluster_Override(c Cluster, scope constructs.Construct, id *string, props *ClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-msk-alpha.Cluster",
		[]interface{}{scope, id, props},
		c,
	)
}

// Reference an existing cluster, defined outside of the CDK code, by name.
// Experimental.
func Cluster_FromClusterArn(scope constructs.Construct, id *string, clusterArn *string) ICluster {
	_init_.Initialize()

	if err := validateCluster_FromClusterArnParameters(scope, id, clusterArn); err != nil {
		panic(err)
	}
	var returns ICluster

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.Cluster",
		"fromClusterArn",
		[]interface{}{scope, id, clusterArn},
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
// Experimental.
func Cluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.Cluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Cluster_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCluster_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.Cluster",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Cluster_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCluster_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.Cluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Cluster_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.Cluster",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_Cluster) AddUser(usernames ...*string) {
	args := []interface{}{}
	for _, a := range usernames {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addUser",
		args,
	)
}

func (c *jsiiProxy_Cluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_Cluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := c.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetResourceNameAttribute(nameAttr *string) *string {
	if err := c.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

