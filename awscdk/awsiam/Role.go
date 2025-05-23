package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// IAM Role.
//
// Defines an IAM role. The role is created with an assume policy document associated with
// the specified AWS service principal defined in `serviceAssumeRole`.
//
// Example:
//   // Option 3: Create a new role that allows the account root principal to assume. Add this role in the `system:masters` and witch to this role from the AWS console.
//   var cluster cluster
//
//
//   consoleReadOnlyRole := iam.NewRole(this, jsii.String("ConsoleReadOnlyRole"), &RoleProps{
//   	AssumedBy: iam.NewArnPrincipal(jsii.String("arn_for_trusted_principal")),
//   })
//   consoleReadOnlyRole.AddToPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("eks:AccessKubernetesApi"),
//   		jsii.String("eks:Describe*"),
//   		jsii.String("eks:List*"),
//   	},
//   	Resources: []*string{
//   		cluster.ClusterArn,
//   	},
//   }))
//
//   // Add this role to system:masters RBAC group
//   cluster.awsAuth.AddMastersRole(consoleReadOnlyRole)
//
type Role interface {
	awscdk.Resource
	IRole
	// When this Principal is used in an AssumeRole policy, the action to use.
	AssumeRoleAction() *string
	// The assume role policy document associated with this role.
	AssumeRolePolicy() PolicyDocument
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
	// Returns the permissions boundary attached to this role.
	PermissionsBoundary() IManagedPolicy
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// Returns the role.
	PolicyFragment() PrincipalPolicyFragment
	// The AWS account ID of this principal.
	//
	// Can be undefined when the account is not known
	// (for example, for service principals).
	// Can be a Token - in that case,
	// it's assumed to be AWS::AccountId.
	PrincipalAccount() *string
	// Returns the ARN of this role.
	RoleArn() *string
	// Returns the stable and unique string identifying the role.
	//
	// For example,
	// AIDAJQABLZS4A3QDU576Q.
	RoleId() *string
	// Returns the name of the role.
	RoleName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Attaches a managed policy to this role.
	AddManagedPolicy(policy IManagedPolicy)
	// Add to the policy of this principal.
	AddToPolicy(statement PolicyStatement) *bool
	// Adds a permission to the role's default policy document.
	//
	// If there is no default policy attached to this role, it will be created.
	AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult
	// Skip applyRemovalPolicy if role synthesis is prevented by customizeRoles.
	//
	// Because in this case, this construct does not have a CfnResource in the tree.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Attaches a policy to this role.
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
	// Grant the actions defined in actions to the identity Principal on this resource.
	Grant(grantee IPrincipal, actions ...*string) Grant
	// Grant permissions to the given principal to assume this role.
	GrantAssumeRole(identity IPrincipal) Grant
	// Grant permissions to the given principal to pass this role.
	GrantPassRole(identity IPrincipal) Grant
	// Returns a string representation of this construct.
	ToString() *string
	// Return a copy of this Role object whose Policies will not be updated.
	//
	// Use the object returned by this method if you want this Role to be used by
	// a construct without it automatically updating the Role's Policies.
	//
	// If you do, you are responsible for adding the correct statements to the
	// Role's policies yourself.
	WithoutPolicyUpdates(options *WithoutPolicyUpdatesOptions) IRole
}

// The jsii proxy struct for Role
type jsiiProxy_Role struct {
	internal.Type__awscdkResource
	jsiiProxy_IRole
}

