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
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	vpc: vpc,
//   })
//
//   // Either add default capacity
//   cluster.addCapacity(jsii.String("DefaultAutoScalingGroupCapacity"), &addCapacityOptions{
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
//   	desiredCapacity: jsii.Number(3),
//   })
//
//   // Or add customized capacity. Be sure to start the Amazon ECS-optimized AMI.
//   autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
//   	machineImage: ecs.ecsOptimizedImage.amazonLinux(),
//   	// Or use Amazon ECS-Optimized Amazon Linux 2 AMI
//   	// machineImage: EcsOptimizedImage.amazonLinux2(),
//   	desiredCapacity: jsii.Number(3),
//   })
//
//   capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &asgCapacityProviderProps{
//   	autoScalingGroup: autoScalingGroup,
//   })
//   cluster.addAsgCapacityProvider(capacityProvider)
//
// Experimental.
type EcsOptimizedImage interface {
	awsec2.IMachineImage
	// Return the correct image.
	// Experimental.
	GetImage(scope awscdk.Construct) *awsec2.MachineImageConfig
}

// The jsii proxy struct for EcsOptimizedImage
type jsiiProxy_EcsOptimizedImage struct {
	internal.Type__awsec2IMachineImage
}

// Construct an Amazon Linux AMI image from the latest ECS Optimized AMI published in SSM.
// Experimental.
func EcsOptimizedImage_AmazonLinux(options *EcsOptimizedImageOptions) EcsOptimizedImage {
	_init_.Initialize()

	if err := validateEcsOptimizedImage_AmazonLinuxParameters(options); err != nil {
		panic(err)
	}
	var returns EcsOptimizedImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.EcsOptimizedImage",
		"amazonLinux",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct an Amazon Linux 2 image from the latest ECS Optimized AMI published in SSM.
// Experimental.
func EcsOptimizedImage_AmazonLinux2(hardwareType AmiHardwareType, options *EcsOptimizedImageOptions) EcsOptimizedImage {
	_init_.Initialize()

	if err := validateEcsOptimizedImage_AmazonLinux2Parameters(options); err != nil {
		panic(err)
	}
	var returns EcsOptimizedImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.EcsOptimizedImage",
		"amazonLinux2",
		[]interface{}{hardwareType, options},
		&returns,
	)

	return returns
}

// Construct a Windows image from the latest ECS Optimized AMI published in SSM.
// Experimental.
func EcsOptimizedImage_Windows(windowsVersion WindowsOptimizedVersion, options *EcsOptimizedImageOptions) EcsOptimizedImage {
	_init_.Initialize()

	if err := validateEcsOptimizedImage_WindowsParameters(windowsVersion, options); err != nil {
		panic(err)
	}
	var returns EcsOptimizedImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.EcsOptimizedImage",
		"windows",
		[]interface{}{windowsVersion, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsOptimizedImage) GetImage(scope awscdk.Construct) *awsec2.MachineImageConfig {
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

