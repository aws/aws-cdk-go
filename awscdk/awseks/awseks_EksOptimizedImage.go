package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awseks/internal"
)

// Construct an Amazon Linux 2 image from the latest EKS Optimized AMI published in SSM.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksOptimizedImage := awscdk.Aws_eks.NewEksOptimizedImage(&eksOptimizedImageProps{
//   	cpuArch: awscdk.*Aws_eks.cpuArch_ARM_64,
//   	kubernetesVersion: jsii.String("kubernetesVersion"),
//   	nodeType: awscdk.*Aws_eks.nodeType_STANDARD,
//   })
//
// Experimental.
type EksOptimizedImage interface {
	awsec2.IMachineImage
	// Return the correct image.
	// Experimental.
	GetImage(scope awscdk.Construct) *awsec2.MachineImageConfig
}

// The jsii proxy struct for EksOptimizedImage
type jsiiProxy_EksOptimizedImage struct {
	internal.Type__awsec2IMachineImage
}

// Constructs a new instance of the EcsOptimizedAmi class.
// Experimental.
func NewEksOptimizedImage(props *EksOptimizedImageProps) EksOptimizedImage {
	_init_.Initialize()

	if err := validateNewEksOptimizedImageParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EksOptimizedImage{}

	_jsii_.Create(
		"monocdk.aws_eks.EksOptimizedImage",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the EcsOptimizedAmi class.
// Experimental.
func NewEksOptimizedImage_Override(e EksOptimizedImage, props *EksOptimizedImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.EksOptimizedImage",
		[]interface{}{props},
		e,
	)
}

func (e *jsiiProxy_EksOptimizedImage) GetImage(scope awscdk.Construct) *awsec2.MachineImageConfig {
	if err := e.validateGetImageParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsec2.MachineImageConfig

	_jsii_.Invoke(
		e,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

