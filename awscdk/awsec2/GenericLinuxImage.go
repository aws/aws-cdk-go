package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Construct a Linux machine image from an AMI map.
//
// Linux images IDs are not published to SSM parameter store yet, so you'll have to
// manually specify an AMI map.
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
type GenericLinuxImage interface {
	IMachineImage
	// Return the image to use in the given context.
	GetImage(scope constructs.Construct) *MachineImageConfig
}

// The jsii proxy struct for GenericLinuxImage
type jsiiProxy_GenericLinuxImage struct {
	jsiiProxy_IMachineImage
}

func NewGenericLinuxImage(amiMap *map[string]*string, props *GenericLinuxImageProps) GenericLinuxImage {
	_init_.Initialize()

	if err := validateNewGenericLinuxImageParameters(amiMap, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenericLinuxImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.GenericLinuxImage",
		[]interface{}{amiMap, props},
		&j,
	)

	return &j
}

func NewGenericLinuxImage_Override(g GenericLinuxImage, amiMap *map[string]*string, props *GenericLinuxImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.GenericLinuxImage",
		[]interface{}{amiMap, props},
		g,
	)
}

func (g *jsiiProxy_GenericLinuxImage) GetImage(scope constructs.Construct) *MachineImageConfig {
	if err := g.validateGetImageParameters(scope); err != nil {
		panic(err)
	}
	var returns *MachineImageConfig

	_jsii_.Invoke(
		g,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

