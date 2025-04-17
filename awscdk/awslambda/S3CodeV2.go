package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Lambda code from an S3 archive.
//
// With option to set KMSKey for encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var key key
//
//   s3CodeV2 := awscdk.Aws_lambda.NewS3CodeV2(bucket, jsii.String("key"), &BucketOptions{
//   	ObjectVersion: jsii.String("objectVersion"),
//   	SourceKMSKey: key,
//   })
//
type S3CodeV2 interface {
	Code
	// Determines whether this Code is inline code or not.
	IsInline() *bool
	// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(_scope constructs.Construct) *CodeConfig
	// Called after the CFN function resource has been created to allow the code class to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions)
}

// The jsii proxy struct for S3CodeV2
type jsiiProxy_S3CodeV2 struct {
	jsiiProxy_Code
}

func (j *jsiiProxy_S3CodeV2) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}


func NewS3CodeV2(bucket awss3.IBucket, key *string, options *BucketOptions) S3CodeV2 {
	_init_.Initialize()

	if err := validateNewS3CodeV2Parameters(bucket, key, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3CodeV2{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.S3CodeV2",
		[]interface{}{bucket, key, options},
		&j,
	)

	return &j
}

func NewS3CodeV2_Override(s S3CodeV2, bucket awss3.IBucket, key *string, options *BucketOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.S3CodeV2",
		[]interface{}{bucket, key, options},
		s,
	)
}

// Loads the function code from a local disk path.
func S3CodeV2_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateS3CodeV2_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3CodeV2",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func S3CodeV2_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	if err := validateS3CodeV2_FromAssetImageParameters(directory, props); err != nil {
		panic(err)
	}
	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3CodeV2",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func S3CodeV2_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	if err := validateS3CodeV2_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3CodeV2",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func S3CodeV2_FromBucketV2(bucket awss3.IBucket, key *string, options *BucketOptions) S3CodeV2 {
	_init_.Initialize()

	if err := validateS3CodeV2_FromBucketV2Parameters(bucket, key, options); err != nil {
		panic(err)
	}
	var returns S3CodeV2

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3CodeV2",
		"fromBucketV2",
		[]interface{}{bucket, key, options},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`.
func S3CodeV2_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	if err := validateS3CodeV2_FromCfnParametersParameters(props); err != nil {
		panic(err)
	}
	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3CodeV2",
		"fromCfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Runs a command to build the code asset that will be used.
func S3CodeV2_FromCustomCommand(output *string, command *[]*string, options *CustomCommandOptions) AssetCode {
	_init_.Initialize()

	if err := validateS3CodeV2_FromCustomCommandParameters(output, command, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3CodeV2",
		"fromCustomCommand",
		[]interface{}{output, command, options},
		&returns,
	)

	return returns
}

// Loads the function code from an asset created by a Docker build.
//
// By default, the asset is expected to be located at `/asset` in the
// image.
func S3CodeV2_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateS3CodeV2_FromDockerBuildParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3CodeV2",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func S3CodeV2_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	if err := validateS3CodeV2_FromEcrImageParameters(repository, props); err != nil {
		panic(err)
	}
	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3CodeV2",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func S3CodeV2_FromInline(code *string) InlineCode {
	_init_.Initialize()

	if err := validateS3CodeV2_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3CodeV2",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3CodeV2) Bind(_scope constructs.Construct) *CodeConfig {
	if err := s.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3CodeV2) BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions) {
	if err := s.validateBindToResourceParameters(_resource, _options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bindToResource",
		[]interface{}{_resource, _options},
	)
}

