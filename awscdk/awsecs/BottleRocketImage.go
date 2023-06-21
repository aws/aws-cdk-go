package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Construct an Bottlerocket image from the latest AMI published in SSM.
//
// Example:
//   var cluster cluster
//
//
//   cluster.AddCapacity(jsii.String("bottlerocket-asg"), &AddCapacityOptions{
//   	MinCapacity: jsii.Number(2),
//   	InstanceType: ec2.NewInstanceType(jsii.String("c5.large")),
//   	MachineImage: ecs.NewBottleRocketImage(),
//   })
//
type BottleRocketImage interface {
	awsec2.IMachineImage
	// Return the correct image.
	GetImage(scope constructs.Construct) *awsec2.MachineImageConfig
}

// The jsii proxy struct for BottleRocketImage
type jsiiProxy_BottleRocketImage struct {
	internal.Type__awsec2IMachineImage
}

// Constructs a new instance of the BottleRocketImage class.
func NewBottleRocketImage(props *BottleRocketImageProps) BottleRocketImage {
	_init_.Initialize()

	if err := validateNewBottleRocketImageParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BottleRocketImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.BottleRocketImage",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the BottleRocketImage class.
func NewBottleRocketImage_Override(b BottleRocketImage, props *BottleRocketImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.BottleRocketImage",
		[]interface{}{props},
		b,
	)
}

// Return whether the given object is a BottleRocketImage.
func BottleRocketImage_IsBottleRocketImage(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBottleRocketImage_IsBottleRocketImageParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.BottleRocketImage",
		"isBottleRocketImage",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BottleRocketImage) GetImage(scope constructs.Construct) *awsec2.MachineImageConfig {
	if err := b.validateGetImageParameters(scope); err != nil {
		panic(err)
	}
	var returns *awsec2.MachineImageConfig

	_jsii_.Invoke(
		b,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

