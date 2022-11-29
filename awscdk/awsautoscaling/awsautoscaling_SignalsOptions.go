package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Customization options for Signal handling.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//   var machineImage iMachineImage
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	instanceType: instanceType,
//   	machineImage: machineImage,
//
//   	// ...
//
//   	init: ec2.cloudFormationInit.fromElements(ec2.initFile.fromString(jsii.String("/etc/my_instance"), jsii.String("This got written during instance startup"))),
//   	signals: autoscaling.signals.waitForAll(&signalsOptions{
//   		timeout: awscdk.Duration.minutes(jsii.Number(10)),
//   	}),
//   })
//
// Experimental.
type SignalsOptions struct {
	// The percentage of signals that need to be successful.
	//
	// If this number is less than 100, a percentage of signals may be failure
	// signals while still succeeding the creation or update in CloudFormation.
	// Experimental.
	MinSuccessPercentage *float64 `field:"optional" json:"minSuccessPercentage" yaml:"minSuccessPercentage"`
	// How long to wait for the signals to be sent.
	//
	// This should reflect how long it takes your instances to start up
	// (including instance start time and instance initialization time).
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

