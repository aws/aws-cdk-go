package awss3deployment

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// `DeployTimeSubstitutedFile` is an extension of `BucketDeployment` that allows users to upload individual files and specify to make substitutions in the file.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var myLambdaFunction function
//   var destinationBucket bucket
//   //(Optional) if provided, the resulting processed file would be uploaded to the destinationBucket under the destinationKey name.
//   var destinationKey string
//   var role role
//
//
//   s3deploy.NewDeployTimeSubstitutedFile(this, jsii.String("MyFile"), &DeployTimeSubstitutedFileProps{
//   	Source: jsii.String("my-file.yaml"),
//   	DestinationKey: destinationKey,
//   	DestinationBucket: destinationBucket,
//   	Substitutions: map[string]*string{
//   		"variableName": myLambdaFunction.functionName,
//   	},
//   	Role: role,
//   })
//
type DeployTimeSubstitutedFile interface {
	BucketDeployment
	Bucket() awss3.IBucket
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
	ObjectKey() *string
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
	//   declare const websiteBucket: s3.IBucket;
	//   const deployment = new s3deploy.BucketDeployment(this, 'Deployment', {
	//     sources: [s3deploy.Source.asset('./website-dist')],
	//     destinationBucket: websiteBucket,
	//   });
	//
	//   deployment.addSource(s3deploy.Source.asset('./another-asset'));
	//
	AddSource(source ISource)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for DeployTimeSubstitutedFile
type jsiiProxy_DeployTimeSubstitutedFile struct {
	jsiiProxy_BucketDeployment
}

func (j *jsiiProxy_DeployTimeSubstitutedFile) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployTimeSubstitutedFile) DeployedBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"deployedBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployTimeSubstitutedFile) HandlerRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"handlerRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployTimeSubstitutedFile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployTimeSubstitutedFile) ObjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployTimeSubstitutedFile) ObjectKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objectKeys",
		&returns,
	)
	return returns
}


func NewDeployTimeSubstitutedFile(scope constructs.Construct, id *string, props *DeployTimeSubstitutedFileProps) DeployTimeSubstitutedFile {
	_init_.Initialize()

	if err := validateNewDeployTimeSubstitutedFileParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeployTimeSubstitutedFile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_deployment.DeployTimeSubstitutedFile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDeployTimeSubstitutedFile_Override(d DeployTimeSubstitutedFile, scope constructs.Construct, id *string, props *DeployTimeSubstitutedFileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_deployment.DeployTimeSubstitutedFile",
		[]interface{}{scope, id, props},
		d,
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
func DeployTimeSubstitutedFile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDeployTimeSubstitutedFile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.DeployTimeSubstitutedFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeployTimeSubstitutedFile) AddSource(source ISource) {
	if err := d.validateAddSourceParameters(source); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addSource",
		[]interface{}{source},
	)
}

func (d *jsiiProxy_DeployTimeSubstitutedFile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