func (j *jsiiProxy_Role) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) AssumeRolePolicy() PolicyDocument {
	var returns PolicyDocument
	_jsii_.Get(
		j,
		"assumeRolePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) PermissionsBoundary() IManagedPolicy {
	var returns IManagedPolicy
	_jsii_.Get(
		j,
		"permissionsBoundary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) RoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewRole(scope constructs.Construct, id *string, props *RoleProps) Role {
	_init_.Initialize()

	if err := validateNewRoleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Role{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.Role",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRole_Override(r Role, scope constructs.Construct, id *string, props *RoleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.Role",
		[]interface{}{scope, id, props},
		r,
	)
}

// Customize the creation of IAM roles within the given scope.
//
// It is recommended that you **do not** use this method and instead allow
// CDK to manage role creation. This should only be used
// in environments where CDK applications are not allowed to created IAM roles.
//
// This can be used to prevent the CDK application from creating roles
// within the given scope and instead replace the references to the roles with
// precreated role names. A report will be synthesized in the cloud assembly (i.e. cdk.out)
// that will contain the list of IAM roles that would have been created along with the
// IAM policy statements that the role should contain. This report can then be used
// to create the IAM roles outside of CDK and then the created role names can be provided
// in `usePrecreatedRoles`.
//
// Example:
//   var app app
//
//   iam.Role_CustomizeRoles(app, &CustomizeRolesOptions{
//   	UsePrecreatedRoles: map[string]*string{
//   		"ConstructPath/To/Role": jsii.String("my-precreated-role-name"),
//   	},
//   })
//
func Role_CustomizeRoles(scope constructs.Construct, options *CustomizeRolesOptions) {
	_init_.Initialize()

	if err := validateRole_CustomizeRolesParameters(scope, options); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_iam.Role",
		"customizeRoles",
		[]interface{}{scope, options},
	)
}

// Lookup an existing Role.
func Role_FromLookup(scope constructs.Construct, id *string, options *RoleLookupOptions) IRole {
	_init_.Initialize()

	if err := validateRole_FromLookupParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns IRole

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Role",
		"fromLookup",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Import an external role by ARN.
//
// If the imported Role ARN is a Token (such as a
// `CfnParameter.valueAsString` or a `Fn.importValue()`) *and* the referenced
// role has a `path` (like `arn:...:role/AdminRoles/Alice`), the
// `roleName` property will not resolve to the correct value. Instead it
// will resolve to the first path component. We unfortunately cannot express
// the correct calculation of the full path name as a CloudFormation
// expression. In this scenario the Role ARN should be supplied without the
// `path` in order to resolve the correct role resource.
func Role_FromRoleArn(scope constructs.Construct, id *string, roleArn *string, options *FromRoleArnOptions) IRole {
	_init_.Initialize()

	if err := validateRole_FromRoleArnParameters(scope, id, roleArn, options); err != nil {
		panic(err)
	}
	var returns IRole

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Role",
		"fromRoleArn",
		[]interface{}{scope, id, roleArn, options},
		&returns,
	)

	return returns
}

// Import an external role by name.
//
// The imported role is assumed to exist in the same account as the account
// the scope's containing Stack is being deployed to.
func Role_FromRoleName(scope constructs.Construct, id *string, roleName *string, options *FromRoleNameOptions) IRole {
	_init_.Initialize()

	if err := validateRole_FromRoleNameParameters(scope, id, roleName, options); err != nil {
		panic(err)
	}
	var returns IRole

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Role",
		"fromRoleName",
		[]interface{}{scope, id, roleName, options},
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
func Role_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Role",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Role_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRole_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Role",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Role_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRole_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Role",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Role.
func Role_IsRole(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRole_IsRoleParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Role",
		"isRole",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Role_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iam.Role",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Role) AddManagedPolicy(policy IManagedPolicy) {
	if err := r.validateAddManagedPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addManagedPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_Role) AddToPolicy(statement PolicyStatement) *bool {
	if err := r.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		r,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := r.validateAddToPrincipalPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		r,
		"addToPrincipalPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := r.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_Role) AttachInlinePolicy(policy Policy) {
	if err := r.validateAttachInlinePolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"attachInlinePolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_Role) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := r.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) GetResourceNameAttribute(nameAttr *string) *string {
	if err := r.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) Grant(grantee IPrincipal, actions ...*string) Grant {
	if err := r.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns Grant

	_jsii_.Invoke(
		r,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) GrantAssumeRole(identity IPrincipal) Grant {
	if err := r.validateGrantAssumeRoleParameters(identity); err != nil {
		panic(err)
	}
	var returns Grant

	_jsii_.Invoke(
		r,
		"grantAssumeRole",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) GrantPassRole(identity IPrincipal) Grant {
	if err := r.validateGrantPassRoleParameters(identity); err != nil {
		panic(err)
	}
	var returns Grant

	_jsii_.Invoke(
		r,
		"grantPassRole",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) WithoutPolicyUpdates(options *WithoutPolicyUpdatesOptions) IRole {
	if err := r.validateWithoutPolicyUpdatesParameters(options); err != nil {
		panic(err)
	}
	var returns IRole

	_jsii_.Invoke(
		r,
		"withoutPolicyUpdates",
		[]interface{}{options},
		&returns,
	)

	return returns
}

