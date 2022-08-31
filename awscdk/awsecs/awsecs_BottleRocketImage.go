package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs/internal"
)

// Construct an Bottlerocket image from the latest AMI published in SSM.
//
// Example:
//   var cluster cluster
//
//
//   cluster.addCapacity(jsii.String("bottlerocket-asg"), &addCapacityOptions{
//   	minCapacity: jsii.Number(2),
//   	instanceType: ec2.NewInstanceType(jsii.String("c5.large")),
//   	machineImage: ecs.NewBottleRocketImage(),
//   })
//
// Experimental.
type BottleRocketImage interface {
	awsec2.IMachineImage
	// Return the correct image.
	// Experimental.
	GetImage(scope awscdk.Construct) *awsec2.MachineImageConfig
}

// The jsii proxy struct for BottleRocketImage
type jsiiProxy_BottleRocketImage struct {
	internal.Type__awsec2IMachineImage
}

// Constructs a new instance of the BottleRocketImage class.
// Experimental.
func NewBottleRocketImage(props *BottleRocketImageProps) BottleRocketImage {
	_init_.Initialize()

	j := jsiiProxy_BottleRocketImage{}

	_jsii_.Create(
		"monocdk.aws_ecs.BottleRocketImage",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the BottleRocketImage class.
// Experimental.
func NewBottleRocketImage_Override(b BottleRocketImage, props *BottleRocketImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.BottleRocketImage",
		[]interface{}{props},
		b,
	)
}

func (b *jsiiProxy_BottleRocketImage) GetImage(scope awscdk.Construct) *awsec2.MachineImageConfig {
	var returns *awsec2.MachineImageConfig

	_jsii_.Invoke(
		b,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

