package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Customization options for Signal handling.
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
type SignalsOptions struct {
	// The percentage of signals that need to be successful.
	//
	// If this number is less than 100, a percentage of signals may be failure
	// signals while still succeeding the creation or update in CloudFormation.
	// Default: 100.
	//
	MinSuccessPercentage *float64 `field:"optional" json:"minSuccessPercentage" yaml:"minSuccessPercentage"`
	// How long to wait for the signals to be sent.
	//
	// This should reflect how long it takes your instances to start up
	// (including instance start time and instance initialization time).
	// Default: Duration.minutes(5)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

