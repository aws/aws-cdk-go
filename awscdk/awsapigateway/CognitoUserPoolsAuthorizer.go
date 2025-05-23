package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Cognito user pools based custom authorizer.
//
// Example:
//   var books resource
//   userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
//
//   auth := apigateway.NewCognitoUserPoolsAuthorizer(this, jsii.String("booksAuthorizer"), &CognitoUserPoolsAuthorizerProps{
//   	CognitoUserPools: []iUserPool{
//   		userPool,
//   	},
//   })
//   books.AddMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &MethodOptions{
//   	Authorizer: auth,
//   	AuthorizationType: apigateway.AuthorizationType_COGNITO,
//   })
//
type CognitoUserPoolsAuthorizer interface {
	Authorizer
	IAuthorizer
	// The authorization type of this authorizer.
	AuthorizationType() AuthorizationType
	// The ARN of the authorizer to be used in permission policies, such as IAM and resource-based grants.
	AuthorizerArn() *string
	// The id of the authorizer.
	AuthorizerId() *string
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
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CognitoUserPoolsAuthorizer
type jsiiProxy_CognitoUserPoolsAuthorizer struct {
	jsiiProxy_Authorizer
	jsiiProxy_IAuthorizer
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) AuthorizationType() AuthorizationType {
	var returns AuthorizationType
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) AuthorizerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolsAuthorizer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCognitoUserPoolsAuthorizer(scope constructs.Construct, id *string, props *CognitoUserPoolsAuthorizerProps) CognitoUserPoolsAuthorizer {
	_init_.Initialize()

	if err := validateNewCognitoUserPoolsAuthorizerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUserPoolsAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CognitoUserPoolsAuthorizer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCognitoUserPoolsAuthorizer_Override(c CognitoUserPoolsAuthorizer, scope constructs.Construct, id *string, props *CognitoUserPoolsAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.CognitoUserPoolsAuthorizer",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is an Authorizer.
func CognitoUserPoolsAuthorizer_IsAuthorizer(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoUserPoolsAuthorizer_IsAuthorizerParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CognitoUserPoolsAuthorizer",
		"isAuthorizer",
		[]interface{}{x},
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
func CognitoUserPoolsAuthorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoUserPoolsAuthorizer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CognitoUserPoolsAuthorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func CognitoUserPoolsAuthorizer_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCognitoUserPoolsAuthorizer_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CognitoUserPoolsAuthorizer",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CognitoUserPoolsAuthorizer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCognitoUserPoolsAuthorizer_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.CognitoUserPoolsAuthorizer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func CognitoUserPoolsAuthorizer_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigateway.CognitoUserPoolsAuthorizer",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CognitoUserPoolsAuthorizer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CognitoUserPoolsAuthorizer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolsAuthorizer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (c *jsiiProxy_CognitoUserPoolsAuthorizer) GetResourceNameAttribute(nameAttr *string) *string {
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

func (c *jsiiProxy_CognitoUserPoolsAuthorizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

