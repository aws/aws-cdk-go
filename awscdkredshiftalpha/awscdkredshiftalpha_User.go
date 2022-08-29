// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkredshiftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdkredshiftalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A user in a Redshift cluster.
//
// Example:
//   user := awscdkredshiftalpha.NewUser(this, jsii.String("User"), &userProps{
//   	cluster: cluster,
//   	databaseName: jsii.String("databaseName"),
//   })
//   cluster.addRotationMultiUser(jsii.String("MultiUserRotation"), &rotationMultiUserOptions{
//   	secret: user.secret,
//   })
//
// Experimental.
type User interface {
	constructs.Construct
	IUser
	// The cluster where the table is located.
	// Experimental.
	Cluster() ICluster
	// The name of the database where the table is located.
	// Experimental.
	DatabaseName() *string
	// Experimental.
	DatabaseProps() *DatabaseOptions
	// Experimental.
	SetDatabaseProps(val *DatabaseOptions)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The password of the user.
	// Experimental.
	Password() awscdk.SecretValue
	// The Secrets Manager secret of the user.
	// Experimental.
	Secret() awssecretsmanager.ISecret
	// The name of the user.
	// Experimental.
	Username() *string
	// Grant this user privilege to access a table.
	// Experimental.
	AddTablePrivileges(table ITable, actions ...TableAction)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be destroyed (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	//
	// This resource is destroyed by default.
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for User
type jsiiProxy_User struct {
	internal.Type__constructsConstruct
	jsiiProxy_IUser
}

func (j *jsiiProxy_User) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DatabaseProps() *DatabaseOptions {
	var returns *DatabaseOptions
	_jsii_.Get(
		j,
		"databaseProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Password() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}


// Experimental.
func NewUser(scope constructs.Construct, id *string, props *UserProps) User {
	_init_.Initialize()

	j := jsiiProxy_User{}

	_jsii_.Create(
		"@aws-cdk/aws-redshift-alpha.User",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewUser_Override(u User, scope constructs.Construct, id *string, props *UserProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-redshift-alpha.User",
		[]interface{}{scope, id, props},
		u,
	)
}

func (j *jsiiProxy_User) SetDatabaseProps(val *DatabaseOptions) {
	_jsii_.Set(
		j,
		"databaseProps",
		val,
	)
}

// Specify a Redshift user using credentials that already exist.
// Experimental.
func User_FromUserAttributes(scope constructs.Construct, id *string, attrs *UserAttributes) IUser {
	_init_.Initialize()

	var returns IUser

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-redshift-alpha.User",
		"fromUserAttributes",
		[]interface{}{scope, id, attrs},
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
func User_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-redshift-alpha.User",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) AddTablePrivileges(table ITable, actions ...TableAction) {
	args := []interface{}{table}
	for _, a := range actions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		u,
		"addTablePrivileges",
		args,
	)
}

func (u *jsiiProxy_User) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_User) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

