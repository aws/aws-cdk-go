package awss3deployment

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

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
//   deployment := s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &bucketDeploymentProps{
//   	sources: []iSource{
//   		s3deploy.source.asset(path.join(__dirname, jsii.String("my-website"))),
//   	},
//   	destinationBucket: websiteBucket,
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
	// Doing this replaces calling `otherResource.node.addDependency(deployment)`.
	DeployedBucket() awss3.IBucket
	// The tree node.
	Node() constructs.Node
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

func (j *jsiiProxy_BucketDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

