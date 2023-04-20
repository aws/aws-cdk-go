package awss3deployment

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3deployment/internal"
	"github.com/aws/constructs-go/constructs/v3"
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
// Experimental.
type BucketDeployment interface {
	awscdk.Construct
	// The bucket after the deployment.
	//
	// If you want to reference the destination bucket in another construct and make sure the
	// bucket deployment has happened before the next operation is started, pass the other construct
	// a reference to `deployment.deployedBucket`.
	//
	// Doing this replaces calling `otherResource.node.addDependency(deployment)`.
	// Experimental.
	DeployedBucket() awss3.IBucket
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for BucketDeployment
type jsiiProxy_BucketDeployment struct {
	internal.Type__awscdkConstruct
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

func (j *jsiiProxy_BucketDeployment) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewBucketDeployment(scope constructs.Construct, id *string, props *BucketDeploymentProps) BucketDeployment {
	_init_.Initialize()

	if err := validateNewBucketDeploymentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BucketDeployment{}

	_jsii_.Create(
		"monocdk.aws_s3_deployment.BucketDeployment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBucketDeployment_Override(b BucketDeployment, scope constructs.Construct, id *string, props *BucketDeploymentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_s3_deployment.BucketDeployment",
		[]interface{}{scope, id, props},
		b,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func BucketDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBucketDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.BucketDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketDeployment) OnPrepare() {
	_jsii_.InvokeVoid(
		b,
		"onPrepare",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BucketDeployment) OnSynthesize(session constructs.ISynthesisSession) {
	if err := b.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (b *jsiiProxy_BucketDeployment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketDeployment) Prepare() {
	_jsii_.InvokeVoid(
		b,
		"prepare",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BucketDeployment) Synthesize(session awscdk.ISynthesisSession) {
	if err := b.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		[]interface{}{session},
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

func (b *jsiiProxy_BucketDeployment) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

