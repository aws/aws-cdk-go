package awss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The bucket policy for an Amazon S3 bucket.
//
// Policies define the operations that are allowed on this resource.
//
// You almost never need to define this construct directly.
//
// All AWS resources that support resource policies have a method called
// `addToResourcePolicy()`, which will automatically create a new resource
// policy if one doesn't exist yet, otherwise it will add to the existing
// policy.
//
// The bucket policy method is implemented differently than `addToResourcePolicy()`
// as `BucketPolicy()` creates a new policy without knowing one earlier existed.
// e.g. if during Bucket creation, if `autoDeleteObject:true`, these policies are
// added to the bucket policy:
//    ["s3:DeleteObject*", "s3:GetBucket*", "s3:List*", "s3:PutBucketPolicy"],
// and when you add a new BucketPolicy with ["s3:GetObject", "s3:ListBucket"] on
// this existing bucket, invoking `BucketPolicy()` will create a new Policy
// without knowing one earlier exists already, so it creates a new one.
// In this case, the custom resource handler will not have access to
// `s3:GetBucketTagging` action which will cause failure during deletion of stack.
//
// Hence its strongly recommended to use `addToResourcePolicy()` method to add
// new permissions to existing policy.
//
// Example:
//   bucketName := "amzn-s3-demo-bucket"
//   accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"), &BucketProps{
//   	ObjectOwnership: s3.ObjectOwnership_BUCKET_OWNER_ENFORCED,
//   	BucketName: jsii.String(BucketName),
//   })
//
//   bucketPolicy := s3.NewCfnBucketPolicy(this, jsii.String("BucketPolicy"), &CfnBucketPolicyProps{
//   	Bucket: bucketName,
//   	PolicyDocument: map[string]interface{}{
//   		"Statement": []map[string]interface{}{
//   			map[string]interface{}{
//   				"Action": jsii.String("s3:*"),
//   				"Effect": jsii.String("Deny"),
//   				"Principal": map[string]*string{
//   					"AWS": jsii.String("*"),
//   				},
//   				"Resource": []*string{
//   					accessLogsBucket.bucketArn,
//   					fmt.Sprintf("%v/*", accessLogsBucket.bucketArn),
//   				},
//   			},
//   		},
//   		"Version": jsii.String("2012-10-17"),
//   	},
//   })
//
//   // Wrap L1 Construct with L2 Bucket Policy Construct. Subsequent
//   // generated bucket policy to allow access log delivery would append
//   // to the current policy.
//   s3.BucketPolicy_FromCfnBucketPolicy(bucketPolicy)
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
//   	ServerAccessLogsBucket: accessLogsBucket,
//   	ServerAccessLogsPrefix: jsii.String("logs"),
//   })
//
type BucketPolicy interface {
	awscdk.Resource
	// The Bucket this Policy applies to.
	Bucket() IBucket
	// A policy document containing permissions to add to the specified bucket.
	//
	// For more information, see Access Policy Language Overview in the Amazon
	// Simple Storage Service Developer Guide.
	Document() awsiam.PolicyDocument
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
	// Sets the removal policy for the BucketPolicy.
	ApplyRemovalPolicy(removalPolicy awscdk.RemovalPolicy)
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

// The jsii proxy struct for BucketPolicy
type jsiiProxy_BucketPolicy struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_BucketPolicy) Bucket() IBucket {
	var returns IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketPolicy) Document() awsiam.PolicyDocument {
	var returns awsiam.PolicyDocument
	_jsii_.Get(
		j,
		"document",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketPolicy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketPolicy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewBucketPolicy(scope constructs.Construct, id *string, props *BucketPolicyProps) BucketPolicy {
	_init_.Initialize()

	if err := validateNewBucketPolicyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BucketPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.BucketPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewBucketPolicy_Override(b BucketPolicy, scope constructs.Construct, id *string, props *BucketPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.BucketPolicy",
		[]interface{}{scope, id, props},
		b,
	)
}

// Create a mutable `BucketPolicy` from a `CfnBucketPolicy`.
func BucketPolicy_FromCfnBucketPolicy(cfnBucketPolicy CfnBucketPolicy) BucketPolicy {
	_init_.Initialize()

	if err := validateBucketPolicy_FromCfnBucketPolicyParameters(cfnBucketPolicy); err != nil {
		panic(err)
	}
	var returns BucketPolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.BucketPolicy",
		"fromCfnBucketPolicy",
		[]interface{}{cfnBucketPolicy},
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
func BucketPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBucketPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.BucketPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func BucketPolicy_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateBucketPolicy_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.BucketPolicy",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func BucketPolicy_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateBucketPolicy_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.BucketPolicy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func BucketPolicy_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.BucketPolicy",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BucketPolicy) ApplyRemovalPolicy(removalPolicy awscdk.RemovalPolicy) {
	if err := b.validateApplyRemovalPolicyParameters(removalPolicy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{removalPolicy},
	)
}

func (b *jsiiProxy_BucketPolicy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketPolicy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := b.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketPolicy) GetResourceNameAttribute(nameAttr *string) *string {
	if err := b.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

