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
//   var vpc iVpc
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &LoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   })
//
//   // instance to add as the target for load balancer.
//   instance := ec2.NewInstance(this, jsii.String("targetInstance"), &InstanceProps{
//   	Vpc: vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),
//   	MachineImage: ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   })
//   lb.AddTarget(elb.NewInstanceTarget(instance))
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

