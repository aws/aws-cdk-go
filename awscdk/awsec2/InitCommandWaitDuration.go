package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Represents a duration to wait after a command has finished, in case of a reboot (Windows only).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   initCommandWaitDuration := awscdk.Aws_ec2.InitCommandWaitDuration_Forever()
//
type InitCommandWaitDuration interface {
}

// The jsii proxy struct for InitCommandWaitDuration
type jsiiProxy_InitCommandWaitDuration struct {
	_ byte // padding
}

func NewInitCommandWaitDuration_Override(i InitCommandWaitDuration) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitCommandWaitDuration",
		nil, // no parameters
		i,
	)
}

// cfn-init will exit and resume only after a reboot.
func InitCommandWaitDuration_Forever() InitCommandWaitDuration {
	_init_.Initialize()

	var returns InitCommandWaitDuration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitCommandWaitDuration",
		"forever",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Do not wait for this command.
func InitCommandWaitDuration_None() InitCommandWaitDuration {
	_init_.Initialize()

	var returns InitCommandWaitDuration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitCommandWaitDuration",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Wait for a specified duration after a command.
func InitCommandWaitDuration_Of(duration awscdk.Duration) InitCommandWaitDuration {
	_init_.Initialize()

	if err := validateInitCommandWaitDuration_OfParameters(duration); err != nil {
		panic(err)
	}
	var returns InitCommandWaitDuration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitCommandWaitDuration",
		"of",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

