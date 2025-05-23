package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a new IAM user.
//
// Example:
//   var definition iChainable
//   user := iam.NewUser(this, jsii.String("MyUser"))
//   stateMachine := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
//   	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
//   })
//
//   //give user permission to send task success to the state machine
//   stateMachine.grant(user, jsii.String("states:SendTaskSuccess"))
//
type User interface {
	awscdk.Resource
	IIdentity
	IUser
	// When this Principal is used in an AssumeRole policy, the action to use.
	AssumeRoleAction() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The principal to grant permissions to.
	GrantPrincipal() IPrincipal
	// The tree node.
	Node() constructs.Node
	// Returns the permissions boundary attached  to this user.
	PermissionsBoundary() IManagedPolicy
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// Return the policy fragment that identifies this principal in a Policy.
	PolicyFragment() PrincipalPolicyFragment
	// The AWS account ID of this principal.
	//
	// Can be undefined when the account is not known
	// (for example, for service principals).
	// Can be a Token - in that case,
	// it's assumed to be AWS::AccountId.
	PrincipalAccount() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// An attribute that represents the user's ARN.
	UserArn() *string
	// An attribute that represents the user name.
	UserName() *string
	// Attaches a managed policy to the user.
	AddManagedPolicy(policy IManagedPolicy)
	// Adds this user to a group.
	AddToGroup(group IGroup)
	// Add to the policy of this principal.
	AddToPolicy(statement PolicyStatement) *bool
	// Adds an IAM statement to the default policy.
	//
	// Returns: true.
	AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult
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
	// Attaches a policy to this user.
	AttachInlinePolicy(policy Policy)
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

// The jsii proxy struct for User
type jsiiProxy_User struct {
	internal.Type__awscdkResource
	jsiiProxy_IIdentity
	jsiiProxy_IUser
}

func (j *jsiiProxy_User) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
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

func (j *jsiiProxy_User) PermissionsBoundary() IManagedPolicy {
	var returns IManagedPolicy
	_jsii_.Get(
		j,
		"permissionsBoundary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}


func NewUser(scope constructs.Construct, id *string, props *UserProps) User {
	_init_.Initialize()

	if err := validateNewUserParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_User{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.User",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewUser_Override(u User, scope constructs.Construct, id *string, props *UserProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.User",
		[]interface{}{scope, id, props},
		u,
	)
}

// Import an existing user given a user ARN.
//
// If the ARN comes from a Token, the User cannot have a path; if so, any attempt
// to reference its username will fail.
func User_FromUserArn(scope constructs.Construct, id *string, userArn *string) IUser {
	_init_.Initialize()

	if err := validateUser_FromUserArnParameters(scope, id, userArn); err != nil {
		panic(err)
	}
	var returns IUser

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.User",
		"fromUserArn",
		[]interface{}{scope, id, userArn},
		&returns,
	)

	return returns
}

// Import an existing user given user attributes.
//
// If the ARN comes from a Token, the User cannot have a path; if so, any attempt
// to reference its username will fail.
func User_FromUserAttributes(scope constructs.Construct, id *string, attrs *UserAttributes) IUser {
	_init_.Initialize()

	if err := validateUser_FromUserAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IUser

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.User",
		"fromUserAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing user given a username.
func User_FromUserName(scope constructs.Construct, id *string, userName *string) IUser {
	_init_.Initialize()

	if err := validateUser_FromUserNameParameters(scope, id, userName); err != nil {
		panic(err)
	}
	var returns IUser

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.User",
		"fromUserName",
		[]interface{}{scope, id, userName},
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
func User_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.User",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func User_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateUser_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.User",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func User_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateUser_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.User",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func User_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iam.User",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_User) AddManagedPolicy(policy IManagedPolicy) {
	if err := u.validateAddManagedPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addManagedPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_User) AddToGroup(group IGroup) {
	if err := u.validateAddToGroupParameters(group); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addToGroup",
		[]interface{}{group},
	)
}

func (u *jsiiProxy_User) AddToPolicy(statement PolicyStatement) *bool {
	if err := u.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		u,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := u.validateAddToPrincipalPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		u,
		"addToPrincipalPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := u.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_User) AttachInlinePolicy(policy Policy) {
	if err := u.validateAttachInlinePolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"attachInlinePolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_User) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := u.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetResourceNameAttribute(nameAttr *string) *string {
	if err := u.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
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

