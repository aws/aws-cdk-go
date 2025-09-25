package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configure whether the AutoScalingGroup waits for signals.
//
// If you do configure waiting for signals, you should make sure the instances
// invoke `cfn-signal` somewhere in their UserData to signal that they have
// started up (either successfully or unsuccessfully).
//
// Signals are used both during intial creation and subsequent updates.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//   var machineImage iMachineImage
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//
//   	// ...
//
//   	Init: ec2.CloudFormationInit_FromElements(ec2.InitFile_FromString(jsii.String("/etc/my_instance"), jsii.String("This got written during instance startup"))),
//   	Signals: autoscaling.Signals_WaitForAll(&SignalsOptions{
//   		Timeout: awscdk.Duration_Minutes(jsii.Number(10)),
//   	}),
//   })
//
type Signals interface {
	// Helper to render the actual creation policy, as the logic between them is quite similar.
	DoRender(options *SignalsOptions, count *float64) *awscdk.CfnCreationPolicy
	// Render the ASG's CreationPolicy.
	RenderCreationPolicy(renderOptions *RenderSignalsOptions) *awscdk.CfnCreationPolicy
}

// The jsii proxy struct for Signals
type jsiiProxy_Signals struct {
	_ byte // padding
}

func NewSignals_Override(s Signals) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.Signals",
		nil, // no parameters
		s,
	)
}

// Wait for the desiredCapacity of the AutoScalingGroup amount of signals to have been received.
//
// If no desiredCapacity has been configured, wait for minCapacity signals intead.
//
// This number is used during initial creation and during replacing updates.
// During rolling updates, all updated instances must send a signal.
func Signals_WaitForAll(options *SignalsOptions) Signals {
	_init_.Initialize()

	if err := validateSignals_WaitForAllParameters(options); err != nil {
		panic(err)
	}
	var returns Signals

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.Signals",
		"waitForAll",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Wait for a specific amount of signals to have been received.
//
// You should send one signal per instance, so this represents the number of
// instances to wait for.
//
// This number is used during initial creation and during replacing updates.
// During rolling updates, all updated instances must send a signal.
func Signals_WaitForCount(count *float64, options *SignalsOptions) Signals {
	_init_.Initialize()

	if err := validateSignals_WaitForCountParameters(count, options); err != nil {
		panic(err)
	}
	var returns Signals

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.Signals",
		"waitForCount",
		[]interface{}{count, options},
		&returns,
	)

	return returns
}

// Wait for the minCapacity of the AutoScalingGroup amount of signals to have been received.
//
// This number is used during initial creation and during replacing updates.
// During rolling updates, all updated instances must send a signal.
func Signals_WaitForMinCapacity(options *SignalsOptions) Signals {
	_init_.Initialize()

	if err := validateSignals_WaitForMinCapacityParameters(options); err != nil {
		panic(err)
	}
	var returns Signals

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.Signals",
		"waitForMinCapacity",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Signals) DoRender(options *SignalsOptions, count *float64) *awscdk.CfnCreationPolicy {
	if err := s.validateDoRenderParameters(options); err != nil {
		panic(err)
	}
	var returns *awscdk.CfnCreationPolicy

	_jsii_.Invoke(
		s,
		"doRender",
		[]interface{}{options, count},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Signals) RenderCreationPolicy(renderOptions *RenderSignalsOptions) *awscdk.CfnCreationPolicy {
	if err := s.validateRenderCreationPolicyParameters(renderOptions); err != nil {
		panic(err)
	}
	var returns *awscdk.CfnCreationPolicy

	_jsii_.Invoke(
		s,
		"renderCreationPolicy",
		[]interface{}{renderOptions},
		&returns,
	)

	return returns
}

