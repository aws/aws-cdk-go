package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
	"github.com/aws/constructs-go/constructs/v3"
)

// Represents a Glue Job's Code assets (an asset can be a scripts, a jar, a python file or any other file).
//
// Example:
//   var bucket bucket
//
//   glue.NewJob(this, jsii.String("PythonShellJob"), &jobProps{
//   	executable: glue.jobExecutable.pythonShell(&pythonShellExecutableProps{
//   		glueVersion: glue.glueVersion_V1_0(),
//   		pythonVersion: glue.pythonVersion_THREE,
//   		script: glue.code.fromBucket(bucket, jsii.String("script.py")),
//   	}),
//   	description: jsii.String("an example Python Shell job"),
//   })
//
// Experimental.
type Code interface {
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig
}

// The jsii proxy struct for Code
type jsiiProxy_Code struct {
	_ byte // padding
}

// Experimental.
func NewCode_Override(c Code) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_glue.Code",
		nil, // no parameters
		c,
	)
}

// Job code from a local disk path.
// Experimental.
func Code_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"monocdk.aws_glue.Code",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Job code as an S3 object.
// Experimental.
func Code_FromBucket(bucket awss3.IBucket, key *string) S3Code {
	_init_.Initialize()

	if err := validateCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"monocdk.aws_glue.Code",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Code) Bind(scope constructs.Construct, grantable awsiam.IGrantable) *CodeConfig {
	if err := c.validateBindParameters(scope, grantable); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, grantable},
		&returns,
	)

	return returns
}

