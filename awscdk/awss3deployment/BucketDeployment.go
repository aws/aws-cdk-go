package awss3deployment

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// `BucketDeployment` populates an S3 bucket with the contents of .zip files from other S3 buckets or from local disk.
//
// Example:
//   var websiteBucket bucket
//
//
//   deployment := s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &BucketDeploymentProps{
//   	Sources: []iSource{
//   		s3deploy.Source_Asset(path.join(__dirname, jsii.String("my-website"))),
//   	},
//   	DestinationBucket: websiteBucket,
//   })
//
//   NewConstructThatReadsFromTheBucket(this, jsii.String("Consumer"), map[string]iBucket{
//   	// Use 'deployment.deployedBucket' instead of 'websiteBucket' here
//   	"bucket": deployment.deployedBucket,
//   })
//
type BucketDeployment interface {
	constructs.Construct
	// The bucket after the deployment.
	//
	// If you want to reference the destination bucket in another construct and make sure the
	// bucket deployment has happened before the next operation is started, pass the other construct
	// a reference to `deployment.deployedBucket`.
	//
	// Note that this only returns an immutable reference to the destination bucket.
	// If sequenced access to the original destination bucket is required, you may add a dependency
	// on the bucket deployment instead: `otherResource.node.addDependency(deployment)`
	DeployedBucket() awss3.IBucket
	// Execution role of the Lambda function behind the custom CloudFormation resource of type `Custom::CDKBucketDeployment`.
	HandlerRole() awsiam.IRole
	// The tree node.
	Node() constructs.Node
	// The object keys for the sources deployed to the S3 bucket.
	//
	// This returns a list of tokenized object keys for source files that are deployed to the bucket.
	//
	// This can be useful when using `BucketDeployment` with `extract` set to `false` and you need to reference
	// the object key that resides in the bucket for that zip source file somewhere else in your CDK
	// application, such as in a CFN output.
	//
	// For example, use `Fn.select(0, myBucketDeployment.objectKeys)` to reference the object key of the
	// first source file in your bucket deployment.
	ObjectKeys() *[]*string
	// Add an additional source to the bucket deployment.
	//
	// Example:
	//   var websiteBucket iBucket
	//
	//   deployment := s3deploy.NewBucketDeployment(this, jsii.String("Deployment"), &BucketDeploymentProps{
	//   	Sources: []iSource{
	//   		s3deploy.Source_Asset(jsii.String("./website-dist")),
	//   	},
	//   	DestinationBucket: websiteBucket,
	//   })
	//
	//   deployment.AddSource(s3deploy.Source_Asset(jsii.String("./another-asset")))
	//
	AddSource(source ISource)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for BucketDeployment
type jsiiProxy_BucketDeployment struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_BucketDeployment) DeployedBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"deployedBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketDeployment) HandlerRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"handlerRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketDeployment) ObjectKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objectKeys",
		&returns,
	)
	return returns
}


func NewBucketDeployment(scope constructs.Construct, id *string, props *BucketDeploymentProps) BucketDeployment {
	_init_.Initialize()

	if err := validateNewBucketDeploymentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BucketDeployment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_deployment.BucketDeployment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewBucketDeployment_Override(b BucketDeployment, scope constructs.Construct, id *string, props *BucketDeploymentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_deployment.BucketDeployment",
		[]interface{}{scope, id, props},
		b,
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
func BucketDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBucketDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.BucketDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketDeployment) AddSource(source ISource) {
	if err := b.validateAddSourceParameters(source); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addSource",
		[]interface{}{source},
	)
}

func (b *jsiiProxy_BucketDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

