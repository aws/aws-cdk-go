package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A new receipt rule set.
//
// Example:
//   ruleSet := ses.NewReceiptRuleSet(this, jsii.String("RuleSet"))
//
//   awsRule := ruleSet.addRule(jsii.String("Aws"), &ReceiptRuleOptions{
//   	Recipients: []*string{
//   		jsii.String("aws.com"),
//   	},
//   })
//
type ReceiptRuleSet interface {
	awscdk.Resource
	IReceiptRuleSet
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
	// The receipt rule set name.
	ReceiptRuleSetName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds a drop spam rule.
	AddDropSpamRule()
	// Adds a new receipt rule in this rule set.
	//
	// The new rule is added after
	// the last added rule unless `after` is specified.
	AddRule(id *string, options *ReceiptRuleOptions) ReceiptRule
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

// The jsii proxy struct for ReceiptRuleSet
type jsiiProxy_ReceiptRuleSet struct {
	internal.Type__awscdkResource
	jsiiProxy_IReceiptRuleSet
}

func (j *jsiiProxy_ReceiptRuleSet) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRuleSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRuleSet) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRuleSet) ReceiptRuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiptRuleSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRuleSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewReceiptRuleSet(scope constructs.Construct, id *string, props *ReceiptRuleSetProps) ReceiptRuleSet {
	_init_.Initialize()

	if err := validateNewReceiptRuleSetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReceiptRuleSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.ReceiptRuleSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewReceiptRuleSet_Override(r ReceiptRuleSet, scope constructs.Construct, id *string, props *ReceiptRuleSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.ReceiptRuleSet",
		[]interface{}{scope, id, props},
		r,
	)
}

// Import an exported receipt rule set.
func ReceiptRuleSet_FromReceiptRuleSetName(scope constructs.Construct, id *string, receiptRuleSetName *string) IReceiptRuleSet {
	_init_.Initialize()

	if err := validateReceiptRuleSet_FromReceiptRuleSetNameParameters(scope, id, receiptRuleSetName); err != nil {
		panic(err)
	}
	var returns IReceiptRuleSet

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ReceiptRuleSet",
		"fromReceiptRuleSetName",
		[]interface{}{scope, id, receiptRuleSetName},
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
func ReceiptRuleSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateReceiptRuleSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ReceiptRuleSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ReceiptRuleSet_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateReceiptRuleSet_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ReceiptRuleSet",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ReceiptRuleSet_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateReceiptRuleSet_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ReceiptRuleSet",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func ReceiptRuleSet_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses.ReceiptRuleSet",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_ReceiptRuleSet) AddDropSpamRule() {
	_jsii_.InvokeVoid(
		r,
		"addDropSpamRule",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReceiptRuleSet) AddRule(id *string, options *ReceiptRuleOptions) ReceiptRule {
	if err := r.validateAddRuleParameters(id, options); err != nil {
		panic(err)
	}
	var returns ReceiptRule

	_jsii_.Invoke(
		r,
		"addRule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReceiptRuleSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := r.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_ReceiptRuleSet) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReceiptRuleSet) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (r *jsiiProxy_ReceiptRuleSet) GetResourceNameAttribute(nameAttr *string) *string {
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

func (r *jsiiProxy_ReceiptRuleSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

