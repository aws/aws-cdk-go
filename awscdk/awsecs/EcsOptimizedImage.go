package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Construct a Linux or Windows machine image from the latest ECS Optimized AMI published in SSM.
//
// Example:
//   var vpc vpc
//
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   // Either add default capacity
//   cluster.AddCapacity(jsii.String("DefaultAutoScalingGroupCapacity"), &AddCapacityOptions{
//   	InstanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
//   	DesiredCapacity: jsii.Number(3),
//   })
//
//   // Or add customized capacity. Be sure to start the Amazon ECS-optimized AMI.
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
//   	MachineImage: ecs.EcsOptimizedImage_AmazonLinux(),
//   	// Or use Amazon ECS-Optimized Amazon Linux 2 AMI
//   	// machineImage: EcsOptimizedImage.amazonLinux2(),
//   	DesiredCapacity: jsii.Number(3),
//   })
//
//   capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &AsgCapacityProviderProps{
//   	AutoScalingGroup: AutoScalingGroup,
//   })
//   cluster.AddAsgCapacityProvider(capacityProvider)
//
type EcsOptimizedImage interface {
	awsec2.IMachineImage
	// Return the correct image.
	GetImage(scope constructs.Construct) *awsec2.MachineImageConfig
}

// The jsii proxy struct for EcsOptimizedImage
type jsiiProxy_EcsOptimizedImage struct {
	internal.Type__awsec2IMachineImage
}

// Construct an Amazon Linux AMI image from the latest ECS Optimized AMI published in SSM.
func EcsOptimizedImage_AmazonLinux(options *EcsOptimizedImageOptions) EcsOptimizedImage {
	_init_.Initialize()

	if err := validateEcsOptimizedImage_AmazonLinuxParameters(options); err != nil {
		panic(err)
	}
	var returns EcsOptimizedImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcsOptimizedImage",
		"amazonLinux",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct an Amazon Linux 2 image from the latest ECS Optimized AMI published in SSM.
func EcsOptimizedImage_AmazonLinux2(hardwareType AmiHardwareType, options *EcsOptimizedImageOptions) EcsOptimizedImage {
	_init_.Initialize()

	if err := validateEcsOptimizedImage_AmazonLinux2Parameters(options); err != nil {
		panic(err)
	}
	var returns EcsOptimizedImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcsOptimizedImage",
		"amazonLinux2",
		[]interface{}{hardwareType, options},
		&returns,
	)

	return returns
}

// Construct an Amazon Linux 2023 image from the latest ECS Optimized AMI published in SSM.
func EcsOptimizedImage_AmazonLinux2023(hardwareType AmiHardwareType, options *EcsOptimizedImageOptions) EcsOptimizedImage {
	_init_.Initialize()

	if err := validateEcsOptimizedImage_AmazonLinux2023Parameters(options); err != nil {
		panic(err)
	}
	var returns EcsOptimizedImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcsOptimizedImage",
		"amazonLinux2023",
		[]interface{}{hardwareType, options},
		&returns,
	)

	return returns
}

// Construct a Windows image from the latest ECS Optimized AMI published in SSM.
func EcsOptimizedImage_Windows(windowsVersion WindowsOptimizedVersion, options *EcsOptimizedImageOptions) EcsOptimizedImage {
	_init_.Initialize()

	if err := validateEcsOptimizedImage_WindowsParameters(windowsVersion, options); err != nil {
		panic(err)
	}
	var returns EcsOptimizedImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcsOptimizedImage",
		"windows",
		[]interface{}{windowsVersion, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsOptimizedImage) GetImage(scope constructs.Construct) *awsec2.MachineImageConfig {
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

