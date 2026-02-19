package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Manages mapping between IAM users and roles to Kubernetes RBAC configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster Cluster
//
//   awsAuth := awscdk.Aws_eks.NewAwsAuth(this, jsii.String("MyAwsAuth"), &AwsAuthProps{
//   	Cluster: cluster,
//   })
//
// See: https://docs.aws.amazon.com/en_us/eks/latest/userguide/add-user-role.html
//
type AwsAuth interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Additional AWS account to add to the aws-auth configmap.
	AddAccount(accountId *string)
	// Adds the specified IAM role to the `system:masters` RBAC group, which means that anyone that can assume it will be able to administer this Kubernetes system.
	AddMastersRole(role awsiam.IRole, username *string)
	// Adds a mapping between an IAM role to a Kubernetes user and groups.
	AddRoleMapping(role awsiam.IRole, mapping *AwsAuthMapping)
	// Adds a mapping between an IAM user to a Kubernetes user and groups.
	AddUserMapping(user awsiam.IUser, mapping *AwsAuthMapping)
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

// The jsii proxy struct for AwsAuth
type jsiiProxy_AwsAuth struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AwsAuth) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewAwsAuth(scope constructs.Construct, id *string, props *AwsAuthProps) AwsAuth {
	_init_.Initialize()

	if err := validateNewAwsAuthParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAuth{}

	_jsii_.Create(
		"aws-cdk-lib.aws_eks.AwsAuth",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAwsAuth_Override(a AwsAuth, scope constructs.Construct, id *string, props *AwsAuthProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_eks.AwsAuth",
		[]interface{}{scope, id, props},
		a,
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
func AwsAuth_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAuth_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.AwsAuth",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuth) AddAccount(accountId *string) {
	if err := a.validateAddAccountParameters(accountId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addAccount",
		[]interface{}{accountId},
	)
}

func (a *jsiiProxy_AwsAuth) AddMastersRole(role awsiam.IRole, username *string) {
	if err := a.validateAddMastersRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMastersRole",
		[]interface{}{role, username},
	)
}

func (a *jsiiProxy_AwsAuth) AddRoleMapping(role awsiam.IRole, mapping *AwsAuthMapping) {
	if err := a.validateAddRoleMappingParameters(role, mapping); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addRoleMapping",
		[]interface{}{role, mapping},
	)
}

func (a *jsiiProxy_AwsAuth) AddUserMapping(user awsiam.IUser, mapping *AwsAuthMapping) {
	if err := a.validateAddUserMappingParameters(user, mapping); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addUserMapping",
		[]interface{}{user, mapping},
	)
}

func (a *jsiiProxy_AwsAuth) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuth) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

