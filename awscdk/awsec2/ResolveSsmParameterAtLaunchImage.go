package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Select the image based on a given SSM parameter at instance launch time.
//
// This Machine Image comes with an imageId as `resolve:ssm:parameter-name` or `resolve:ssm:parameter-name:version` format
// as described in the document:.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userData userData
//
//   resolveSsmParameterAtLaunchImage := awscdk.Aws_ec2.NewResolveSsmParameterAtLaunchImage(jsii.String("parameterName"), &SsmParameterImageOptions{
//   	CachedInContext: jsii.Boolean(false),
//   	Os: awscdk.*Aws_ec2.OperatingSystemType_LINUX,
//   	ParameterVersion: jsii.String("parameterVersion"),
//   	UserData: userData,
//   })
//
// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/using-systems-manager-parameters.html
//
// The AMI ID would be selected at instance launch time.
//
type ResolveSsmParameterAtLaunchImage interface {
	IMachineImage
	// Name of the SSM parameter we're looking up.
	ParameterName() *string
	// Return the image to use in the given context.
	GetImage(_arg constructs.Construct) *MachineImageConfig
}

// The jsii proxy struct for ResolveSsmParameterAtLaunchImage
type jsiiProxy_ResolveSsmParameterAtLaunchImage struct {
	jsiiProxy_IMachineImage
}

func (j *jsiiProxy_ResolveSsmParameterAtLaunchImage) ParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterName",
		&returns,
	)
	return returns
}


func NewResolveSsmParameterAtLaunchImage(parameterName *string, props *SsmParameterImageOptions) ResolveSsmParameterAtLaunchImage {
	_init_.Initialize()

	if err := validateNewResolveSsmParameterAtLaunchImageParameters(parameterName, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResolveSsmParameterAtLaunchImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.ResolveSsmParameterAtLaunchImage",
		[]interface{}{parameterName, props},
		&j,
	)

	return &j
}

func NewResolveSsmParameterAtLaunchImage_Override(r ResolveSsmParameterAtLaunchImage, parameterName *string, props *SsmParameterImageOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.ResolveSsmParameterAtLaunchImage",
		[]interface{}{parameterName, props},
		r,
	)
}

func (r *jsiiProxy_ResolveSsmParameterAtLaunchImage) GetImage(_arg constructs.Construct) *MachineImageConfig {
	if err := r.validateGetImageParameters(_arg); err != nil {
		panic(err)
	}
	var returns *MachineImageConfig

	_jsii_.Invoke(
		r,
		"getImage",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

