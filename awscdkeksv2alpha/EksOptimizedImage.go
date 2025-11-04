package awscdkeksv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkeksv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdkeksv2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Construct an Amazon Linux 2 image from the latest EKS Optimized AMI published in SSM.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkeksv2alpha"
//
//   eksOptimizedImage := eks_v2_alpha.NewEksOptimizedImage(&EksOptimizedImageProps{
//   	CpuArch: eks_v2_alpha.CpuArch_ARM_64,
//   	KubernetesVersion: jsii.String("kubernetesVersion"),
//   	NodeType: eks_v2_alpha.NodeType_STANDARD,
//   })
//
// Experimental.
type EksOptimizedImage interface {
	awsec2.IMachineImage
	// Return the correct image.
	// Experimental.
	GetImage(scope constructs.Construct) *awsec2.MachineImageConfig
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
		"@aws-cdk/aws-eks-v2-alpha.EksOptimizedImage",
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
		"@aws-cdk/aws-eks-v2-alpha.EksOptimizedImage",
		[]interface{}{props},
		e,
	)
}

func (e *jsiiProxy_EksOptimizedImage) GetImage(scope constructs.Construct) *awsec2.MachineImageConfig {
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

