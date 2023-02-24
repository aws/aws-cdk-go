package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Select the latest version of the indicated Windows version.
//
// This Machine Image automatically updates to the latest version on every
// deployment. Be aware this will cause your instances to be replaced when a
// new version of the image becomes available. Do not store stateful information
// on the instance if you are using this image.
//
// The AMI ID is selected using the values published to the SSM parameter store.
//
// https://aws.amazon.com/blogs/mt/query-for-the-latest-windows-ami-using-systems-manager-parameter-store/
//
// Example:
//   // Pick a Windows edition to use
//   windows := ec2.NewWindowsImage(ec2.WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_BASE)
//
//   // Pick the right Amazon Linux edition. All arguments shown are optional
//   // and will default to these values when omitted.
//   amznLinux := ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
//   	Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX,
//   	Edition: ec2.AmazonLinuxEdition_STANDARD,
//   	Virtualization: ec2.AmazonLinuxVirt_HVM,
//   	Storage: ec2.AmazonLinuxStorage_GENERAL_PURPOSE,
//   })
//
//   // For other custom (Linux) images, instantiate a `GenericLinuxImage` with
//   // a map giving the AMI to in for each region:
//
//   linux := ec2.NewGenericLinuxImage(map[string]*string{
//   	"us-east-1": jsii.String("ami-97785bed"),
//   	"eu-west-1": jsii.String("ami-12345678"),
//   })
//
type WindowsImage interface {
	GenericSSMParameterImage
	// Name of the SSM parameter we're looking up.
	ParameterName() *string
	// Return the image to use in the given context.
	GetImage(scope constructs.Construct) *MachineImageConfig
}

// The jsii proxy struct for WindowsImage
type jsiiProxy_WindowsImage struct {
	jsiiProxy_GenericSSMParameterImage
}

func (j *jsiiProxy_WindowsImage) ParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterName",
		&returns,
	)
	return returns
}


func NewWindowsImage(version WindowsVersion, props *WindowsImageProps) WindowsImage {
	_init_.Initialize()

	if err := validateNewWindowsImageParameters(version, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.WindowsImage",
		[]interface{}{version, props},
		&j,
	)

	return &j
}

func NewWindowsImage_Override(w WindowsImage, version WindowsVersion, props *WindowsImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.WindowsImage",
		[]interface{}{version, props},
		w,
	)
}

func (w *jsiiProxy_WindowsImage) GetImage(scope constructs.Construct) *MachineImageConfig {
	if err := w.validateGetImageParameters(scope); err != nil {
		panic(err)
	}
	var returns *MachineImageConfig

	_jsii_.Invoke(
		w,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

