package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Selects the latest version of Amazon Linux.
//
// This Machine Image automatically updates to the latest version on every
// deployment. Be aware this will cause your instances to be replaced when a
// new version of the image becomes available. Do not store stateful information
// on the instance if you are using this image.
//
// The AMI ID is selected using the values published to the SSM parameter store.
//
// Example:
//   // Pick a Windows edition to use
//   windows := ec2.NewWindowsImage(ec2.windowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_BASE)
//
//   // Pick the right Amazon Linux edition. All arguments shown are optional
//   // and will default to these values when omitted.
//   amznLinux := ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
//   	generation: ec2.amazonLinuxGeneration_AMAZON_LINUX,
//   	edition: ec2.amazonLinuxEdition_STANDARD,
//   	virtualization: ec2.amazonLinuxVirt_HVM,
//   	storage: ec2.amazonLinuxStorage_GENERAL_PURPOSE,
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
type AmazonLinuxImage interface {
	GenericSSMParameterImage
	// Name of the SSM parameter we're looking up.
	ParameterName() *string
	// Return the image to use in the given context.
	GetImage(scope constructs.Construct) *MachineImageConfig
}

// The jsii proxy struct for AmazonLinuxImage
type jsiiProxy_AmazonLinuxImage struct {
	jsiiProxy_GenericSSMParameterImage
}

func (j *jsiiProxy_AmazonLinuxImage) ParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterName",
		&returns,
	)
	return returns
}


func NewAmazonLinuxImage(props *AmazonLinuxImageProps) AmazonLinuxImage {
	_init_.Initialize()

	if err := validateNewAmazonLinuxImageParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmazonLinuxImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinuxImage",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAmazonLinuxImage_Override(a AmazonLinuxImage, props *AmazonLinuxImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinuxImage",
		[]interface{}{props},
		a,
	)
}

// Return the SSM parameter name that will contain the Amazon Linux image with the given attributes.
func AmazonLinuxImage_SsmParameterName(props *AmazonLinuxImageProps) *string {
	_init_.Initialize()

	if err := validateAmazonLinuxImage_SsmParameterNameParameters(props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.AmazonLinuxImage",
		"ssmParameterName",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmazonLinuxImage) GetImage(scope constructs.Construct) *MachineImageConfig {
	if err := a.validateGetImageParameters(scope); err != nil {
		panic(err)
	}
	var returns *MachineImageConfig

	_jsii_.Invoke(
		a,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

