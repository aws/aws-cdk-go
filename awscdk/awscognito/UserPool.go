package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a Cognito User Pool.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//   pool.addClient(jsii.String("app-client"), &UserPoolClientOptions{
//   	OAuth: &OAuthSettings{
//   		Flows: &OAuthFlows{
//   			AuthorizationCodeGrant: jsii.Boolean(true),
//   		},
//   		Scopes: []oAuthScope{
//   			cognito.*oAuthScope_OPENID(),
//   		},
//   		CallbackUrls: []*string{
//   			jsii.String("https://my-app-domain.com/welcome"),
//   		},
//   		LogoutUrls: []*string{
//   			jsii.String("https://my-app-domain.com/signin"),
//   		},
//   	},
//   })
//
type UserPool interface {
	awscdk.Resource
	IUserPool
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Get all identity providers registered with this user pool.
	IdentityProviders() *[]IUserPoolIdentityProvider
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
	// The ARN of the user pool.
	UserPoolArn() *string
	// The physical ID of this user pool resource.
	UserPoolId() *string
	// User pool provider name.
	UserPoolProviderName() *string
	// User pool provider URL.
	UserPoolProviderUrl() *string
	// Add a new app client to this user pool.
	AddClient(id *string, options *UserPoolClientOptions) UserPoolClient
	// Associate a domain to this user pool.
	AddDomain(id *string, options *UserPoolDomainOptions) UserPoolDomain
	// Add a new group to this user pool.
	AddGroup(id *string, options *UserPoolGroupOptions) UserPoolGroup
	// Add a new resource server to this user pool.
	AddResourceServer(id *string, options *UserPoolResourceServerOptions) UserPoolResourceServer
	// Add a lambda trigger to a user pool operation.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html
	//
	AddTrigger(operation UserPoolOperation, fn awslambda.IFunction, lambdaVersion LambdaVersion)
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
	// Adds an IAM policy statement associated with this user pool to an IAM principal's policy.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Register an identity provider with this user pool.
	RegisterIdentityProvider(provider IUserPoolIdentityProvider)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for UserPool
type jsiiProxy_UserPool struct {
	internal.Type__awscdkResource
	jsiiProxy_IUserPool
}

func (j *jsiiProxy_UserPool) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) IdentityProviders() *[]IUserPoolIdentityProvider {
	var returns *[]IUserPoolIdentityProvider
	_jsii_.Get(
		j,
		"identityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) UserPoolArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) UserPoolProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPool) UserPoolProviderUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolProviderUrl",
		&returns,
	)
	return returns
}


func NewUserPool(scope constructs.Construct, id *string, props *UserPoolProps) UserPool {
	_init_.Initialize()

	if err := validateNewUserPoolParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserPool{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.UserPool",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewUserPool_Override(u UserPool, scope constructs.Construct, id *string, props *UserPoolProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.UserPool",
		[]interface{}{scope, id, props},
		u,
	)
}

// Import an existing user pool based on its ARN.
func UserPool_FromUserPoolArn(scope constructs.Construct, id *string, userPoolArn *string) IUserPool {
	_init_.Initialize()

	if err := validateUserPool_FromUserPoolArnParameters(scope, id, userPoolArn); err != nil {
		panic(err)
	}
	var returns IUserPool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPool",
		"fromUserPoolArn",
		[]interface{}{scope, id, userPoolArn},
		&returns,
	)

	return returns
}

// Import an existing user pool based on its id.
func UserPool_FromUserPoolId(scope constructs.Construct, id *string, userPoolId *string) IUserPool {
	_init_.Initialize()

	if err := validateUserPool_FromUserPoolIdParameters(scope, id, userPoolId); err != nil {
		panic(err)
	}
	var returns IUserPool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPool",
		"fromUserPoolId",
		[]interface{}{scope, id, userPoolId},
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
func UserPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUserPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func UserPool_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateUserPool_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPool",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func UserPool_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateUserPool_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPool",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func UserPool_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.UserPool",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_UserPool) AddClient(id *string, options *UserPoolClientOptions) UserPoolClient {
	if err := u.validateAddClientParameters(id, options); err != nil {
		panic(err)
	}
	var returns UserPoolClient

	_jsii_.Invoke(
		u,
		"addClient",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) AddDomain(id *string, options *UserPoolDomainOptions) UserPoolDomain {
	if err := u.validateAddDomainParameters(id, options); err != nil {
		panic(err)
	}
	var returns UserPoolDomain

	_jsii_.Invoke(
		u,
		"addDomain",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) AddGroup(id *string, options *UserPoolGroupOptions) UserPoolGroup {
	if err := u.validateAddGroupParameters(id, options); err != nil {
		panic(err)
	}
	var returns UserPoolGroup

	_jsii_.Invoke(
		u,
		"addGroup",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) AddResourceServer(id *string, options *UserPoolResourceServerOptions) UserPoolResourceServer {
	if err := u.validateAddResourceServerParameters(id, options); err != nil {
		panic(err)
	}
	var returns UserPoolResourceServer

	_jsii_.Invoke(
		u,
		"addResourceServer",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) AddTrigger(operation UserPoolOperation, fn awslambda.IFunction, lambdaVersion LambdaVersion) {
	if err := u.validateAddTriggerParameters(operation, fn); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addTrigger",
		[]interface{}{operation, fn, lambdaVersion},
	)
}

func (u *jsiiProxy_UserPool) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := u.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_UserPool) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (u *jsiiProxy_UserPool) GetResourceNameAttribute(nameAttr *string) *string {
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

func (u *jsiiProxy_UserPool) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := u.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		u,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPool) RegisterIdentityProvider(provider IUserPoolIdentityProvider) {
	if err := u.validateRegisterIdentityProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"registerIdentityProvider",
		[]interface{}{provider},
	)
}

func (u *jsiiProxy_UserPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

