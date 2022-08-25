package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs/internal"
)

// Construct a Linux or Windows machine image from the latest ECS Optimized AMI published in SSM.
//
// Example:
//   var vpc vpc
//
//   myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &computeEnvironmentProps{
//   	computeResources: &computeResources{
//   		image: ecs.NewEcsOptimizedAmi(&ecsOptimizedAmiProps{
//   			generation: ec2.amazonLinuxGeneration_AMAZON_LINUX_2,
//   		}),
//   		vpc: vpc,
//   	},
//   })
//
// Deprecated: see {@link EcsOptimizedImage#amazonLinux}, {@link EcsOptimizedImage#amazonLinux} and {@link EcsOptimizedImage#windows}.
type EcsOptimizedAmi interface {
	awsec2.IMachineImage
	// Return the correct image.
	// Deprecated: see {@link EcsOptimizedImage#amazonLinux}, {@link EcsOptimizedImage#amazonLinux} and {@link EcsOptimizedImage#windows}.
	GetImage(scope awscdk.Construct) *awsec2.MachineImageConfig
}

// The jsii proxy struct for EcsOptimizedAmi
type jsiiProxy_EcsOptimizedAmi struct {
	internal.Type__awsec2IMachineImage
}

// Constructs a new instance of the EcsOptimizedAmi class.
// Deprecated: see {@link EcsOptimizedImage#amazonLinux}, {@link EcsOptimizedImage#amazonLinux} and {@link EcsOptimizedImage#windows}.
func NewEcsOptimizedAmi(props *EcsOptimizedAmiProps) EcsOptimizedAmi {
	_init_.Initialize()

	j := jsiiProxy_EcsOptimizedAmi{}

	_jsii_.Create(
		"monocdk.aws_ecs.EcsOptimizedAmi",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the EcsOptimizedAmi class.
// Deprecated: see {@link EcsOptimizedImage#amazonLinux}, {@link EcsOptimizedImage#amazonLinux} and {@link EcsOptimizedImage#windows}.
func NewEcsOptimizedAmi_Override(e EcsOptimizedAmi, props *EcsOptimizedAmiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.EcsOptimizedAmi",
		[]interface{}{props},
		e,
	)
}

func (e *jsiiProxy_EcsOptimizedAmi) GetImage(scope awscdk.Construct) *awsec2.MachineImageConfig {
	var returns *awsec2.MachineImageConfig

	_jsii_.Invoke(
		e,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

